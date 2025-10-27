package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/rogpeppe/go-internal/gotooltest"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"unparam": main,
	})
}

var update = flag.Bool("u", false, "update testscripts")

func TestScript(t *testing.T) {
	t.Parallel()
	p := testscript.Params{
		Dir:                 filepath.Join("testdata", "script"),
		RequireExplicitExec: true,
		UpdateScripts:       *update,
		Setup: func(env *testscript.Env) error {
			env.Vars = append(env.Vars, "/="+string(os.PathSeparator))
			return nil
		},
	}
	if err := gotooltest.Setup(&p); err != nil {
		t.Fatal(err)
	}
	testscript.Run(t, p)
}
