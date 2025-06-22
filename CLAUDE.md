# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Rule

- MUST NOT edit .gitignore
- MUST NOT edit content of .git directory.

## Project Overview

See [README.md](./README.md)

## Development Commands

### Parser WASM Build Commands

**Note: All Docker commands must be executed from the project root directory.**

Build specific stages of the Docker image:

```bash
# Build ANTLR generator stage (generates Go parser from grammar)
docker build --target antlr-generator -t antlr-editor:antlr-generator -f parser/Dockerfile .

## Architecture

*Architecture details will be documented as the codebase grows*