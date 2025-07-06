#!/bin/bash

# ANTLR Code Generation Script
# This script generates ANTLR parser code from the grammar file

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Get the directory where this script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

# Check if we're in the parser directory
if [ "$(basename "$PWD")" != "parser" ]; then
    echo -e "${YELLOW}Note: This script should be run from the parser directory${NC}"
    echo -e "${YELLOW}Switching to parser directory...${NC}"
    cd "$SCRIPT_DIR"
fi

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}Error: Docker is not installed or not in PATH${NC}"
    echo "Please install Docker to generate the ANTLR parser code"
    exit 1
fi

# Check if Docker daemon is running
if ! docker info &> /dev/null; then
    echo -e "${RED}Error: Docker daemon is not running${NC}"
    echo "Please start Docker and try again"
    exit 1
fi

echo -e "${GREEN}Generating ANTLR parser code...${NC}"
echo "Grammar file: $PROJECT_ROOT/grammar/Expression.g4"
echo "Output directory: $SCRIPT_DIR/gen/parser"

# Create output directory if it doesn't exist
mkdir -p "$SCRIPT_DIR/gen/parser"

# Run the Docker build command from the project root
cd "$PROJECT_ROOT"
if docker build --target antlr-generated --output=type=local,dest=analyzer/gen/parser -f analyzer/Dockerfile .; then
    echo -e "${GREEN}✓ ANTLR parser code generated successfully!${NC}"
    echo ""
    echo "Generated files:"
    cd "$SCRIPT_DIR"
    ls -la gen/parser/
else
    echo -e "${RED}✗ Failed to generate ANTLR parser code${NC}"
    exit 1
fi