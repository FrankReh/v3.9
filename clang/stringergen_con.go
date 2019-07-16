// Code generated by "stringer -output=stringergen_con.go -type=AccessSpecifier,AvailabilityKind,CallingConv,ChildVisitResult,CommentInlineCommandRenderKind,CommentKind,CommentParamPassDirection,CompletionChunkKind,ExceptionSpecification,DiagnosticSeverity,EvalResultKind,IdxAttrKind,IdxEntityCXXTemplateKind,IdxEntityKind,IdxEntityLanguage,IdxEntityRefKind,IdxObjCContainerKind,IndexOptFlags,LanguageKind,LinkageKind,NameRefFlags,RefQualifierKind,Reparse_Flags,Result,SaveTranslationUnit_Flags,StorageClass,TemplateArgumentKind,TUResourceUsageKind,VisibilityKind,VisitorResult"; DO NOT EDIT.

package clang

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AccessSpecifier_Invalid-0]
	_ = x[AccessSpecifier_Public-1]
	_ = x[AccessSpecifier_Protected-2]
	_ = x[AccessSpecifier_Private-3]
}

const _AccessSpecifier_name = "AccessSpecifier_InvalidAccessSpecifier_PublicAccessSpecifier_ProtectedAccessSpecifier_Private"

var _AccessSpecifier_index = [...]uint8{0, 23, 45, 70, 93}

