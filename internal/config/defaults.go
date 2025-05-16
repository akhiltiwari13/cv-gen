package config

import "github.com/akhiltiwari13/cv-gen/internal/logging"

// applyDefaults applies default values to missing configuration settings
func applyDefaults(cfg *Config) {
	logger := logging.GetLogger()
	// General defaults
	if cfg.General.InputFile == "" {
		cfg.General.InputFile = "RESUME.md"
		logger.Debug().Str("input_file", cfg.General.InputFile).Msg("Applied default input file")
	}
	if cfg.General.OutputFile == "" {
		cfg.General.OutputFile = "resume.pdf"
		logger.Debug().Str("output_file", cfg.General.OutputFile).Msg("Applied default output file")
	}

	// Mode default
	if cfg.Mode == "" {
		cfg.Mode = "ats"
		logger.Debug().Str("mode", cfg.Mode).Msg("Applied default mode")
	}

	// Styling defaults
	if cfg.Styling.FontFamily == "" {
		cfg.Styling.FontFamily = "Arial, sans-serif"
		logger.Debug().Str("font_family", cfg.Styling.FontFamily).Msg("Applied default font family")
	}
	if cfg.Styling.FontSize == "" {
		cfg.Styling.FontSize = "12px"
		logger.Debug().Str("font_size", cfg.Styling.FontSize).Msg("Applied default font size")
	}
	if cfg.Styling.MarginSize == "" {
		cfg.Styling.MarginSize = "20px"
		logger.Debug().Str("margin_size", cfg.Styling.MarginSize).Msg("Applied default margin size")
	}
	if cfg.Styling.Theme == "" {
		logger.Debug().Str("theme", cfg.Styling.Theme).Msg("Applied default theme")
		cfg.Styling.Theme = "default"
	}

	// Paths defaults
	if cfg.Paths.StylesDir == "" {
		cfg.Paths.StylesDir = "./styles"
		logger.Debug().Str("styles_dir", cfg.Paths.StylesDir).Msg("Applied default styles directory")
	}
	if cfg.Paths.TemplateFile == "" {
		cfg.Paths.TemplateFile = "./templates/resume.html"
		logger.Debug().Str("template_file", cfg.Paths.TemplateFile).Msg("Applied default template file")
	}

	// Theme files defaults
	if cfg.Paths.ThemeFiles == nil {
		cfg.Paths.ThemeFiles = map[string]string{
			"base":             "base.css",
			"ats":              "ats.css",
			"default":          "default.css",
			"dark":             "dark.css",
			"catppuccin_mocha": "catppuccin-mocha.css",
			"catppuccin_latte": "catppuccin-latte.css",
			"nord":             "nord.css",
			"github_dark":      "github-dark.css",
			"professional":     "professional.css",
		}
		logger.Debug().Msg("Applied default theme files map")
	} else {
		// Ensure all required theme files are defined
		if _, exists := cfg.Paths.ThemeFiles["base"]; !exists {
			cfg.Paths.ThemeFiles["base"] = "base.css"
			logger.Debug().Msg("Added missing base theme file")
		}
		if _, exists := cfg.Paths.ThemeFiles["ats"]; !exists {
			cfg.Paths.ThemeFiles["ats"] = "ats.css"
			logger.Debug().Msg("Added missing ats theme file")
		}
		if _, exists := cfg.Paths.ThemeFiles["default"]; !exists {
			cfg.Paths.ThemeFiles["default"] = "default.css"
			logger.Debug().Msg("Added missing default theme file")
		}
		if _, exists := cfg.Paths.ThemeFiles["professional"]; !exists {
			logger.Debug().Msg("Added missing professional theme file")
			cfg.Paths.ThemeFiles["professional"] = "professional.css"
		}
	}

	// PDF defaults
	if cfg.PDF.DPI == 0 {
		cfg.PDF.DPI = 96
		logger.Debug().Int("dpi", cfg.PDF.DPI).Msg("Applied default DPI")
	}
	if cfg.PDF.PageSize == "" {
		cfg.PDF.PageSize = "A4"
		logger.Debug().Str("page_size", cfg.PDF.PageSize).Msg("Applied default page size")
	}
	if cfg.PDF.MarginTop == 0 {
		cfg.PDF.MarginTop = 10
		logger.Debug().Int("margin_top", cfg.PDF.MarginTop).Msg("Applied default top margin")
	}
	if cfg.PDF.MarginBottom == 0 {
		cfg.PDF.MarginBottom = 10
		logger.Debug().Int("margin_bottom", cfg.PDF.MarginBottom).Msg("Applied default bottom margin")
	}
	if cfg.PDF.MarginLeft == 0 {
		cfg.PDF.MarginLeft = 10
		logger.Debug().Int("margin_left", cfg.PDF.MarginLeft).Msg("Applied default left margin")
	}
	if cfg.PDF.MarginRight == 0 {
		cfg.PDF.MarginRight = 10
		logger.Debug().Int("margin_right", cfg.PDF.MarginRight).Msg("Applied default right margin")
	}
	logger.Debug().Msg("Default values applied to configuration")
}
