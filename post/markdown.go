package post

import (
	bf "github.com/russross/blackfriday"
	//blackfriday2 "gopkg.in/russross/blackfriday.v2"
)

// These are copied from
// https://github.com/russross/blackfriday/blob/5b2fb1b893850a20a483145b676af75eaad51a5d/markdown.go#L48
const (
	commonHtmlFlags = 0 |
		bf.HTML_USE_XHTML |
		bf.HTML_USE_SMARTYPANTS |
		bf.HTML_SMARTYPANTS_FRACTIONS |
		bf.HTML_SMARTYPANTS_DASHES |
		bf.HTML_SMARTYPANTS_LATEX_DASHES

	commonExtensions = 0 |
		bf.EXTENSION_NO_INTRA_EMPHASIS |
		bf.EXTENSION_TABLES |
		bf.EXTENSION_FENCED_CODE |
		bf.EXTENSION_AUTOLINK |
		bf.EXTENSION_STRIKETHROUGH |
		bf.EXTENSION_SPACE_HEADERS |
		bf.EXTENSION_HEADER_IDS |
		bf.EXTENSION_BACKSLASH_LINE_BREAK |
		bf.EXTENSION_DEFINITION_LISTS |
		bf.EXTENSION_AUTO_HEADER_IDS
)

func compileMarkdownToHTML(input []byte) []byte {
	// Use the default extensions
	//return bf.MarkdownCommon(input);

	// Use custom extensions (defaults + extra)
	renderer := bf.HtmlRenderer(commonHtmlFlags, "", "")
	return bf.Markdown(input, renderer, commonExtensions)
}
