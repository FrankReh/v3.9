package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"

// A parsed comment.
type Comment struct {
	c C.CXComment
}

/*
	Parameter Comment AST node of any kind.

	Returns the type of the AST node.
*/
func (c Comment) Kind() CommentKind {
	return CommentKind(C.clang_Comment_getKind(c.c))
}

/*
	Parameter Comment AST node of any kind.

	Returns number of children of the AST node.
*/
func (c Comment) NumChildren() uint32 {
	return uint32(C.clang_Comment_getNumChildren(c.c))
}

/*
	Parameter Comment AST node of any kind.

	Parameter ChildIdx child index (zero-based).

	Returns the specified child of the AST node.
*/
func (c Comment) Child(childIdx uint32) Comment {
	return Comment{C.clang_Comment_getChild(c.c, C.uint(childIdx))}
}

/*
	A CXComment_Paragraph node is considered whitespace if it contains
	only CXComment_Text nodes that are empty or whitespace.

	Other AST nodes (except CXComment_Paragraph and CXComment_Text) are
	never considered whitespace.

	Returns non-zero if Comment is whitespace.
*/
func (c Comment) IsWhitespace() bool {
	return C.clang_Comment_isWhitespace(c.c) != 0
}

// Returns non-zero if Comment is inline content and has a newline immediately following it in the comment text. Newlines between paragraphs do not count.
func (c Comment) InlineContentComment_HasTrailingNewline() bool {
	return C.clang_InlineContentComment_hasTrailingNewline(c.c) != 0
}

/*
	Parameter Comment a CXComment_Text AST node.

	Returns text contained in the AST node.
*/
func (c Comment) TextComment_getText() string {
	return cx2GoString(C.clang_TextComment_getText(c.c))
}

/*
	Parameter Comment a CXComment_InlineCommand AST node.

	Returns name of the inline command.
*/
func (c Comment) InlineCommandComment_getCommandName() string {
	return cx2GoString(C.clang_InlineCommandComment_getCommandName(c.c))
}

/*
	Parameter Comment a CXComment_InlineCommand AST node.

	Returns the most appropriate rendering mode, chosen on command
	semantics in Doxygen.
*/
func (c Comment) InlineCommandComment_getRenderKind() CommentInlineCommandRenderKind {
	return CommentInlineCommandRenderKind(C.clang_InlineCommandComment_getRenderKind(c.c))
}

// The most appropriate rendering mode for an inline command, chosen on command semantics in Doxygen.
type CommentInlineCommandRenderKind uint32

const (
	// Command argument should be rendered in a normal font.
	CommentInlineCommandRenderKind_Normal CommentInlineCommandRenderKind = C.CXCommentInlineCommandRenderKind_Normal

	// Command argument should be rendered in a bold font.
	CommentInlineCommandRenderKind_Bold CommentInlineCommandRenderKind = C.CXCommentInlineCommandRenderKind_Bold

	// Command argument should be rendered in a monospaced font.
	CommentInlineCommandRenderKind_Monospaced CommentInlineCommandRenderKind = C.CXCommentInlineCommandRenderKind_Monospaced

	// Command argument should be rendered emphasized (typically italic font).
	CommentInlineCommandRenderKind_Emphasized CommentInlineCommandRenderKind = C.CXCommentInlineCommandRenderKind_Emphasized

	// Command argument should not be rendered (since it only defines an anchor).
	CommentInlineCommandRenderKind_Anchor CommentInlineCommandRenderKind = C.CXCommentInlineCommandRenderKind_Anchor
)

/*
	Parameter Comment a CXComment_InlineCommand AST node.

	Returns number of command arguments.
*/
func (c Comment) InlineCommandComment_getNumArgs() uint32 {
	return uint32(C.clang_InlineCommandComment_getNumArgs(c.c))
}

/*
	Parameter Comment a CXComment_InlineCommand AST node.

	Parameter ArgIdx argument index (zero-based).

	Returns text of the specified argument.
*/
func (c Comment) InlineCommandComment_getArgText(argIdx uint32) string {
	return cx2GoString(C.clang_InlineCommandComment_getArgText(c.c, C.uint(argIdx)))
}

/*
	Parameter Comment a CXComment_HTMLStartTag or CXComment_HTMLEndTag AST
	node.

	Returns HTML tag name.
*/
func (c Comment) HTMLTagComment_getTagName() string {
	return cx2GoString(C.clang_HTMLTagComment_getTagName(c.c))
}

/*
	Parameter Comment a CXComment_HTMLStartTag AST node.

	Returns non-zero if tag is self-closing (for example, &lt;br /&gt;).
*/
func (c Comment) HTMLStartTagComment_IsSelfClosing() bool {
	return C.clang_HTMLStartTagComment_isSelfClosing(c.c) != 0
}

