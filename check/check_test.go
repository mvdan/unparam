// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package check

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

var (
	write = flag.Bool("w", false, "write test outputs")
	debug = flag.Bool("debug", false, "debug prints")
)

func TestCHA(t *testing.T) {
	got, err := UnusedParams(true, "cha", false, *debug,
		"./testdata",
		"./testdata/main",
	)
	if err != nil {
		t.Fatal(err)
	}
	logPath := filepath.Join("testdata", "log")
	if *write {
		body := strings.Join(got, "\n") + "\n"
		err := ioutil.WriteFile(logPath, []byte(body), 0644)
		if err != nil {
			t.Fatal(err)
		}
		return
	}
	wantBs, err := ioutil.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}
	want := strings.Split(string(wantBs), "\n")
	if want[len(want)-1] == "" { // for the trailing newline
		want = want[:len(want)-1]
	}
	if diff := lineDiff(want, got); diff != "" {
		t.Fatalf("Unexpected output:\n%s", diff)
	}
}

// lineDiff returns a diff between two lists of lines. Position information is
// not recorded, and the differences are always reported in lexicographic order.
func lineDiff(want, got []string) string {
	diff := make(map[string]int)
	for _, line := range want {
		diff[line]--
	}
	for _, line := range got {
		diff[line]++
	}
	var lines []string
	for line, val := range diff {
		op := "+"
		if val < 0 {
			op = "-"
			val = -val
		}
		for i := 0; i < val; i++ {
			lines = append(lines, op+line)
		}
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i][1:] < lines[j][1:]
	})
	return strings.Join(lines, "\n")
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
