"""
Python bindings for the ANTLR expression analyzer.
"""
import ctypes
import platform
from pathlib import Path

from .models import TokenType, TokenInfo, ErrorInfo, AnalysisResult

# C struct definitions
class CTokenInfo(ctypes.Structure):
    """C struct for token information."""
    _fields_ = [
        ("token_type", ctypes.c_int),
        ("text", ctypes.c_char_p),
        ("start", ctypes.c_int32),
        ("end", ctypes.c_int32),
        ("line", ctypes.c_int32),
        ("column", ctypes.c_int32),
        ("is_valid", ctypes.c_int32),
    ]


class CErrorInfo(ctypes.Structure):
    """C struct for error information."""
    _fields_ = [
        ("message", ctypes.c_char_p),
        ("line", ctypes.c_int32),
        ("column", ctypes.c_int32),
        ("start", ctypes.c_int32),
        ("end", ctypes.c_int32),
    ]


class CAnalysisResult(ctypes.Structure):
    """C struct for analysis result."""
    _fields_ = [
        ("tokens", ctypes.POINTER(CTokenInfo)),
        ("token_count", ctypes.c_int32),
        ("errors", ctypes.POINTER(CErrorInfo)),
        ("error_count", ctypes.c_int32),
    ]


class Analyzer:
    """Python interface to the ANTLR expression analyzer."""
    
    def __init__(self, lib_path: Path | None = None):
        """
        Initialize the analyzer with the shared library.
        
        Args:
            lib_path: Path to the shared library. If None, will search in default locations.
        """
        self._lib = self._load_library(lib_path)
        self._setup_functions()
    
    def _load_library(self, lib_path: Path | None = None) -> ctypes.CDLL:
        """Load the shared library."""

        if lib_path is None:
            # Try to find library in standard locations
            lib_dir = Path(__file__).parent
            system = platform.system()
            
            if system == "Darwin":
                lib_name = "libanalyzer.dylib"
            elif system == "Linux":
                lib_name = "libanalyzer.so"
            elif system == "Windows":
                lib_name = "analyzer.dll"
            else:
                raise OSError(f"Unsupported platform: {system}")
            
            # Try package directory first
            lib_path = lib_dir / lib_name
            if not lib_path.exists():
                # Try parent directory
                lib_path = lib_dir.parent / lib_name
            
            if not lib_path.exists():
                raise OSError(f"Could not find {lib_name}")
        
        return ctypes.CDLL(str(lib_path))
    
    def _setup_functions(self):
        """Setup function signatures for the C library."""
        # ValidateFFI
        self._lib.ValidateFFI.argtypes = [ctypes.c_char_p, ctypes.c_int]
        self._lib.ValidateFFI.restype = ctypes.c_int
        
        # AnalyzeFFI
        self._lib.AnalyzeFFI.argtypes = [ctypes.c_char_p, ctypes.c_int]
        self._lib.AnalyzeFFI.restype = ctypes.POINTER(CAnalysisResult)
        
        # FreeAnalysisResult
        self._lib.FreeAnalysisResult.argtypes = [ctypes.POINTER(CAnalysisResult)]
        self._lib.FreeAnalysisResult.restype = None
    
    def validate(self, expression: str) -> bool:
        """
        Validate an expression.
        
        Args:
            expression: The expression to validate.
            
        Returns:
            True if the expression is valid, False otherwise.
        """
        if not expression:
            return False
        
        expr_bytes = expression.encode('utf-8')
        result = self._lib.ValidateFFI(expr_bytes, len(expr_bytes))
        return bool(result)
    
    def analyze(self, expression: str) -> AnalysisResult:
        """
        Analyze an expression and return detailed information.
        
        Args:
            expression: The expression to analyze.
            
        Returns:
            AnalysisResult containing tokens and errors.
        """
        if not expression:
            return AnalysisResult(tokens=[], errors=[])
        
        expr_bytes = expression.encode('utf-8')
        c_result_ptr = self._lib.AnalyzeFFI(expr_bytes, len(expr_bytes))
        
        if not c_result_ptr:
            return AnalysisResult(tokens=[], errors=[])
        
        try:
            c_result = c_result_ptr.contents
            
            # Convert tokens
            tokens = []
            for i in range(c_result.token_count):
                c_token = c_result.tokens[i]
                token = TokenInfo(
                    token_type=TokenType(c_token.token_type),
                    text=c_token.text.decode('utf-8') if c_token.text else "",
                    start=c_token.start,
                    end=c_token.end,
                    line=c_token.line,
                    column=c_token.column,
                    is_valid=bool(c_token.is_valid)
                )
                tokens.append(token)
            
            # Convert errors
            errors = []
            for i in range(c_result.error_count):
                c_error = c_result.errors[i]
                error = ErrorInfo(
                    message=c_error.message.decode('utf-8') if c_error.message else "",
                    line=c_error.line,
                    column=c_error.column,
                    start=c_error.start,
                    end=c_error.end
                )
                errors.append(error)
            
            return AnalysisResult(tokens=tokens, errors=errors)
        finally:
            # Free the C memory
            self._lib.FreeAnalysisResult(c_result_ptr)