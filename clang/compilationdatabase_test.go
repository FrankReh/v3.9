package clang_test

import (
	"testing"

	"github.com/frankreh/go-clang/clang"
)

func TestCompilationDatabaseError(t *testing.T) {
	// This isn't a long test, but the error message that libclang puts out to the console
	// is very distracting, so skip it during short tests.
	if testing.Short() {
		t.Skip("skipping test inn short mode.")
	}
	_, err := clang.FromDirectory("../testdata-not-there")
	if err != clang.CanNotLoadDatabaseErr {
		t.Fatalf("expected %v", clang.CanNotLoadDatabaseErr)
	}
}

func TestCompilationDatabase(t *testing.T) {
	db, err := clang.FromDirectory("../testdata")
	if err != nil {
		t.Fatalf("error loading compilation database: %v", err)
	}
	defer db.Dispose()

	table := []struct {
		directory string
		args      []string
	}{
		{
			directory: "/home/user/llvm/build",
			args: []string{
				"/usr/bin/clang++",
				"--driver-mode=g++", // as of clang 10.0, this is added automatically anyway
				"-Irelative",
				//FIXME: bug in clang ?
				//`-DSOMEDEF="With spaces, quotes and \-es.`,
				"-DSOMEDEF=With spaces, quotes and -es.",
				"-c",
				"-o",
				"file.o",
				"file.cc",
			},
		},
		{
			directory: "@TESTDIR@",
			args:      []string{"g++", "--driver-mode=g++", "-c", "-DMYMACRO=a", "subdir/a.cpp"},
		},
	}

	cmds := db.AllCompileCommands()
	if len(cmds) != len(table) {
		t.Fatalf("expected #cmds=%d. got=%d", len(table), len(cmds))
	}

	for i, cmd := range cmds {
		if cmd.Directory != table[i].directory {
			t.Fatalf("expected dir=%q. got=%q", table[i].directory, cmd.Directory)
		}

		nargs := len(cmd.Args)
		if nargs != len(table[i].args) {
			t.Logf("i %d", i)
			t.Logf("table[i].args: %q", table[i].args)
			t.Logf("cmd.Args: %q", cmd.Args)
			t.Fatalf("expected #args=%d. got=%d", len(table[i].args), nargs)
		}
		if nargs > len(table[i].args) {
			nargs = len(table[i].args)
		}
		for j, arg := range cmd.Args {
			if arg != table[i].args[j] {
				t.Fatalf("expected arg[%d]=%q. got=%q", j, table[i].args[j], arg)
			}
		}
	}
}
