from dataclasses import dataclass

from .token import TokenInfo
from .error import ErrorInfo


@dataclass(frozen=True)
class TokenizeResult:
    """Result of tokenizing an expression."""

    tokens: list[TokenInfo]
    errors: list[ErrorInfo]

    @property
    def is_valid(self) -> bool:
        """Check if the expression is valid (no errors)."""
        return len(self.errors) == 0
