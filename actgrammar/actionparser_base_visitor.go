// Code generated from ActionParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package actgrammar // ActionParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseActionParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseActionParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitSql_stmt(ctx *Sql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitCall_stmt(ctx *Call_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitLiteral_value(ctx *Literal_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitNumber_value(ctx *Number_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitVariable_name(ctx *Variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitAction_name(ctx *Action_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitAction_literal_value(ctx *Action_literal_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
