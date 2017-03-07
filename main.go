// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"

	"github.com/kisielk/gotool"
)

func main() {
	flag.Parse()
	warns, err := unusedParams(flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, warn := range warns {
		fmt.Println(warn)
	}
}

func unusedParams(args ...string) ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	l := &linter{wd: wd}
	return l.warns(args...)
}

type linter struct {
	wd  string
	buf bytes.Buffer

	funcSigns map[string]bool
	seenTypes map[types.Type]bool
}

func (l *linter) warns(args ...string) ([]string, error) {
	paths := gotool.ImportPaths(args)
	var conf loader.Config
	if _, err := conf.FromArgs(paths, false); err != nil {
		return nil, err
	}
	lprog, err := conf.Load()
	if err != nil {
		return nil, err
	}
	wantPkg := make(map[*types.Package]bool)
	for _, info := range lprog.InitialPackages() {
		wantPkg[info.Pkg] = true
	}
	prog := ssautil.CreateProgram(lprog, 0)
	prog.Build()

	var potential []*ssa.Parameter
	for fn := range ssautil.AllFunctions(prog) {
		if fn.Pkg == nil { // builtin?
			continue
		}
		if len(fn.Blocks) == 0 { // stub
			continue
		}
		if !wantPkg[fn.Pkg.Pkg] { // not part of given pkgs
			continue
		}
		if dummyImpl(fn.Blocks[0]) { // panic implementation
			continue
		}
		for i, par := range fn.Params {
			if i == 0 && fn.Signature.Recv() != nil { // receiver
				continue
			}
			switch par.Object().Name() {
			case "", "_": // unnamed
				continue
			}
			if len(*par.Referrers()) > 0 { // used
				continue
			}
			potential = append(potential, par)
		}

	}
	sort.Slice(potential, func(i, j int) bool {
		return potential[i].Pos() < potential[j].Pos()
	})

	addSigns := func(pkg *ssa.Package, onlyExported bool) {
		for _, mb := range pkg.Members {
			if onlyExported && !ast.IsExported(mb.Name()) {
				continue
			}
			l.addSign(mb.Type(), mb.Token() == token.FUNC)
		}
	}

	var curPkg *types.Package
	warns := make([]string, 0, len(potential))
	for _, par := range potential {
		pkg := par.Parent().Pkg
		// since they are sorted by position, we will see all
		// the warnings for any package contiguously
		if tpkg := pkg.Pkg; tpkg != curPkg {
			curPkg = tpkg
			l.funcSigns = make(map[string]bool)
			l.seenTypes = make(map[types.Type]bool)
			addSigns(pkg, false)
			for _, imp := range tpkg.Imports() {
				addSigns(prog.Package(imp), true)
			}
		}
		sign := par.Parent().Signature
		if l.funcSigns[l.signString(sign)] { // could be required
			continue
		}
		line := prog.Fset.Position(par.Pos()).String()
		if strings.HasPrefix(line, l.wd) {
			line = line[len(l.wd)+1:]
		}
		warns = append(warns, fmt.Sprintf("%s: %s is unused",
			line, par.Name()))
	}
	return warns, nil
}

func (l *linter) addSign(t types.Type, ignoreSign bool) {
	if l.seenTypes[t] {
		return
	}
	l.seenTypes[t] = true
	switch x := t.(type) {
	case *types.Signature:
		params := x.Params()
		if params.Len() == 0 {
			break
		}
		if !ignoreSign { // otherwise funcs block themselves
			l.funcSigns[l.signString(x)] = true
		}
		for i := 0; i < params.Len(); i++ {
			l.addSign(params.At(i).Type(), false)
		}
	case *types.Struct:
		for i := 0; i < x.NumFields(); i++ {
			l.addSign(x.Field(i).Type(), false)
		}
	case *types.Named:
		for i := 0; i < x.NumMethods(); i++ {
			l.addSign(x.Method(i).Type(), true)
		}
	case *types.Interface:
		for i := 0; i < x.NumMethods(); i++ {
			l.addSign(x.Method(i).Type(), false)
		}
	}
	l.addSign(t.Underlying(), false)
}

var rxHarmlessCall = regexp.MustCompile(`(?i)\blog(ger)?\b|\bf?print`)

// dummyImpl reports whether a block is a dummy implementation. This is
// true if the block will almost immediately panic, throw or return
// constants only.
func dummyImpl(blk *ssa.BasicBlock) bool {
	for _, instr := range blk.Instrs {
		if !harmlessIfOps(instr) {
			return false
		}
		switch x := instr.(type) {
		case *ssa.Alloc, *ssa.Store, *ssa.UnOp, *ssa.BinOp,
			*ssa.MakeInterface, *ssa.MakeMap, *ssa.Extract,
			*ssa.IndexAddr, *ssa.FieldAddr, *ssa.Slice,
			*ssa.Lookup, *ssa.ChangeType, *ssa.TypeAssert,
			*ssa.Convert, *ssa.ChangeInterface:
			// non-trivial expressions in panic/log/print
			// calls
		case *ssa.Return, *ssa.Panic:
			return true
		case *ssa.Call:
			if rxHarmlessCall.MatchString(x.Call.Value.String()) {
				continue
			}
			return x.Call.Value.Name() == "throw" // runtime's panic
		default:
			return false
		}
	}
	return false
}

func harmlessIfOps(instr ssa.Instruction) bool {
	withOps, ok := instr.(interface{
		Operands([]*ssa.Value) []*ssa.Value
	})
	if !ok {
		return true
	}
	var ops [8]*ssa.Value
	for _, val := range withOps.Operands(ops[:0]) {
		switch (*val).(type) {
		case nil, *ssa.Const, *ssa.ChangeType, *ssa.Alloc,
			*ssa.MakeInterface, *ssa.Function,
			*ssa.Global, *ssa.IndexAddr, *ssa.Slice:
		default:
			return false
		}
	}
	return true
}

func tupleJoin(buf *bytes.Buffer, t *types.Tuple) {
	buf.WriteByte('(')
	for i := 0; i < t.Len(); i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(t.At(i).Type().String())
	}
	buf.WriteByte(')')
}

// signString is similar to Signature.String(), but it ignores
// param/result names.
func (l *linter) signString(sign *types.Signature) string {
	l.buf.Reset()
	tupleJoin(&l.buf, sign.Params())
	tupleJoin(&l.buf, sign.Results())
	return l.buf.String()
}
