#!/usr/bin/env python3
"""
Example usage of the ANTLR Expression Analyzer Python bindings
"""

from analyzer_ffi import AnalyzerFFI, validate, analyze

def main():
    print("ANTLR Expression Analyzer - Python FFI Example")
    print("=" * 50)
    
    # Initialize the analyzer
    try:
        analyzer = AnalyzerFFI()
        print("✓ Analyzer initialized successfully")
    except Exception as e:
        print(f"✗ Failed to initialize analyzer: {e}")
        return
    
    # Test expressions
    test_cases = [
        ("2 + 3 * 4", "Simple arithmetic expression"),
        ("func(a, b, c)", "Function call with parameters"),
        ("obj.property", "Object property access"),
        ("array[0]", "Array indexing"),
        ("x + y * (z - 1)", "Complex arithmetic with parentheses"),
        ("invalid ! syntax @", "Invalid expression"),
        ("", "Empty expression"),
        ("true && false", "Boolean expression"),
        ("x > 5 ? 'yes' : 'no'", "Ternary operator"),
    ]
    
    print("\nTesting expressions:")
    print("-" * 50)
    
    for expr, description in test_cases:
        print(f"\nExpression: '{expr}'")
        print(f"Description: {description}")
        
        # Test validation
        is_valid = analyzer.validate(expr)
        print(f"Valid: {is_valid}")
        
        # Get detailed analysis
        result = analyzer.analyze(expr)
        tokens = result.get('tokens', [])
        errors = result.get('errors', [])
        
        print(f"Tokens found: {len(tokens)}")
        if tokens:
            for i, token in enumerate(tokens[:5]):  # Show first 5 tokens
                token_type = token.get('type', 'Unknown')
                text = token.get('text', '')
                start = token.get('start', 0)
                end = token.get('end', 0)
                print(f"  Token {i+1}: '{text}' ({token_type}) [{start}:{end}]")
            if len(tokens) > 5:
                print(f"  ... and {len(tokens) - 5} more tokens")
        
        if errors:
            print(f"Errors found: {len(errors)}")
            for error in errors:
                message = error.get('message', 'Unknown error')
                line = error.get('line', 1)
                column = error.get('column', 0)
                print(f"  Error at line {line}, column {column}: {message}")
        else:
            print("No errors found")
    
    print("\n" + "=" * 50)
    print("Example completed")

def test_convenience_functions():
    """Test the convenience functions"""
    print("\nTesting convenience functions:")
    print("-" * 30)
    
    expr = "2 + 3 * 4"
    
    print(f"Expression: '{expr}'")
    print(f"validate(expr): {validate(expr)}")
    
    tokens = get_tokens(expr)
    print(f"get_tokens(expr): {len(tokens)} tokens")
    
    errors = get_errors(expr)
    print(f"get_errors(expr): {len(errors)} errors")
    
    has_err = has_errors(expr)
    print(f"has_errors(expr): {has_err}")

if __name__ == "__main__":
    main()
    
    # Import convenience functions for testing
    from analyzer_ffi import get_tokens, get_errors, has_errors
    test_convenience_functions()