// Code generated from ActionLexer.g4 by ANTLR 4.12.0. DO NOT EDIT.

package actgrammar

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ActionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var actionlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func actionlexerLexerInit() {
	staticData := &actionlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"SCOL", "L_PAREN", "R_PAREN", "COMMA", "DOLLAR", "AT", "EQ", "PLUS",
		"PERIOD", "SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_", "SQL_KEYWORDS",
		"SQL_STMT", "IDENTIFIER", "VARIABLE", "UNSIGNED_NUMBER_LITERAL", "SIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
		"WSNL", "DIGIT", "DOUBLE_QUOTE_STRING_CHAR", "SINGLE_QUOTE_STRING_CHAR",
		"DOUBLE_QUOTE_STRING", "SINGLE_QUOTE_STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 228, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 120, 8, 14, 1, 15, 1, 15, 4, 15, 124, 8, 15, 11, 15, 12, 15, 125,
		1, 16, 1, 16, 5, 16, 130, 8, 16, 10, 16, 12, 16, 133, 9, 16, 1, 17, 1,
		17, 1, 17, 1, 18, 4, 18, 139, 8, 18, 11, 18, 12, 18, 140, 1, 19, 3, 19,
		144, 8, 19, 1, 19, 4, 19, 147, 8, 19, 11, 19, 12, 19, 148, 1, 20, 1, 20,
		3, 20, 153, 8, 20, 1, 21, 4, 21, 156, 8, 21, 11, 21, 12, 21, 157, 1, 21,
		1, 21, 1, 22, 4, 22, 163, 8, 22, 11, 22, 12, 22, 164, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 5, 23, 173, 8, 23, 10, 23, 12, 23, 176, 9, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 187,
		8, 24, 10, 24, 12, 24, 190, 9, 24, 1, 24, 1, 24, 1, 25, 4, 25, 195, 8,
		25, 11, 25, 12, 25, 196, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 3, 27, 204,
		8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 209, 8, 28, 1, 29, 1, 29, 5, 29, 213,
		8, 29, 10, 29, 12, 29, 216, 9, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 222,
		8, 30, 10, 30, 12, 30, 225, 9, 30, 1, 30, 1, 30, 1, 174, 0, 31, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 0, 53, 0, 55, 0, 57, 0, 59, 0, 61,
		0, 1, 0, 24, 2, 0, 83, 83, 115, 115, 2, 0, 69, 69, 101, 101, 2, 0, 76,
		76, 108, 108, 2, 0, 67, 67, 99, 99, 2, 0, 84, 84, 116, 116, 2, 0, 73, 73,
		105, 105, 2, 0, 78, 78, 110, 110, 2, 0, 82, 82, 114, 114, 2, 0, 85, 85,
		117, 117, 2, 0, 80, 80, 112, 112, 2, 0, 68, 68, 100, 100, 2, 0, 65, 65,
		97, 97, 2, 0, 87, 87, 119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 59, 59, 125,
		125, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0,
		43, 43, 45, 45, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10,
		13, 13, 32, 32, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 4,
		0, 10, 10, 13, 13, 39, 39, 92, 92, 240, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		1, 63, 1, 0, 0, 0, 3, 65, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 69, 1, 0, 0,
		0, 9, 71, 1, 0, 0, 0, 11, 73, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15, 77, 1,
		0, 0, 0, 17, 79, 1, 0, 0, 0, 19, 81, 1, 0, 0, 0, 21, 88, 1, 0, 0, 0, 23,
		95, 1, 0, 0, 0, 25, 102, 1, 0, 0, 0, 27, 109, 1, 0, 0, 0, 29, 119, 1, 0,
		0, 0, 31, 121, 1, 0, 0, 0, 33, 127, 1, 0, 0, 0, 35, 134, 1, 0, 0, 0, 37,
		138, 1, 0, 0, 0, 39, 143, 1, 0, 0, 0, 41, 152, 1, 0, 0, 0, 43, 155, 1,
		0, 0, 0, 45, 162, 1, 0, 0, 0, 47, 168, 1, 0, 0, 0, 49, 182, 1, 0, 0, 0,
		51, 194, 1, 0, 0, 0, 53, 198, 1, 0, 0, 0, 55, 203, 1, 0, 0, 0, 57, 208,
		1, 0, 0, 0, 59, 210, 1, 0, 0, 0, 61, 219, 1, 0, 0, 0, 63, 64, 5, 59, 0,
		0, 64, 2, 1, 0, 0, 0, 65, 66, 5, 40, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 5,
		41, 0, 0, 68, 6, 1, 0, 0, 0, 69, 70, 5, 44, 0, 0, 70, 8, 1, 0, 0, 0, 71,
		72, 5, 36, 0, 0, 72, 10, 1, 0, 0, 0, 73, 74, 5, 64, 0, 0, 74, 12, 1, 0,
		0, 0, 75, 76, 5, 61, 0, 0, 76, 14, 1, 0, 0, 0, 77, 78, 5, 43, 0, 0, 78,
		16, 1, 0, 0, 0, 79, 80, 5, 46, 0, 0, 80, 18, 1, 0, 0, 0, 81, 82, 7, 0,
		0, 0, 82, 83, 7, 1, 0, 0, 83, 84, 7, 2, 0, 0, 84, 85, 7, 1, 0, 0, 85, 86,
		7, 3, 0, 0, 86, 87, 7, 4, 0, 0, 87, 20, 1, 0, 0, 0, 88, 89, 7, 5, 0, 0,
		89, 90, 7, 6, 0, 0, 90, 91, 7, 0, 0, 0, 91, 92, 7, 1, 0, 0, 92, 93, 7,
		7, 0, 0, 93, 94, 7, 4, 0, 0, 94, 22, 1, 0, 0, 0, 95, 96, 7, 8, 0, 0, 96,
		97, 7, 9, 0, 0, 97, 98, 7, 10, 0, 0, 98, 99, 7, 11, 0, 0, 99, 100, 7, 4,
		0, 0, 100, 101, 7, 1, 0, 0, 101, 24, 1, 0, 0, 0, 102, 103, 7, 10, 0, 0,
		103, 104, 7, 1, 0, 0, 104, 105, 7, 2, 0, 0, 105, 106, 7, 1, 0, 0, 106,
		107, 7, 4, 0, 0, 107, 108, 7, 1, 0, 0, 108, 26, 1, 0, 0, 0, 109, 110, 7,
		12, 0, 0, 110, 111, 7, 5, 0, 0, 111, 112, 7, 4, 0, 0, 112, 113, 7, 13,
		0, 0, 113, 28, 1, 0, 0, 0, 114, 120, 3, 19, 9, 0, 115, 120, 3, 21, 10,
		0, 116, 120, 3, 23, 11, 0, 117, 120, 3, 25, 12, 0, 118, 120, 3, 27, 13,
		0, 119, 114, 1, 0, 0, 0, 119, 115, 1, 0, 0, 0, 119, 116, 1, 0, 0, 0, 119,
		117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 30, 1, 0, 0, 0, 121, 123, 3,
		29, 14, 0, 122, 124, 8, 14, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0,
		0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 32, 1, 0, 0, 0,
		127, 131, 7, 15, 0, 0, 128, 130, 7, 16, 0, 0, 129, 128, 1, 0, 0, 0, 130,
		133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 34, 1,
		0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 135, 3, 9, 4, 0, 135, 136, 3, 33, 16,
		0, 136, 36, 1, 0, 0, 0, 137, 139, 3, 53, 26, 0, 138, 137, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 38, 1,
		0, 0, 0, 142, 144, 7, 17, 0, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0,
		0, 144, 146, 1, 0, 0, 0, 145, 147, 3, 53, 26, 0, 146, 145, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		40, 1, 0, 0, 0, 150, 153, 3, 59, 29, 0, 151, 153, 3, 61, 30, 0, 152, 150,
		1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 42, 1, 0, 0, 0, 154, 156, 7, 18,
		0, 0, 155, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		157, 158, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 6, 21, 0, 0, 160,
		44, 1, 0, 0, 0, 161, 163, 7, 19, 0, 0, 162, 161, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0,
		0, 0, 166, 167, 6, 22, 0, 0, 167, 46, 1, 0, 0, 0, 168, 169, 5, 47, 0, 0,
		169, 170, 5, 42, 0, 0, 170, 174, 1, 0, 0, 0, 171, 173, 9, 0, 0, 0, 172,
		171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 174, 172,
		1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 178, 5, 42,
		0, 0, 178, 179, 5, 47, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 6, 23, 0,
		0, 181, 48, 1, 0, 0, 0, 182, 183, 5, 47, 0, 0, 183, 184, 5, 47, 0, 0, 184,
		188, 1, 0, 0, 0, 185, 187, 8, 19, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190,
		1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0,
		0, 0, 190, 188, 1, 0, 0, 0, 191, 192, 6, 24, 0, 0, 192, 50, 1, 0, 0, 0,
		193, 195, 7, 20, 0, 0, 194, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196,
		194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 52, 1, 0, 0, 0, 198, 199, 7,
		21, 0, 0, 199, 54, 1, 0, 0, 0, 200, 204, 8, 22, 0, 0, 201, 202, 5, 92,
		0, 0, 202, 204, 9, 0, 0, 0, 203, 200, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0,
		204, 56, 1, 0, 0, 0, 205, 209, 8, 23, 0, 0, 206, 207, 5, 92, 0, 0, 207,
		209, 9, 0, 0, 0, 208, 205, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 58, 1,
		0, 0, 0, 210, 214, 5, 34, 0, 0, 211, 213, 3, 55, 27, 0, 212, 211, 1, 0,
		0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 217, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 218, 5, 34, 0, 0, 218,
		60, 1, 0, 0, 0, 219, 223, 5, 39, 0, 0, 220, 222, 3, 57, 28, 0, 221, 220,
		1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0,
		0, 0, 224, 226, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 5, 39, 0, 0,
		227, 62, 1, 0, 0, 0, 17, 0, 119, 125, 131, 140, 143, 148, 152, 157, 164,
		174, 188, 196, 203, 208, 214, 223, 1, 0, 1, 0,
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

// ActionLexerInit initializes any static state used to implement ActionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewActionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ActionLexerInit() {
	staticData := &actionlexerLexerStaticData
	staticData.once.Do(actionlexerLexerInit)
}

// NewActionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewActionLexer(input antlr.CharStream) *ActionLexer {
	ActionLexerInit()
	l := new(ActionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &actionlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "ActionLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ActionLexer tokens.
const (
	ActionLexerSCOL                    = 1
	ActionLexerL_PAREN                 = 2
	ActionLexerR_PAREN                 = 3
	ActionLexerCOMMA                   = 4
	ActionLexerDOLLAR                  = 5
	ActionLexerAT                      = 6
	ActionLexerEQ                      = 7
	ActionLexerPLUS                    = 8
	ActionLexerPERIOD                  = 9
	ActionLexerSELECT_                 = 10
	ActionLexerINSERT_                 = 11
	ActionLexerUPDATE_                 = 12
	ActionLexerDELETE_                 = 13
	ActionLexerWITH_                   = 14
	ActionLexerSQL_KEYWORDS            = 15
	ActionLexerSQL_STMT                = 16
	ActionLexerIDENTIFIER              = 17
	ActionLexerVARIABLE                = 18
	ActionLexerUNSIGNED_NUMBER_LITERAL = 19
	ActionLexerSIGNED_NUMBER_LITERAL   = 20
	ActionLexerSTRING_LITERAL          = 21
	ActionLexerWS                      = 22
	ActionLexerTERMINATOR              = 23
	ActionLexerBLOCK_COMMENT           = 24
	ActionLexerLINE_COMMENT            = 25
)
