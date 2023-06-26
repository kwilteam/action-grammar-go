// Code generated from ActionParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package actgrammar // ActionParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by ActionParser.
type ActionParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ActionParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ActionParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by ActionParser#sql_stmt.
	VisitSql_stmt(ctx *Sql_stmtContext) interface{}

	// Visit a parse tree produced by ActionParser#call_stmt.
	VisitCall_stmt(ctx *Call_stmtContext) interface{}

	// Visit a parse tree produced by ActionParser#call_receivers.
	VisitCall_receivers(ctx *Call_receiversContext) interface{}

	// Visit a parse tree produced by ActionParser#literal_value.
	VisitLiteral_value(ctx *Literal_valueContext) interface{}

	// Visit a parse tree produced by ActionParser#variable_name.
	VisitVariable_name(ctx *Variable_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#block_variable_name.
	VisitBlock_variable_name(ctx *Block_variable_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#fn_name.
	VisitFn_name(ctx *Fn_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#call_body.
	VisitCall_body(ctx *Call_bodyContext) interface{}

	// Visit a parse tree produced by ActionParser#fn_arg.
	VisitFn_arg(ctx *Fn_argContext) interface{}

	// Visit a parse tree produced by ActionParser#fn_arg_list.
	VisitFn_arg_list(ctx *Fn_arg_listContext) interface{}
}
