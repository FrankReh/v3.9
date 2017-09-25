package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// Target information for a given translation unit.
type TargetInfo struct {
	Triple       string
	PointerWidth int
}

func (tu TranslationUnit) TargetInfo() TargetInfo {
	ti := C.clang_getTranslationUnitTargetInfo(tu.c)
	defer C.clang_TargetInfo_dispose(ti)

	triple := cxstring{C.clang_TargetInfo_getTriple(ti)}
	defer triple.Dispose()

	return TargetInfo{
		Triple:       triple.String(),
		PointerWidth: int(C.clang_TargetInfo_getPointerWidth(ti)),
	}
}
