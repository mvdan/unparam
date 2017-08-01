// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package check

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

var (
	write = flag.Bool("w", false, "write test outputs")
	debug = flag.Bool("debug", false, "debug prints")
)

func TestUnusedParams(t *testing.T) {
	warns, err := UnusedParams(true, *debug, "./testdata")
	if err != nil {
		t.Fatal(err)
	}
	logPath := filepath.Join("testdata", "log")
	got := strings.Join(warns, "\n") + "\n"
	if *write {
		err := ioutil.WriteFile(logPath, []byte(got), 0644)
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	wantBs, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBs)
	if got != want {
		t.Fatalf("Unexpected output. Want:\n%sGot:\n%s", want, got)
	}
}
