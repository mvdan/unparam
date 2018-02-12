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

func TestCHA(t *testing.T) {
	warns, err := UnusedParams(true, "cha", false, *debug,
		"./testdata",
		"./testdata/main",
	)
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

func TestRTA(t *testing.T) {
	warns, err := UnusedParams(true, "rta", false, *debug,
		"./testdata/main",
	)
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(warns, "\n")
	want := strings.TrimSpace(`
testdata/main/main.go:3:19: OneUnused - b is unused
testdata/main/main.go:10:24: mightImplement - b is unused
	`)
	if got != want {
		t.Fatalf("want:\n%s\ngot:\n%s", want, got)
	}
}
