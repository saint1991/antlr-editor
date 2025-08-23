#!/usr/bin/env python3
"""
Example usage of the ANTLR Expression Analyzer Python bindings.
"""

from analyzer import Analyzer, TokenType


def example_validate(analyzer: Analyzer) -> None:
    """
    Example function to demonstrate simple validation of expressions.
    """

    print("===== Simple Validation =====")
    expressions = [
        "[age] > 18",
        "AND([name] == 'John', [age] >= 21)",
        "price * quantity - discount",
        "invalid expression >",
    ]

    for expr in expressions:
        is_valid = analyzer.validate(expr)
        status = "✓ Valid" if is_valid else "✗ Invalid"
        print(f"{status}: {expr}")


def example_analyze(analyzer: Analyzer) -> None:
    """
    Example function to demonstrate detailed analysis of expressions.
    """
    print("===== Detailed Analysis =====")

    expression = "AND([age] > 18, OR([status] == 'active', [status] == 'premium'))"
    result = analyzer.tokenize(expression)

    print(f"Expression: {expression}")
    print(f"Valid: {result.is_valid}")

    if result.is_valid:
        print("Tokens:")
        for token in result.tokens:
            if token.token_type != TokenType.WHITESPACE:
                print(f"  {token.text} -> {token.token_type.name}")
    else:
        print("Errors:")
        for error in result.errors:
            print(f"  {error.message} at line {error.line}, column {error.column}")


def example_error_detection(analyzer: Analyzer) -> None:
    """
    Example function to demonstrate error detection in expressions.
    """
    print("===== Error Detection =====")

    invalid_expression = (
        "AND([age > 18, OR([status] == 'active', [status] == 'premium'))"
    )
    result = analyzer.tokenize(invalid_expression)

    print(f"  Expression: {invalid_expression}")
    print(f"  Valid: {result.is_valid}")

    if result.errors:
        print("  Errors found:")
        for error in result.errors:
            print(f"    - {error.message} @ line {error.line}, column {error.column}")


def main() -> None:
    # Create an analyzer instance
    analyzer = Analyzer()

    # Example 1: Simple validation
    example_validate(analyzer)
    print()

    # Example 2: Detailed analysis
    example_analyze(analyzer)
    print()

    # Example 3: Error detection
    example_error_detection(analyzer)
    print()


if __name__ == "__main__":
    main()
