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
		"", "';'", "'('", "')'", "','", "'$'", "'@'", "'='", "'.'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'~'", "'||'", "'<<'", "'>>'", "'&'", "'|'", "'=='",
		"'<'", "'<='", "'>'", "'>='", "'!='", "'<>'", "", "", "", "", "", "'not'",
		"'and'", "'or'",
	}
	staticData.symbolicNames = []string{
		"", "SCOL", "L_PAREN", "R_PAREN", "COMMA", "DOLLAR", "AT", "ASSIGN",
		"PERIOD", "PLUS", "MINUS", "STAR", "DIV", "MOD", "TILDE", "PIPE2", "LT2",
		"GT2", "AMP", "PIPE", "EQ", "LT", "LT_EQ", "GT", "GT_EQ", "SQL_NOT_EQ1",
		"SQL_NOT_EQ2", "SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_",
		"NOT_", "AND_", "OR_", "SQL_KEYWORDS", "SQL_STMT", "IDENTIFIER", "VARIABLE",
		"BLOCK_VARIABLE", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"statement", "literal_value", "action_name", "stmt", "sql_stmt", "call_stmt",
		"call_receivers", "call_body", "variable", "block_var", "extension_call_name",
		"fn_name", "sfn_name", "fn_arg_list", "fn_arg_expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 161, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 4, 0,
		32, 8, 0, 11, 0, 12, 0, 33, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 42,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 50, 8, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 5, 6, 58, 8, 6, 10, 6, 12, 6, 61, 9, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 3, 11, 78, 8, 11, 1, 12, 1, 12, 1, 13, 3, 13, 83, 8, 13, 1,
		13, 1, 13, 5, 13, 87, 8, 13, 10, 13, 12, 13, 90, 9, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 5, 14, 108, 8, 14, 10, 14, 12, 14, 111, 9, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 120, 8, 14, 10, 14,
		12, 14, 123, 9, 14, 1, 14, 3, 14, 126, 8, 14, 1, 14, 1, 14, 3, 14, 130,
		8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 156, 8, 14, 10, 14, 12, 14, 159, 9,
		14, 1, 14, 0, 1, 28, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 0, 7, 2, 0, 40, 40, 42, 42, 2, 0, 9, 10, 14, 14, 1, 0, 11, 13,
		1, 0, 9, 10, 1, 0, 16, 19, 1, 0, 21, 24, 3, 0, 7, 7, 20, 20, 25, 26, 171,
		0, 31, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0, 6, 41, 1, 0, 0,
		0, 8, 43, 1, 0, 0, 0, 10, 49, 1, 0, 0, 0, 12, 54, 1, 0, 0, 0, 14, 62, 1,
		0, 0, 0, 16, 67, 1, 0, 0, 0, 18, 69, 1, 0, 0, 0, 20, 71, 1, 0, 0, 0, 22,
		77, 1, 0, 0, 0, 24, 79, 1, 0, 0, 0, 26, 82, 1, 0, 0, 0, 28, 129, 1, 0,
		0, 0, 30, 32, 3, 6, 3, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 31,
		1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 1, 1, 0, 0, 0, 35, 36, 7, 0, 0, 0,
		36, 3, 1, 0, 0, 0, 37, 38, 5, 37, 0, 0, 38, 5, 1, 0, 0, 0, 39, 42, 3, 8,
		4, 0, 40, 42, 3, 10, 5, 0, 41, 39, 1, 0, 0, 0, 41, 40, 1, 0, 0, 0, 42,
		7, 1, 0, 0, 0, 43, 44, 5, 36, 0, 0, 44, 45, 5, 1, 0, 0, 45, 9, 1, 0, 0,
		0, 46, 47, 3, 12, 6, 0, 47, 48, 5, 7, 0, 0, 48, 50, 1, 0, 0, 0, 49, 46,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 3, 14, 7, 0,
		52, 53, 5, 1, 0, 0, 53, 11, 1, 0, 0, 0, 54, 59, 3, 16, 8, 0, 55, 56, 5,
		4, 0, 0, 56, 58, 3, 16, 8, 0, 57, 55, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59,
		57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 13, 1, 0, 0, 0, 61, 59, 1, 0, 0,
		0, 62, 63, 3, 22, 11, 0, 63, 64, 5, 2, 0, 0, 64, 65, 3, 26, 13, 0, 65,
		66, 5, 3, 0, 0, 66, 15, 1, 0, 0, 0, 67, 68, 5, 38, 0, 0, 68, 17, 1, 0,
		0, 0, 69, 70, 5, 39, 0, 0, 70, 19, 1, 0, 0, 0, 71, 72, 5, 37, 0, 0, 72,
		73, 5, 8, 0, 0, 73, 74, 5, 37, 0, 0, 74, 21, 1, 0, 0, 0, 75, 78, 3, 20,
		10, 0, 76, 78, 3, 4, 2, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78,
		23, 1, 0, 0, 0, 79, 80, 5, 37, 0, 0, 80, 25, 1, 0, 0, 0, 81, 83, 3, 28,
		14, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 88, 1, 0, 0, 0, 84,
		85, 5, 4, 0, 0, 85, 87, 3, 28, 14, 0, 86, 84, 1, 0, 0, 0, 87, 90, 1, 0,
		0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 27, 1, 0, 0, 0, 90, 88,
		1, 0, 0, 0, 91, 92, 6, 14, -1, 0, 92, 130, 3, 2, 1, 0, 93, 130, 3, 16,
		8, 0, 94, 130, 3, 18, 9, 0, 95, 96, 5, 2, 0, 0, 96, 97, 3, 28, 14, 0, 97,
		98, 5, 3, 0, 0, 98, 130, 1, 0, 0, 0, 99, 100, 7, 1, 0, 0, 100, 130, 3,
		28, 14, 12, 101, 102, 5, 32, 0, 0, 102, 130, 3, 28, 14, 5, 103, 104, 5,
		2, 0, 0, 104, 109, 3, 28, 14, 0, 105, 106, 5, 4, 0, 0, 106, 108, 3, 28,
		14, 0, 107, 105, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0,
		109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112,
		113, 5, 3, 0, 0, 113, 130, 1, 0, 0, 0, 114, 115, 3, 24, 12, 0, 115, 125,
		5, 2, 0, 0, 116, 121, 3, 28, 14, 0, 117, 118, 5, 4, 0, 0, 118, 120, 3,
		28, 14, 0, 119, 117, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0,
		0, 0, 121, 122, 1, 0, 0, 0, 122, 126, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0,
		124, 126, 5, 11, 0, 0, 125, 116, 1, 0, 0, 0, 125, 124, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 5, 3, 0, 0, 128, 130,
		1, 0, 0, 0, 129, 91, 1, 0, 0, 0, 129, 93, 1, 0, 0, 0, 129, 94, 1, 0, 0,
		0, 129, 95, 1, 0, 0, 0, 129, 99, 1, 0, 0, 0, 129, 101, 1, 0, 0, 0, 129,
		103, 1, 0, 0, 0, 129, 114, 1, 0, 0, 0, 130, 157, 1, 0, 0, 0, 131, 132,
		10, 11, 0, 0, 132, 133, 5, 15, 0, 0, 133, 156, 3, 28, 14, 12, 134, 135,
		10, 10, 0, 0, 135, 136, 7, 2, 0, 0, 136, 156, 3, 28, 14, 11, 137, 138,
		10, 9, 0, 0, 138, 139, 7, 3, 0, 0, 139, 156, 3, 28, 14, 10, 140, 141, 10,
		8, 0, 0, 141, 142, 7, 4, 0, 0, 142, 156, 3, 28, 14, 9, 143, 144, 10, 7,
		0, 0, 144, 145, 7, 5, 0, 0, 145, 156, 3, 28, 14, 8, 146, 147, 10, 6, 0,
		0, 147, 148, 7, 6, 0, 0, 148, 156, 3, 28, 14, 7, 149, 150, 10, 4, 0, 0,
		150, 151, 5, 33, 0, 0, 151, 156, 3, 28, 14, 5, 152, 153, 10, 3, 0, 0, 153,
		154, 5, 34, 0, 0, 154, 156, 3, 28, 14, 4, 155, 131, 1, 0, 0, 0, 155, 134,
		1, 0, 0, 0, 155, 137, 1, 0, 0, 0, 155, 140, 1, 0, 0, 0, 155, 143, 1, 0,
		0, 0, 155, 146, 1, 0, 0, 0, 155, 149, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		29, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 13, 33, 41, 49, 59, 77, 82, 88, 109,
		121, 125, 129, 155, 157,
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
	ActionParserASSIGN                  = 7
	ActionParserPERIOD                  = 8
	ActionParserPLUS                    = 9
	ActionParserMINUS                   = 10
	ActionParserSTAR                    = 11
	ActionParserDIV                     = 12
	ActionParserMOD                     = 13
	ActionParserTILDE                   = 14
	ActionParserPIPE2                   = 15
	ActionParserLT2                     = 16
	ActionParserGT2                     = 17
	ActionParserAMP                     = 18
	ActionParserPIPE                    = 19
	ActionParserEQ                      = 20
	ActionParserLT                      = 21
	ActionParserLT_EQ                   = 22
	ActionParserGT                      = 23
	ActionParserGT_EQ                   = 24
	ActionParserSQL_NOT_EQ1             = 25
	ActionParserSQL_NOT_EQ2             = 26
	ActionParserSELECT_                 = 27
	ActionParserINSERT_                 = 28
	ActionParserUPDATE_                 = 29
	ActionParserDELETE_                 = 30
	ActionParserWITH_                   = 31
	ActionParserNOT_                    = 32
	ActionParserAND_                    = 33
	ActionParserOR_                     = 34
	ActionParserSQL_KEYWORDS            = 35
	ActionParserSQL_STMT                = 36
	ActionParserIDENTIFIER              = 37
	ActionParserVARIABLE                = 38
	ActionParserBLOCK_VARIABLE          = 39
	ActionParserUNSIGNED_NUMBER_LITERAL = 40
	ActionParserSIGNED_NUMBER_LITERAL   = 41
	ActionParserSTRING_LITERAL          = 42
	ActionParserWS                      = 43
	ActionParserTERMINATOR              = 44
	ActionParserBLOCK_COMMENT           = 45
	ActionParserLINE_COMMENT            = 46
)

