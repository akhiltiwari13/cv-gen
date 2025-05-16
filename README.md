# Enhanced Markdown-to-PDF Resume Generator

This fork enhances the original Markdown-to-PDF Resume Generator with advanced styling options, custom themes (including Professional, Catppuccin and Nord), font customization, improved Unicode character support, and comprehensive logging.
[original readme](./README_original.md)

## Features

- **Dual Modes**: Create both ATS-friendly and visually styled resumes from the same Markdown source
- **Configuration File**: All settings can be managed through a YAML configuration file
- **External CSS**: Theme styling is stored in separate CSS files for easy customization
- **Custom Themes**: Built-in support for various themes including Catppuccin Mocha, Catppuccin Latte, and Nord
- **Font Customization**: Use any font installed on your system
- **Unicode Support**: Properly renders international characters and symbols
- **Detailed Logging**: Comprehensive logging throughout the application for easier debugging
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
| `-log-level`         | Logging level| info   |
| `-pretty-log`         | Use pretty logging format| true   |
| `-list-themes` | List available themes      |               |
| `-list-fonts`  | List common system fonts   |               |
| `-version`     | Show version information   |               |

### Configuration File

The `config.yaml` file allows you to set all options:

```yaml
# General settings
general:
  input_file: "RESUME.md" # Input markdown resume file
  output_file: "resume.pdf" # Output PDF file

# Rendering mode: "ats" for ATS-friendly or "custom" for styled
mode: "ats"

# Styling configuration
styling:
  font_family: "Arial, sans-serif"
  font_size: "12px"
  margin_size: "20px"
  theme: "default"
  custom_css_path: ""
# Paths and additional settings are also available
```

## Available Themes

- `professional` - Modern professional resume style (ATS compatible)
- `default` - Clean, professional light theme
- `dark` - Dark background with light text
- `catppuccin-mocha` - Dark Catppuccin theme with vibrant accents
- `catppuccin-latte` - Light Catppuccin theme with pastel accents
- `nord` - Cool blue-based dark theme
- `github-dark` - GitHub's dark mode theme

## Customizing Themes

You can easily create your own themes:

1. Create a new CSS file in the `styles/` directory
2. Add the path to your theme in `config.yaml` under `paths.theme_files`
3. Select your theme using `-theme your-theme-name`

## Unicode Support

This fork has enhanced Unicode character support, making it suitable for resumes in multiple languages or with special symbols. Proper UTF-8 encoding is ensured throughout the conversion pipeline.

## ATS Optimization Tips

When using `-mode ats`:

- Keep your Markdown sections straightforward
- Ensure contact info appears as plain text near the top
- Avoid tables or columns
- Don't rely on complex formatting
