#!/bin/bash

# --- Configuration ---
# Get the directory where the script is located
SCRIPT_PATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

# Calculate paths (mimicking the PS Resolve-Path logic)
# ProjectRoot is 2 levels up (../../)
PROJECT_ROOT="$(realpath "$SCRIPT_PATH/../..")"
# ScanTarget is 1 level up (../) aka the service folder
SCAN_TARGET="$(realpath "$SCRIPT_PATH/..")"

TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
OUTPUT_DIR="$SCAN_TARGET/.dumps"
OUTPUT_FILE="$OUTPUT_DIR/codebase_$TIMESTAMP.txt"

# Folders to completely skip
EXCLUDED_DIRS=(".git" ".idea" ".vscode" ".dumps" ".scripts" ".trash" "tmp" "bin" "obj" "node_modules" ".terraform" "terraform.tfstate.d")

# Files to skip (Extensions without dot)
EXCLUDED_EXTENSIONS=("png" "jpg" "jpeg" "exe" "dll" "zip" "7z" "iso")
# Specific filenames to skip
EXCLUDED_FILES=("go.sum" "package-lock.json" "dump_codebase.ps1" "dump_codebase.sh" "codebase.txt")

# --- Setup Output ---
mkdir -p "$OUTPUT_DIR"
# Create/Clear the file
> "$OUTPUT_FILE"

# ANSI Colors
MAGENTA='\033[0;35m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo -e "Snapshot Root: ${MAGENTA}$PROJECT_ROOT${NC}"
echo -e "Scanning Only: ${YELLOW}$SCAN_TARGET${NC}"
echo -e "Output File:   ${CYAN}$OUTPUT_FILE${NC}"

# --- Recursive Function ---
walk_directory() {
    local current_path="$1"

    # Enable dotglob to see hidden files (.git, .vscode)
    # Enable nullglob to handle empty directories gracefully
    shopt -s dotglob nullglob
    local items=("$current_path"/*)
    shopt -u dotglob nullglob

    for item in "${items[@]}"; do
        local name=$(basename "$item")

        if [ -d "$item" ]; then
            # 1. Handle Directory: Check exclusions
            local skip=false
            for excl in "${EXCLUDED_DIRS[@]}"; do
                if [ "$name" == "$excl" ]; then skip=true; break; fi
            done

            if [ "$skip" == "true" ]; then
                # echo "Skipping dir: $name"
                continue
            fi

            # Recurse
            walk_directory "$item"
        elif [ -f "$item" ]; then
            # 2. Handle File: Check extensions/names

            # Get extension
            local ext="${name##*.}"
            if [ "$ext" == "$name" ]; then ext=""; fi # handle files with no extension

            # Check extension exclusion
            local skip_ext=false
            for excl in "${EXCLUDED_EXTENSIONS[@]}"; do
                if [ "$ext" == "$excl" ]; then skip_ext=true; break; fi
            done
            if [ "$skip_ext" == "true" ]; then continue; fi

            # Check filename exclusion
            local skip_name=false
            for excl in "${EXCLUDED_FILES[@]}"; do
                if [ "$name" == "$excl" ]; then skip_name=true; break; fi
            done
            if [ "$skip_name" == "true" ]; then continue; fi

            # Check _templ.go specifically
            if [[ "$name" == *"_templ.go" ]]; then continue; fi

            # 3. Append to Dump
            # Calculate path relative to PROJECT_ROOT
            local relative_path="${item#$PROJECT_ROOT/}"

            {
                echo "=================================================="
                echo "FILE: $relative_path"
                echo "=================================================="
            } >> "$OUTPUT_FILE"

            # Append content
            # We check if it's a binary file roughly by using grep for NUL bytes
            # If it's binary, we skip printing content to avoid messing up the text file
            if grep -qI "." "$item" || [ ! -s "$item" ]; then
                cat "$item" >> "$OUTPUT_FILE"
            else
                echo "[BINARY OR UNREADABLE CONTENT]" >> "$OUTPUT_FILE"
            fi

            echo -e "\n" >> "$OUTPUT_FILE"
        fi
    done
}

# --- Execute ---
echo -e "${YELLOW}Starting scan...${NC}"
walk_directory "$SCAN_TARGET"
echo -e "${GREEN}Done!${NC}"
