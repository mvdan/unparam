// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// Package check implements the unparam linter. Note that its API is not
// stable.
package check

import (
	"bytes"
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
	"github.com/mvdan/lint"
)

func UnusedParams(tests bool, args ...string) ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	c := &Checker{wd: wd, tests: tests}
	return c.lines(args...)
}

type Checker struct {
	lprog *loader.Program
	prog  *ssa.Program

	wd  string
	buf bytes.Buffer

	tests bool

	funcSigns map[string]bool
	seenTypes map[types.Type]bool
}

var (
	_ lint.Checker = (*Checker)(nil)
	_ lint.WithSSA = (*Checker)(nil)
)

func (c *Checker) lines(args ...string) ([]string, error) {
	paths := gotool.ImportPaths(args)
	var conf loader.Config
	if _, err := conf.FromArgs(paths, c.tests); err != nil {
		return nil, err
	}
	lprog, err := conf.Load()
	if err != nil {
		return nil, err
	}
	prog := ssautil.CreateProgram(lprog, 0)
	prog.Build()
	c.Program(lprog)
	c.ProgramSSA(prog)
	issues, err := c.Check()
	if err != nil {
		return nil, err
	}
	lines := make([]string, len(issues))
	for i, issue := range issues {
		fpos := prog.Fset.Position(issue.Pos()).String()
		if strings.HasPrefix(fpos, c.wd) {
			fpos = fpos[len(c.wd)+1:]
		}
		lines[i] = fmt.Sprintf("%s: %s", fpos, issue.Message())
	}
	return lines, nil
}

type Issue struct {
	pos token.Pos
	msg string
}

func (i Issue) Pos() token.Pos  { return i.pos }
func (i Issue) Message() string { return i.msg }

func (c *Checker) Program(lprog *loader.Program) {
	c.lprog = lprog
}

func (c *Checker) ProgramSSA(prog *ssa.Program) {
	c.prog = prog
}

func (c *Checker) Check() ([]lint.Issue, error) {
	wantPkg := make(map[*types.Package]bool)
	for _, info := range c.lprog.InitialPackages() {
		wantPkg[info.Pkg] = true
	}

	var potential []*ssa.Parameter
funcLoop:
	for fn := range ssautil.AllFunctions(c.prog) {
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
		if refs := fn.Referrers(); refs != nil {
			for _, instr := range *refs {
				switch instr.(type) {
				case *ssa.Store:
					continue funcLoop
				}
			}
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
	// TODO: replace by sort.Slice once we drop Go 1.7 support
	sort.Sort(byPos(potential))

	addSigns := func(pkg *ssa.Package, onlyExported bool) {
		for _, mb := range pkg.Members {
			if onlyExported && !ast.IsExported(mb.Name()) {
				continue
			}
			c.addSign(mb.Type(), mb.Token() == token.FUNC)
		}
	}

	var curPkg *types.Package
	issues := make([]lint.Issue, 0, len(potential))
	for _, par := range potential {
		pkg := par.Parent().Pkg
		// since they are sorted by position, we will see all
		// the warnings for any package contiguously
		if tpkg := pkg.Pkg; tpkg != curPkg {
			curPkg = tpkg
			c.funcSigns = make(map[string]bool)
			c.seenTypes = make(map[types.Type]bool)
			addSigns(pkg, false)
			for _, imp := range tpkg.Imports() {
				addSigns(c.prog.Package(imp), true)
			}
		}
		sign := par.Parent().Signature
		if c.funcSigns[c.signString(sign)] { // could be required
			continue
		}
		issues = append(issues, Issue{
			pos: par.Pos(),
			msg: fmt.Sprintf("%s is unused", par.Name()),
		})
	}
	return issues, nil
}

type byPos []*ssa.Parameter

func (p byPos) Len() int           { return len(p) }
func (p byPos) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byPos) Less(i, j int) bool { return p[i].Pos() < p[j].Pos() }

func (c *Checker) addSign(t types.Type, ignoreSign bool) {
	if c.seenTypes[t] {
		return
	}
	c.seenTypes[t] = true
	switch x := t.(type) {
	case *types.Signature:
		params := x.Params()
		if params.Len() == 0 {
			break
		}
		if !ignoreSign { // otherwise funcs block themselves
			c.funcSigns[c.signString(x)] = true
		}
		for i := 0; i < params.Len(); i++ {
			c.addSign(params.At(i).Type(), false)
		}
	case *types.Struct:
		for i := 0; i < x.NumFields(); i++ {
			c.addSign(x.Field(i).Type(), false)
		}
	case *types.Named:
		for i := 0; i < x.NumMethods(); i++ {
			c.addSign(x.Method(i).Type(), true)
		}
		c.addSign(t.Underlying(), false)
	case *types.Interface:
		for i := 0; i < x.NumMethods(); i++ {
			c.addSign(x.Method(i).Type(), false)
		}
	case withElem:
		c.addSign(x.Elem(), false)
	}
}

type withElem interface {
	Elem() types.Type
}

var rxHarmlessCall = regexp.MustCompile(`(?i)\b(log(ger)?|errors)\b|\bf?print`)

// dummyImpl reports whether a block is a dummy implementation. This is
// true if the block will almost immediately panic, throw or return
// constants only.
func dummyImpl(blk *ssa.BasicBlock) bool {
	var ops [8]*ssa.Value
	for _, instr := range blk.Instrs {
		for _, val := range instr.Operands(ops[:0]) {
			switch x := (*val).(type) {
			case nil, *ssa.Const, *ssa.ChangeType, *ssa.Alloc,
				*ssa.MakeInterface, *ssa.Function,
				*ssa.Global, *ssa.IndexAddr, *ssa.Slice:
			case *ssa.Call:
				if rxHarmlessCall.MatchString(x.Call.Value.String()) {
					continue
				}
			default:
				return false
			}
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
func (c *Checker) signString(sign *types.Signature) string {
	c.buf.Reset()
	tupleJoin(&c.buf, sign.Params())
	tupleJoin(&c.buf, sign.Results())
	return c.buf.String()
}
