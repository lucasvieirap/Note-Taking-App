package util

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MarkdownToHTML(data []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	htmlFlags := html.CommonFlags | html.HrefTargetBlank

	markdownParser := parser.NewWithExtensions(extensions)
	document := markdownParser.Parse(data)

	options := html.RendererOptions{ Flags: htmlFlags }
	htmlRenderer := html.NewRenderer(options)

	return markdown.Render(document, htmlRenderer)
}
