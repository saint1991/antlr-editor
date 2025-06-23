grammar Expression;

// Parser Rules
expression
    : literal                                          # LiteralExpr
    | columnReference                                  # ColumnRefExpr
    | functionCall                                     # FunctionCallExpr
    | IDENTIFIER                                       # IdentifierExpr
    | LPAREN expression RPAREN                         # ParenExpr
    | SUB expression                                   # UnaryMinusExpr
    | <assoc=right> expression POW expression          # PowerExpr
    | expression (MUL | DIV) expression                # MulDivExpr
    | expression (ADD | SUB) expression                # AddSubExpr
    | expression (EQ | NEQ) expression                 # ComparisonExpr
    | expression AND expression                        # AndExpr
    | expression OR expression                         # OrExpr
    ;

literal
    : STRING_LITERAL
    | INTEGER_LITERAL
    | FLOAT_LITERAL
    | BOOLEAN_LITERAL
    ;

columnReference
    : LBRACKET IDENTIFIER RBRACKET
    ;

functionCall
    : FUNCTION_NAME LPAREN argumentList? RPAREN
    ;

argumentList
    : expression (COMMA expression)*
    ;

// Lexer Rules

// Operators
ADD : '+' ;
SUB : '-' ;
MUL : '*' ;
DIV : '/' ;
POW : '^' ;
EQ  : '==' ;
NEQ : '!=' ;
OR  : '||' ;
AND : '&&' ;

// Delimiters
LPAREN   : '(' ;
RPAREN   : ')' ;
LBRACKET : '[' ;
RBRACKET : ']' ;
COMMA    : ',' ;

// Literals
STRING_LITERAL
    : '\'' ( ~['\r\n\\] | '\\' . )* '\''
    | '"'  ( ~["\r\n\\] | '\\' . )* '"'
    ;

INTEGER_LITERAL
    : [0-9]+
    ;

FLOAT_LITERAL
    : [0-9]+ '.' [0-9]+
    | [0-9]+ ('.' [0-9]+)? [eE] [+-]? [0-9]+
    ;

BOOLEAN_LITERAL
    : 'true'
    | 'false'
    ;

// Function names (uppercase only)
FUNCTION_NAME
    : [A-Z]+
    ;

// Identifiers for column references (letters, digits, underscore)
IDENTIFIER
    : [a-zA-Z_][a-zA-Z0-9_]*
    ;

// Skip whitespace
WS : [ \t\r\n]+ -> skip ;