// ActionParser rules.
const (
	ActionParserRULE_statement           = 0
	ActionParserRULE_literal_value       = 1
	ActionParserRULE_action_name         = 2
	ActionParserRULE_stmt                = 3
	ActionParserRULE_sql_stmt            = 4
	ActionParserRULE_call_stmt           = 5
	ActionParserRULE_call_receivers      = 6
	ActionParserRULE_call_body           = 7
	ActionParserRULE_variable            = 8
	ActionParserRULE_block_var           = 9
	ActionParserRULE_extension_call_name = 10
	ActionParserRULE_fn_name             = 11
	ActionParserRULE_sfn_name            = 12
	ActionParserRULE_fn_arg_list         = 13
	ActionParserRULE_fn_arg_expr         = 14
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036337152) != 0) {
		{
			p.SetState(30)
			p.Stmt()
		}

		p.SetState(33)
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
	p.EnterRule(localctx, 2, ActionParserRULE_literal_value)
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
		p.SetState(35)
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
	p.EnterRule(localctx, 4, ActionParserRULE_action_name)

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
		p.SetState(37)
		p.Match(ActionParserIDENTIFIER)
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
	p.EnterRule(localctx, 6, ActionParserRULE_stmt)

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ActionParserSQL_STMT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Sql_stmt()
		}

	case ActionParserIDENTIFIER, ActionParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
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
	p.EnterRule(localctx, 8, ActionParserRULE_sql_stmt)

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
		p.Match(ActionParserSQL_STMT)
	}
	{
		p.SetState(44)
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
	ASSIGN() antlr.TerminalNode

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

func (s *Call_stmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ActionParserASSIGN, 0)
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
	p.EnterRule(localctx, 10, ActionParserRULE_call_stmt)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ActionParserVARIABLE {
		{
			p.SetState(46)
			p.Call_receivers()
		}
		{
			p.SetState(47)
			p.Match(ActionParserASSIGN)
		}

	}
	{
		p.SetState(51)
		p.Call_body()
	}
	{
		p.SetState(52)
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
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext
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

func (s *Call_receiversContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *Call_receiversContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
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
	p.EnterRule(localctx, 12, ActionParserRULE_call_receivers)
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
		p.SetState(54)
		p.Variable()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ActionParserCOMMA {
		{
			p.SetState(55)
			p.Match(ActionParserCOMMA)
		}
		{
			p.SetState(56)
			p.Variable()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 14, ActionParserRULE_call_body)

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
		p.SetState(62)
		p.Fn_name()
	}
	{
		p.SetState(63)
		p.Match(ActionParserL_PAREN)
	}
	{
		p.SetState(64)
		p.Fn_arg_list()
	}
	{
		p.SetState(65)
		p.Match(ActionParserR_PAREN)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIABLE() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(ActionParserVARIABLE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ActionParserRULE_variable)

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
		p.SetState(67)
		p.Match(ActionParserVARIABLE)
	}

	return localctx
}

// IBlock_varContext is an interface to support dynamic dispatch.
type IBlock_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_VARIABLE() antlr.TerminalNode

	// IsBlock_varContext differentiates from other interfaces.
	IsBlock_varContext()
}

type Block_varContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_varContext() *Block_varContext {
	var p = new(Block_varContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_block_var
	return p
}

func (*Block_varContext) IsBlock_varContext() {}

func NewBlock_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_varContext {
	var p = new(Block_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_block_var

	return p
}

func (s *Block_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_varContext) BLOCK_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ActionParserBLOCK_VARIABLE, 0)
}

func (s *Block_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitBlock_var(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Block_var() (localctx IBlock_varContext) {
	this := p
	_ = this

	localctx = NewBlock_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ActionParserRULE_block_var)

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
		p.SetState(69)
		p.Match(ActionParserBLOCK_VARIABLE)
	}

	return localctx
}

// IExtension_call_nameContext is an interface to support dynamic dispatch.
type IExtension_call_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	PERIOD() antlr.TerminalNode

	// IsExtension_call_nameContext differentiates from other interfaces.
	IsExtension_call_nameContext()
}

