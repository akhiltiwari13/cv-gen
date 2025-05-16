package converter

import (
	"github.com/russross/blackfriday/v2"
	"github.com/akhiltiwari13/cv-gen/internal/logging"
)

// MarkdownToHTML converts markdown content to HTML
func MarkdownToHTML(content []byte) ([]byte, error) {
	logger := logging.GetLogger()
	logger.Debug().Msg("Starting markdow to HTML conversion")

	// Configure Blackfriday options
	extensions := blackfriday.CommonExtensions |
		blackfriday.AutoHeadingIDs |
		blackfriday.NoEmptyLineBeforeBlock |
		blackfriday.HardLineBreak

	logger.Debug().Int("extension_flags", int(extensions)).Msg("Configured markdown extensions")
	// Create renderer with options
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags | blackfriday.HrefTargetBlank,
	})

	logger.Debug().Int("html_flags", int(blackfriday.CommonHTMLFlags | blackfriday.HrefTargetBlank)).Msg("Configured html renderer")

	// Run markdown to HTML conversion
	logger.Debug().Msg("Converting markdown to HTML")

	html := blackfriday.Run(content, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(renderer),)
	logger.Debug().Int("html_length", len(html)).Msg("Markdown conversion complete")

	return html, nil
}
