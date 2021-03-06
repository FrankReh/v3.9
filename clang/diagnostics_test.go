package clang_test

import (
	"strings"
	"testing"

	"github.com/frankreh/go-clang/clang"
)

func TestDiagnostics(t *testing.T) {
	idx := clang.NewIndex(0, 0)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit("go-clang.c", nil, nil, 0)
	assertTrue(t, tu.IsValid())
	defer tu.Dispose()

	diags := tu.Diagnostics()

	ok := false
	for _, d := range diags {
		if strings.Contains(d.Spelling(), "_cgo_export.h") {
			ok = true
		}
		t.Log(d)
		t.Log(d.Severity(), d.Spelling())
		t.Log(d.FormatDiagnostic(clang.Diagnostic_DisplayCategoryName | clang.Diagnostic_DisplaySourceLocation))
	}
	assertTrue(t, ok)

	// Test that unexpected values don't break us.
	max := ^clang.DiagnosticDisplayOptions(0)
	assertStringNotEmpty(t, max.String())
}
