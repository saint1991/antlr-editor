# ANTLR code generation stage
ARG AMAZON_CORRETTO_TAG=24-alpine3.21
ARG BUSYBOX_TAG=1.37
ARG GOLANG_TAG=1.24-bookworm


FROM amazoncorretto:${AMAZON_CORRETTO_TAG} AS antlr-generator

# Install wget for downloading ANTLR4
RUN apk update && apk add --no-cache wget

# Install latest ANTLR4 tool
ARG ANTLR_VERSION=4.13.2
RUN wget https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar -O /opt/antlr4.jar

# Set working directory for ANTLR generation
WORKDIR /antlr

# Copy grammar files (rename to match grammar name)
COPY grammar/Expression.g4 ./grammar/Expression.g4

# Generate Go parser from grammar (visitor pattern only)
RUN java -jar /opt/antlr4.jar -Dlanguage=Go -o parser -package parser -visitor -no-listener grammar/Expression.g4


FROM scratch AS antlr-generated

COPY --from=antlr-generator /antlr/parser/grammar/* /


# WASM build stage with TinyGo
FROM golang:${GOLANG_TAG} AS wasm-builder

# Install TinyGo and WASM build dependencies
RUN apt-get update && apt-get install -y \
    wget \
    binaryen \
    && rm -rf /var/lib/apt/lists/*

# Install TinyGo
ARG TINYGO_VERSION=0.38.0
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb \
    && dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb \
    && rm tinygo_${TINYGO_VERSION}_amd64.deb

# Create working directories
RUN mkdir -p /go/dist /go/src

WORKDIR /go/src

# Copy go module files
COPY analyzer/go.mod analyzer/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY analyzer/ ./
COPY --from=antlr-generator /antlr/parser/grammar/* ./gen/parser/

# Build WASM module with standard Go (for comparison)
RUN GOOS=js GOARCH=wasm go build -o analyzer-go.wasm ./wasm/analyzer.go

# Build WASM module with TinyGo
RUN tinygo build -o analyzer-tinygo.wasm -target wasm ./wasm/analyzer.go

# Optimize both WASM binaries
RUN wasm-opt -O3 -o /go/dist/analyzer-go.wasm analyzer-go.wasm
RUN wasm-opt -O3 -o /go/dist/analyzer-tinygo.wasm analyzer-tinygo.wasm
RUN wasm-opt -Oz --strip-debug -o /go/dist/analyzer.wasm analyzer-tinygo.wasm


# Export WASM artifact
FROM scratch AS wasm-output

COPY --from=wasm-builder /go/dist/analyzer.wasm /analyzer.wasm
COPY --from=wasm-builder /go/dist/analyzer-go.wasm /analyzer-go.wasm
COPY --from=wasm-builder /go/dist/analyzer-tinygo.wasm /analyzer-tinygo.wasm