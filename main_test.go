// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestUnusedParams(t *testing.T) {
	var buf bytes.Buffer
	if err := unusedParams(&buf, "./testdata"); err != nil {
		t.Fatal(err)
	}
	wantBs, err := ioutil.ReadFile(filepath.Join("testdata", "log"))
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBs)
	if got := buf.String(); got != want {
		t.Fatalf("Unexpected output. Want:\n%sGot:\n%s", want, got)
	}
}
