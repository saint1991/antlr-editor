grammar Expression;

// Parser Rules
expression
    : literal                                          # LiteralExpr
    | columnReference                                  # ColumnRefExpr
    | functionCall                                     # FunctionCallExpr
    | LPAREN expression RPAREN                         # ParenExpr
    | SUB expression                                   # UnaryMinusExpr
    | <assoc=right> expression POW expression          # PowerExpr
    | expression (MUL | DIV) expression                # MulDivExpr
    | expression (ADD | SUB) expression                # AddSubExpr
    | expression (LT | LE | GT | GE | EQ | NEQ) expression  # ComparisonExpr
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
LT  : '<' ;
LE  : '<=' ;
GT  : '>' ;
GE  : '>=' ;
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

// Literals - Order matters for proper tokenization
BOOLEAN_LITERAL
    : 'true'
    | 'false'
    ;

FLOAT_LITERAL
    : [0-9]+ '.' [0-9]+
    | [0-9]+ ('.' [0-9]+)? [eE] [+-]? [0-9]+
    ;

INTEGER_LITERAL
    : [0-9]+
    ;

STRING_LITERAL
    : '\'' ( ~['\r\n\\] | '\\' . )* '\''
    | '"'  ( ~["\r\n\\] | '\\' . )* '"'
    ;

// Function names (uppercase only) - must come before IDENTIFIER
FUNCTION_NAME
    : [A-Z]+
    ;

// Identifiers for column references (letters, digits, underscore)
IDENTIFIER
    : [a-zA-Z_][a-zA-Z0-9_]*
    ;

// Skip whitespace
WS : [ \t\r\n]+ -> skip ;