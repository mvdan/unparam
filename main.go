// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"flag"
	"fmt"
	"go/types"
	"os"

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
	pkgs := make(map[*types.Package]bool)
	for _, pinfo := range iprog.InitialPackages() {
		pkgs[pinfo.Pkg] = true
	}

	prog := ssautil.CreateProgram(iprog, 0)

	prog.Build()
	for fn := range ssautil.AllFunctions(prog) {
		if fn.Pkg == nil || !pkgs[fn.Pkg.Pkg] {
			continue
		}
		fmt.Println(fn)
	}
	return nil
}