type Extension_call_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtension_call_nameContext() *Extension_call_nameContext {
	var p = new(Extension_call_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_extension_call_name
	return p
}

func (*Extension_call_nameContext) IsExtension_call_nameContext() {}

func NewExtension_call_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extension_call_nameContext {
	var p = new(Extension_call_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_extension_call_name

	return p
}

func (s *Extension_call_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Extension_call_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ActionParserIDENTIFIER)
}

func (s *Extension_call_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ActionParserIDENTIFIER, i)
}

func (s *Extension_call_nameContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(ActionParserPERIOD, 0)
}

func (s *Extension_call_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extension_call_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extension_call_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitExtension_call_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Extension_call_name() (localctx IExtension_call_nameContext) {
	this := p
	_ = this

	localctx = NewExtension_call_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ActionParserRULE_extension_call_name)

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
		p.SetState(71)
		p.Match(ActionParserIDENTIFIER)
	}
	{
		p.SetState(72)
		p.Match(ActionParserPERIOD)
	}
	{
		p.SetState(73)
		p.Match(ActionParserIDENTIFIER)
	}

	return localctx
}

// IFn_nameContext is an interface to support dynamic dispatch.
type IFn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Extension_call_name() IExtension_call_nameContext
	Action_name() IAction_nameContext

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

