#!/bin/bash
echo "Creating test file structure..."
mkdir -p styles
echo "/* Base styling */" >styles/base.css
echo "/* ATS styling */" >styles/ats.css
echo "/* Default theme */" >styles/default.css
echo "/* Nord theme */" >styles/nord.css
echo "/* Catppuccin Mocha theme */" >styles/catppuccin-mocha.css
echo "/* Catppuccin Latte theme */" >styles/catppuccin-latte.css
echo "/* Dark theme */" >styles/dark.css
