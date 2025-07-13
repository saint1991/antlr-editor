"""
ANTLR Expression Analyzer Python Bindings

This package provides Python bindings for the Go-based ANTLR expression analyzer.
"""

from .analyzer_ffi import (
    AnalyzerFFI,
    validate,
    analyze,
    get_tokens,
    get_errors,
    has_errors,
    get_analyzer
)

__all__ = [
    'AnalyzerFFI',
    'validate',
    'analyze', 
    'get_tokens',
    'get_errors',
    'has_errors',
    'get_analyzer'
]

__version__ = '1.0.0'