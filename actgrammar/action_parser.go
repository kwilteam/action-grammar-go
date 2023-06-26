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
		"SQL_STMT", "IDENTIFIER", "VARIABLE", "BLOCK_VARIABLE", "UNSIGNED_NUMBER_LITERAL",
		"SIGNED_NUMBER_LITERAL", "STRING_LITERAL", "WS", "TERMINATOR", "BLOCK_COMMENT",
		"LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"statement", "stmt", "sql_stmt", "call_stmt", "call_receivers", "literal_value",
		"variable_name", "block_variable_name", "fn_name", "call_body", "fn_arg",
		"fn_arg_list",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 84, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 4, 0, 26, 8, 0, 11, 0, 12, 0, 27, 1, 1, 1, 1, 3,
		1, 32, 8, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 40, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 48, 8, 4, 10, 4, 12, 4, 51, 9, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 62, 8, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 72, 8, 10, 1, 11, 3,
		11, 75, 8, 11, 1, 11, 1, 11, 5, 11, 79, 8, 11, 10, 11, 12, 11, 82, 9, 11,
		1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 1, 2, 0,
		20, 20, 22, 22, 80, 0, 25, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 33, 1, 0,
		0, 0, 6, 39, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 52, 1, 0, 0, 0, 12, 54,
		1, 0, 0, 0, 14, 56, 1, 0, 0, 0, 16, 58, 1, 0, 0, 0, 18, 63, 1, 0, 0, 0,
		20, 71, 1, 0, 0, 0, 22, 74, 1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1,
		0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28,
		1, 1, 0, 0, 0, 29, 32, 3, 4, 2, 0, 30, 32, 3, 6, 3, 0, 31, 29, 1, 0, 0,
		0, 31, 30, 1, 0, 0, 0, 32, 3, 1, 0, 0, 0, 33, 34, 5, 16, 0, 0, 34, 35,
		5, 1, 0, 0, 35, 5, 1, 0, 0, 0, 36, 37, 3, 8, 4, 0, 37, 38, 5, 7, 0, 0,
		38, 40, 1, 0, 0, 0, 39, 36, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 1,
		0, 0, 0, 41, 42, 3, 18, 9, 0, 42, 43, 5, 1, 0, 0, 43, 7, 1, 0, 0, 0, 44,
		49, 3, 12, 6, 0, 45, 46, 5, 4, 0, 0, 46, 48, 3, 12, 6, 0, 47, 45, 1, 0,
		0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 9,
		1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 53, 7, 0, 0, 0, 53, 11, 1, 0, 0, 0,
		54, 55, 5, 18, 0, 0, 55, 13, 1, 0, 0, 0, 56, 57, 5, 19, 0, 0, 57, 15, 1,
		0, 0, 0, 58, 61, 5, 17, 0, 0, 59, 60, 5, 9, 0, 0, 60, 62, 5, 17, 0, 0,
		61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 17, 1, 0, 0, 0, 63, 64, 3,
		16, 8, 0, 64, 65, 5, 2, 0, 0, 65, 66, 3, 22, 11, 0, 66, 67, 5, 3, 0, 0,
		67, 19, 1, 0, 0, 0, 68, 72, 3, 10, 5, 0, 69, 72, 3, 12, 6, 0, 70, 72, 3,
		14, 7, 0, 71, 68, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72,
		21, 1, 0, 0, 0, 73, 75, 3, 20, 10, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0,
		0, 0, 75, 80, 1, 0, 0, 0, 76, 77, 5, 4, 0, 0, 77, 79, 3, 20, 10, 0, 78,
		76, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0,
		0, 81, 23, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 8, 27, 31, 39, 49, 61, 71, 74,
		80,
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
	ActionParserBLOCK_VARIABLE          = 19
	ActionParserUNSIGNED_NUMBER_LITERAL = 20
	ActionParserSIGNED_NUMBER_LITERAL   = 21
	ActionParserSTRING_LITERAL          = 22
	ActionParserWS                      = 23
	ActionParserTERMINATOR              = 24
	ActionParserBLOCK_COMMENT           = 25
	ActionParserLINE_COMMENT            = 26
)

// ActionParser rules.
const (
	ActionParserRULE_statement           = 0
	ActionParserRULE_stmt                = 1
	ActionParserRULE_sql_stmt            = 2
	ActionParserRULE_call_stmt           = 3
	ActionParserRULE_call_receivers      = 4
	ActionParserRULE_literal_value       = 5
	ActionParserRULE_variable_name       = 6
	ActionParserRULE_block_variable_name = 7
	ActionParserRULE_fn_name             = 8
	ActionParserRULE_call_body           = 9
	ActionParserRULE_fn_arg              = 10
	ActionParserRULE_fn_arg_list         = 11
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

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

func (s *StatementContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&458752) != 0) {
		{
			p.SetState(24)
			p.Stmt()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sql_stmt() ISql_stmtContext
	Call_stmt() ICall_stmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Sql_stmt() ISql_stmtContext {
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

func (s *StmtContext) Call_stmt() ICall_stmtContext {
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

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ActionParserRULE_stmt)

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

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ActionParserSQL_STMT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Sql_stmt()
		}

	case ActionParserIDENTIFIER, ActionParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
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
	p.EnterRule(localctx, 4, ActionParserRULE_sql_stmt)

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
		p.SetState(33)
		p.Match(ActionParserSQL_STMT)
	}
	{
		p.SetState(34)
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
	Call_body() ICall_bodyContext
	SCOL() antlr.TerminalNode
	Call_receivers() ICall_receiversContext
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

func (s *Call_stmtContext) Call_body() ICall_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_bodyContext)
}

