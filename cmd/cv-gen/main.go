package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/akhiltiwari13/cv-gen/internal/config"
	"github.com/akhiltiwari13/cv-gen/internal/converter"
	"github.com/akhiltiwari13/cv-gen/internal/theme"
	"github.com/akhiltiwari13/cv-gen/internal/logging"
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
	logLevel := flag.String("log-level", "info", "Log level: debug, info, warn, error")
	prettyLog := flag.Bool("pretty-log", true, "Use pretty logging format")
	// resumeStyle := flag.String("resume-style", "professional", "Resume style: basic, professional")

	// Utility flags
	listThemes := flag.Bool("list-themes", false, "List available themes")
	listFonts := flag.Bool("list-fonts", false, "List common system fonts")
	version := flag.Bool("version", false, "Display version information")

	flag.Parse()

	// Initialize logger
	logging.InitLogger(*logLevel, *prettyLog)
	logger := logging.GetLogger()

	logger.Info().Msg("Starting cv-gen resume generator")
	// Display version if requested
	if *version {
		fmt.Println("Resume Generator v1.0.0")
		fmt.Println("A tool to generate ATS-friendly or custom-styled resumes from Markdown")
		logger.Info().Msg("Displayed version information")
		return
	}

	// List themes if requested
	if *listThemes {
		logger.Info().Msg("Displaying available themes")
		theme.DisplayAvailableThemes()
		return
	}

	// List fonts if requested
	if *listFonts {
		logger.Info().Msg("Displaying available fonts")
		theme.DisplayAvailableFonts()
		return
	}


	// Load configuration
	logger.Debug().Str("config_file", *configFile).Msg("Loading configuration")
	cfg, err := config.LoadConfig(*configFile)
	log.Printf("Loaded config: Mode=%s, Theme=%s", cfg.Mode, cfg.Styling.Theme)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	logger.Debug().
		Str("input_file", cfg.General.InputFile).
		Str("output_file", cfg.General.OutputFile).
		Str("mode", cfg.Mode).
		Str("theme", cfg.Styling.Theme).
		Msg("Configuration loaded")

	// Override config with command line flags if provided
	if *inputFile != "" {
		logger.Debug().Str("input_file", *inputFile).Msg("Overriding input file from command line")
		cfg.General.InputFile = *inputFile
	}
	if *outputFile != "" {
		logger.Debug().Str("output_file", *outputFile).Msg("Overriding output file from command line")
		cfg.General.OutputFile = *outputFile
	}
	if *mode != "" {
		logger.Debug().Str("mode", *mode).Msg("Overriding mode from command line")
		cfg.Mode = *mode
	}
	if *themeOption != "" {
		logger.Debug().Str("theme", *themeOption).Msg("Overriding theme from command line")
		cfg.Styling.Theme = *themeOption
	}
	if *fontFamily != "" {
		logger.Debug().Str("font_family", *fontFamily).Msg("Overriding font family from command line")
		cfg.Styling.FontFamily = *fontFamily
	}
	if *fontSize != "" {
		logger.Debug().Str("font_size", *fontSize).Msg("Overriding font size from command line")
		cfg.Styling.FontSize = *fontSize
	}
	if *margin != "" {
		logger.Debug().Str("margin_size", *margin).Msg("Overriding margin size from command line")
		cfg.Styling.MarginSize = *margin
	}
	if *customCSS != "" {
		logger.Debug().Str("custom_css_path", *customCSS).Msg("Overriding custom CSS path from command line")
		cfg.Styling.CustomCSSPath = *customCSS
	}

	logger.Info().Str("config_file", *configFile).Msg("Starting resume generation")

	// Ensure absolute paths for CSS files
	logger.Debug().Msg("Resolving configuration paths")
	err = resolveConfigPaths(cfg)
	if err != nil {
		log.Fatalf("Error resolving paths: %v", err)
	}

	// Read markdown content
	logger.Debug().Str("input_file", cfg.General.InputFile).Msg("Reading markdown content")
	markdownContent, err := os.ReadFile(cfg.General.InputFile)
	if err != nil {
		log.Fatalf("Error reading %s: %v", cfg.General.InputFile, err)
		logger.Fatal().Err(err).Str("input_file", cfg.General.InputFile).Msg("Error reading markdown file")
	}
	logger.Debug().Int("markdown_length", len(markdownContent)).Msg("Markdown content read successfully")

	// Convert markdown to HTML
	logger.Debug().Msg("Converting markdown to HTML")
	htmlContent, err := converter.MarkdownToHTML(markdownContent)
	if err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
		logger.Fatal().Err(err).Msg("Error converting markdown to HTML")
	}
	logger.Debug().Int("html_length", len(htmlContent)).Msg("Converted markdown to HTML")

	// Generate final HTML with styling
	logger.Debug().Msg("Applying styling to HTML")
	finalHTML, err := theme.ApplyStyling(htmlContent, cfg)
	if err != nil {
		log.Fatalf("Error applying styling: %v", err)
		logger.Fatal().Err(err).Msg("Error applying styling")
	}
	logger.Debug().Int("final_html_length", len(finalHTML)).Msg("Applied styling to HTML")

	// Convert to PDF
	logger.Debug().Str("output_file", cfg.General.OutputFile).Msg("Converting HTML to PDF")
	err = converter.HTMLToPDF(finalHTML, cfg)
	if err != nil {
		log.Fatalf("Error generating PDF: %v", err)
		logger.Fatal().Err(err).Msg("Error generating PDF")
	}
	logger.Info().Str("output_file", cfg.General.OutputFile).Msg("Successfully generated PDF")
	log.Printf("Successfully generated %s\n", cfg.General.OutputFile)
	if cfg.Mode == "custom" {
		log.Printf("Generated with custom styling using theme: %s\n", cfg.Styling.Theme)
		logger.Info().Str("theme", cfg.Styling.Theme).Msg("Generated with custom styling using theme")
	} else {
		log.Printf("Generated in ATS-friendly mode for optimal compatibility with applicant tracking systems\n")
		logger.Info().Msg("Generated in ATS-friendly mode for optimal compatibility with applicant tracking systems")
	}
}

