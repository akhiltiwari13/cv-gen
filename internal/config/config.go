package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	General GeneralConfig `yaml:"general"`
	Mode    string        `yaml:"mode"`
	Styling StylingConfig `yaml:"styling"`
	Paths   PathsConfig   `yaml:"paths"`
	PDF     PDFConfig     `yaml:"pdf"`
}

// GeneralConfig contains general application settings
type GeneralConfig struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

// StylingConfig contains styling-related configurations
type StylingConfig struct {
	FontFamily    string `yaml:"font_family"`
	FontSize      string `yaml:"font_size"`
	MarginSize    string `yaml:"margin_size"`
	Theme         string `yaml:"theme"`
	CustomCSSPath string `yaml:"custom_css_path"`
}

// PathsConfig contains path configurations for resources
type PathsConfig struct {
	StylesDir    string            `yaml:"styles_dir"`
	TemplateFile string            `yaml:"template_file"`
	ThemeFiles   map[string]string `yaml:"theme_files"`
}

// PDFConfig contains settings for PDF generation
type PDFConfig struct {
	DPI            int    `yaml:"dpi"`
	PageSize       string `yaml:"page_size"`
	MarginTop      int    `yaml:"margin_top"`
	MarginBottom   int    `yaml:"margin_bottom"`
	MarginLeft     int    `yaml:"margin_left"`
	MarginRight    int    `yaml:"margin_right"`
	UnicodeEnabled bool   `yaml:"unicode_enabled"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(path string) (*Config, error) {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file not found: %s", path)
	}

	// Read file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Parse YAML
	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	// Apply defaults for missing values
	applyDefaults(&cfg)

	// Validate configuration
	if err = validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// validateConfig validates the configuration
func validateConfig(cfg *Config) error {
	// Validate mode
	if cfg.Mode != "ats" && cfg.Mode != "custom" {
		return fmt.Errorf("invalid mode: %s (must be 'ats' or 'custom')", cfg.Mode)
	}

	// Validate input file
	if cfg.General.InputFile == "" {
		return fmt.Errorf("input file not specified")
	}

	// Validate theme
	if cfg.Mode == "custom" {
		validThemes := []string{"default", "dark", "catppuccin-mocha", "catppuccin-latte", "nord", "github-dark"}
		themeValid := false
		for _, t := range validThemes {
			if cfg.Styling.Theme == t {
				themeValid = true
				break
			}
		}
		if !themeValid {
			return fmt.Errorf("invalid theme: %s", cfg.Styling.Theme)
		}
	}

	// Ensure required paths exist
	if cfg.Paths.StylesDir == "" {
		return fmt.Errorf("styles directory not specified")
	}

	if cfg.Paths.TemplateFile == "" {
		return fmt.Errorf("template file not specified")
	}

	// Convert any relative paths to absolute
	if !filepath.IsAbs(cfg.General.InputFile) {
		absPath, err := filepath.Abs(cfg.General.InputFile)
		if err != nil {
			return fmt.Errorf("error resolving input file path: %v", err)
		}
		cfg.General.InputFile = absPath
	}

	if !filepath.IsAbs(cfg.General.OutputFile) {
		absPath, err := filepath.Abs(cfg.General.OutputFile)
		if err != nil {
			return fmt.Errorf("error resolving output file path: %v", err)
		}
		cfg.General.OutputFile = absPath
	}

	return nil
}

// GetThemeCSSPath returns the path to the CSS file for the specified theme
func (cfg *Config) GetThemeCSSPath(themeName string) string {
	// Normalize theme name to match config keys
	normalizedTheme := themeName
	switch themeName {
	case "catppuccin-mocha":
		normalizedTheme = "catppuccin_mocha"
	case "catppuccin-latte":
		normalizedTheme = "catppuccin_latte"
	case "github-dark":
		normalizedTheme = "github_dark"
	}

	// Get theme file name from config
	fileName, exists := cfg.Paths.ThemeFiles[normalizedTheme]
	if !exists {
		// Fallback to default if theme not found
		fileName = cfg.Paths.ThemeFiles["default"]
	}

	// Return full path
	return filepath.Join(cfg.Paths.StylesDir, fileName)
}

// GetBaseStylePath returns the path to the base CSS file
func (cfg *Config) GetBaseStylePath() string {
	return filepath.Join(cfg.Paths.StylesDir, cfg.Paths.ThemeFiles["base"])
}

// GetATSStylePath returns the path to the ATS mode CSS file
func (cfg *Config) GetATSStylePath() string {
	return filepath.Join(cfg.Paths.StylesDir, cfg.Paths.ThemeFiles["ats"])
}
