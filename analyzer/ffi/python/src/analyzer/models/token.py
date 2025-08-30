from dataclasses import dataclass
from enum import IntEnum


class TokenType(IntEnum):
    """Token types for expression analysis."""

    STRING = 0
    INTEGER = 1
    FLOAT = 2
    BOOLEAN = 3
    COLUMN_REFERENCE = 4
    FUNCTION = 5
    OPERATOR = 6
    COMMA = 7
    LEFT_PAREN = 8
    RIGHT_PAREN = 9
    LEFT_BRACKET = 10
    RIGHT_BRACKET = 11
    WHITESPACE = 12
    ERROR = 13
    EOF = 14


@dataclass(frozen=True)
class TokenInfo:
    """Information about a token in the expression."""

    token_type: TokenType
    text: str
    start: int
    end: int
    line: int
    column: int
