#!/bin/bash

# Function to display help message
usage() {
    echo "Usage: $0 --c <old_string> --to <new_string>"
    exit 1
}

# Parse command-line arguments
while [[ "$#" -gt 0 ]]; do
    case "$1" in
        --c)
            OLD_STRING="$2"
            shift 2
            ;;
        --to)
            NEW_STRING="$2"
            shift 2
            ;;
        *)
            usage
            ;;
    esac
done

# Validate input
if [[ -z "$OLD_STRING" || -z "$NEW_STRING" ]]; then
    usage
fi

# Find and replace in all files within the current directory
find . -type f -exec sed -i "s|$OLD_STRING|$NEW_STRING|g" {} +

echo "Replaced all occurrences of '$OLD_STRING' with '$NEW_STRING'."