func (s *Fn_nameContext) Extension_call_name() IExtension_call_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtension_call_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtension_call_nameContext)
}

func (s *Fn_nameContext) Action_name() IAction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_nameContext)
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
	p.EnterRule(localctx, 22, ActionParserRULE_fn_name)

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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Extension_call_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Action_name()
		}

	}

	return localctx
}

// ISfn_nameContext is an interface to support dynamic dispatch.
type ISfn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsSfn_nameContext differentiates from other interfaces.
	IsSfn_nameContext()
}

type Sfn_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySfn_nameContext() *Sfn_nameContext {
	var p = new(Sfn_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_sfn_name
	return p
}

func (*Sfn_nameContext) IsSfn_nameContext() {}

func NewSfn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sfn_nameContext {
	var p = new(Sfn_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_sfn_name

	return p
}

func (s *Sfn_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Sfn_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ActionParserIDENTIFIER, 0)
}

func (s *Sfn_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sfn_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sfn_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitSfn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Sfn_name() (localctx ISfn_nameContext) {
	this := p
	_ = this

	localctx = NewSfn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ActionParserRULE_sfn_name)

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
		p.SetState(79)
		p.Match(ActionParserIDENTIFIER)
	}

	return localctx
}

// IFn_arg_listContext is an interface to support dynamic dispatch.
type IFn_arg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFn_arg_expr() []IFn_arg_exprContext
	Fn_arg_expr(i int) IFn_arg_exprContext
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

func (s *Fn_arg_listContext) AllFn_arg_expr() []IFn_arg_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
			len++
		}
	}

	tst := make([]IFn_arg_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFn_arg_exprContext); ok {
			tst[i] = t.(IFn_arg_exprContext)
			i++
		}
	}

	return tst
}

