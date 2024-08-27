#!/bin/bash

# Define variables
OUTPUT_FILE="bootstrap"
ZIP_FILE="url-shortener.zip"
CMD_DIR="./cmd/main.go"

ZIP_PATTERN="*.zip"

# Borrar archivos existentes que coincidan con el patrón
echo "Deleting old ZIP files matching pattern: $ZIP_PATTERN"
rm -f $ZIP_PATTERN

# Compile the Go project for Linux
echo "Compiling the project..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_FILE $CMD_DIR

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo "Build successful!"

# Create the ZIP file
echo "Creating ZIP file..."
zip -j $ZIP_FILE $OUTPUT_FILE

# Check if the ZIP creation was successful
if [ $? -ne 0 ]; then
    echo "Failed to create ZIP file!"
    exit 1
fi

echo "ZIP file created successfully: $ZIP_FILE"

# Clean up the compiled binary
echo "Cleaning up..."
rm $OUTPUT_FILE

echo "Done!"
