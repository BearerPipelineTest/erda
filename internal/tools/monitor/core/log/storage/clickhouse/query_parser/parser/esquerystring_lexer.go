// Code generated from EsQueryString.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 81, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 6, 10, 55, 10, 10, 13, 10, 14, 10, 56, 3,
	10, 3, 10, 3, 11, 3, 11, 6, 11, 63, 10, 11, 13, 11, 14, 11, 64, 3, 11,
	3, 11, 3, 12, 6, 12, 70, 10, 12, 13, 12, 14, 12, 71, 3, 13, 3, 13, 6, 13,
	76, 10, 13, 13, 13, 14, 13, 77, 3, 13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 3,
	2, 7, 5, 2, 11, 12, 15, 15, 34, 34, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124,
	7, 2, 47, 48, 50, 59, 67, 92, 97, 97, 99, 124, 8, 2, 11, 12, 15, 15, 34,
	34, 36, 36, 42, 43, 60, 60, 3, 2, 36, 36, 2, 84, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2,
	5, 31, 3, 2, 2, 2, 7, 35, 3, 2, 2, 2, 9, 39, 3, 2, 2, 2, 11, 43, 3, 2,
	2, 2, 13, 46, 3, 2, 2, 2, 15, 49, 3, 2, 2, 2, 17, 51, 3, 2, 2, 2, 19, 54,
	3, 2, 2, 2, 21, 60, 3, 2, 2, 2, 23, 69, 3, 2, 2, 2, 25, 73, 3, 2, 2, 2,
	27, 28, 7, 80, 2, 2, 28, 29, 7, 81, 2, 2, 29, 30, 7, 86, 2, 2, 30, 4, 3,
	2, 2, 2, 31, 32, 7, 112, 2, 2, 32, 33, 7, 113, 2, 2, 33, 34, 7, 118, 2,
	2, 34, 6, 3, 2, 2, 2, 35, 36, 7, 67, 2, 2, 36, 37, 7, 80, 2, 2, 37, 38,
	7, 70, 2, 2, 38, 8, 3, 2, 2, 2, 39, 40, 7, 99, 2, 2, 40, 41, 7, 112, 2,
	2, 41, 42, 7, 102, 2, 2, 42, 10, 3, 2, 2, 2, 43, 44, 7, 81, 2, 2, 44, 45,
	7, 84, 2, 2, 45, 12, 3, 2, 2, 2, 46, 47, 7, 113, 2, 2, 47, 48, 7, 116,
	2, 2, 48, 14, 3, 2, 2, 2, 49, 50, 7, 42, 2, 2, 50, 16, 3, 2, 2, 2, 51,
	52, 7, 43, 2, 2, 52, 18, 3, 2, 2, 2, 53, 55, 9, 2, 2, 2, 54, 53, 3, 2,
	2, 2, 55, 56, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 59, 8, 10, 2, 2, 59, 20, 3, 2, 2, 2, 60, 62, 9, 3, 2, 2,
	61, 63, 9, 4, 2, 2, 62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3,
	2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 7, 60, 2, 2, 67,
	22, 3, 2, 2, 2, 68, 70, 10, 5, 2, 2, 69, 68, 3, 2, 2, 2, 70, 71, 3, 2,
	2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 24, 3, 2, 2, 2, 73, 75,
	7, 36, 2, 2, 74, 76, 10, 6, 2, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2,
	2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80,
	7, 36, 2, 2, 80, 26, 3, 2, 2, 2, 7, 2, 56, 64, 71, 77, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'NOT'", "'not'", "'AND'", "'and'", "'OR'", "'or'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "GROUP_BEGIN", "GROUP_END", "WHITESPACE", "FIELD",
	"TERM", "PHRASE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "GROUP_BEGIN", "GROUP_END",
	"WHITESPACE", "FIELD", "TERM", "PHRASE",
}

type EsQueryStringLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewEsQueryStringLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *EsQueryStringLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEsQueryStringLexer(input antlr.CharStream) *EsQueryStringLexer {
	l := new(EsQueryStringLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "EsQueryString.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EsQueryStringLexer tokens.
const (
	EsQueryStringLexerT__0        = 1
	EsQueryStringLexerT__1        = 2
	EsQueryStringLexerT__2        = 3
	EsQueryStringLexerT__3        = 4
	EsQueryStringLexerT__4        = 5
	EsQueryStringLexerT__5        = 6
	EsQueryStringLexerGROUP_BEGIN = 7
	EsQueryStringLexerGROUP_END   = 8
	EsQueryStringLexerWHITESPACE  = 9
	EsQueryStringLexerFIELD       = 10
	EsQueryStringLexerTERM        = 11
	EsQueryStringLexerPHRASE      = 12
)