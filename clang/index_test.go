package clang_test

import (
	"strings"
	"testing"

	"github.com/frankreh/go-clang/clang"
)

func TestGlobalOptions(t *testing.T) {
	idx := clang.NewIndex(0, 1)
	defer idx.Dispose()

	globalOptFlags := idx.GlobalOptions()

	assertEqualInt(t, int(globalOptFlags), 0)

	idx.SetGlobalOptions(clang.GlobalOpt_ThreadBackgroundPriorityForIndexing |
		clang.GlobalOpt_ThreadBackgroundPriorityForEditing)
	globalOptFlags = idx.GlobalOptions()

	// Just make sure there are two parts to the string value now.
	assertEqualInt(t, 2, len(strings.Split(globalOptFlags.String(), "|")))

	// Test that unexpected values don't break us.
	max := ^clang.GlobalOptFlags(0)
	assertStringNotEmpty(t, max.String())
}
