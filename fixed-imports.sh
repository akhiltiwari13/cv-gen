#!/bin/bash
# Script to replace import paths in all Go files

# Define the find and replace strings
FIND="github.com/akhiltiwari13/cv-gen"
REPLACE="github.com/akhiltiwari13/resume"

# Find all .go files and perform the replacement
find . -name "*.go" -type f -exec sed -i '' "s|${FIND}|${REPLACE}|g" {} \;

echo "Fixed import paths in Go files"
