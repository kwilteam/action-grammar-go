/*
 * A ANTLR4 grammar for Action. ONLY temporary.
 * Developed by the Kuneiform team.
*/

parser grammar ActionParser;

options {
    tokenVocab=ActionLexer;
}

statement:
    stmt+
;

stmt:
    sql_stmt
    | call_stmt
;

sql_stmt:
    SQL_STMT SCOL
;

call_stmt:
    (variable_name EQ)? expr SCOL
;

literal_value:
    STRING_LITERAL
    | UNSIGNED_NUMBER_LITERAL
;

number_value:
    UNSIGNED_NUMBER_LITERAL
;

variable_name:
    VARIABLE
;

table_name:
    IDENTIFIER
;

action_name:
    IDENTIFIER
;

column_name:
    IDENTIFIER
;

action_literal_value:
    STRING_LITERAL
    | UNSIGNED_NUMBER_LITERAL
;

function_name:
    IDENTIFIER (PERIOD IDENTIFIER)?
;

expr:
    action_literal_value
    | variable_name
    | L_PAREN elevate_expr=expr R_PAREN
    | expr PLUS expr
    | function_name L_PAREN expr? (COMMA expr)* R_PAREN
;