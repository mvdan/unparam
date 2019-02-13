// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"os"

	"mvdan.cc/unparam/passes/unparam"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	os.Exit(main1())
}

func main1() int {
	// TODO: make singlechecker.Main return an int instead of using os.Exit.
	singlechecker.Main(unparam.Analyzer)
	return 0
}