/*
	Parameter Comment a CXComment_HTMLStartTag AST node.

	Returns number of attributes (name-value pairs) attached to the start tag.
*/
func (c Comment) HTMLStartTag_getNumAttrs() uint32 {
	return uint32(C.clang_HTMLStartTag_getNumAttrs(c.c))
}

/*
	Parameter Comment a CXComment_HTMLStartTag AST node.

	Parameter AttrIdx attribute index (zero-based).

	Returns name of the specified attribute.
*/
func (c Comment) HTMLStartTag_getAttrName(attrIdx uint32) string {
	return cx2GoString(C.clang_HTMLStartTag_getAttrName(c.c, C.uint(attrIdx)))
}

/*
	Parameter Comment a CXComment_HTMLStartTag AST node.

	Parameter AttrIdx attribute index (zero-based).

	Returns value of the specified attribute.
*/
func (c Comment) HTMLStartTag_getAttrValue(attrIdx uint32) string {
	return cx2GoString(C.clang_HTMLStartTag_getAttrValue(c.c, C.uint(attrIdx)))
}

/*
	Parameter Comment a CXComment_BlockCommand AST node.

	Returns name of the block command.
*/
func (c Comment) BlockCommandComment_getCommandName() string {
	return cx2GoString(C.clang_BlockCommandComment_getCommandName(c.c))
}

/*
	Parameter Comment a CXComment_BlockCommand AST node.

	Returns number of word-like arguments.
*/
func (c Comment) BlockCommandComment_getNumArgs() uint32 {
	return uint32(C.clang_BlockCommandComment_getNumArgs(c.c))
}

/*
	Parameter Comment a CXComment_BlockCommand AST node.

	Parameter ArgIdx argument index (zero-based).

	Returns text of the specified word-like argument.
*/
func (c Comment) BlockCommandComment_getArgText(argIdx uint32) string {
	return cx2GoString(C.clang_BlockCommandComment_getArgText(c.c, C.uint(argIdx)))
}

/*
	Parameter Comment a CXComment_BlockCommand or
	CXComment_VerbatimBlockCommand AST node.

	Returns paragraph argument of the block command.
*/
func (c Comment) BlockCommandComment_getParagraph() Comment {
	return Comment{C.clang_BlockCommandComment_getParagraph(c.c)}
}

/*
	Parameter Comment a CXComment_ParamCommand AST node.

	Returns parameter name.
*/
func (c Comment) ParamCommandComment_getParamName() string {
	return cx2GoString(C.clang_ParamCommandComment_getParamName(c.c))
}

/*
	Parameter Comment a CXComment_ParamCommand AST node.

	Returns non-zero if the parameter that this AST node represents was found
	in the function prototype and clang_ParamCommandComment_getParamIndex
	function will return a meaningful value.
*/
func (c Comment) ParamCommandComment_IsParamIndexValid() bool {
	return C.clang_ParamCommandComment_isParamIndexValid(c.c) != 0
}

/*
	Parameter Comment a CXComment_ParamCommand AST node.

	Returns zero-based parameter index in function prototype.
*/
func (c Comment) ParamCommandComment_getParamIndex() uint32 {
	return uint32(C.clang_ParamCommandComment_getParamIndex(c.c))
}

/*
	Parameter Comment a CXComment_ParamCommand AST node.

	Returns non-zero if parameter passing direction was specified explicitly in
	the comment.
*/
func (c Comment) ParamCommandComment_IsDirectionExplicit() bool {
	return C.clang_ParamCommandComment_isDirectionExplicit(c.c) != 0
}

/*
	Parameter Comment a CXComment_ParamCommand AST node.

	Returns parameter passing direction.
*/
func (c Comment) ParamCommandComment_getDirection() CommentParamPassDirection {
	return CommentParamPassDirection(C.clang_ParamCommandComment_getDirection(c.c))
}

// Describes parameter passing direction for Parameter or arg command.
type CommentParamPassDirection uint32

const (
	// The parameter is an input parameter.
	CommentParamPassDirection_In CommentParamPassDirection = C.CXCommentParamPassDirection_In

	// The parameter is an output parameter.
	CommentParamPassDirection_Out CommentParamPassDirection = C.CXCommentParamPassDirection_Out

	// The parameter is an input and output parameter.
	CommentParamPassDirection_InOut CommentParamPassDirection = C.CXCommentParamPassDirection_InOut
)

/*
	Parameter Comment a CXComment_TParamCommand AST node.

	Returns template parameter name.
*/
func (c Comment) TParamCommandComment_getParamName() string {
	return cx2GoString(C.clang_TParamCommandComment_getParamName(c.c))
}

