#!/bin/bash

# ==========================================
#  Tutoring Platform - Environment Doctor
# ==========================================

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "\n${BLUE}🏥 Running System Diagnosis...${NC}\n"

ERRORS=0

# Helper Function to check a command
check_cmd() {
    NAME=$1
    CMD=$2
    INSTALL_TIP=$3

    if command -v $CMD &> /dev/null; then
        VERSION=$($CMD --version 2>/dev/null | head -n 1)
        if [ -z "$VERSION" ]; then VERSION="Detected"; fi
        echo -e "${GREEN}[OK]${NC} $NAME found"
    else
        echo -e "${RED}[MISSING]${NC} $NAME is not installed."
        echo -e "      ${YELLOW}👉 Fix: $INSTALL_TIP${NC}"
        ERRORS=$((ERRORS+1))
    fi
}

# --- 1. Core System Tools ---
echo -e "${YELLOW}--- Core Tools ---${NC}"
check_cmd "Git" "git" "sudo apt install git"
check_cmd "Make" "make" "sudo apt install make"
check_cmd "Curl" "curl" "sudo apt install curl"
check_cmd "Docker" "docker" "Follow official Docker docs for Ubuntu"

# Check Docker Permissions
if command -v docker &> /dev/null; then
    if docker ps &> /dev/null; then
        echo -e "${GREEN}[OK]${NC} Docker Daemon is accessible"
    else
        echo -e "${RED}[ERROR]${NC} Cannot connect to Docker daemon."
        echo -e "       ${YELLOW}👉 Fix: sudo usermod -aG docker \$USER && newgrp docker${NC}"
        ERRORS=$((ERRORS+1))
    fi
fi

# --- 2. Go Ecosystem ---
echo -e "\n${YELLOW}--- Go Environment ---${NC}"
check_cmd "Go" "go" "sudo snap install go --classic"

# Check GOPATH/bin is in PATH
if [[ ":$PATH:" != *":$(go env GOPATH)/bin:"* ]]; then
    echo -e "${RED}[WARNING]${NC} $(go env GOPATH)/bin is NOT in your PATH."
    echo -e "          You might not be able to run 'air' or 'templ'."
    echo -e "          ${YELLOW}👉 Fix: export PATH=\$PATH:\$(go env GOPATH)/bin${NC}"
fi

check_cmd "Air (Hot Reload)" "air" "go install github.com/air-verse/air@latest"
check_cmd "Templ (UI Gen)" "templ" "go install github.com/a-h/templ/cmd/templ@latest"
check_cmd "Migrate (DB)" "migrate" "Download from github.com/golang-migrate/migrate"

# --- 3. Third-Party Integrations ---
echo -e "\n${YELLOW}--- Cloud Services ---${NC}"
check_cmd "Stripe CLI" "stripe" "sudo apt install stripe"
check_cmd "Google Cloud SDK" "gcloud" "Follow Google Cloud SDK docs"

# --- 4. Configuration Files ---
echo -e "\n${YELLOW}--- Configuration ---${NC}"

if [ -f ".env" ]; then
    echo -e "${GREEN}[OK]${NC} .env file exists"
else
    echo -e "${RED}[MISSING]${NC} .env file not found."
    echo -e "      ${YELLOW}👉 Fix: cp .env.example .env${NC}"
    ERRORS=$((ERRORS+1))
fi

ADC_PATH="$HOME/.config/gcloud/application_default_credentials.json"
if [ -f "$ADC_PATH" ]; then
    echo -e "${GREEN}[OK]${NC} Google Credentials found"
else
    echo -e "${RED}[MISSING]${NC} Google Credentials not found at default path."
    echo -e "      ${YELLOW}👉 Fix: gcloud auth application-default login${NC}"
    ERRORS=$((ERRORS+1))
fi

# --- Summary ---
echo -e "\n================================="
if [ $ERRORS -eq 0 ]; then
    echo -e "${GREEN}✅ You are ready to code!${NC}"
else
    echo -e "${RED}❌ Found $ERRORS issue(s) to fix.${NC}"
fi
echo -e "=================================\n"
