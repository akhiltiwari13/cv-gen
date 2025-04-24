package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/russross/blackfriday/v2"
)

// ThemeConfig holds styling configuration
type ThemeConfig struct {
	Mode          string // "ats" or "custom"
	FontFamily    string
	CustomCSSPath string
	ColorScheme   string
	MarginSize    string
	FontSize      string
}

func main() {
	// Define command line flags
	resumeFile := flag.String("input", "RESUME.md", "Input markdown resume file")
	outputFile := flag.String("output", "resume.pdf", "Output PDF file")
	mode := flag.String("mode", "ats", "Mode: 'ats' for ATS-friendly or 'custom' for styled")
	fontFamily := flag.String("font", "Arial, sans-serif", "Font family to use")
	cssPath := flag.String("css", "", "Path to custom CSS file")
	colorScheme := flag.String("theme", "default", "Color scheme: default, dark, light, catppuccin-mocha, catppuccin-latte, etc.")
	margin := flag.String("margin", "20px", "Margin size")
	fontSize := flag.String("fontsize", "12px", "Base font size")
	listFonts := flag.Bool("list-fonts", false, "List available system fonts")
	listThemes := flag.Bool("list-themes", false, "List available themes")
	
	flag.Parse()

	// Handle utility commands
	if *listFonts {
		displayAvailableFonts()
		return
	}
	
	if *listThemes {
		displayAvailableThemes()
		return
	}

	// Configure theme
	config := ThemeConfig{
		Mode:          *mode,
		FontFamily:    *fontFamily,
		CustomCSSPath: *cssPath,
		ColorScheme:   *colorScheme,
		MarginSize:    *margin,
		FontSize:      *fontSize,
	}

	// Read markdown content
	data, err := os.ReadFile(*resumeFile)
	if err != nil {
		log.Fatalf("Error reading %s: %v", *resumeFile, err)
	}

	// Convert markdown to HTML
	html := blackfriday.Run(data)
	
	// Generate final HTML with styling
	finalHTML := generateHTML(html, config)

	// Convert to PDF
	if err := generatePDF(finalHTML, *outputFile, config); err != nil {
		log.Fatalf("Error generating PDF: %v", err)
	}

	log.Printf("Successfully generated %s\n", *outputFile)
	if config.Mode == "custom" {
		log.Printf("Generated with custom styling using theme: %s\n", config.ColorScheme)
	} else {
		log.Printf("Generated in ATS-friendly mode for optimal compatibility with applicant tracking systems\n")
	}
}

