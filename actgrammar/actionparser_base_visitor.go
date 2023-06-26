// Code generated from ActionParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package actgrammar // ActionParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseActionParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseActionParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitSql_stmt(ctx *Sql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitCall_stmt(ctx *Call_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitCall_receivers(ctx *Call_receiversContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitLiteral_value(ctx *Literal_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitVariable_name(ctx *Variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitBlock_variable_name(ctx *Block_variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitFn_name(ctx *Fn_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitCall_body(ctx *Call_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitFn_arg(ctx *Fn_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseActionParserVisitor) VisitFn_arg_list(ctx *Fn_arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}
