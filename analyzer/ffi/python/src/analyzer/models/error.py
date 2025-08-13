from dataclasses import dataclass


@dataclass(frozen=True)
class ErrorInfo:
    """Information about an error in the expression."""

    message: str
    line: int
    column: int
    start: int
    end: int