/*
	Parameter Comment a CXComment_TParamCommand AST node.

	Returns non-zero if the parameter that this AST node represents was found
	in the template parameter list and
	clang_TParamCommandComment_getDepth and
	clang_TParamCommandComment_getIndex functions will return a meaningful
	value.
*/
func (c Comment) TParamCommandComment_IsParamPositionValid() bool {
	return C.clang_TParamCommandComment_isParamPositionValid(c.c) != 0
}

/*
	Parameter Comment a CXComment_TParamCommand AST node.

	Returns zero-based nesting depth of this parameter in the template parameter list.

	For example,
	\verbatim
	template<typename C, template<typename T> class TT>
	void test(TT<int> aaa);
	\endverbatim
	for C and TT nesting depth is 0,
	for T nesting depth is 1.
*/
func (c Comment) TParamCommandComment_getDepth() uint32 {
	return uint32(C.clang_TParamCommandComment_getDepth(c.c))
}

/*
	Parameter Comment a CXComment_TParamCommand AST node.

	Returns zero-based parameter index in the template parameter list at a
	given nesting depth.

	For example,
	\verbatim
	template<typename C, template<typename T> class TT>
	void test(TT<int> aaa);
	\endverbatim
	for C and TT nesting depth is 0, so we can ask for index at depth 0:
	at depth 0 C's index is 0, TT's index is 1.

	For T nesting depth is 1, so we can ask for index at depth 0 and 1:
	at depth 0 T's index is 1 (same as TT's),
	at depth 1 T's index is 0.
*/
func (c Comment) TParamCommandComment_getIndex(depth uint32) uint32 {
	return uint32(C.clang_TParamCommandComment_getIndex(c.c, C.uint(depth)))
}

/*
	Parameter Comment a CXComment_VerbatimBlockLine AST node.

	Returns text contained in the AST node.
*/
func (c Comment) VerbatimBlockLineComment_getText() string {
	return cx2GoString(C.clang_VerbatimBlockLineComment_getText(c.c))
}

/*
	Parameter Comment a CXComment_VerbatimLine AST node.

	Returns text contained in the AST node.
*/
func (c Comment) VerbatimLineComment_getText() string {
	return cx2GoString(C.clang_VerbatimLineComment_getText(c.c))
}

/*
	Convert an HTML tag AST node to string.

	Parameter Comment a CXComment_HTMLStartTag or CXComment_HTMLEndTag AST
	node.

	Returns string containing an HTML tag.
*/
func (c Comment) HTMLTagComment_getAsString() string {
	return cx2GoString(C.clang_HTMLTagComment_getAsString(c.c))
}

/*
	Convert a given full parsed comment to an HTML fragment.

	Specific details of HTML layout are subject to change. Don't try to parse
	this HTML back into an AST, use other APIs instead.

	Currently the following CSS classes are used:
	\li "para-brief" for \paragraph and equivalent commands;
	\li "para-returns" for \Returns paragraph and equivalent commands;
	\li "word-returns" for the "Returns" word in \Returns paragraph.

	Function argument documentation is rendered as a \<dl\> list with arguments
	sorted in function prototype order. CSS classes used:
	\li "param-name-index-NUMBER" for parameter name (\<dt\>);
	\li "param-descr-index-NUMBER" for parameter description (\<dd\>);
	\li "param-name-index-invalid" and "param-descr-index-invalid" are used if
	parameter index is invalid.

	Template parameter documentation is rendered as a \<dl\> list with
	parameters sorted in template parameter list order. CSS classes used:
	\li "tparam-name-index-NUMBER" for parameter name (\<dt\>);
	\li "tparam-descr-index-NUMBER" for parameter description (\<dd\>);
	\li "tparam-name-index-other" and "tparam-descr-index-other" are used for
	names inside template template parameters;
	\li "tparam-name-index-invalid" and "tparam-descr-index-invalid" are used if
	parameter position is invalid.

	Parameter Comment a CXComment_FullComment AST node.

	Returns string containing an HTML fragment.
*/
func (c Comment) FullComment_getAsHTML() string {
	return cx2GoString(C.clang_FullComment_getAsHTML(c.c))
}

/*
	Convert a given full parsed comment to an XML document.

	A Relax NG schema for the XML can be found in comment-xml-schema.rng file
	inside clang source tree.

	Parameter Comment a CXComment_FullComment AST node.

	Returns string containing an XML document.
*/
func (c Comment) FullComment_getAsXML() string {
	return cx2GoString(C.clang_FullComment_getAsXML(c.c))
}

func (c Comment) TranslationUnit() TranslationUnit {
	return TranslationUnit{c.c.TranslationUnit}
}
