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
		"", "';'", "'('", "')'", "','", "'$'", "'@'", "'='", "'.'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='", "'!='", "'<>'", "",
		"", "", "", "", "'not'", "'and'", "'or'",
	}
	staticData.symbolicNames = []string{
		"", "SCOL", "L_PAREN", "R_PAREN", "COMMA", "DOLLAR", "AT", "ASSIGN",
		"PERIOD", "PLUS", "MINUS", "STAR", "DIV", "MOD", "LT", "LT_EQ", "GT",
		"GT_EQ", "SQL_NOT_EQ1", "SQL_NOT_EQ2", "SELECT_", "INSERT_", "UPDATE_",
		"DELETE_", "WITH_", "NOT_", "AND_", "OR_", "SQL_KEYWORDS", "SQL_STMT",
		"IDENTIFIER", "VARIABLE", "BLOCK_VARIABLE", "UNSIGNED_NUMBER_LITERAL",
		"STRING_LITERAL", "WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"SCOL", "L_PAREN", "R_PAREN", "COMMA", "DOLLAR", "AT", "ASSIGN", "PERIOD",
		"PLUS", "MINUS", "STAR", "DIV", "MOD", "LT", "LT_EQ", "GT", "GT_EQ",
		"SQL_NOT_EQ1", "SQL_NOT_EQ2", "SELECT_", "INSERT_", "UPDATE_", "DELETE_",
		"WITH_", "NOT_", "AND_", "OR_", "SQL_KEYWORDS", "SQL_STMT", "IDENTIFIER",
		"VARIABLE", "BLOCK_VARIABLE", "UNSIGNED_NUMBER_LITERAL", "STRING_LITERAL",
		"WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT", "WSNL", "DIGIT",
		"DOUBLE_QUOTE_STRING_CHAR", "SINGLE_QUOTE_STRING_CHAR", "SINGLE_QUOTE_STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 271, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 179, 8, 27, 1,
		28, 1, 28, 4, 28, 183, 8, 28, 11, 28, 12, 28, 184, 1, 29, 1, 29, 5, 29,
		189, 8, 29, 10, 29, 12, 29, 192, 9, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 32, 4, 32, 201, 8, 32, 11, 32, 12, 32, 202, 1, 33, 1, 33,
		1, 34, 4, 34, 208, 8, 34, 11, 34, 12, 34, 209, 1, 34, 1, 34, 1, 35, 4,
		35, 215, 8, 35, 11, 35, 12, 35, 216, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 5, 36, 225, 8, 36, 10, 36, 12, 36, 228, 9, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 239, 8, 37, 10, 37,
		12, 37, 242, 9, 37, 1, 37, 1, 37, 1, 38, 4, 38, 247, 8, 38, 11, 38, 12,
		38, 248, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 3, 40, 256, 8, 40, 1, 41, 1,
		41, 1, 41, 3, 41, 261, 8, 41, 1, 42, 1, 42, 5, 42, 265, 8, 42, 10, 42,
		12, 42, 268, 9, 42, 1, 42, 1, 42, 1, 226, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 0, 79, 0, 81, 0, 83,
		0, 85, 0, 1, 0, 23, 2, 0, 83, 83, 115, 115, 2, 0, 69, 69, 101, 101, 2,
		0, 76, 76, 108, 108, 2, 0, 67, 67, 99, 99, 2, 0, 84, 84, 116, 116, 2, 0,
		73, 73, 105, 105, 2, 0, 78, 78, 110, 110, 2, 0, 82, 82, 114, 114, 2, 0,
		85, 85, 117, 117, 2, 0, 80, 80, 112, 112, 2, 0, 68, 68, 100, 100, 2, 0,
		65, 65, 97, 97, 2, 0, 87, 87, 119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 59,
		59, 125, 125, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32,
		32, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 4, 0, 10, 10, 13,
		13, 39, 39, 92, 92, 280, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 1, 87, 1, 0, 0, 0, 3, 89, 1, 0, 0, 0, 5, 91, 1, 0, 0, 0,
		7, 93, 1, 0, 0, 0, 9, 95, 1, 0, 0, 0, 11, 97, 1, 0, 0, 0, 13, 99, 1, 0,
		0, 0, 15, 101, 1, 0, 0, 0, 17, 103, 1, 0, 0, 0, 19, 105, 1, 0, 0, 0, 21,
		107, 1, 0, 0, 0, 23, 109, 1, 0, 0, 0, 25, 111, 1, 0, 0, 0, 27, 113, 1,
		0, 0, 0, 29, 115, 1, 0, 0, 0, 31, 118, 1, 0, 0, 0, 33, 120, 1, 0, 0, 0,
		35, 123, 1, 0, 0, 0, 37, 126, 1, 0, 0, 0, 39, 129, 1, 0, 0, 0, 41, 136,
		1, 0, 0, 0, 43, 143, 1, 0, 0, 0, 45, 150, 1, 0, 0, 0, 47, 157, 1, 0, 0,
		0, 49, 162, 1, 0, 0, 0, 51, 166, 1, 0, 0, 0, 53, 170, 1, 0, 0, 0, 55, 178,
		1, 0, 0, 0, 57, 180, 1, 0, 0, 0, 59, 186, 1, 0, 0, 0, 61, 193, 1, 0, 0,
		0, 63, 196, 1, 0, 0, 0, 65, 200, 1, 0, 0, 0, 67, 204, 1, 0, 0, 0, 69, 207,
		1, 0, 0, 0, 71, 214, 1, 0, 0, 0, 73, 220, 1, 0, 0, 0, 75, 234, 1, 0, 0,
		0, 77, 246, 1, 0, 0, 0, 79, 250, 1, 0, 0, 0, 81, 255, 1, 0, 0, 0, 83, 260,
		1, 0, 0, 0, 85, 262, 1, 0, 0, 0, 87, 88, 5, 59, 0, 0, 88, 2, 1, 0, 0, 0,
		89, 90, 5, 40, 0, 0, 90, 4, 1, 0, 0, 0, 91, 92, 5, 41, 0, 0, 92, 6, 1,
		0, 0, 0, 93, 94, 5, 44, 0, 0, 94, 8, 1, 0, 0, 0, 95, 96, 5, 36, 0, 0, 96,
		10, 1, 0, 0, 0, 97, 98, 5, 64, 0, 0, 98, 12, 1, 0, 0, 0, 99, 100, 5, 61,
		0, 0, 100, 14, 1, 0, 0, 0, 101, 102, 5, 46, 0, 0, 102, 16, 1, 0, 0, 0,
		103, 104, 5, 43, 0, 0, 104, 18, 1, 0, 0, 0, 105, 106, 5, 45, 0, 0, 106,
		20, 1, 0, 0, 0, 107, 108, 5, 42, 0, 0, 108, 22, 1, 0, 0, 0, 109, 110, 5,
		47, 0, 0, 110, 24, 1, 0, 0, 0, 111, 112, 5, 37, 0, 0, 112, 26, 1, 0, 0,
		0, 113, 114, 5, 60, 0, 0, 114, 28, 1, 0, 0, 0, 115, 116, 5, 60, 0, 0, 116,
		117, 5, 61, 0, 0, 117, 30, 1, 0, 0, 0, 118, 119, 5, 62, 0, 0, 119, 32,
		1, 0, 0, 0, 120, 121, 5, 62, 0, 0, 121, 122, 5, 61, 0, 0, 122, 34, 1, 0,
		0, 0, 123, 124, 5, 33, 0, 0, 124, 125, 5, 61, 0, 0, 125, 36, 1, 0, 0, 0,
		126, 127, 5, 60, 0, 0, 127, 128, 5, 62, 0, 0, 128, 38, 1, 0, 0, 0, 129,
		130, 7, 0, 0, 0, 130, 131, 7, 1, 0, 0, 131, 132, 7, 2, 0, 0, 132, 133,
		7, 1, 0, 0, 133, 134, 7, 3, 0, 0, 134, 135, 7, 4, 0, 0, 135, 40, 1, 0,
		0, 0, 136, 137, 7, 5, 0, 0, 137, 138, 7, 6, 0, 0, 138, 139, 7, 0, 0, 0,
		139, 140, 7, 1, 0, 0, 140, 141, 7, 7, 0, 0, 141, 142, 7, 4, 0, 0, 142,
		42, 1, 0, 0, 0, 143, 144, 7, 8, 0, 0, 144, 145, 7, 9, 0, 0, 145, 146, 7,
		10, 0, 0, 146, 147, 7, 11, 0, 0, 147, 148, 7, 4, 0, 0, 148, 149, 7, 1,
		0, 0, 149, 44, 1, 0, 0, 0, 150, 151, 7, 10, 0, 0, 151, 152, 7, 1, 0, 0,
		152, 153, 7, 2, 0, 0, 153, 154, 7, 1, 0, 0, 154, 155, 7, 4, 0, 0, 155,
		156, 7, 1, 0, 0, 156, 46, 1, 0, 0, 0, 157, 158, 7, 12, 0, 0, 158, 159,
		7, 5, 0, 0, 159, 160, 7, 4, 0, 0, 160, 161, 7, 13, 0, 0, 161, 48, 1, 0,
		0, 0, 162, 163, 5, 110, 0, 0, 163, 164, 5, 111, 0, 0, 164, 165, 5, 116,
		0, 0, 165, 50, 1, 0, 0, 0, 166, 167, 5, 97, 0, 0, 167, 168, 5, 110, 0,
		0, 168, 169, 5, 100, 0, 0, 169, 52, 1, 0, 0, 0, 170, 171, 5, 111, 0, 0,
		171, 172, 5, 114, 0, 0, 172, 54, 1, 0, 0, 0, 173, 179, 3, 39, 19, 0, 174,
		179, 3, 41, 20, 0, 175, 179, 3, 43, 21, 0, 176, 179, 3, 45, 22, 0, 177,
		179, 3, 47, 23, 0, 178, 173, 1, 0, 0, 0, 178, 174, 1, 0, 0, 0, 178, 175,
		1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 177, 1, 0, 0, 0, 179, 56, 1, 0,
		0, 0, 180, 182, 3, 55, 27, 0, 181, 183, 8, 14, 0, 0, 182, 181, 1, 0, 0,
		0, 183, 184, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185,
		58, 1, 0, 0, 0, 186, 190, 7, 15, 0, 0, 187, 189, 7, 16, 0, 0, 188, 187,
		1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 191, 60, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 194, 3, 9, 4, 0,
		194, 195, 3, 59, 29, 0, 195, 62, 1, 0, 0, 0, 196, 197, 3, 11, 5, 0, 197,
		198, 3, 59, 29, 0, 198, 64, 1, 0, 0, 0, 199, 201, 3, 79, 39, 0, 200, 199,
		1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0,
		0, 0, 203, 66, 1, 0, 0, 0, 204, 205, 3, 85, 42, 0, 205, 68, 1, 0, 0, 0,
		206, 208, 7, 17, 0, 0, 207, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212,
		6, 34, 0, 0, 212, 70, 1, 0, 0, 0, 213, 215, 7, 18, 0, 0, 214, 213, 1, 0,
		0, 0, 215, 216, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0,
		217, 218, 1, 0, 0, 0, 218, 219, 6, 35, 0, 0, 219, 72, 1, 0, 0, 0, 220,
		221, 5, 47, 0, 0, 221, 222, 5, 42, 0, 0, 222, 226, 1, 0, 0, 0, 223, 225,
		9, 0, 0, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 227, 1, 0,
		0, 0, 226, 224, 1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0,
		229, 230, 5, 42, 0, 0, 230, 231, 5, 47, 0, 0, 231, 232, 1, 0, 0, 0, 232,
		233, 6, 36, 0, 0, 233, 74, 1, 0, 0, 0, 234, 235, 5, 47, 0, 0, 235, 236,
		5, 47, 0, 0, 236, 240, 1, 0, 0, 0, 237, 239, 8, 18, 0, 0, 238, 237, 1,
		0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0,
		0, 241, 243, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 244, 6, 37, 0, 0, 244,
		76, 1, 0, 0, 0, 245, 247, 7, 19, 0, 0, 246, 245, 1, 0, 0, 0, 247, 248,
		1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 78, 1, 0,
		0, 0, 250, 251, 7, 20, 0, 0, 251, 80, 1, 0, 0, 0, 252, 256, 8, 21, 0, 0,
		253, 254, 5, 92, 0, 0, 254, 256, 9, 0, 0, 0, 255, 252, 1, 0, 0, 0, 255,
		253, 1, 0, 0, 0, 256, 82, 1, 0, 0, 0, 257, 261, 8, 22, 0, 0, 258, 259,
		5, 92, 0, 0, 259, 261, 9, 0, 0, 0, 260, 257, 1, 0, 0, 0, 260, 258, 1, 0,
		0, 0, 261, 84, 1, 0, 0, 0, 262, 266, 5, 39, 0, 0, 263, 265, 3, 83, 41,
		0, 264, 263, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266,
		267, 1, 0, 0, 0, 267, 269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 270,
		5, 39, 0, 0, 270, 86, 1, 0, 0, 0, 13, 0, 178, 184, 190, 202, 209, 216,
		226, 240, 248, 255, 260, 266, 1, 0, 1, 0,
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
	ActionLexerASSIGN                  = 7
	ActionLexerPERIOD                  = 8
	ActionLexerPLUS                    = 9
	ActionLexerMINUS                   = 10
	ActionLexerSTAR                    = 11
	ActionLexerDIV                     = 12
	ActionLexerMOD                     = 13
	ActionLexerLT                      = 14
	ActionLexerLT_EQ                   = 15
	ActionLexerGT                      = 16
	ActionLexerGT_EQ                   = 17
	ActionLexerSQL_NOT_EQ1             = 18
	ActionLexerSQL_NOT_EQ2             = 19
	ActionLexerSELECT_                 = 20
	ActionLexerINSERT_                 = 21
	ActionLexerUPDATE_                 = 22
	ActionLexerDELETE_                 = 23
	ActionLexerWITH_                   = 24
	ActionLexerNOT_                    = 25
	ActionLexerAND_                    = 26
	ActionLexerOR_                     = 27
	ActionLexerSQL_KEYWORDS            = 28
	ActionLexerSQL_STMT                = 29
	ActionLexerIDENTIFIER              = 30
	ActionLexerVARIABLE                = 31
	ActionLexerBLOCK_VARIABLE          = 32
	ActionLexerUNSIGNED_NUMBER_LITERAL = 33
	ActionLexerSTRING_LITERAL          = 34
	ActionLexerWS                      = 35
	ActionLexerTERMINATOR              = 36
	ActionLexerBLOCK_COMMENT           = 37
	ActionLexerLINE_COMMENT            = 38
)
