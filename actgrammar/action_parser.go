// Code generated from ActionParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package actgrammar // ActionParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ActionParser struct {
	*antlr.BaseParser
}

var actionparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func actionparserParserInit() {
	staticData := &actionparserParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'('", "')'", "','", "'$'", "'@'", "'='", "'+'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "SCOL", "L_PAREN", "R_PAREN", "COMMA", "DOLLAR", "AT", "EQ", "PLUS",
		"PERIOD", "SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_", "SQL_KEYWORDS",
		"SQL_STMT", "IDENTIFIER", "VARIABLE", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"statement", "sql_stmt", "call_stmt", "literal_value", "number_value",
		"variable_name", "table_name", "action_name", "column_name", "action_literal_value",
		"function_name", "expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 90, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 3, 0, 27, 8, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 3, 10, 57, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 69, 8, 11, 1, 11, 1, 11, 5, 11, 73, 8, 11,
		10, 11, 12, 11, 76, 9, 11, 1, 11, 1, 11, 3, 11, 80, 8, 11, 1, 11, 1, 11,
		1, 11, 5, 11, 85, 8, 11, 10, 11, 12, 11, 88, 9, 11, 1, 11, 0, 1, 22, 12,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 1, 2, 0, 19, 19, 21, 21,
		86, 0, 26, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 34, 1, 0, 0, 0, 6, 39, 1,
		0, 0, 0, 8, 41, 1, 0, 0, 0, 10, 43, 1, 0, 0, 0, 12, 45, 1, 0, 0, 0, 14,
		47, 1, 0, 0, 0, 16, 49, 1, 0, 0, 0, 18, 51, 1, 0, 0, 0, 20, 53, 1, 0, 0,
		0, 22, 79, 1, 0, 0, 0, 24, 27, 3, 2, 1, 0, 25, 27, 3, 4, 2, 0, 26, 24,
		1, 0, 0, 0, 26, 25, 1, 0, 0, 0, 27, 1, 1, 0, 0, 0, 28, 29, 5, 16, 0, 0,
		29, 30, 5, 1, 0, 0, 30, 3, 1, 0, 0, 0, 31, 32, 3, 10, 5, 0, 32, 33, 5,
		7, 0, 0, 33, 35, 1, 0, 0, 0, 34, 31, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35,
		36, 1, 0, 0, 0, 36, 37, 3, 22, 11, 0, 37, 38, 5, 1, 0, 0, 38, 5, 1, 0,
		0, 0, 39, 40, 7, 0, 0, 0, 40, 7, 1, 0, 0, 0, 41, 42, 5, 19, 0, 0, 42, 9,
		1, 0, 0, 0, 43, 44, 5, 18, 0, 0, 44, 11, 1, 0, 0, 0, 45, 46, 5, 17, 0,
		0, 46, 13, 1, 0, 0, 0, 47, 48, 5, 17, 0, 0, 48, 15, 1, 0, 0, 0, 49, 50,
		5, 17, 0, 0, 50, 17, 1, 0, 0, 0, 51, 52, 7, 0, 0, 0, 52, 19, 1, 0, 0, 0,
		53, 56, 5, 17, 0, 0, 54, 55, 5, 9, 0, 0, 55, 57, 5, 17, 0, 0, 56, 54, 1,
		0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 21, 1, 0, 0, 0, 58, 59, 6, 11, -1, 0,
		59, 80, 3, 18, 9, 0, 60, 80, 3, 10, 5, 0, 61, 62, 5, 2, 0, 0, 62, 63, 3,
		22, 11, 0, 63, 64, 5, 3, 0, 0, 64, 80, 1, 0, 0, 0, 65, 66, 3, 20, 10, 0,
		66, 68, 5, 2, 0, 0, 67, 69, 3, 22, 11, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1,
		0, 0, 0, 69, 74, 1, 0, 0, 0, 70, 71, 5, 4, 0, 0, 71, 73, 3, 22, 11, 0,
		72, 70, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1,
		0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 78, 5, 3, 0, 0, 78,
		80, 1, 0, 0, 0, 79, 58, 1, 0, 0, 0, 79, 60, 1, 0, 0, 0, 79, 61, 1, 0, 0,
		0, 79, 65, 1, 0, 0, 0, 80, 86, 1, 0, 0, 0, 81, 82, 10, 2, 0, 0, 82, 83,
		5, 8, 0, 0, 83, 85, 3, 22, 11, 3, 84, 81, 1, 0, 0, 0, 85, 88, 1, 0, 0,
		0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 23, 1, 0, 0, 0, 88, 86,
		1, 0, 0, 0, 7, 26, 34, 56, 68, 74, 79, 86,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ActionParserInit initializes any static state used to implement ActionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewActionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ActionParserInit() {
	staticData := &actionparserParserStaticData
	staticData.once.Do(actionparserParserInit)
}

// NewActionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewActionParser(input antlr.TokenStream) *ActionParser {
	ActionParserInit()
	this := new(ActionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &actionparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ActionParser.g4"

	return this
}

// ActionParser tokens.
const (
	ActionParserEOF                     = antlr.TokenEOF
	ActionParserSCOL                    = 1
	ActionParserL_PAREN                 = 2
	ActionParserR_PAREN                 = 3
	ActionParserCOMMA                   = 4
	ActionParserDOLLAR                  = 5
	ActionParserAT                      = 6
	ActionParserEQ                      = 7
	ActionParserPLUS                    = 8
	ActionParserPERIOD                  = 9
	ActionParserSELECT_                 = 10
	ActionParserINSERT_                 = 11
	ActionParserUPDATE_                 = 12
	ActionParserDELETE_                 = 13
	ActionParserWITH_                   = 14
	ActionParserSQL_KEYWORDS            = 15
	ActionParserSQL_STMT                = 16
	ActionParserIDENTIFIER              = 17
	ActionParserVARIABLE                = 18
	ActionParserUNSIGNED_NUMBER_LITERAL = 19
	ActionParserSIGNED_NUMBER_LITERAL   = 20
	ActionParserSTRING_LITERAL          = 21
	ActionParserWS                      = 22
	ActionParserTERMINATOR              = 23
	ActionParserBLOCK_COMMENT           = 24
	ActionParserLINE_COMMENT            = 25
)

// ActionParser rules.
const (
	ActionParserRULE_statement            = 0
	ActionParserRULE_sql_stmt             = 1
	ActionParserRULE_call_stmt            = 2
	ActionParserRULE_literal_value        = 3
	ActionParserRULE_number_value         = 4
	ActionParserRULE_variable_name        = 5
	ActionParserRULE_table_name           = 6
	ActionParserRULE_action_name          = 7
	ActionParserRULE_column_name          = 8
	ActionParserRULE_action_literal_value = 9
	ActionParserRULE_function_name        = 10
	ActionParserRULE_expr                 = 11
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sql_stmt() ISql_stmtContext
	Call_stmt() ICall_stmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Sql_stmt() ISql_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISql_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISql_stmtContext)
}