// generateHTML applies the appropriate styling based on the ThemeConfig
func generateHTML(htmlContent []byte, config ThemeConfig) string {
	// Base CSS for all modes
	baseCSS := fmt.Sprintf(`
		body {
			font-family: %s;
			font-size: %s;
			margin: %s;
			line-height: 1.5;
		}
		h1, h2, h3 {
			margin-bottom: 0.5em;
			line-height: 1.2;
		}
		p, ul, ol {
			line-height: 1.4;
			margin-bottom: 1em;
		}
		/* Better list formatting */
		ul, ol {
			padding-left: 2em;
		}
		li {
			margin-bottom: 0.5em;
		}
		/* Better horizontal rules */
		hr {
			border: none;
			height: 1px;
			margin: 1.5em 0;
		}
	`, config.FontFamily, config.FontSize, config.MarginSize)
	
	// Additional styling based on mode
	var styleContent string
	
	if config.Mode == "ats" {
		// ATS-friendly mode - minimal styling
		styleContent = baseCSS + `
			body {
				background-color: #FFF;
				color: #000;
			}
			h1, h2, h3 {
				color: #000;
			}
			hr {
				background-color: #000;
			}
			/* Ensure links are readable by ATS */
			a {
				color: #000;
				text-decoration: underline;
			}
		`
	} else {
		// Custom styling mode
		styleContent = baseCSS

		// Add color scheme CSS
		themeCSS := getColorSchemeCSS(config.ColorScheme)
		styleContent += themeCSS
		
		// Add custom CSS file if provided
		if config.CustomCSSPath != "" {
			customCSS, err := os.ReadFile(config.CustomCSSPath)
			if err != nil {
				log.Printf("Warning: Could not load custom CSS file: %v", err)
			} else {
				styleContent += string(customCSS)
			}
		}
	}

	// Construct HTML template with proper Unicode handling
	htmlTemplate := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title>Resume</title>
  <style>
    %s
    /* Ensure proper Unicode rendering */
    body {
      text-rendering: optimizeLegibility;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
    }
    /* Added for better print rendering */
    @page {
      margin: 0;
    }
    @media print {
      body {
        -webkit-print-color-adjust: exact;
        print-color-adjust: exact;
      }
    }
  </style>
</head>
<body>
  %s
</body>
</html>`, styleContent, string(htmlContent))

	return htmlTemplate
}

// getColorSchemeCSS returns CSS for predefined color schemes
func getColorSchemeCSS(scheme string) string {
	switch strings.ToLower(scheme) {
	case "dark":
		return `
			body {
				background-color: #222;
				color: #eee;
			}
			h1, h2, h3 {
				color: #fff;
			}
			a {
				color: #6bf;
			}
			hr {
				background-color: #444;
			}
			strong {
				color: #fff;
			}
		`
	case "catppuccin-mocha":
		return `
			body {
				background-color: #1e1e2e;
				color: #cdd6f4;
			}
			h1 {
				color: #f5c2e7;
			}
			h2 {
				color: #cba6f7;
			}
			h3 {
				color: #89b4fa;
			}
			a {
				color: #89dceb;
				text-decoration: none;
				border-bottom: 1px dotted #89dceb;
			}
			hr {
				background-color: #313244;
			}
			strong {
				color: #f9e2af;
				font-weight: bold;
			}
			em {
				color: #a6e3a1;
			}
			code {
				background-color: #313244;
				padding: 0.2em 0.4em;
				border-radius: 3px;
				font-family: "JetBrains Mono", monospace;
				font-size: 0.9em;
			}
			ul li::marker {
				color: #f5c2e7;
			}
		`
	case "catppuccin-latte":
		return `
			body {
				background-color: #eff1f5;
				color: #4c4f69;
			}
			h1 {
				color: #d20f39;
			}
			h2 {
				color: #8839ef;
			}
			h3 {
				color: #1e66f5;
			}
			a {
				color: #179299;
				text-decoration: none;
				border-bottom: 1px dotted #179299;
			}
			hr {
				background-color: #ccd0da;
			}
			strong {
				color: #df8e1d;
				font-weight: bold;
			}
			em {
				color: #40a02b;
			}
			code {
				background-color: #ccd0da;
				padding: 0.2em 0.4em;
				border-radius: 3px;
				font-family: "JetBrains Mono", monospace;
				font-size: 0.9em;
			}
			ul li::marker {
				color: #d20f39;
			}
		`
	case "nord":
		return `
			body {
				background-color: #2e3440;
				color: #d8dee9;
			}
			h1 {
				color: #88c0d0;
			}
			h2 {
				color: #81a1c1;
			}
			h3 {
				color: #5e81ac;
			}
			a {
				color: #8fbcbb;
				text-decoration: none;
				border-bottom: 1px dotted #8fbcbb;
			}
			hr {
				background-color: #4c566a;
			}
			strong {
				color: #ebcb8b;
				font-weight: bold;
			}
			em {
				color: #a3be8c;
			}
			code {
				background-color: #3b4252;
				padding: 0.2em 0.4em;
				border-radius: 3px;
				font-family: "Fira Code", monospace;
				font-size: 0.9em;
			}
			ul li::marker {
				color: #88c0d0;
			}
		`
	case "github-dark":
		return `
			body {
				background-color: #0d1117;
				color: #c9d1d9;
			}
			h1 {
				color: #58a6ff;
			}
			h2 {
				color: #79c0ff;
			}
			h3 {
				color: #a5d6ff;
			}
			a {
				color: #58a6ff;
				text-decoration: none;
				border-bottom: 1px dotted #58a6ff;
			}
			hr {
				background-color: #30363d;
			}
			strong {
				color: #e3b341;
				font-weight: bold;
			}
			em {
				color: #7ee787;
			}
			code {
				background-color: #21262d;
				padding: 0.2em 0.4em;
				border-radius: 3px;
				font-family: "SFMono-Regular", Consolas, monospace;
				font-size: 0.9em;
			}
			ul li::marker {
				color: #58a6ff;
			}
		`
	default: // "light" or default
		return `
			body {
				background-color: #fff;
				color: #333;
			}
			h1, h2, h3 {
				color: #111;
			}
			a {
				color: #0077cc;
				text-decoration: none;
				border-bottom: 1px dotted #0077cc;
			}
			hr {
				background-color: #ddd;
			}
			strong {
				color: #333;
				font-weight: bold;
			}
		`
	}
}

// generatePDF converts HTML to PDF using wkhtmltopdf
func generatePDF(htmlContent string, outputPath string, config ThemeConfig) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return fmt.Errorf("failed to create PDF generator: %v", err)
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent)))
	page.Encoding.Set("UTF-8")
	pdfg.Dpi.Set(96)
	
	// Enable additional options for custom styling
	if filepath.Ext(outputPath) == ".pdf" {
		pdfg.PageSize.Set("A4")
		pdfg.MarginBottom.Set(10)
		pdfg.MarginTop.Set(10)
		pdfg.MarginLeft.Set(10)
		pdfg.MarginRight.Set(10)
		// Only set grayscale if generating in ATS mode
		if config.Mode == "ats" {
			pdfg.Grayscale.Set(true)
		}
	}

	pdfg.AddPage(page)

	if err := pdfg.Create(); err != nil {
		return fmt.Errorf("error generating PDF: %v", err)
	}

	if err := pdfg.WriteFile(outputPath); err != nil {
		return fmt.Errorf("error writing PDF file: %v", err)
	}

	return nil
}

// Helper function to load an entire file
func loadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// displayAvailableFonts lists fonts that might be available on the system
// This is a simplified version as actual font detection would require more complex libraries
func displayAvailableFonts() {
	fmt.Println("Common System Fonts (availability depends on your operating system):")
	fmt.Println("\nSans-Serif Fonts:")
	fmt.Println("  - Arial, sans-serif")
	fmt.Println("  - Helvetica, Arial, sans-serif")
	fmt.Println("  - Verdana, Geneva, sans-serif")
	fmt.Println("  - Tahoma, Geneva, sans-serif")
	fmt.Println("  - 'Open Sans', sans-serif")
	fmt.Println("  - 'Roboto', sans-serif")
	
	fmt.Println("\nSerif Fonts:")
	fmt.Println("  - 'Times New Roman', Times, serif")
	fmt.Println("  - Georgia, Times, serif")
	fmt.Println("  - Garamond, serif")
	
	fmt.Println("\nMonospace Fonts (great for 'hacker' aesthetics):")
	fmt.Println("  - 'Courier New', Courier, monospace")
	fmt.Println("  - 'Fira Code', monospace")
	fmt.Println("  - 'JetBrains Mono', monospace")
	fmt.Println("  - 'Source Code Pro', monospace")
	fmt.Println("  - 'Inconsolata', monospace")
	fmt.Println("  - 'Consolas', monospace")
	
	fmt.Println("\nTo use a font, specify it with the -font flag:")
	fmt.Println("  go run main.go -font \"JetBrains Mono, monospace\"")
	fmt.Println("\nNote: Some fonts may not be available on your system.")
}

// displayAvailableThemes shows all available built-in themes
func displayAvailableThemes() {
	fmt.Println("Available Themes:")
	fmt.Println("  - default: Clean, professional light theme")
	fmt.Println("  - dark: Dark background with light text")
	fmt.Println("  - catppuccin-mocha: Dark Catppuccin theme with vibrant accents")
	fmt.Println("  - catppuccin-latte: Light Catppuccin theme with pastel accents")
	fmt.Println("  - nord: Cool blue-based dark theme")
	fmt.Println("  - github-dark: GitHub's dark mode theme")
	
	fmt.Println("\nTo use a theme, specify it with the -theme flag:")
	fmt.Println("  go run main.go -mode custom -theme catppuccin-mocha")
	
	fmt.Println("\nYou can also create your own theme with a CSS file:")
	fmt.Println("  go run main.go -mode custom -css my-custom-theme.css")
}
