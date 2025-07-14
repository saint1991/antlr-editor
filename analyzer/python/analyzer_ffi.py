#!/usr/bin/env python3
"""
Python FFI bindings for the ANTLR Expression Analyzer

This module provides Python bindings for the Go-based ANTLR expression analyzer
using ctypes to interface with the compiled shared library.
"""

import ctypes
import json
import platform
import os
from pathlib import Path
from typing import Dict, List, Optional, Union


class AnalyzerFFI:
    """Python FFI wrapper for the ANTLR Expression Analyzer"""
    
    def __init__(self, library_path: Optional[Union[str, Path]] = None):
        """
        Initialize the analyzer FFI wrapper
        
        Args:
            library_path: Path to the shared library. If None, attempts to auto-detect.
        """
        self._lib = None
        self._load_library(library_path)
        self._setup_function_signatures()
    
    def _get_default_library_path(self) -> str:
        """Determine the default library path based on the current platform"""
        system = platform.system().lower()
        
        # Map platform info to library names
        if system == "linux":
            return "libanalyzer.so"
        elif system == "darwin":
            return "libanalyzer.dylib"
        elif system == "windows":
            return "libanalyzer.dll"
        
        # Fallback to .so for Unix-like systems
        return "libanalyzer.so"
    
    def _load_library(self, library_path: Optional[Union[str, Path]]):
        """Load the shared library"""
        if library_path is None:
            # Try to find the library in the same directory as this script
            script_dir = Path(__file__).parent
            library_name = self._get_default_library_path()
            library_path = script_dir / library_name
            
            if not library_path.exists():
                # Try looking in common locations
                for search_path in [
                    Path.cwd(),
                    script_dir.parent / "dist",
                    Path("/usr/local/lib"),
                    Path("/usr/lib"),
                ]:
                    potential_path = search_path / library_name
                    if potential_path.exists():
                        library_path = potential_path
                        break
                else:
                    raise FileNotFoundError(f"Could not find library: {library_name}")
        
        try:
            self._lib = ctypes.CDLL(str(library_path))
        except OSError as e:
            raise RuntimeError(f"Failed to load library {library_path}: {e}")
    
    def _setup_function_signatures(self):
        """Set up ctypes function signatures for the FFI functions"""
        # ValidateFFIString(expression *C.char) C.int
        self._lib.ValidateFFIString.argtypes = [ctypes.c_char_p]
        self._lib.ValidateFFIString.restype = ctypes.c_int
        
        # AnalyzeFFI(expression *C.char) *C.char
        self._lib.AnalyzeFFI.argtypes = [ctypes.c_char_p]
        self._lib.AnalyzeFFI.restype = ctypes.c_char_p
        
        # FreeString(s *C.char)
        self._lib.FreeString.argtypes = [ctypes.c_char_p]
        self._lib.FreeString.restype = None
    
    def validate(self, expression: str) -> bool:
        """
        Validate an expression for syntax correctness
        
        Args:
            expression: The expression string to validate
            
        Returns:
            True if the expression is syntactically valid, False otherwise
        """
        if not expression:
            return False
        
        expression_bytes = expression.encode('utf-8')
        result = self._lib.ValidateFFIString(expression_bytes)
        return result == 1
    
    def analyze(self, expression: str) -> Dict:
        """
        Analyze an expression and return detailed token information
        
        Args:
            expression: The expression string to analyze
            
        Returns:
            Dictionary containing tokens and errors information
        """
        if not expression:
            return {"tokens": [], "errors": []}
        
        expression_bytes = expression.encode('utf-8')
        result_ptr = self._lib.AnalyzeFFI(expression_bytes)
        
        if not result_ptr:
            return {"tokens": [], "errors": [{"message": "Analysis failed", "line": 1, "column": 0, "start": 0, "end": len(expression)}]}
        
        try:
            # Convert C string to Python string
            result_json = ctypes.string_at(result_ptr).decode('utf-8')
            result_data = json.loads(result_json)
            return result_data
        except (json.JSONDecodeError, UnicodeDecodeError) as e:
            return {"tokens": [], "errors": [{"message": f"Failed to parse result: {e}", "line": 1, "column": 0, "start": 0, "end": len(expression)}]}
        finally:
            # Always free the allocated string
            self._lib.FreeString(result_ptr)
    
    def get_tokens(self, expression: str) -> List[Dict]:
        """
        Get just the token information from an expression
        
        Args:
            expression: The expression string to analyze
            
        Returns:
            List of token dictionaries
        """
        result = self.analyze(expression)
        return result.get("tokens", [])
    
    def get_errors(self, expression: str) -> List[Dict]:
        """
        Get just the error information from an expression
        
        Args:
            expression: The expression string to analyze
            
        Returns:
            List of error dictionaries
        """
        result = self.analyze(expression)
        return result.get("errors", [])
    
    def has_errors(self, expression: str) -> bool:
        """
        Check if an expression has any syntax errors
        
        Args:
            expression: The expression string to check
            
        Returns:
            True if there are syntax errors, False otherwise
        """
        errors = self.get_errors(expression)
        return len(errors) > 0


# Convenience functions for direct usage
_default_analyzer = None

def get_analyzer() -> AnalyzerFFI:
    """Get the default analyzer instance"""
    global _default_analyzer
    if _default_analyzer is None:
        _default_analyzer = AnalyzerFFI()
    return _default_analyzer

def validate(expression: str) -> bool:
    """Validate an expression using the default analyzer"""
    return get_analyzer().validate(expression)

def analyze(expression: str) -> Dict:
    """Analyze an expression using the default analyzer"""
    return get_analyzer().analyze(expression)

def get_tokens(expression: str) -> List[Dict]:
    """Get tokens from an expression using the default analyzer"""
    return get_analyzer().get_tokens(expression)

def get_errors(expression: str) -> List[Dict]:
    """Get errors from an expression using the default analyzer"""
    return get_analyzer().get_errors(expression)

def has_errors(expression: str) -> bool:
    """Check if an expression has errors using the default analyzer"""
    return get_analyzer().has_errors(expression)


if __name__ == "__main__":
    # Example usage
    analyzer = AnalyzerFFI()
    
    test_expressions = [
        "2 + 3 * 4",
        "func(a, b)",
        "invalid syntax !@#",
        "x.field",
        "array[index]"
    ]
    
    for expr in test_expressions:
        print(f"Expression: {expr}")
        print(f"  Valid: {analyzer.validate(expr)}")
        
        result = analyzer.analyze(expr)
        print(f"  Tokens: {len(result.get('tokens', []))}")
        print(f"  Errors: {len(result.get('errors', []))}")
        
        if result.get('errors'):
            for error in result['errors']:
                print(f"    Error: {error.get('message', 'Unknown error')}")
        
        print()