func (i AccessSpecifier) String() string {
	if i >= AccessSpecifier(len(_AccessSpecifier_index)-1) {
		return "AccessSpecifier(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AccessSpecifier_name[_AccessSpecifier_index[i]:_AccessSpecifier_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Availability_Available-0]
	_ = x[Availability_Deprecated-1]
	_ = x[Availability_NotAvailable-2]
	_ = x[Availability_NotAccessible-3]
}

const _AvailabilityKind_name = "Availability_AvailableAvailability_DeprecatedAvailability_NotAvailableAvailability_NotAccessible"

var _AvailabilityKind_index = [...]uint8{0, 22, 45, 70, 96}

func (i AvailabilityKind) String() string {
	if i >= AvailabilityKind(len(_AvailabilityKind_index)-1) {
		return "AvailabilityKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AvailabilityKind_name[_AvailabilityKind_index[i]:_AvailabilityKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CallingConv_Default-0]
	_ = x[CallingConv_C-1]
	_ = x[CallingConv_X86StdCall-2]
	_ = x[CallingConv_X86FastCall-3]
	_ = x[CallingConv_X86ThisCall-4]
	_ = x[CallingConv_X86Pascal-5]
	_ = x[CallingConv_AAPCS-6]
	_ = x[CallingConv_AAPCS_VFP-7]
	_ = x[CallingConv_X86RegCall-8]
	_ = x[CallingConv_IntelOclBicc-9]
	_ = x[CallingConv_Win64-10]
	_ = x[CallingConv_X86_64Win64-10]
	_ = x[CallingConv_X86_64SysV-11]
	_ = x[CallingConv_X86VectorCall-12]
	_ = x[CallingConv_Swift-13]
	_ = x[CallingConv_PreserveMost-14]
	_ = x[CallingConv_PreserveAll-15]
	_ = x[CallingConv_AArch64VectorCall-16]
	_ = x[CallingConv_Invalid-100]
	_ = x[CallingConv_Unexposed-200]
}

const (
	_CallingConv_name_0 = "CallingConv_DefaultCallingConv_CCallingConv_X86StdCallCallingConv_X86FastCallCallingConv_X86ThisCallCallingConv_X86PascalCallingConv_AAPCSCallingConv_AAPCS_VFPCallingConv_X86RegCallCallingConv_IntelOclBiccCallingConv_Win64CallingConv_X86_64SysVCallingConv_X86VectorCallCallingConv_SwiftCallingConv_PreserveMostCallingConv_PreserveAllCallingConv_AArch64VectorCall"
	_CallingConv_name_1 = "CallingConv_Invalid"
	_CallingConv_name_2 = "CallingConv_Unexposed"
)

var (
	_CallingConv_index_0 = [...]uint16{0, 19, 32, 54, 77, 100, 121, 138, 159, 181, 205, 222, 244, 269, 286, 310, 333, 362}
)

func (i CallingConv) String() string {
	switch {
	case 0 <= i && i <= 16:
		return _CallingConv_name_0[_CallingConv_index_0[i]:_CallingConv_index_0[i+1]]
	case i == 100:
		return _CallingConv_name_1
	case i == 200:
		return _CallingConv_name_2
	default:
		return "CallingConv(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ChildVisit_Break-0]
	_ = x[ChildVisit_Continue-1]
	_ = x[ChildVisit_Recurse-2]
}

const _ChildVisitResult_name = "ChildVisit_BreakChildVisit_ContinueChildVisit_Recurse"

var _ChildVisitResult_index = [...]uint8{0, 16, 35, 53}

func (i ChildVisitResult) String() string {
	if i >= ChildVisitResult(len(_ChildVisitResult_index)-1) {
		return "ChildVisitResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ChildVisitResult_name[_ChildVisitResult_index[i]:_ChildVisitResult_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CommentInlineCommandRenderKind_Normal-0]
	_ = x[CommentInlineCommandRenderKind_Bold-1]
	_ = x[CommentInlineCommandRenderKind_Monospaced-2]
	_ = x[CommentInlineCommandRenderKind_Emphasized-3]
}

const _CommentInlineCommandRenderKind_name = "CommentInlineCommandRenderKind_NormalCommentInlineCommandRenderKind_BoldCommentInlineCommandRenderKind_MonospacedCommentInlineCommandRenderKind_Emphasized"

var _CommentInlineCommandRenderKind_index = [...]uint8{0, 37, 72, 113, 154}

func (i CommentInlineCommandRenderKind) String() string {
	if i >= CommentInlineCommandRenderKind(len(_CommentInlineCommandRenderKind_index)-1) {
		return "CommentInlineCommandRenderKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommentInlineCommandRenderKind_name[_CommentInlineCommandRenderKind_index[i]:_CommentInlineCommandRenderKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Comment_Null-0]
	_ = x[Comment_Text-1]
	_ = x[Comment_InlineCommand-2]
	_ = x[Comment_HTMLStartTag-3]
	_ = x[Comment_HTMLEndTag-4]
	_ = x[Comment_Paragraph-5]
	_ = x[Comment_BlockCommand-6]
	_ = x[Comment_ParamCommand-7]
	_ = x[Comment_TParamCommand-8]
	_ = x[Comment_VerbatimBlockCommand-9]
	_ = x[Comment_VerbatimBlockLine-10]
	_ = x[Comment_VerbatimLine-11]
	_ = x[Comment_FullComment-12]
}

const _CommentKind_name = "Comment_NullComment_TextComment_InlineCommandComment_HTMLStartTagComment_HTMLEndTagComment_ParagraphComment_BlockCommandComment_ParamCommandComment_TParamCommandComment_VerbatimBlockCommandComment_VerbatimBlockLineComment_VerbatimLineComment_FullComment"

var _CommentKind_index = [...]uint8{0, 12, 24, 45, 65, 83, 100, 120, 140, 161, 189, 214, 234, 253}

func (i CommentKind) String() string {
	if i >= CommentKind(len(_CommentKind_index)-1) {
		return "CommentKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommentKind_name[_CommentKind_index[i]:_CommentKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CommentParamPassDirection_In-0]
	_ = x[CommentParamPassDirection_Out-1]
	_ = x[CommentParamPassDirection_InOut-2]
}

const _CommentParamPassDirection_name = "CommentParamPassDirection_InCommentParamPassDirection_OutCommentParamPassDirection_InOut"

var _CommentParamPassDirection_index = [...]uint8{0, 28, 57, 88}

func (i CommentParamPassDirection) String() string {
	if i >= CommentParamPassDirection(len(_CommentParamPassDirection_index)-1) {
		return "CommentParamPassDirection(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommentParamPassDirection_name[_CommentParamPassDirection_index[i]:_CommentParamPassDirection_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CompletionChunk_Optional-0]
	_ = x[CompletionChunk_TypedText-1]
	_ = x[CompletionChunk_Text-2]
	_ = x[CompletionChunk_Placeholder-3]
	_ = x[CompletionChunk_Informative-4]
	_ = x[CompletionChunk_CurrentParameter-5]
	_ = x[CompletionChunk_LeftParen-6]
	_ = x[CompletionChunk_RightParen-7]
	_ = x[CompletionChunk_LeftBracket-8]
	_ = x[CompletionChunk_RightBracket-9]
	_ = x[CompletionChunk_LeftBrace-10]
	_ = x[CompletionChunk_RightBrace-11]
	_ = x[CompletionChunk_LeftAngle-12]
	_ = x[CompletionChunk_RightAngle-13]
	_ = x[CompletionChunk_Comma-14]
	_ = x[CompletionChunk_ResultType-15]
	_ = x[CompletionChunk_Colon-16]
	_ = x[CompletionChunk_SemiColon-17]
	_ = x[CompletionChunk_Equal-18]
	_ = x[CompletionChunk_HorizontalSpace-19]
	_ = x[CompletionChunk_VerticalSpace-20]
}

const _CompletionChunkKind_name = "CompletionChunk_OptionalCompletionChunk_TypedTextCompletionChunk_TextCompletionChunk_PlaceholderCompletionChunk_InformativeCompletionChunk_CurrentParameterCompletionChunk_LeftParenCompletionChunk_RightParenCompletionChunk_LeftBracketCompletionChunk_RightBracketCompletionChunk_LeftBraceCompletionChunk_RightBraceCompletionChunk_LeftAngleCompletionChunk_RightAngleCompletionChunk_CommaCompletionChunk_ResultTypeCompletionChunk_ColonCompletionChunk_SemiColonCompletionChunk_EqualCompletionChunk_HorizontalSpaceCompletionChunk_VerticalSpace"

var _CompletionChunkKind_index = [...]uint16{0, 24, 49, 69, 96, 123, 155, 180, 206, 233, 261, 286, 312, 337, 363, 384, 410, 431, 456, 477, 508, 537}

func (i CompletionChunkKind) String() string {
	if i >= CompletionChunkKind(len(_CompletionChunkKind_index)-1) {
		return "CompletionChunkKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompletionChunkKind_name[_CompletionChunkKind_index[i]:_CompletionChunkKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ExceptionSpecification_NonFunction - -1]
	_ = x[ExceptionSpecification_None-0]
	_ = x[ExceptionSpecification_DynamicNone-1]
	_ = x[ExceptionSpecification_Dynamic-2]
	_ = x[ExceptionSpecification_MSAny-3]
	_ = x[ExceptionSpecification_BasicNoexcept-4]
	_ = x[ExceptionSpecification_ComputedNoexcept-5]
	_ = x[ExceptionSpecification_Unevaluated-6]
	_ = x[ExceptionSpecification_Uninstantiated-7]
	_ = x[ExceptionSpecification_Unparsed-8]
}

const _ExceptionSpecification_name = "ExceptionSpecification_NonFunctionExceptionSpecification_NoneExceptionSpecification_DynamicNoneExceptionSpecification_DynamicExceptionSpecification_MSAnyExceptionSpecification_BasicNoexceptExceptionSpecification_ComputedNoexceptExceptionSpecification_UnevaluatedExceptionSpecification_UninstantiatedExceptionSpecification_Unparsed"

var _ExceptionSpecification_index = [...]uint16{0, 34, 61, 95, 125, 153, 189, 228, 262, 299, 330}

func (i ExceptionSpecification) String() string {
	i -= -1
	if i < 0 || i >= ExceptionSpecification(len(_ExceptionSpecification_index)-1) {
		return "ExceptionSpecification(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _ExceptionSpecification_name[_ExceptionSpecification_index[i]:_ExceptionSpecification_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Diagnostic_Ignored-0]
	_ = x[Diagnostic_Note-1]
	_ = x[Diagnostic_Warning-2]
	_ = x[Diagnostic_Error-3]
	_ = x[Diagnostic_Fatal-4]
}

const _DiagnosticSeverity_name = "Diagnostic_IgnoredDiagnostic_NoteDiagnostic_WarningDiagnostic_ErrorDiagnostic_Fatal"

var _DiagnosticSeverity_index = [...]uint8{0, 18, 33, 51, 67, 83}

func (i DiagnosticSeverity) String() string {
	if i >= DiagnosticSeverity(len(_DiagnosticSeverity_index)-1) {
		return "DiagnosticSeverity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DiagnosticSeverity_name[_DiagnosticSeverity_index[i]:_DiagnosticSeverity_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Eval_Int-1]
	_ = x[Eval_Float-2]
	_ = x[Eval_ObjCStrLiteral-3]
	_ = x[Eval_StrLiteral-4]
	_ = x[Eval_CFStr-5]
	_ = x[Eval_Other-6]
	_ = x[Eval_UnExposed-0]
}

const _EvalResultKind_name = "Eval_UnExposedEval_IntEval_FloatEval_ObjCStrLiteralEval_StrLiteralEval_CFStrEval_Other"

var _EvalResultKind_index = [...]uint8{0, 14, 22, 32, 51, 66, 76, 86}

func (i EvalResultKind) String() string {
	if i >= EvalResultKind(len(_EvalResultKind_index)-1) {
		return "EvalResultKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EvalResultKind_name[_EvalResultKind_index[i]:_EvalResultKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IdxAttr_Unexposed-0]
	_ = x[IdxAttr_IBAction-1]
	_ = x[IdxAttr_IBOutlet-2]
	_ = x[IdxAttr_IBOutletCollection-3]
}

const _IdxAttrKind_name = "IdxAttr_UnexposedIdxAttr_IBActionIdxAttr_IBOutletIdxAttr_IBOutletCollection"

var _IdxAttrKind_index = [...]uint8{0, 17, 33, 49, 75}

func (i IdxAttrKind) String() string {
	if i >= IdxAttrKind(len(_IdxAttrKind_index)-1) {
		return "IdxAttrKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxAttrKind_name[_IdxAttrKind_index[i]:_IdxAttrKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IdxEntity_NonTemplate-0]
	_ = x[IdxEntity_Template-1]
	_ = x[IdxEntity_TemplatePartialSpecialization-2]
	_ = x[IdxEntity_TemplateSpecialization-3]
}

const _IdxEntityCXXTemplateKind_name = "IdxEntity_NonTemplateIdxEntity_TemplateIdxEntity_TemplatePartialSpecializationIdxEntity_TemplateSpecialization"

var _IdxEntityCXXTemplateKind_index = [...]uint8{0, 21, 39, 78, 110}

func (i IdxEntityCXXTemplateKind) String() string {
	if i >= IdxEntityCXXTemplateKind(len(_IdxEntityCXXTemplateKind_index)-1) {
		return "IdxEntityCXXTemplateKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxEntityCXXTemplateKind_name[_IdxEntityCXXTemplateKind_index[i]:_IdxEntityCXXTemplateKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IdxEntity_Unexposed-0]
	_ = x[IdxEntity_Typedef-1]
	_ = x[IdxEntity_Function-2]
	_ = x[IdxEntity_Variable-3]
	_ = x[IdxEntity_Field-4]
	_ = x[IdxEntity_EnumConstant-5]
	_ = x[IdxEntity_ObjCClass-6]
	_ = x[IdxEntity_ObjCProtocol-7]
	_ = x[IdxEntity_ObjCCategory-8]
	_ = x[IdxEntity_ObjCInstanceMethod-9]
	_ = x[IdxEntity_ObjCClassMethod-10]
	_ = x[IdxEntity_ObjCProperty-11]
	_ = x[IdxEntity_ObjCIvar-12]
	_ = x[IdxEntity_Enum-13]
	_ = x[IdxEntity_Struct-14]
	_ = x[IdxEntity_Union-15]
	_ = x[IdxEntity_CXXClass-16]
	_ = x[IdxEntity_CXXNamespace-17]
	_ = x[IdxEntity_CXXNamespaceAlias-18]
	_ = x[IdxEntity_CXXStaticVariable-19]
	_ = x[IdxEntity_CXXStaticMethod-20]
	_ = x[IdxEntity_CXXInstanceMethod-21]
	_ = x[IdxEntity_CXXConstructor-22]
	_ = x[IdxEntity_CXXDestructor-23]
	_ = x[IdxEntity_CXXConversionFunction-24]
	_ = x[IdxEntity_CXXTypeAlias-25]
	_ = x[IdxEntity_CXXInterface-26]
}

const _IdxEntityKind_name = "IdxEntity_UnexposedIdxEntity_TypedefIdxEntity_FunctionIdxEntity_VariableIdxEntity_FieldIdxEntity_EnumConstantIdxEntity_ObjCClassIdxEntity_ObjCProtocolIdxEntity_ObjCCategoryIdxEntity_ObjCInstanceMethodIdxEntity_ObjCClassMethodIdxEntity_ObjCPropertyIdxEntity_ObjCIvarIdxEntity_EnumIdxEntity_StructIdxEntity_UnionIdxEntity_CXXClassIdxEntity_CXXNamespaceIdxEntity_CXXNamespaceAliasIdxEntity_CXXStaticVariableIdxEntity_CXXStaticMethodIdxEntity_CXXInstanceMethodIdxEntity_CXXConstructorIdxEntity_CXXDestructorIdxEntity_CXXConversionFunctionIdxEntity_CXXTypeAliasIdxEntity_CXXInterface"

var _IdxEntityKind_index = [...]uint16{0, 19, 36, 54, 72, 87, 109, 128, 150, 172, 200, 225, 247, 265, 279, 295, 310, 328, 350, 377, 404, 429, 456, 480, 503, 534, 556, 578}

func (i IdxEntityKind) String() string {
	if i >= IdxEntityKind(len(_IdxEntityKind_index)-1) {
		return "IdxEntityKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxEntityKind_name[_IdxEntityKind_index[i]:_IdxEntityKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IdxEntityLang_None-0]
	_ = x[IdxEntityLang_C-1]
	_ = x[IdxEntityLang_ObjC-2]
	_ = x[IdxEntityLang_CXX-3]
	_ = x[IdxEntityLang_Swift-4]
}

const _IdxEntityLanguage_name = "IdxEntityLang_NoneIdxEntityLang_CIdxEntityLang_ObjCIdxEntityLang_CXXIdxEntityLang_Swift"

var _IdxEntityLanguage_index = [...]uint8{0, 18, 33, 51, 68, 87}

func (i IdxEntityLanguage) String() string {
	if i >= IdxEntityLanguage(len(_IdxEntityLanguage_index)-1) {
		return "IdxEntityLanguage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxEntityLanguage_name[_IdxEntityLanguage_index[i]:_IdxEntityLanguage_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IdxEntityRef_Direct-1]
	_ = x[IdxEntityRef_Implicit-2]
}

const _IdxEntityRefKind_name = "IdxEntityRef_DirectIdxEntityRef_Implicit"

var _IdxEntityRefKind_index = [...]uint8{0, 19, 40}

func (i IdxEntityRefKind) String() string {
	i -= 1
	if i >= IdxEntityRefKind(len(_IdxEntityRefKind_index)-1) {
		return "IdxEntityRefKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _IdxEntityRefKind_name[_IdxEntityRefKind_index[i]:_IdxEntityRefKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IdxObjCContainer_ForwardRef-0]
	_ = x[IdxObjCContainer_Interface-1]
	_ = x[IdxObjCContainer_Implementation-2]
}

const _IdxObjCContainerKind_name = "IdxObjCContainer_ForwardRefIdxObjCContainer_InterfaceIdxObjCContainer_Implementation"

var _IdxObjCContainerKind_index = [...]uint8{0, 27, 53, 84}

func (i IdxObjCContainerKind) String() string {
	if i >= IdxObjCContainerKind(len(_IdxObjCContainerKind_index)-1) {
		return "IdxObjCContainerKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxObjCContainerKind_name[_IdxObjCContainerKind_index[i]:_IdxObjCContainerKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IndexOpt_None-0]
	_ = x[IndexOpt_SuppressRedundantRefs-1]
	_ = x[IndexOpt_IndexFunctionLocalSymbols-2]
	_ = x[IndexOpt_IndexImplicitTemplateInstantiations-4]
	_ = x[IndexOpt_SuppressWarnings-8]
	_ = x[IndexOpt_SkipParsedBodiesInSession-16]
}

const (
	_IndexOptFlags_name_0 = "IndexOpt_NoneIndexOpt_SuppressRedundantRefsIndexOpt_IndexFunctionLocalSymbols"
	_IndexOptFlags_name_1 = "IndexOpt_IndexImplicitTemplateInstantiations"
	_IndexOptFlags_name_2 = "IndexOpt_SuppressWarnings"
	_IndexOptFlags_name_3 = "IndexOpt_SkipParsedBodiesInSession"
)

var (
	_IndexOptFlags_index_0 = [...]uint8{0, 13, 43, 77}
)

func (i IndexOptFlags) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _IndexOptFlags_name_0[_IndexOptFlags_index_0[i]:_IndexOptFlags_index_0[i+1]]
	case i == 4:
		return _IndexOptFlags_name_1
	case i == 8:
		return _IndexOptFlags_name_2
	case i == 16:
		return _IndexOptFlags_name_3
	default:
		return "IndexOptFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Language_Invalid-0]
	_ = x[Language_C-1]
	_ = x[Language_ObjC-2]
	_ = x[Language_CPlusPlus-3]
}

const _LanguageKind_name = "Language_InvalidLanguage_CLanguage_ObjCLanguage_CPlusPlus"

var _LanguageKind_index = [...]uint8{0, 16, 26, 39, 57}

func (i LanguageKind) String() string {
	if i >= LanguageKind(len(_LanguageKind_index)-1) {
		return "LanguageKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LanguageKind_name[_LanguageKind_index[i]:_LanguageKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Linkage_Invalid-0]
	_ = x[Linkage_NoLinkage-1]
	_ = x[Linkage_Internal-2]
	_ = x[Linkage_UniqueExternal-3]
	_ = x[Linkage_External-4]
}

const _LinkageKind_name = "Linkage_InvalidLinkage_NoLinkageLinkage_InternalLinkage_UniqueExternalLinkage_External"

var _LinkageKind_index = [...]uint8{0, 15, 32, 48, 70, 86}

func (i LinkageKind) String() string {
	if i >= LinkageKind(len(_LinkageKind_index)-1) {
		return "LinkageKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LinkageKind_name[_LinkageKind_index[i]:_LinkageKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NameRange_WantQualifier-1]
	_ = x[NameRange_WantTemplateArgs-2]
	_ = x[NameRange_WantSinglePiece-4]
}

const (
	_NameRefFlags_name_0 = "NameRange_WantQualifierNameRange_WantTemplateArgs"
	_NameRefFlags_name_1 = "NameRange_WantSinglePiece"
)

var (
	_NameRefFlags_index_0 = [...]uint8{0, 23, 49}
)

func (i NameRefFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _NameRefFlags_name_0[_NameRefFlags_index_0[i]:_NameRefFlags_index_0[i+1]]
	case i == 4:
		return _NameRefFlags_name_1
	default:
		return "NameRefFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RefQualifier_None-0]
	_ = x[RefQualifier_LValue-1]
	_ = x[RefQualifier_RValue-2]
}

const _RefQualifierKind_name = "RefQualifier_NoneRefQualifier_LValueRefQualifier_RValue"

var _RefQualifierKind_index = [...]uint8{0, 17, 36, 55}

func (i RefQualifierKind) String() string {
	if i >= RefQualifierKind(len(_RefQualifierKind_index)-1) {
		return "RefQualifierKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RefQualifierKind_name[_RefQualifierKind_index[i]:_RefQualifierKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Reparse_None-0]
}

const _Reparse_Flags_name = "Reparse_None"

var _Reparse_Flags_index = [...]uint8{0, 12}

func (i Reparse_Flags) String() string {
	if i >= Reparse_Flags(len(_Reparse_Flags_index)-1) {
		return "Reparse_Flags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Reparse_Flags_name[_Reparse_Flags_index[i]:_Reparse_Flags_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Result_Success-0]
	_ = x[Result_Invalid-1]
	_ = x[Result_VisitBreak-2]
}

const _Result_name = "Result_SuccessResult_InvalidResult_VisitBreak"

var _Result_index = [...]uint8{0, 14, 28, 45}

func (i Result) String() string {
	if i >= Result(len(_Result_index)-1) {
		return "Result(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Result_name[_Result_index[i]:_Result_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SaveTranslationUnit_None-0]
}

const _SaveTranslationUnit_Flags_name = "SaveTranslationUnit_None"

var _SaveTranslationUnit_Flags_index = [...]uint8{0, 24}

func (i SaveTranslationUnit_Flags) String() string {
	if i >= SaveTranslationUnit_Flags(len(_SaveTranslationUnit_Flags_index)-1) {
		return "SaveTranslationUnit_Flags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SaveTranslationUnit_Flags_name[_SaveTranslationUnit_Flags_index[i]:_SaveTranslationUnit_Flags_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SC_Invalid-0]
	_ = x[SC_None-1]
	_ = x[SC_Extern-2]
	_ = x[SC_Static-3]
	_ = x[SC_PrivateExtern-4]
	_ = x[SC_OpenCLWorkGroupLocal-5]
	_ = x[SC_Auto-6]
	_ = x[SC_Register-7]
}

const _StorageClass_name = "SC_InvalidSC_NoneSC_ExternSC_StaticSC_PrivateExternSC_OpenCLWorkGroupLocalSC_AutoSC_Register"

var _StorageClass_index = [...]uint8{0, 10, 17, 26, 35, 51, 74, 81, 92}

func (i StorageClass) String() string {
	if i >= StorageClass(len(_StorageClass_index)-1) {
		return "StorageClass(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StorageClass_name[_StorageClass_index[i]:_StorageClass_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TemplateArgumentKind_Null-0]
	_ = x[TemplateArgumentKind_Type-1]
	_ = x[TemplateArgumentKind_Declaration-2]
	_ = x[TemplateArgumentKind_NullPtr-3]
	_ = x[TemplateArgumentKind_Integral-4]
	_ = x[TemplateArgumentKind_Template-5]
	_ = x[TemplateArgumentKind_TemplateExpansion-6]
	_ = x[TemplateArgumentKind_Expression-7]
	_ = x[TemplateArgumentKind_Pack-8]
	_ = x[TemplateArgumentKind_Invalid-9]
}

const _TemplateArgumentKind_name = "TemplateArgumentKind_NullTemplateArgumentKind_TypeTemplateArgumentKind_DeclarationTemplateArgumentKind_NullPtrTemplateArgumentKind_IntegralTemplateArgumentKind_TemplateTemplateArgumentKind_TemplateExpansionTemplateArgumentKind_ExpressionTemplateArgumentKind_PackTemplateArgumentKind_Invalid"

var _TemplateArgumentKind_index = [...]uint16{0, 25, 50, 82, 110, 139, 168, 206, 237, 262, 290}

func (i TemplateArgumentKind) String() string {
	if i >= TemplateArgumentKind(len(_TemplateArgumentKind_index)-1) {
		return "TemplateArgumentKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TemplateArgumentKind_name[_TemplateArgumentKind_index[i]:_TemplateArgumentKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TUResourceUsage_AST-1]
	_ = x[TUResourceUsage_Identifiers-2]
	_ = x[TUResourceUsage_Selectors-3]
	_ = x[TUResourceUsage_GlobalCompletionResults-4]
	_ = x[TUResourceUsage_SourceManagerContentCache-5]
	_ = x[TUResourceUsage_AST_SideTables-6]
	_ = x[TUResourceUsage_SourceManager_Membuffer_Malloc-7]
	_ = x[TUResourceUsage_SourceManager_Membuffer_MMap-8]
	_ = x[TUResourceUsage_ExternalASTSource_Membuffer_Malloc-9]
	_ = x[TUResourceUsage_ExternalASTSource_Membuffer_MMap-10]
	_ = x[TUResourceUsage_Preprocessor-11]
	_ = x[TUResourceUsage_PreprocessingRecord-12]
	_ = x[TUResourceUsage_SourceManager_DataStructures-13]
	_ = x[TUResourceUsage_Preprocessor_HeaderSearch-14]
	_ = x[TUResourceUsage_MEMORY_IN_BYTES_BEGIN-1]
	_ = x[TUResourceUsage_MEMORY_IN_BYTES_END-14]
	_ = x[TUResourceUsage_First-1]
	_ = x[TUResourceUsage_Last-14]
}

const _TUResourceUsageKind_name = "TUResourceUsage_ASTTUResourceUsage_IdentifiersTUResourceUsage_SelectorsTUResourceUsage_GlobalCompletionResultsTUResourceUsage_SourceManagerContentCacheTUResourceUsage_AST_SideTablesTUResourceUsage_SourceManager_Membuffer_MallocTUResourceUsage_SourceManager_Membuffer_MMapTUResourceUsage_ExternalASTSource_Membuffer_MallocTUResourceUsage_ExternalASTSource_Membuffer_MMapTUResourceUsage_PreprocessorTUResourceUsage_PreprocessingRecordTUResourceUsage_SourceManager_DataStructuresTUResourceUsage_Preprocessor_HeaderSearch"

var _TUResourceUsageKind_index = [...]uint16{0, 19, 46, 71, 110, 151, 181, 227, 271, 321, 369, 397, 432, 476, 517}

func (i TUResourceUsageKind) String() string {
	i -= 1
	if i >= TUResourceUsageKind(len(_TUResourceUsageKind_index)-1) {
		return "TUResourceUsageKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TUResourceUsageKind_name[_TUResourceUsageKind_index[i]:_TUResourceUsageKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Visibility_Invalid-0]
	_ = x[Visibility_Hidden-1]
	_ = x[Visibility_Protected-2]
	_ = x[Visibility_Default-3]
}

const _VisibilityKind_name = "Visibility_InvalidVisibility_HiddenVisibility_ProtectedVisibility_Default"

var _VisibilityKind_index = [...]uint8{0, 18, 35, 55, 73}

func (i VisibilityKind) String() string {
	if i >= VisibilityKind(len(_VisibilityKind_index)-1) {
		return "VisibilityKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VisibilityKind_name[_VisibilityKind_index[i]:_VisibilityKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Visit_Break-0]
	_ = x[Visit_Continue-1]
}

const _VisitorResult_name = "Visit_BreakVisit_Continue"

var _VisitorResult_index = [...]uint8{0, 11, 25}

func (i VisitorResult) String() string {
	if i >= VisitorResult(len(_VisitorResult_index)-1) {
		return "VisitorResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VisitorResult_name[_VisitorResult_index[i]:_VisitorResult_index[i+1]]
}
