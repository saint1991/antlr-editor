#ifndef TOKEN_H
#define TOKEN_H

#include <stdint.h>

enum TokenType {
	TOKEN_TYPE_STRING = 0,
	TOKEN_TYPE_INTEGER,
	TOKEN_TYPE_FLOAT,
	TOKEN_TYPE_BOOLEAN,
	TOKEN_TYPE_COLUMN_REFERENCE,
	TOKEN_TYPE_FUNCTION,
	TOKEN_TYPE_OPERATOR,
	TOKEN_TYPE_COMMA,
	TOKEN_TYPE_LEFT_PAREN,
	TOKEN_TYPE_RIGHT_PAREN,
	TOKEN_TYPE_LEFT_BRACKET,
	TOKEN_TYPE_RIGHT_BRACKET,
	TOKEN_TYPE_WHITESPACE,
	TOKEN_TYPE_ERROR,
	TOKEN_TYPE_EOF
};

typedef struct {
    enum TokenType token_type;  // TokenType enum value
    char* text;                 // Token text
    int32_t start;              // Start position
    int32_t end;                // End position
    int32_t line;               // Line number (1-based)
    int32_t column;             // Column number (0-based)
    int32_t is_valid;           // 0 or 1 for boolean
} CTokenInfo;

#endif // TOKEN_H
