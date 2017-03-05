// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"flag"
	"fmt"
	"go/token"
	"go/types"
	"os"
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
	paths := gotool.ImportPaths(args)
	var conf loader.Config
	if _, err := conf.FromArgs(paths, false); err != nil {
		return nil, err
	}
	iprog, err := conf.Load()
	if err != nil {
		return nil, err
	}
	pkgInfos := make(map[*types.Package]*types.Info)
	for _, pinfo := range iprog.InitialPackages() {
		pkgInfos[pinfo.Pkg] = &pinfo.Info
	}
	prog := ssautil.CreateProgram(iprog, 0)
	prog.Build()

	ifaceFuncs := make(map[string]bool)
	addSign := func(sign *types.Signature) {
		if sign.Params().Len() == 0 {
			return
		}
		ifaceFuncs[signString(sign)] = true
	}
	for _, pkg := range prog.AllPackages() {
		for _, member := range pkg.Members {
			switch member.Token() {
			case token.FUNC:
				params := member.Type().(*types.Signature).Params()
				for i := 0; i < params.Len(); i++ {
					p := params.At(i)
					sign, ok := p.Type().(*types.Signature)
					if ok {
						addSign(sign)
					}
				}
				continue
			case token.TYPE:
			default:
				continue
			}
			switch x := member.Type().Underlying().(type) {
			case *types.Struct:
				for i := 0; i < x.NumFields(); i++ {
					f := x.Field(i)
					sign, ok := f.Type().(*types.Signature)
					if ok {
						addSign(sign)
					}
				}
			case *types.Interface:
				for i := 0; i < x.NumMethods(); i++ {
					m := x.Method(i)
					addSign(m.Type().(*types.Signature))
				}
			case *types.Signature:
				addSign(x)
			}
		}
	}

	var warns []string
	for fn := range ssautil.AllFunctions(prog) {
		if fn.Pkg == nil { // builtin?
			continue
		}
		if len(fn.Blocks) == 0 { // stub
			continue
		}
		info := pkgInfos[fn.Pkg.Pkg]
		if info == nil { // not part of given pkgs
			continue
		}
		sign := fn.Signature
		if ifaceFuncs[signString(sign)] { // could implement iface
			continue
		}
		blk := fn.Blocks[0]
		if ret, ok := blk.Instrs[0].(*ssa.Return); ok &&
			len(ret.Results) == 0 { // dummy implementation
			continue
		}
		if willPanic(blk) { // panic implementation
			continue
		}
		for i, param := range fn.Params {
			if i == 0 && sign.Recv() != nil { // receiver, not param
				continue
			}
			switch param.Object().Name() {
			case "", "_": // unnamed
				continue
			}
			if len(*param.Referrers()) > 0 { // used
				continue
			}
			pos := prog.Fset.Position(param.Pos())
			line := pos.String()
			if strings.HasPrefix(line, wd) {
				line = line[len(wd)+1:]
			}
			warns = append(warns, fmt.Sprintf("%s: %s is unused",
				line, param.Name()))
		}
	}
	sort.Strings(warns)
	return warns, nil
}

// willPanic reports whether a block will just panic. We can't simply
// use the first instruction as there might be others before it, like
// MakeInterface.
func willPanic(blk *ssa.BasicBlock) bool {
	for _, instr := range blk.Instrs {
		switch x := instr.(type) {
		case *ssa.Panic:
			return true
		case *ssa.Call:
			if x.Call.Value.Name() == "throw" { // runtime's panic
				return true
			}
		}
	}
	return false
}

func tupleStrs(t *types.Tuple) []string {
	l := make([]string, t.Len())
	for i := 0; i < t.Len(); i++ {
		l[i] = t.At(i).Type().String()
	}
	return l
}

// signString is similar to Signature.String(), but it ignores
// param/result names.
func signString(sign *types.Signature) string {
	return fmt.Sprintf("(%s) (%s)",
		strings.Join(tupleStrs(sign.Params()), ", "),
		strings.Join(tupleStrs(sign.Results()), ", "))
}
