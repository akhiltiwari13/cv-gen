package converter

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/akhiltiwari13/cv-gen/internal/config"
	"github.com/akhiltiwari13/cv-gen/internal/logging"
)

// HTMLToPDF converts HTML content to PDF using wkhtmltopdf
func HTMLToPDF(htmlContent []byte, cfg *config.Config) error {
	logger := logging.GetLogger()
	logger.Debug().Str("output_file", cfg.General.OutputFile).Int("html_length", len(htmlContent)).Msg("Starting HTML to PDF conversion")

	// Create a new PDF generator
	logger.Debug().Msg("Creating new PDF generator")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create PDF generator")
		return fmt.Errorf("failed to create PDF generator: %v", err)
	}

	// Set up a page from the HTML content
	logger.Debug().Msg("Creating page from HTML content")
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlContent))

	// Configure page options
	page.Encoding.Set("UTF-8")
	logger.Debug().Msg("Set page encoding to UTF-8")

	// Enable advanced settings for custom mode
	if cfg.Mode == "custom" {
		// Note: Some options might not be available in your version of wkhtmltopdf
		page.LoadErrorHandling.Set("ignore")
	}

	// Set PDF generator options
	pdfg.Dpi.Set(uint(cfg.PDF.DPI))
	logger.Debug().Int("dpi", cfg.PDF.DPI).Msg("Set PDF DPI")
	pdfg.PageSize.Set(cfg.PDF.PageSize)
	logger.Debug().Str("page_size", cfg.PDF.PageSize).Msg("Set PDF page size")
	pdfg.MarginTop.Set(uint(cfg.PDF.MarginTop))
	pdfg.MarginBottom.Set(uint(cfg.PDF.MarginBottom))
	pdfg.MarginLeft.Set(uint(cfg.PDF.MarginLeft))
	pdfg.MarginRight.Set(uint(cfg.PDF.MarginRight))
	logger.Debug().
		Int("margin_top", cfg.PDF.MarginTop).
		Int("margin_bottom", cfg.PDF.MarginBottom).
		Int("margin_left", cfg.PDF.MarginLeft).
		Int("margin_right", cfg.PDF.MarginRight).
		Msg("Set PDF margins")

	pdfg.NoCollate.Set(false)
	// Set grayscale only for ATS mode
	if cfg.Mode == "ats" {
		pdfg.Grayscale.Set(true)
	}

	// Add the page to the generator
	pdfg.AddPage(page)
	logger.Debug().Msg("Added page to PDF generator")

	// Create the PDF with a timeout context to prevent hanging
	logger.Debug().Msg("Starting PDF cration with timeout context")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger.Debug().Str("output_file", cfg.General.OutputFile).Msg("Writing PDF to file")
	if err := pdfg.CreateContext(ctx); err != nil {
		logger.Error().Err(err).Str("path", cfg.General.OutputFile).Msg("Error writing PDF file")
		return fmt.Errorf("error generating PDF: %v", err)
	}

	// PageDivisions and NoPageBreaks handling removed to fix incompatibility

	// Write the PDF to file
	if err := pdfg.WriteFile(cfg.General.OutputFile); err != nil {
		return fmt.Errorf("error writing PDF file: %v", err)
	}

	logger.Info().
		Str("output_file", cfg.General.OutputFile).
		Int("pdf_size_bytes", len(pdfg.Bytes())).
		Msg("PDF generation complete")
	return nil
}
