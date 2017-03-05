// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestUnusedParams(t *testing.T) {
	warns, err := unusedParams("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	wantBs, err := ioutil.ReadFile(filepath.Join("testdata", "log"))
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBs)
	got := strings.Join(warns, "\n") + "\n"
	if got != want {
		t.Fatalf("Unexpected output. Want:\n%sGot:\n%s", want, got)
	}
}
