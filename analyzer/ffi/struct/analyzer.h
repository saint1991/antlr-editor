#ifndef ANALYZER_H
#define ANALYZER_H

#include <stdint.h>

#include "error.h"
#include "token.h"

typedef struct {
    CTokenInfo* tokens;  // Array of tokens
    int32_t token_count; // Number of tokens
    CErrorInfo* errors;  // Array of errors
    int32_t error_count; // Number of errors
} CTokenizeResult;

#endif // ANALYZER_H
