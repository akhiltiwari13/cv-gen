package converter

import (
	"github.com/russross/blackfriday/v2"
)

// MarkdownToHTML converts markdown content to HTML
func MarkdownToHTML(content []byte) ([]byte, error) {
	// Configure Blackfriday options
	extensions := blackfriday.CommonExtensions |
		blackfriday.AutoHeadingIDs |
		blackfriday.NoEmptyLineBeforeBlock |
		blackfriday.HardLineBreak

	// Create renderer with options
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags | blackfriday.HrefTargetBlank,
	})

	// Run markdown to HTML conversion
	return blackfriday.Run(content,
		blackfriday.WithExtensions(extensions),
		blackfriday.WithRenderer(renderer),
	), nil
}