func (s *StatementContext) Call_stmt() ICall_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_stmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ActionParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(26)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ActionParserSQL_STMT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Sql_stmt()
		}

	case ActionParserL_PAREN, ActionParserIDENTIFIER, ActionParserVARIABLE, ActionParserUNSIGNED_NUMBER_LITERAL, ActionParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Call_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISql_stmtContext is an interface to support dynamic dispatch.
type ISql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SQL_STMT() antlr.TerminalNode
	SCOL() antlr.TerminalNode

	// IsSql_stmtContext differentiates from other interfaces.
	IsSql_stmtContext()
}

type Sql_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmtContext() *Sql_stmtContext {
	var p = new(Sql_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_sql_stmt
	return p
}

func (*Sql_stmtContext) IsSql_stmtContext() {}

func NewSql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmtContext {
	var p = new(Sql_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_sql_stmt

	return p
}

func (s *Sql_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_stmtContext) SQL_STMT() antlr.TerminalNode {
	return s.GetToken(ActionParserSQL_STMT, 0)
}

func (s *Sql_stmtContext) SCOL() antlr.TerminalNode {
	return s.GetToken(ActionParserSCOL, 0)
}

func (s *Sql_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitSql_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Sql_stmt() (localctx ISql_stmtContext) {
	this := p
	_ = this

	localctx = NewSql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ActionParserRULE_sql_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(ActionParserSQL_STMT)
	}
	{
		p.SetState(29)
		p.Match(ActionParserSCOL)
	}

	return localctx
}

// ICall_stmtContext is an interface to support dynamic dispatch.
type ICall_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	SCOL() antlr.TerminalNode
	Variable_name() IVariable_nameContext
	EQ() antlr.TerminalNode

	// IsCall_stmtContext differentiates from other interfaces.
	IsCall_stmtContext()
}

type Call_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_stmtContext() *Call_stmtContext {
	var p = new(Call_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_call_stmt
	return p
}

func (*Call_stmtContext) IsCall_stmtContext() {}

func NewCall_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_stmtContext {
	var p = new(Call_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_call_stmt

	return p
}

func (s *Call_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Call_stmtContext) SCOL() antlr.TerminalNode {
	return s.GetToken(ActionParserSCOL, 0)
}

func (s *Call_stmtContext) Variable_name() IVariable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_nameContext)
}

func (s *Call_stmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(ActionParserEQ, 0)
}

func (s *Call_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitCall_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Call_stmt() (localctx ICall_stmtContext) {
	this := p
	_ = this

	localctx = NewCall_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ActionParserRULE_call_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(31)
			p.Variable_name()
		}
		{
			p.SetState(32)
			p.Match(ActionParserEQ)
		}

	}
	{
		p.SetState(36)
		p.expr(0)
	}
	{
		p.SetState(37)
		p.Match(ActionParserSCOL)
	}

	return localctx
}