func (s *Fn_arg_listContext) Fn_arg_expr(i int) IFn_arg_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
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

	return t.(IFn_arg_exprContext)
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
	p.EnterRule(localctx, 26, ActionParserRULE_fn_arg_list)
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
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6463925798404) != 0 {
		{
			p.SetState(81)
			p.fn_arg_expr(0)
		}

	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ActionParserCOMMA {
		{
			p.SetState(84)
			p.Match(ActionParserCOMMA)
		}
		{
			p.SetState(85)
			p.fn_arg_expr(0)
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFn_arg_exprContext is an interface to support dynamic dispatch.
type IFn_arg_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetElevate_expr returns the elevate_expr rule contexts.
	GetElevate_expr() IFn_arg_exprContext

	// GetUnary_expr returns the unary_expr rule contexts.
	GetUnary_expr() IFn_arg_exprContext

	// Get_fn_arg_expr returns the _fn_arg_expr rule contexts.
	Get_fn_arg_expr() IFn_arg_exprContext

	// SetElevate_expr sets the elevate_expr rule contexts.
	SetElevate_expr(IFn_arg_exprContext)

	// SetUnary_expr sets the unary_expr rule contexts.
	SetUnary_expr(IFn_arg_exprContext)

	// Set_fn_arg_expr sets the _fn_arg_expr rule contexts.
	Set_fn_arg_expr(IFn_arg_exprContext)

	// GetExpr_list returns the expr_list rule context list.
	GetExpr_list() []IFn_arg_exprContext

	// SetExpr_list sets the expr_list rule context list.
	SetExpr_list([]IFn_arg_exprContext)

	// Getter signatures
	Literal_value() ILiteral_valueContext
	Variable() IVariableContext
	Block_var() IBlock_varContext
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllFn_arg_expr() []IFn_arg_exprContext
	Fn_arg_expr(i int) IFn_arg_exprContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	TILDE() antlr.TerminalNode
	NOT_() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Sfn_name() ISfn_nameContext
	STAR() antlr.TerminalNode
	PIPE2() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LT2() antlr.TerminalNode
	GT2() antlr.TerminalNode
	AMP() antlr.TerminalNode
	PIPE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LT_EQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	GT_EQ() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	EQ() antlr.TerminalNode
	SQL_NOT_EQ1() antlr.TerminalNode
	SQL_NOT_EQ2() antlr.TerminalNode
	AND_() antlr.TerminalNode
	OR_() antlr.TerminalNode

	// IsFn_arg_exprContext differentiates from other interfaces.
	IsFn_arg_exprContext()
}

type Fn_arg_exprContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	elevate_expr IFn_arg_exprContext
	unary_expr   IFn_arg_exprContext
	_fn_arg_expr IFn_arg_exprContext
	expr_list    []IFn_arg_exprContext
}

func NewEmptyFn_arg_exprContext() *Fn_arg_exprContext {
	var p = new(Fn_arg_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActionParserRULE_fn_arg_expr
	return p
}

func (*Fn_arg_exprContext) IsFn_arg_exprContext() {}

func NewFn_arg_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_arg_exprContext {
	var p = new(Fn_arg_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActionParserRULE_fn_arg_expr

	return p
}

func (s *Fn_arg_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_arg_exprContext) GetElevate_expr() IFn_arg_exprContext { return s.elevate_expr }

func (s *Fn_arg_exprContext) GetUnary_expr() IFn_arg_exprContext { return s.unary_expr }

func (s *Fn_arg_exprContext) Get_fn_arg_expr() IFn_arg_exprContext { return s._fn_arg_expr }

func (s *Fn_arg_exprContext) SetElevate_expr(v IFn_arg_exprContext) { s.elevate_expr = v }

func (s *Fn_arg_exprContext) SetUnary_expr(v IFn_arg_exprContext) { s.unary_expr = v }

func (s *Fn_arg_exprContext) Set_fn_arg_expr(v IFn_arg_exprContext) { s._fn_arg_expr = v }

func (s *Fn_arg_exprContext) GetExpr_list() []IFn_arg_exprContext { return s.expr_list }

func (s *Fn_arg_exprContext) SetExpr_list(v []IFn_arg_exprContext) { s.expr_list = v }

func (s *Fn_arg_exprContext) Literal_value() ILiteral_valueContext {
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

func (s *Fn_arg_exprContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Fn_arg_exprContext) Block_var() IBlock_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_varContext)
}

func (s *Fn_arg_exprContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ActionParserL_PAREN, 0)
}

func (s *Fn_arg_exprContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ActionParserR_PAREN, 0)
}

