#!/bin/bash

# build-wasm.sh - Build WASM modules with TinyGo and optimization

set -e

# Configuration
WASM_DIR="dist"
WASM_SOURCE="./wasm/analyzer.go"

# Create output directory
mkdir -p "${WASM_DIR}"

echo "Building WASM modules..."

# Build with standard Go compiler (for comparison)
echo "Building with standard Go..."
GOOS=js GOARCH=wasm go build -o "${WASM_DIR}/analyzer-go.wasm" "${WASM_SOURCE}"

# Check if TinyGo is available
if command -v tinygo &> /dev/null; then
    echo "Building with TinyGo..."
    tinygo build -o "${WASM_DIR}/analyzer-tinygo.wasm" -target wasm "${WASM_SOURCE}"
else
    echo "Warning: TinyGo not found. Skipping TinyGo build."
    echo "Install TinyGo from: https://tinygo.org/getting-started/install/"
    exit 1
fi

# Check if wasm-opt is available for optimization
if command -v wasm-opt &> /dev/null; then
    echo "Optimizing WASM modules with wasm-opt..."
    
    # Optimize standard Go WASM
    wasm-opt -O3 -o "${WASM_DIR}/analyzer-go-optimized.wasm" "${WASM_DIR}/analyzer-go.wasm"
    
    # Optimize TinyGo WASM with different optimization levels
    wasm-opt -O3 -o "${WASM_DIR}/analyzer-tinygo-O3.wasm" "${WASM_DIR}/analyzer-tinygo.wasm"
    wasm-opt -Oz --strip-debug -o "${WASM_DIR}/analyzer.wasm" "${WASM_DIR}/analyzer-tinygo.wasm"
    
    echo "Final optimized WASM: ${WASM_DIR}/analyzer.wasm"
else
    echo "Warning: wasm-opt not found. WASM optimization skipped."
    echo "Install from: https://github.com/WebAssembly/binaryen"
    cp "${WASM_DIR}/analyzer-tinygo.wasm" "${WASM_DIR}/analyzer.wasm"
fi

# Display file sizes
echo
echo "File sizes:"
ls -lh "${WASM_DIR}"/*.wasm | awk '{print $5 "\t" $9}'

echo
echo "Build completed successfully!"