// ILiteral_valueContext is an interface to support dynamic dispatch.
type ILiteral_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsLiteral_valueContext differentiates from other interfaces.
	IsLiteral_valueContext()
}

type Literal_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_valueContext() *Literal_valueContext {
	var p = new(Literal_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_literal_value
	return p
}

func (*Literal_valueContext) IsLiteral_valueContext() {}

func NewLiteral_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_valueContext {
	var p = new(Literal_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_literal_value

	return p
}

func (s *Literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ActionParserSTRING_LITERAL, 0)
}

func (s *Literal_valueContext) UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ActionParserUNSIGNED_NUMBER_LITERAL, 0)
}

func (s *Literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitLiteral_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Literal_value() (localctx ILiteral_valueContext) {
	this := p
	_ = this

	localctx = NewLiteral_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ActionParserRULE_literal_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ActionParserUNSIGNED_NUMBER_LITERAL || _la == ActionParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumber_valueContext is an interface to support dynamic dispatch.
type INumber_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsNumber_valueContext differentiates from other interfaces.
	IsNumber_valueContext()
}

type Number_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_valueContext() *Number_valueContext {
	var p = new(Number_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_number_value
	return p
}

func (*Number_valueContext) IsNumber_valueContext() {}

func NewNumber_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_valueContext {
	var p = new(Number_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_number_value

	return p
}

func (s *Number_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_valueContext) UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ActionParserUNSIGNED_NUMBER_LITERAL, 0)
}

func (s *Number_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitNumber_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Number_value() (localctx INumber_valueContext) {
	this := p
	_ = this

	localctx = NewNumber_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ActionParserRULE_number_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(ActionParserUNSIGNED_NUMBER_LITERAL)
	}

	return localctx
}

// IVariable_nameContext is an interface to support dynamic dispatch.
type IVariable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIABLE() antlr.TerminalNode

	// IsVariable_nameContext differentiates from other interfaces.
	IsVariable_nameContext()
}

type Variable_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_nameContext() *Variable_nameContext {
	var p = new(Variable_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_variable_name
	return p
}

func (*Variable_nameContext) IsVariable_nameContext() {}

func NewVariable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_nameContext {
	var p = new(Variable_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_variable_name

	return p
}

func (s *Variable_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_nameContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ActionParserVARIABLE, 0)
}

func (s *Variable_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitVariable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Variable_name() (localctx IVariable_nameContext) {
	this := p
	_ = this

	localctx = NewVariable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ActionParserRULE_variable_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(ActionParserVARIABLE)
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ActionParserIDENTIFIER, 0)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitTable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Table_name() (localctx ITable_nameContext) {
	this := p
	_ = this

	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ActionParserRULE_table_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(ActionParserIDENTIFIER)
	}

	return localctx
}

// IAction_nameContext is an interface to support dynamic dispatch.
type IAction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsAction_nameContext differentiates from other interfaces.
	IsAction_nameContext()
}

type Action_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_nameContext() *Action_nameContext {
	var p = new(Action_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_action_name
	return p
}

func (*Action_nameContext) IsAction_nameContext() {}

func NewAction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_nameContext {
	var p = new(Action_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_action_name

	return p
}

func (s *Action_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ActionParserIDENTIFIER, 0)
}

func (s *Action_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitAction_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Action_name() (localctx IAction_nameContext) {
	this := p
	_ = this

	localctx = NewAction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ActionParserRULE_action_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(ActionParserIDENTIFIER)
	}

	return localctx
}

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_column_name
	return p
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ActionParserIDENTIFIER, 0)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitColumn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Column_name() (localctx IColumn_nameContext) {
	this := p
	_ = this

	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ActionParserRULE_column_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(ActionParserIDENTIFIER)
	}

	return localctx
}

