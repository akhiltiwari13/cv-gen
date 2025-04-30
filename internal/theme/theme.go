package theme

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/akhiltiwari13/resume/internal/config"
	"github.com/akhiltiwari13/resume/internal/utils"
)

// ApplyStyling applies the appropriate styling to the HTML content
func ApplyStyling(htmlContent []byte, cfg *config.Config) ([]byte, error) {
	// Read HTML template
	templateContent, err := os.ReadFile(cfg.Paths.TemplateFile)
	if err != nil {
		return nil, fmt.Errorf("error reading HTML template: %v", err)
	}

	// Parse template
	tmpl, err := template.New("resume").Parse(string(templateContent))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML template: %v", err)
	}

	// Load CSS styles
	styles, err := loadStyles(cfg)
	if err != nil {
		return nil, fmt.Errorf("error loading styles: %v", err)
	}

	// Prepare template data
	data := map[string]interface{}{
		"Content":        template.HTML(htmlContent),
		"Styles":         template.CSS(styles),
		"Title":          "Resume",
		"FontFamily":     cfg.Styling.FontFamily,
		"FontSize":       cfg.Styling.FontSize,
		"MarginSize":     cfg.Styling.MarginSize,
		"UnicodeEnabled": cfg.PDF.UnicodeEnabled,
	}

	// Execute template
	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		return nil, fmt.Errorf("error executing HTML template: %v", err)
	}

	// Save HTML for debugging
	debugFile := "debug-output.html"
	if err := os.WriteFile(debugFile, result.Bytes(), 0644); err != nil {
		log.Printf("Warning: Could not save debug HTML: %v", err)
	} else {
		log.Printf("Saved debug HTML to %s", debugFile)
	}
	return result.Bytes(), nil
}

// loadStyles loads and concatenates the CSS styles based on configuration
func loadStyles(cfg *config.Config) (string, error) {
	var allStyles strings.Builder

	// Load base styles (always included)
	log.Printf("Loading base styles from: %s", cfg.GetBaseStylePath())
	baseStyles, err := utils.LoadFile(cfg.GetBaseStylePath())
	if err != nil {
		return "", fmt.Errorf("error loading base styles: %v", err)
	}
	allStyles.WriteString(baseStyles)
	allStyles.WriteString("\n\n")

	log.Printf("Using mode: %s", cfg.Mode)
	if cfg.Mode == "custom" {
		log.Printf("Loading theme: %s from path: %s", cfg.Styling.Theme, cfg.GetThemeCSSPath(cfg.Styling.Theme))
	}

	if cfg.Mode == "ats" {
		// ATS mode styles
		log.Printf("Loading ATS styles")
		atsStyles, err := utils.LoadFile(cfg.GetATSStylePath())
		if err != nil {
			return "", fmt.Errorf("error loading ATS styles: %v", err)
		}
		allStyles.WriteString(atsStyles)
	} else {
		// Custom mode styles - load theme
		log.Printf("Loading custom theme styles (NOT loading ATS styles)")
		themeStyles, err := utils.LoadFile(cfg.GetThemeCSSPath(cfg.Styling.Theme))
		if err != nil {
			return "", fmt.Errorf("error loading theme styles: %v", err)
		}
		allStyles.WriteString(themeStyles)
	}
	// Add custom CSS if provided
	if cfg.Styling.CustomCSSPath != "" {
		customStyles, err := utils.LoadFile(cfg.Styling.CustomCSSPath)
		if err != nil {
			return "", fmt.Errorf("error loading custom styles: %v", err)
		}
		allStyles.WriteString("\n\n/* Custom CSS */\n")
		allStyles.WriteString(customStyles)
	}

	return allStyles.String(), nil
}

// DisplayAvailableThemes shows all available built-in themes
func DisplayAvailableThemes() {
	fmt.Println("Available Themes:")
	fmt.Println("  - default: Clean, professional light theme")
	fmt.Println("  - dark: Dark background with light text")
	fmt.Println("  - catppuccin-mocha: Dark Catppuccin theme with vibrant accents")
	fmt.Println("  - catppuccin-latte: Light Catppuccin theme with pastel accents")
	fmt.Println("  - nord: Cool blue-based dark theme")
	fmt.Println("  - github-dark: GitHub's dark mode theme")

	fmt.Println("\nTo use a theme, specify it in your config.yaml file:")
	fmt.Println("  mode: \"custom\"")
	fmt.Println("  styling:")
	fmt.Println("    theme: \"catppuccin-mocha\"")

	fmt.Println("\nOr use the command line flag:")
	fmt.Println("  go run main.go -mode custom -theme catppuccin-mocha")
}

// DisplayAvailableFonts lists fonts that might be available on the system
func DisplayAvailableFonts() {
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

	fmt.Println("\nTo use a font, specify it in your config.yaml file:")
	fmt.Println("  styling:")
	fmt.Println("    font_family: \"JetBrains Mono, monospace\"")

	fmt.Println("\nOr use the command line flag:")
	fmt.Println("  go run main.go -font \"JetBrains Mono, monospace\"")
	fmt.Println("\nNote: Some fonts may not be available on your system.")
}