// resolveConfigPaths ensures all paths in the config are absolute
func resolveConfigPaths(cfg *config.Config) error {
	logger := logging.GetLogger()

	// Get executable directory
	exePath, err := os.Executable()
	if err != nil {
		logger.Error().Err(err).Msg("Error getting executable path")
		return fmt.Errorf("error getting executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	logger.Debug().Str("exe_dir", exeDir).Msg("Resolved executable directory")

	// Resolve styles directory if it's relative
	if !filepath.IsAbs(cfg.Paths.StylesDir) {
		origPath := cfg.Paths.StylesDir
		cfg.Paths.StylesDir = filepath.Join(exeDir, cfg.Paths.StylesDir)
		logger.Debug().Str("original", origPath).Str("resolved", cfg.Paths.StylesDir).Msg("Resolved styles directory path")
	}

	// Resolve template file if it's relative
	if !filepath.IsAbs(cfg.Paths.TemplateFile) {
		origPath := cfg.Paths.TemplateFile
		cfg.Paths.TemplateFile = filepath.Join(exeDir, cfg.Paths.TemplateFile)
		logger.Debug().Str("original", origPath).Str("resolved", cfg.Paths.TemplateFile).Msg("Resolved template file path")
	}

	// Resolve custom CSS path if provided and it's relative
	if cfg.Styling.CustomCSSPath != "" && !filepath.IsAbs(cfg.Styling.CustomCSSPath) {
		origPath := cfg.Styling.CustomCSSPath
		cfg.Styling.CustomCSSPath = filepath.Join(exeDir, cfg.Styling.CustomCSSPath)
		logger.Debug().Str("original", origPath).Str("resolved", cfg.Styling.CustomCSSPath).Msg("Resolved custom CSS path")
	}

	logger.Debug().Msg("All config paths resolved")
	return nil
}