func (s *Call_stmtContext) SCOL() antlr.TerminalNode {
	return s.GetToken(ActionParserSCOL, 0)
}

func (s *Call_stmtContext) Call_receivers() ICall_receiversContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_receiversContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_receiversContext)
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
	p.EnterRule(localctx, 6, ActionParserRULE_call_stmt)
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ActionParserVARIABLE {
		{
			p.SetState(36)
			p.Call_receivers()
		}
		{
			p.SetState(37)
			p.Match(ActionParserEQ)
		}

	}
	{
		p.SetState(41)
		p.Call_body()
	}
	{
		p.SetState(42)
		p.Match(ActionParserSCOL)
	}

	return localctx
}

// ICall_receiversContext is an interface to support dynamic dispatch.
type ICall_receiversContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariable_name() []IVariable_nameContext
	Variable_name(i int) IVariable_nameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCall_receiversContext differentiates from other interfaces.
	IsCall_receiversContext()
}

type Call_receiversContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_receiversContext() *Call_receiversContext {
	var p = new(Call_receiversContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_call_receivers
	return p
}

func (*Call_receiversContext) IsCall_receiversContext() {}

func NewCall_receiversContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_receiversContext {
	var p = new(Call_receiversContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_call_receivers

	return p
}

func (s *Call_receiversContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_receiversContext) AllVariable_name() []IVariable_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariable_nameContext); ok {
			len++
		}
	}

	tst := make([]IVariable_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariable_nameContext); ok {
			tst[i] = t.(IVariable_nameContext)
			i++
		}
	}

	return tst
}

func (s *Call_receiversContext) Variable_name(i int) IVariable_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_nameContext); ok {
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

	return t.(IVariable_nameContext)
}

func (s *Call_receiversContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ActionParserCOMMA)
}

func (s *Call_receiversContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ActionParserCOMMA, i)
}

func (s *Call_receiversContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_receiversContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_receiversContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitCall_receivers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Call_receivers() (localctx ICall_receiversContext) {
	this := p
	_ = this

	localctx = NewCall_receiversContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ActionParserRULE_call_receivers)
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
		p.SetState(44)
		p.Variable_name()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ActionParserCOMMA {
		{
			p.SetState(45)
			p.Match(ActionParserCOMMA)
		}
		{
			p.SetState(46)
			p.Variable_name()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 10, ActionParserRULE_literal_value)
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
		p.SetState(52)
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
	p.EnterRule(localctx, 12, ActionParserRULE_variable_name)

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
		p.SetState(54)
		p.Match(ActionParserVARIABLE)
	}

	return localctx
}

// IBlock_variable_nameContext is an interface to support dynamic dispatch.
type IBlock_variable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_VARIABLE() antlr.TerminalNode

	// IsBlock_variable_nameContext differentiates from other interfaces.
	IsBlock_variable_nameContext()
}

type Block_variable_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_variable_nameContext() *Block_variable_nameContext {
	var p = new(Block_variable_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_block_variable_name
	return p
}

func (*Block_variable_nameContext) IsBlock_variable_nameContext() {}

func NewBlock_variable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_variable_nameContext {
	var p = new(Block_variable_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_block_variable_name

	return p
}

func (s *Block_variable_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_variable_nameContext) BLOCK_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ActionParserBLOCK_VARIABLE, 0)
}

func (s *Block_variable_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_variable_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_variable_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitBlock_variable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Block_variable_name() (localctx IBlock_variable_nameContext) {
	this := p
	_ = this

	localctx = NewBlock_variable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ActionParserRULE_block_variable_name)

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
		p.SetState(56)
		p.Match(ActionParserBLOCK_VARIABLE)
	}

	return localctx
}

// IFn_nameContext is an interface to support dynamic dispatch.
type IFn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	PERIOD() antlr.TerminalNode

	// IsFn_nameContext differentiates from other interfaces.
	IsFn_nameContext()
}

type Fn_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_nameContext() *Fn_nameContext {
	var p = new(Fn_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_fn_name
	return p
}

func (*Fn_nameContext) IsFn_nameContext() {}

func NewFn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_nameContext {
	var p = new(Fn_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_fn_name

	return p
}

func (s *Fn_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ActionParserIDENTIFIER)
}

func (s *Fn_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ActionParserIDENTIFIER, i)
}

