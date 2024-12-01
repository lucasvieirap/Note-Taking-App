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
	htmlDocument := markdownParser.Parse(data)

	htmlOptions := html.RendererOptions{ Flags: htmlFlags }
	htmlRenderer := html.NewRenderer(htmlOptions)

	return markdown.Render(htmlDocument, htmlRenderer)
}
