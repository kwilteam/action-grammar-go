// Code generated from ActionParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package actgrammar // ActionParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by ActionParser.
type ActionParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ActionParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ActionParser#sql_stmt.
	VisitSql_stmt(ctx *Sql_stmtContext) interface{}

	// Visit a parse tree produced by ActionParser#call_stmt.
	VisitCall_stmt(ctx *Call_stmtContext) interface{}

	// Visit a parse tree produced by ActionParser#literal_value.
	VisitLiteral_value(ctx *Literal_valueContext) interface{}

	// Visit a parse tree produced by ActionParser#number_value.
	VisitNumber_value(ctx *Number_valueContext) interface{}

	// Visit a parse tree produced by ActionParser#variable_name.
	VisitVariable_name(ctx *Variable_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#action_name.
	VisitAction_name(ctx *Action_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#action_literal_value.
	VisitAction_literal_value(ctx *Action_literal_valueContext) interface{}

	// Visit a parse tree produced by ActionParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by ActionParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