func (s *Fn_arg_exprContext) AllFn_arg_expr() []IFn_arg_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
			len++
		}
	}

	tst := make([]IFn_arg_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFn_arg_exprContext); ok {
			tst[i] = t.(IFn_arg_exprContext)
			i++
		}
	}

	return tst
}

func (s *Fn_arg_exprContext) Fn_arg_expr(i int) IFn_arg_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
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

	return t.(IFn_arg_exprContext)
}

func (s *Fn_arg_exprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ActionParserMINUS, 0)
}

func (s *Fn_arg_exprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ActionParserPLUS, 0)
}

func (s *Fn_arg_exprContext) TILDE() antlr.TerminalNode {
	return s.GetToken(ActionParserTILDE, 0)
}

func (s *Fn_arg_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(ActionParserNOT_, 0)
}

func (s *Fn_arg_exprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ActionParserCOMMA)
}

func (s *Fn_arg_exprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ActionParserCOMMA, i)
}

func (s *Fn_arg_exprContext) Sfn_name() ISfn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISfn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISfn_nameContext)
}

func (s *Fn_arg_exprContext) STAR() antlr.TerminalNode {
	return s.GetToken(ActionParserSTAR, 0)
}

func (s *Fn_arg_exprContext) PIPE2() antlr.TerminalNode {
	return s.GetToken(ActionParserPIPE2, 0)
}

func (s *Fn_arg_exprContext) DIV() antlr.TerminalNode {
	return s.GetToken(ActionParserDIV, 0)
}

func (s *Fn_arg_exprContext) MOD() antlr.TerminalNode {
	return s.GetToken(ActionParserMOD, 0)
}

func (s *Fn_arg_exprContext) LT2() antlr.TerminalNode {
	return s.GetToken(ActionParserLT2, 0)
}

func (s *Fn_arg_exprContext) GT2() antlr.TerminalNode {
	return s.GetToken(ActionParserGT2, 0)
}

func (s *Fn_arg_exprContext) AMP() antlr.TerminalNode {
	return s.GetToken(ActionParserAMP, 0)
}

func (s *Fn_arg_exprContext) PIPE() antlr.TerminalNode {
	return s.GetToken(ActionParserPIPE, 0)
}

func (s *Fn_arg_exprContext) LT() antlr.TerminalNode {
	return s.GetToken(ActionParserLT, 0)
}

func (s *Fn_arg_exprContext) LT_EQ() antlr.TerminalNode {
	return s.GetToken(ActionParserLT_EQ, 0)
}

func (s *Fn_arg_exprContext) GT() antlr.TerminalNode {
	return s.GetToken(ActionParserGT, 0)
}

func (s *Fn_arg_exprContext) GT_EQ() antlr.TerminalNode {
	return s.GetToken(ActionParserGT_EQ, 0)
}

func (s *Fn_arg_exprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ActionParserASSIGN, 0)
}

func (s *Fn_arg_exprContext) EQ() antlr.TerminalNode {
	return s.GetToken(ActionParserEQ, 0)
}

func (s *Fn_arg_exprContext) SQL_NOT_EQ1() antlr.TerminalNode {
	return s.GetToken(ActionParserSQL_NOT_EQ1, 0)
}

func (s *Fn_arg_exprContext) SQL_NOT_EQ2() antlr.TerminalNode {
	return s.GetToken(ActionParserSQL_NOT_EQ2, 0)
}

func (s *Fn_arg_exprContext) AND_() antlr.TerminalNode {
	return s.GetToken(ActionParserAND_, 0)
}

func (s *Fn_arg_exprContext) OR_() antlr.TerminalNode {
	return s.GetToken(ActionParserOR_, 0)
}

func (s *Fn_arg_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_arg_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_arg_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ActionParserVisitor:
		return t.VisitFn_arg_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ActionParser) Fn_arg_expr() (localctx IFn_arg_exprContext) {
	return p.fn_arg_expr(0)
}

