/*
 * A ANTLR4 grammar for Action. ONLY temporary.
 * Developed by the Kuneiform team.
*/

lexer grammar ActionLexer;

// symbols
SCOL:      ';';
L_PAREN:   '(';
R_PAREN:   ')';
COMMA:     ',';
DOLLAR:    '$';
AT:        '@';
EQ:        '=';
PLUS:      '+';
PERIOD:    '.';

// keywords
// sql keywords
SELECT_:   [sS][eE][lL][eE][cC][tT];
INSERT_:   [iI][nN][sS][eE][rR][tT];
UPDATE_:   [uU][pP][dD][aA][tT][eE];
DELETE_:   [dD][eE][lL][eE][tT][eE];
WITH_:     [wW][iI][tT][hH]        ;

SQL_KEYWORDS: SELECT_ | INSERT_ | UPDATE_ | DELETE_ | WITH_;
// we only need sql statement as a whole, sql-parser will parse it
SQL_STMT: SQL_KEYWORDS ~[;}]+;

// literals
IDENTIFIER:
    [a-zA-Z] [a-zA-Z_0-9]*
;

VARIABLE: DOLLAR IDENTIFIER;
BLOCK_VARIABLE: AT IDENTIFIER;

UNSIGNED_NUMBER_LITERAL:
    DIGIT+
;

SIGNED_NUMBER_LITERAL:
    [+-]? DIGIT+
;

STRING_LITERAL:
    DOUBLE_QUOTE_STRING
    | SINGLE_QUOTE_STRING
;

WS:            [ \t]+        -> channel(HIDDEN);
TERMINATOR:    [\r\n]+       -> channel(HIDDEN);
BLOCK_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT:  '//' ~[\r\n]* -> channel(HIDDEN);

// fragments
fragment WSNL: [ \t\r\n]+; // whitespace with new line
fragment DIGIT: [0-9];

fragment DOUBLE_QUOTE_STRING_CHAR: ~["\r\n\\] | ('\\' .);
fragment SINGLE_QUOTE_STRING_CHAR: ~['\r\n\\] | ('\\' .);

fragment DOUBLE_QUOTE_STRING: '"' DOUBLE_QUOTE_STRING_CHAR* '"';
fragment SINGLE_QUOTE_STRING: '\'' SINGLE_QUOTE_STRING_CHAR* '\'';