// IAction_literal_valueContext is an interface to support dynamic dispatch.
type IAction_literal_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsAction_literal_valueContext differentiates from other interfaces.
	IsAction_literal_valueContext()
}

type Action_literal_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_literal_valueContext() *Action_literal_valueContext {
	var p = new(Action_literal_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_action_literal_value
	return p
}

func (*Action_literal_valueContext) IsAction_literal_valueContext() {}

func NewAction_literal_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_literal_valueContext {
	var p = new(Action_literal_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_action_literal_value

	return p
}

func (s *Action_literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_literal_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ActionParserSTRING_LITERAL, 0)
}

func (s *Action_literal_valueContext) UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ActionParserUNSIGNED_NUMBER_LITERAL, 0)
}

func (s *Action_literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_literal_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitAction_literal_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Action_literal_value() (localctx IAction_literal_valueContext) {
	this := p
	_ = this

	localctx = NewAction_literal_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ActionParserRULE_action_literal_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ActionParserUNSIGNED_NUMBER_LITERAL || _la == ActionParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	PERIOD() antlr.TerminalNode

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_function_name
	return p
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ActionParserIDENTIFIER)
}

func (s *Function_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ActionParserIDENTIFIER, i)
}

func (s *Function_nameContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(ActionParserPERIOD, 0)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitFunction_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Function_name() (localctx IFunction_nameContext) {
	this := p
	_ = this

	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ActionParserRULE_function_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(ActionParserIDENTIFIER)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ActionParserPERIOD {
		{
			p.SetState(54)
			p.Match(ActionParserPERIOD)
		}
		{
			p.SetState(55)
			p.Match(ActionParserIDENTIFIER)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetElevate_expr returns the elevate_expr rule contexts.
	GetElevate_expr() IExprContext

	// SetElevate_expr sets the elevate_expr rule contexts.
	SetElevate_expr(IExprContext)

	// Getter signatures
	Action_literal_value() IAction_literal_valueContext
	Variable_name() IVariable_nameContext
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Function_name() IFunction_nameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	PLUS() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	elevate_expr IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetElevate_expr() IExprContext { return s.elevate_expr }

func (s *ExprContext) SetElevate_expr(v IExprContext) { s.elevate_expr = v }

func (s *ExprContext) Action_literal_value() IAction_literal_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_literal_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_literal_valueContext)
}

func (s *ExprContext) Variable_name() IVariable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_nameContext)
}

func (s *ExprContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ActionParserL_PAREN, 0)
}

func (s *ExprContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ActionParserR_PAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Function_name() IFunction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *ExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ActionParserCOMMA)
}

func (s *ExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ActionParserCOMMA, i)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ActionParserPLUS, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ActionParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, ActionParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ActionParserUNSIGNED_NUMBER_LITERAL, ActionParserSTRING_LITERAL:
		{
			p.SetState(59)
			p.Action_literal_value()
		}

	case ActionParserVARIABLE:
		{
			p.SetState(60)
			p.Variable_name()
		}

	case ActionParserL_PAREN:
		{
			p.SetState(61)
			p.Match(ActionParserL_PAREN)
		}
		{
			p.SetState(62)

			var _x = p.expr(0)

			localctx.(*ExprContext).elevate_expr = _x
		}
		{
			p.SetState(63)
			p.Match(ActionParserR_PAREN)
		}

	case ActionParserIDENTIFIER:
		{
			p.SetState(65)
			p.Function_name()
		}
		{
			p.SetState(66)
			p.Match(ActionParserL_PAREN)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3014660) != 0 {
			{
				p.SetState(67)
				p.expr(0)
			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ActionParserCOMMA {
			{
				p.SetState(70)
				p.Match(ActionParserCOMMA)
			}
			{
				p.SetState(71)
				p.expr(0)
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(77)
			p.Match(ActionParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_expr)
			p.SetState(81)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(82)
				p.Match(ActionParserPLUS)
			}
			{
				p.SetState(83)
				p.expr(3)
			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ActionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ActionParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
