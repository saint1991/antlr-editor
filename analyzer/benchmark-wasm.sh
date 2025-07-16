#!/bin/bash

# benchmark-wasm.sh - Compare WASM binary sizes and performance

set -e

echo "=== WASM Size Benchmark ==="
echo

# Generate ANTLR parser if needed
if [ ! -d "gen/parser" ]; then
    echo "Generating ANTLR parser..."
    ./codegen.sh
fi

# Build WASM modules
echo "Building WASM modules..."
./build-wasm.sh

DIST_DIR="dist"

echo
echo "=== Size Comparison ==="

if [ -f "${DIST_DIR}/analyzer-go.wasm" ] && [ -f "${DIST_DIR}/analyzer-tinygo.wasm" ]; then
    GO_SIZE=$(stat -f%z "${DIST_DIR}/analyzer-go.wasm" 2>/dev/null || stat -c%s "${DIST_DIR}/analyzer-go.wasm")
    TINYGO_SIZE=$(stat -f%z "${DIST_DIR}/analyzer-tinygo.wasm" 2>/dev/null || stat -c%s "${DIST_DIR}/analyzer-tinygo.wasm")
    
    echo "Standard Go WASM: $(numfmt --to=iec $GO_SIZE) ($GO_SIZE bytes)"
    echo "TinyGo WASM: $(numfmt --to=iec $TINYGO_SIZE) ($TINYGO_SIZE bytes)"
    
    REDUCTION=$(( (GO_SIZE - TINYGO_SIZE) * 100 / GO_SIZE ))
    echo "TinyGo reduction: $REDUCTION%"
    
    if [ -f "${DIST_DIR}/analyzer.wasm" ]; then
        OPTIMIZED_SIZE=$(stat -f%z "${DIST_DIR}/analyzer.wasm" 2>/dev/null || stat -c%s "${DIST_DIR}/analyzer.wasm")
        echo "TinyGo + wasm-opt: $(numfmt --to=iec $OPTIMIZED_SIZE) ($OPTIMIZED_SIZE bytes)"
        
        TOTAL_REDUCTION=$(( (GO_SIZE - OPTIMIZED_SIZE) * 100 / GO_SIZE ))
        echo "Total reduction: $TOTAL_REDUCTION%"
    fi
    
    echo
    echo "=== Detailed File Sizes ==="
    ls -lh "${DIST_DIR}"/*.wasm | awk '{print $5 "\t" $9}' | sort -k1 -h
else
    echo "WASM files not found. Build may have failed."
    exit 1
fi

echo
echo "=== Compression Analysis ==="

# Test gzip compression (common for web serving)
for file in "${DIST_DIR}"/*.wasm; do
    if [ -f "$file" ]; then
        basename=$(basename "$file")
        original_size=$(stat -f%z "$file" 2>/dev/null || stat -c%s "$file")
        gzip_size=$(gzip -c "$file" | wc -c)
        compression_ratio=$(( (original_size - gzip_size) * 100 / original_size ))
        
        echo "$basename:"
        echo "  Original: $(numfmt --to=iec $original_size)"
        echo "  Gzipped:  $(numfmt --to=iec $gzip_size) (${compression_ratio}% reduction)"
    fi
done

echo
echo "Benchmark completed!"