func (p *ActionParser) fn_arg_expr(_p int) (localctx IFn_arg_exprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFn_arg_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFn_arg_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, ActionParserRULE_fn_arg_expr, _p)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(92)
			p.Literal_value()
		}

	case 2:
		{
			p.SetState(93)
			p.Variable()
		}

	case 3:
		{
			p.SetState(94)
			p.Block_var()
		}

	case 4:
		{
			p.SetState(95)
			p.Match(ActionParserL_PAREN)
		}
		{
			p.SetState(96)

			var _x = p.fn_arg_expr(0)

			localctx.(*Fn_arg_exprContext).elevate_expr = _x
		}
		{
			p.SetState(97)
			p.Match(ActionParserR_PAREN)
		}

	case 5:
		{
			p.SetState(99)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17920) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(100)

			var _x = p.fn_arg_expr(12)

			localctx.(*Fn_arg_exprContext).unary_expr = _x
		}

	case 6:
		{
			p.SetState(101)
			p.Match(ActionParserNOT_)
		}
		{
			p.SetState(102)

			var _x = p.fn_arg_expr(5)

			localctx.(*Fn_arg_exprContext).unary_expr = _x
		}

	case 7:
		{
			p.SetState(103)
			p.Match(ActionParserL_PAREN)
		}
		{
			p.SetState(104)

			var _x = p.fn_arg_expr(0)

			localctx.(*Fn_arg_exprContext)._fn_arg_expr = _x
		}
		localctx.(*Fn_arg_exprContext).expr_list = append(localctx.(*Fn_arg_exprContext).expr_list, localctx.(*Fn_arg_exprContext)._fn_arg_expr)
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ActionParserCOMMA {
			{
				p.SetState(105)
				p.Match(ActionParserCOMMA)
			}
			{
				p.SetState(106)

				var _x = p.fn_arg_expr(0)

				localctx.(*Fn_arg_exprContext)._fn_arg_expr = _x
			}
			localctx.(*Fn_arg_exprContext).expr_list = append(localctx.(*Fn_arg_exprContext).expr_list, localctx.(*Fn_arg_exprContext)._fn_arg_expr)

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			p.Match(ActionParserR_PAREN)
		}

	case 8:
		{
			p.SetState(114)
			p.Sfn_name()
		}
		{
			p.SetState(115)
			p.Match(ActionParserL_PAREN)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ActionParserL_PAREN, ActionParserPLUS, ActionParserMINUS, ActionParserTILDE, ActionParserNOT_, ActionParserIDENTIFIER, ActionParserVARIABLE, ActionParserBLOCK_VARIABLE, ActionParserUNSIGNED_NUMBER_LITERAL, ActionParserSTRING_LITERAL:
			{
				p.SetState(116)
				p.fn_arg_expr(0)
			}
			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ActionParserCOMMA {
				{
					p.SetState(117)
					p.Match(ActionParserCOMMA)
				}
				{
					p.SetState(118)
					p.fn_arg_expr(0)
				}

				p.SetState(123)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		case ActionParserSTAR:
			{
				p.SetState(124)
				p.Match(ActionParserSTAR)
			}

		case ActionParserR_PAREN:

		default:
		}
		{
			p.SetState(127)
			p.Match(ActionParserR_PAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(132)
					p.Match(ActionParserPIPE2)
				}
				{
					p.SetState(133)
					p.fn_arg_expr(12)
				}

			case 2:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(135)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(136)
					p.fn_arg_expr(11)
				}

			case 3:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(138)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ActionParserPLUS || _la == ActionParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(139)
					p.fn_arg_expr(10)
				}

			case 4:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(141)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&983040) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(142)
					p.fn_arg_expr(9)
				}

			case 5:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(144)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31457280) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(145)
					p.fn_arg_expr(8)
				}

			case 6:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(147)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&101712000) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(148)
					p.fn_arg_expr(7)
				}

			case 7:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(150)
					p.Match(ActionParserAND_)
				}
				{
					p.SetState(151)
					p.fn_arg_expr(5)
				}

			case 8:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ActionParserRULE_fn_arg_expr)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(153)
					p.Match(ActionParserOR_)
				}
				{
					p.SetState(154)
					p.fn_arg_expr(4)
				}

			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ActionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *Fn_arg_exprContext = nil
		if localctx != nil {
			t = localctx.(*Fn_arg_exprContext)
		}
		return p.Fn_arg_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ActionParser) Fn_arg_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