func (s *Fn_nameContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(ActionParserPERIOD, 0)
}

func (s *Fn_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitFn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Fn_name() (localctx IFn_nameContext) {
	this := p
	_ = this

	localctx = NewFn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ActionParserRULE_fn_name)
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
		p.SetState(58)
		p.Match(ActionParserIDENTIFIER)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ActionParserPERIOD {
		{
			p.SetState(59)
			p.Match(ActionParserPERIOD)
		}
		{
			p.SetState(60)
			p.Match(ActionParserIDENTIFIER)
		}

	}

	return localctx
}

// ICall_bodyContext is an interface to support dynamic dispatch.
type ICall_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fn_name() IFn_nameContext
	L_PAREN() antlr.TerminalNode
	Fn_arg_list() IFn_arg_listContext
	R_PAREN() antlr.TerminalNode

	// IsCall_bodyContext differentiates from other interfaces.
	IsCall_bodyContext()
}

type Call_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_bodyContext() *Call_bodyContext {
	var p = new(Call_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_call_body
	return p
}

func (*Call_bodyContext) IsCall_bodyContext() {}

func NewCall_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_bodyContext {
	var p = new(Call_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_call_body

	return p
}

func (s *Call_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_bodyContext) Fn_name() IFn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_nameContext)
}

func (s *Call_bodyContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ActionParserL_PAREN, 0)
}

func (s *Call_bodyContext) Fn_arg_list() IFn_arg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_arg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_arg_listContext)
}

func (s *Call_bodyContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ActionParserR_PAREN, 0)
}

func (s *Call_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitCall_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Call_body() (localctx ICall_bodyContext) {
	this := p
	_ = this

	localctx = NewCall_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ActionParserRULE_call_body)

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
		p.SetState(63)
		p.Fn_name()
	}
	{
		p.SetState(64)
		p.Match(ActionParserL_PAREN)
	}
	{
		p.SetState(65)
		p.Fn_arg_list()
	}
	{
		p.SetState(66)
		p.Match(ActionParserR_PAREN)
	}

	return localctx
}

// IFn_argContext is an interface to support dynamic dispatch.
type IFn_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_value() ILiteral_valueContext
	Variable_name() IVariable_nameContext
	Block_variable_name() IBlock_variable_nameContext

	// IsFn_argContext differentiates from other interfaces.
	IsFn_argContext()
}

type Fn_argContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_argContext() *Fn_argContext {
	var p = new(Fn_argContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_fn_arg
	return p
}

func (*Fn_argContext) IsFn_argContext() {}

func NewFn_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_argContext {
	var p = new(Fn_argContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_fn_arg

	return p
}

func (s *Fn_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_argContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Fn_argContext) Variable_name() IVariable_nameContext {
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

func (s *Fn_argContext) Block_variable_name() IBlock_variable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_variable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_variable_nameContext)
}

func (s *Fn_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_argContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitFn_arg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Fn_arg() (localctx IFn_argContext) {
	this := p
	_ = this

	localctx = NewFn_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ActionParserRULE_fn_arg)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ActionParserUNSIGNED_NUMBER_LITERAL, ActionParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Literal_value()
		}

	case ActionParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Variable_name()
		}

	case ActionParserBLOCK_VARIABLE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Block_variable_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFn_arg_listContext is an interface to support dynamic dispatch.
type IFn_arg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFn_arg() []IFn_argContext
	Fn_arg(i int) IFn_argContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFn_arg_listContext differentiates from other interfaces.
	IsFn_arg_listContext()
}

type Fn_arg_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_arg_listContext() *Fn_arg_listContext {
	var p = new(Fn_arg_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_fn_arg_list
	return p
}

func (*Fn_arg_listContext) IsFn_arg_listContext() {}

func NewFn_arg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_arg_listContext {
	var p = new(Fn_arg_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_fn_arg_list

	return p
}

func (s *Fn_arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_arg_listContext) AllFn_arg() []IFn_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFn_argContext); ok {
			len++
		}
	}

	tst := make([]IFn_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFn_argContext); ok {
			tst[i] = t.(IFn_argContext)
			i++
		}
	}

	return tst
}

func (s *Fn_arg_listContext) Fn_arg(i int) IFn_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_argContext); ok {
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

	return t.(IFn_argContext)
}

func (s *Fn_arg_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ActionParserCOMMA)
}

func (s *Fn_arg_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ActionParserCOMMA, i)
}

func (s *Fn_arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_arg_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitFn_arg_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Fn_arg_list() (localctx IFn_arg_listContext) {
	this := p
	_ = this

	localctx = NewFn_arg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ActionParserRULE_fn_arg_list)
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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6029312) != 0 {
		{
			p.SetState(73)
			p.Fn_arg()
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ActionParserCOMMA {
		{
			p.SetState(76)
			p.Match(ActionParserCOMMA)
		}
		{
			p.SetState(77)
			p.Fn_arg()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
