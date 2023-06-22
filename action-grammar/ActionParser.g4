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
    (call_receivers EQ)?
    call_body SCOL
;

call_receivers:
    variable_name (COMMA variable_name)*
;

literal_value:
    STRING_LITERAL
    | UNSIGNED_NUMBER_LITERAL
;

variable_name:
    VARIABLE
;

// `a.b` is only for extension calls for now
fn_name:
    IDENTIFIER (PERIOD IDENTIFIER)?
;

// use expr in the future, limit syntax for now
call_body:
    fn_name
    L_PAREN fn_arg_list R_PAREN
;

fn_arg:
    literal_value
    | variable_name
;

fn_arg_list:
    fn_arg? (COMMA fn_arg)*
;


//expr:
//    literal_value
//    | variable_name
//    | L_PAREN elevate_expr=expr R_PAREN
//    | expr PLUS expr
//    | function_name L_PAREN expr? (COMMA expr)* R_PAREN
//;