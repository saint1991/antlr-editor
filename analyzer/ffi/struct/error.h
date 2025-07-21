#ifndef ERROR_H
#define ERROR_H

#include <stdint.h>

typedef struct {
    char* message;       // Error message
    int32_t line;        // Error line (1-based)
    int32_t column;      // Error column (0-based)
    int32_t start;       // Error start position
    int32_t end;         // Error end position
} CErrorInfo;

#endif // ERROR_H
