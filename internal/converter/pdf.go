package converter

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/codysnider/resume/internal/config"
)

// HTMLToPDF converts HTML content to PDF using wkhtmltopdf
func HTMLToPDF(htmlContent []byte, cfg *config.Config) error {
	// Create a new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return fmt.Errorf("failed to create PDF generator: %v", err)
	}

	// Set up a page from the HTML content
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlContent))

	// Configure page options
	page.Encoding.Set("UTF-8")

	// Set PDF generator options
	pdfg.Dpi.Set(uint(cfg.PDF.DPI))
	pdfg.PageSize.Set(cfg.PDF.PageSize)
	pdfg.MarginTop.Set(uint(cfg.PDF.MarginTop))
	pdfg.MarginBottom.Set(uint(cfg.PDF.MarginBottom))
	pdfg.MarginLeft.Set(uint(cfg.PDF.MarginLeft))
	pdfg.MarginRight.Set(uint(cfg.PDF.MarginRight))

	// Set grayscale only for ATS mode
	if cfg.Mode == "ats" {
		pdfg.Grayscale.Set(true)
	}

	// Add the page to the generator
	pdfg.AddPage(page)

	// Create the PDF with a timeout context to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := pdfg.CreateContext(ctx); err != nil {
		return fmt.Errorf("error generating PDF: %v", err)
	}

	// Write the PDF to file
	if err := pdfg.WriteFile(cfg.General.OutputFile); err != nil {
		return fmt.Errorf("error writing PDF file: %v", err)
	}

	return nil
}
