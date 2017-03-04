// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"flag"
	"fmt"
	"go/types"
	"os"
	"strings"
	"unicode"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa/ssautil"

	"github.com/kisielk/gotool"
)

func main() {
	flag.Parse()
	if err := unusedParams(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func isExported(name string) bool {
	for _, r := range name {
		return unicode.IsUpper(r)
	}
	return false
}

func unusedParams(args []string) error {
	paths := gotool.ImportPaths(args)
	var conf loader.Config
	_, err := conf.FromArgs(paths, false)
	if err != nil {
		return err
	}
	iprog, err := conf.Load()
	if err != nil {
		return err
	}
	pkgInfos := make(map[*types.Package]*types.Info)
	for _, pinfo := range iprog.InitialPackages() {
		pkgInfos[pinfo.Pkg] = &pinfo.Info
	}
	prog := ssautil.CreateProgram(iprog, 0)
	prog.Build()

	ifaceFuncs := make(map[string]bool)
	for _, pkg := range prog.AllPackages() {
		for _, member := range pkg.Members {
			if !isExported(member.Name()) {
				continue
			}
			under := member.Type().Underlying()
			iface, ok := under.(*types.Interface)
			if !ok {
				continue
			}
			for i := 0; i < iface.NumMethods(); i++ {
				m := iface.Method(i)
				sign := m.Type().(*types.Signature)
				if sign.Params().Len() == 0 {
					continue
				}
				ifaceFuncs[signString(sign)] = true
			}
		}
	}

	for fn := range ssautil.AllFunctions(prog) {
		if fn.Pkg == nil { // builtin?
			continue
		}
		if fn.Blocks == nil { // stub
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
		for i, param := range fn.Params {
			if i == 0 && sign.Recv() != nil { // receiver, not param
				continue
			}
			if len(*param.Referrers()) > 0 {
				continue
			}
			pos := prog.Fset.Position(param.Pos())
			fmt.Printf("%v: %s is unused\n", pos, param.Name())
		}
	}
	return nil
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
