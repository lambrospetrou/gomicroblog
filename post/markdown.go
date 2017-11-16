package post

import (
	bf2 "gopkg.in/russross/blackfriday.v2"
)

// These are copied from
// https://github.com/russross/blackfriday/blob/v2/markdown.go#L32
func compileMarkdownToHTML(input []byte) []byte {
	return bf2.Run(input, bf2.WithExtensions(bf2.CommonExtensions|
		bf2.AutoHeadingIDs))
}
