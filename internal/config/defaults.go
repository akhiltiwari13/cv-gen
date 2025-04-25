package config

// applyDefaults applies default values to missing configuration settings
func applyDefaults(cfg *Config) {
// General defaults
if cfg.General.InputFile == "" {
cfg.General.InputFile = "RESUME.md"
}
if cfg.General.OutputFile == "" {
cfg.General.OutputFile = "resume.pdf"
}

// Mode default
if cfg.Mode == "" {
cfg.Mode = "ats"
}

// Styling defaults
if cfg.Styling.FontFamily == "" {
cfg.Styling.FontFamily = "Arial, sans-serif"
}
if cfg.Styling.FontSize == "" {
cfg.Styling.FontSize = "12px"
}
if cfg.Styling.MarginSize == "" {
cfg.Styling.MarginSize = "20px"
}
if cfg.Styling.Theme == "" {
cfg.Styling.Theme = "default"
}

// Paths defaults
if cfg.Paths.StylesDir == "" {
cfg.Paths.StylesDir = "./styles"
}
if cfg.Paths.TemplateFile == "" {
cfg.Paths.TemplateFile = "./templates/resume.html"
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
"modern_clean":     "modern-clean.css",
"nord":             "nord.css",
"github_dark":      "github-dark.css",
}
} else {
// Ensure all required theme files are defined
if _, exists := cfg.Paths.ThemeFiles["base"]; !exists {
cfg.Paths.ThemeFiles["base"] = "base.css"
}
if _, exists := cfg.Paths.ThemeFiles["ats"]; !exists {
cfg.Paths.ThemeFiles["ats"] = "ats.css"
}
if _, exists := cfg.Paths.ThemeFiles["default"]; !exists {
cfg.Paths.ThemeFiles["default"] = "default.css"
}
if _, exists := cfg.Paths.ThemeFiles["modern_clean"]; !exists {
cfg.Paths.ThemeFiles["modern_clean"] = "modern-clean.css"
}
}

// PDF defaults
if cfg.PDF.DPI == 0 {
cfg.PDF.DPI = 96
}
if cfg.PDF.PageSize == "" {
cfg.PDF.PageSize = "A4"
}
if cfg.PDF.MarginTop == 0 {
cfg.PDF.MarginTop = 10
}
if cfg.PDF.MarginBottom == 0 {
cfg.PDF.MarginBottom = 10
}
if cfg.PDF.MarginLeft == 0 {
cfg.PDF.MarginLeft = 10
}
if cfg.PDF.MarginRight == 0 {
cfg.PDF.MarginRight = 10
}
}
