package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/codysnider/resume/internal/config"
	"github.com/codysnider/resume/internal/converter"
	"github.com/codysnider/resume/internal/theme"
)

func main() {
	// Define command line flags
	configFile := flag.String("config", "config.yaml", "Path to configuration file")
	inputFile := flag.String("input", "", "Input markdown resume file (overrides config)")
	outputFile := flag.String("output", "", "Output PDF file (overrides config)")
	mode := flag.String("mode", "", "Mode: 'ats' or 'custom' (overrides config)")
	themeOption := flag.String("theme", "", "Theme for custom mode (overrides config)")
	fontFamily := flag.String("font", "", "Font family (overrides config)")
	fontSize := flag.String("fontsize", "", "Font size (overrides config)")
	margin := flag.String("margin", "", "Margin size (overrides config)")
	customCSS := flag.String("css", "", "Custom CSS file (overrides config)")

	// Utility flags
	listThemes := flag.Bool("list-themes", false, "List available themes")
	listFonts := flag.Bool("list-fonts", false, "List common system fonts")
	version := flag.Bool("version", false, "Display version information")

	flag.Parse()

	// Display version if requested
	if *version {
		fmt.Println("Resume Generator v1.0.0")
		fmt.Println("A tool to generate ATS-friendly or custom-styled resumes from Markdown")
		return
	}

	// List themes if requested
	if *listThemes {
		theme.DisplayAvailableThemes()
		return
	}

	// List fonts if requested
	if *listFonts {
		theme.DisplayAvailableFonts()
		return
	}

	// Load configuration
	cfg, err := config.LoadConfig(*configFile)
	log.Printf("Loaded config: Mode=%s, Theme=%s", cfg.Mode, cfg.Styling.Theme)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Override config with command line flags if provided
	if *inputFile != "" {
		cfg.General.InputFile = *inputFile
	}
	if *outputFile != "" {
		cfg.General.OutputFile = *outputFile
	}
	if *mode != "" {
		cfg.Mode = *mode
	}
	if *themeOption != "" {
		cfg.Styling.Theme = *themeOption
	}
	if *fontFamily != "" {
		cfg.Styling.FontFamily = *fontFamily
	}
	if *fontSize != "" {
		cfg.Styling.FontSize = *fontSize
	}
	if *margin != "" {
		cfg.Styling.MarginSize = *margin
	}
	if *customCSS != "" {
		cfg.Styling.CustomCSSPath = *customCSS
	}

	// Ensure absolute paths for CSS files
	err = resolveConfigPaths(cfg)
	if err != nil {
		log.Fatalf("Error resolving paths: %v", err)
	}

	// Read markdown content
	markdownContent, err := os.ReadFile(cfg.General.InputFile)
	if err != nil {
		log.Fatalf("Error reading %s: %v", cfg.General.InputFile, err)
	}

	// Convert markdown to HTML
	htmlContent, err := converter.MarkdownToHTML(markdownContent)
	if err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
	}

	// Generate final HTML with styling
	finalHTML, err := theme.ApplyStyling(htmlContent, cfg)
	if err != nil {
		log.Fatalf("Error applying styling: %v", err)
	}

	// Convert to PDF
	err = converter.HTMLToPDF(finalHTML, cfg)
	if err != nil {
		log.Fatalf("Error generating PDF: %v", err)
	}

	log.Printf("Successfully generated %s\n", cfg.General.OutputFile)
	if cfg.Mode == "custom" {
		log.Printf("Generated with custom styling using theme: %s\n", cfg.Styling.Theme)
	} else {
		log.Printf("Generated in ATS-friendly mode for optimal compatibility with applicant tracking systems\n")
	}
}

// resolveConfigPaths ensures all paths in the config are absolute
func resolveConfigPaths(cfg *config.Config) error {
	// Get current working directory
	projectDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}

	// Resolve styles directory if it's relative
	if !filepath.IsAbs(cfg.Paths.StylesDir) {
		cfg.Paths.StylesDir = filepath.Join(projectDir, cfg.Paths.StylesDir)
	}

	// Resolve template file if it's relative
	if !filepath.IsAbs(cfg.Paths.TemplateFile) {
		cfg.Paths.TemplateFile = filepath.Join(projectDir, cfg.Paths.TemplateFile)
	}

	// Resolve custom CSS path if provided and it's relative
	if cfg.Styling.CustomCSSPath != "" && !filepath.IsAbs(cfg.Styling.CustomCSSPath) {
		cfg.Styling.CustomCSSPath = filepath.Join(projectDir, cfg.Styling.CustomCSSPath)
	}

	// Resolve input file if it's relative
	if !filepath.IsAbs(cfg.General.InputFile) {
		cfg.General.InputFile = filepath.Join(projectDir, cfg.General.InputFile)
	}

	// Resolve output file if it's relative
	if !filepath.IsAbs(cfg.General.OutputFile) {
		cfg.General.OutputFile = filepath.Join(projectDir, cfg.General.OutputFile)
	}

	// Add this after all command-line overrides
	log.Printf("Final config: Mode=%s, Theme=%s", cfg.Mode, cfg.Styling.Theme)
	return nil
}
