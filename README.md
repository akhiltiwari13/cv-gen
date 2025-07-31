# Enhanced Markdown-to-PDF Resume Generator

This fork enhances the original Markdown-to-PDF Resume Generator with advanced styling options, custom themes optimized for 1-page resumes, font customization, improved Unicode character support, comprehensive logging, and multiple new light themes.
[original readme](./README_original.md)

## Features

- **Dual Modes**: Create both ATS-friendly and visually styled resumes from the same Markdown source
- **Configuration File**: All settings can be managed through a YAML configuration file
- **External CSS**: Theme styling is stored in separate CSS files for easy customization
- **1-Page Optimized Themes**: All themes are optimized to fit content on a single page like the professional theme
- **Custom Themes**: Built-in support for various themes including Catppuccin, Nord, TokyoNight, and multiple light themes
- **New Light Themes**: Added minimal-light, elegant-light, fresh-light, and corporate-light themes
- **Font Customization**: Use any font installed on your system
- **Unicode Support**: Properly renders international characters and symbols
- **Configurable Logging**: Set log level, output to file, and control formatting through configuration
- **Modular Design**: Well-structured codebase following modern Go practices

## Why This Exists

The original project focused solely on ATS compatibility. This fork maintains that capability while adding the option to create stylish, visually appealing resumes with a hacker aesthetic when you don't need strict ATS adherence.

## Project Structure

```
cv-gen/
├── cmd/cv-gen/        # Application entry point
├── internal/          # Internal packages
├── styles/            # CSS theme files
├── templates/         # HTML templates
├── config.yaml        # Default configuration
└── README.md          # Documentation
```

## Requirements

- **Go** (1.18+ recommended)
- **wkhtmltopdf** installed on your system
  (e.g., on Debian-based distros: `sudo apt-get install wkhtmltopdf`; on Mac: `brew install wkhtmltopdf`)
- System fonts for custom font options

## Installation

```bash
# Clone the repository
git clone https://github.com/akhiltiwari13/cv-gen.git
cd cv-gen

# Build the application
go build -o cv-gen ./cmd/cv-gen

# Or run directly
go run ./cmd/cv-gen
```

## Usage

### Basic Usage

```bash
# Generate using config file settings
./cv-gen

# Generate with custom options
./cv-gen -mode custom -theme nord -font "JetBrains Mono, monospace"

# Try different light themes
./cv-gen -mode custom -theme minimal-light
./cv-gen -mode custom -theme elegant-light
./cv-gen -mode custom -theme corporate-light

# Use debug logging with file output
./cv-gen -log-level debug

# List available themes
./cv-gen -list-themes

# List common system fonts
./cv-gen -list-fonts
```


### Command-line Options

| Option         | Description                | Default       |
| -------------- | -------------------------- | ------------- |
| `-config`      | Path to configuration file | `config.yaml` |
| `-input`       | Input markdown resume file | From config   |
| `-output`      | Output PDF file            | From config   |
| `-mode`        | Mode: 'ats' or 'custom'    | From config   |
| `-theme`       | Theme for custom mode      | From config   |
| `-font`        | Font family                | From config   |
| `-fontsize`    | Base font size             | From config   |
| `-margin`      | Margin size                | From config   |
| `-css`         | Custom CSS file            | From config   |
| `-log-level`   | Logging level (debug/info/warn/error) | info   |
| `-pretty-log`  | Use pretty logging format              | true   |
| `-list-themes` | List available themes      |               |
| `-list-fonts`  | List common system fonts   |               |
| `-version`     | Show version information   |               |

### Configuration File

The `config.yaml` file allows you to set all options:

```yaml
# General settings
general:
  input_file: "RESUME.md"     # Input markdown resume file
  output_file: "resume.pdf"   # Output PDF file

# Rendering mode: "ats" for ATS-friendly or "custom" for styled
mode: "custom"

# Styling configuration (optimized for 1-page layout)
styling:
  font_family: "Calibri, Arial, sans-serif"
  font_size: "10.5pt"         # Optimized for single page
  margin_size: "0.5in"        # Consistent margins
  theme: "professional"       # Choose from available themes
  custom_css_path: "./styles/professional.css"

# PDF generation settings
pdf:
  dpi: 300
  page_size: "A4"
  margin_top: 1
  margin_bottom: 1
  margin_left: 1
  margin_right: 1
  unicode_enabled: true

# Logging configuration
logging:
  level: "info"               # debug, info, warn, error
  log_file: ""                # Leave empty for stderr, or specify file path
  pretty: true                # Pretty print logs to console
```

## Available Themes

All themes are optimized for 1-page resumes with consistent formatting and spacing.

### Professional Themes
- `professional` - Modern professional resume style (ATS compatible)
- `default` - Clean, professional light theme
- `ats` - Maximum ATS compatibility with simple formatting

### Light Themes
- `minimal-light` - Clean and simple with serif fonts
- `elegant-light` - Sophisticated with Garamond typography
- `fresh-light` - Modern and vibrant with clean lines
- `corporate-light` - Professional business style
- `modern-clean` - Modern clean theme with Catppuccin Latte colors
- `catppuccin-latte` - Light Catppuccin theme with pastel accents

### Dark Themes
- `catppuccin-mocha` - Dark Catppuccin theme with vibrant accents
- `nord` - Cool blue-based dark theme with JetBrains Mono
- `tokyonight` - Popular dark theme with purple and blue accents
- `github-dark` - GitHub's dark mode theme

## 1-Page Resume Optimization

All themes have been optimized to create professional 1-page resumes:

- **Consistent Font Sizes**: 10.5pt body text, 18pt headers, 11pt section headers
- **Optimized Spacing**: Reduced margins and line spacing for maximum content
- **Smart Layout**: Bullet points, date ranges, and contact info positioned efficiently
- **Professional Margins**: 0.5in margins provide clean, printable layout

## Logging Configuration

The application now supports comprehensive logging:

```yaml
logging:
  level: "debug"              # Set verbosity: debug, info, warn, error
  log_file: "cv-gen.log"      # Output to file (empty = stderr)
  pretty: true                # Human-readable format for console
```

**Log Levels:**
- `debug` - Detailed information for troubleshooting
- `info` - General information about operations
- `warn` - Warning messages for potential issues
- `error` - Error messages only

## Customizing Themes

You can easily create your own themes:

1. Create a new CSS file in the `styles/` directory following the 1-page optimization patterns
2. Add the path to your theme in `config.yaml` under `paths.theme_files`
3. Select your theme using `-theme your-theme-name`

**Key CSS classes for 1-page optimization:**
- Use `font-size: 10.5pt` for body text
- Use `margin-top: 15px; margin-bottom: 10px` for section headers
- Use `margin-bottom: 3px` for list items
- Include `.date-range { float: right; font-size: 10pt; }` for dates

## Unicode Support

This fork has enhanced Unicode character support, making it suitable for resumes in multiple languages or with special symbols. Proper UTF-8 encoding is ensured throughout the conversion pipeline.

## ATS Optimization Tips

When using `-mode ats`:

- Keep your Markdown sections straightforward
- Ensure contact info appears as plain text near the top
- Avoid tables or columns
- Don't rely on complex formatting
