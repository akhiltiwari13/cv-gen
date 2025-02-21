# Markdown-to-PDF Resume Generator

This repo contains a small Go program that converts a Markdown file (e.g., `RESUME.md`) into a PDF. The main objective is to produce an **ATS-friendly** resume where the text is genuinely machine-readable, rather than rendered as images or unselectable text.

## Why This Exists

I wanted to tune and automate my resume, striking a balance between a minimal, professional format and ensuring automated screening systems can parse everything. Many PDF-export tools add excessive styling or embed fonts in ways that can confuse ATS (Applicant Tracking Systems). This project uses [blackfriday](https://github.com/russross/blackfriday) to convert Markdown to HTML and [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) to render that HTML into a clean PDF.

## How It Works

1. **Markdown to HTML:** `blackfriday` converts the contents of `RESUME.md` into HTML.
2. **HTML to PDF:** `wkhtmltopdf` (through `go-wkhtmltopdf`) turns that HTML into a PDF with minimal styling, using only system-safe fonts.
3. **ATS-Focused Output:** The design avoids complicated layouts, columns, or embedded images. The generated PDF has real text with proper spacing and UTF-8 encoding.

## Requirements

- **Go** (1.18+ recommended)
- **wkhtmltopdf** installed on your system  
  (e.g., on Debian-based distros: `sudo apt-get install wkhtmltopdf`; on Mac: `brew install wkhtmltopdf`)

## Usage

1. **Clone** this repo.
2. **Copy** the example resume template `cp RESUME.md.example RESUME.md`
3. **Edit** `RESUME.md` with your resume content.
4. **Install dependencies** `go get .`
5. **Run** `go run .`

## Tuning for ATS

- Keep your Markdown sections straightforward (headings, bullet points, short paragraphs).
- Avoid tables, columns, or heavy images.
- Ensure contact info appears as plain text near the top.
- Don’t rely on external fonts or weird CSS.
