package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/mylang/interpreter/parser"
)

// CustomLexer extends the ANTLR-generated lexer with additional functionality
type CustomLexer struct {
	*parser.MyLangLexer
	Errors []string
}

// NewCustomLexer creates a new custom lexer
func NewCustomLexer(input antlr.CharStream) *CustomLexer {
	lexer := parser.NewMyLangLexer(input)
	return &CustomLexer{
		MyLangLexer: lexer,
		Errors:      make([]string, 0),
	}
}

// SyntaxError is called when a syntax error is encountered
func (l *CustomLexer) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, msg)
}

// ReportAmbiguity is called when an ambiguity is encountered
func (l *CustomLexer) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// Implement if needed
}

// ReportAttemptingFullContext is called when full context parsing is attempted
func (l *CustomLexer) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// Implement if needed
}

// ReportContextSensitivity is called when context sensitivity is encountered
func (l *CustomLexer) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	// Implement if needed
}

// CustomErrorListener implements ANTLR error listener interface
type CustomErrorListener struct {
	antlr.DefaultErrorListener
	Errors []string
}

// NewCustomErrorListener creates a new custom error listener
func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		Errors: make([]string, 0),
	}
}

// SyntaxError is called when a syntax error is encountered
func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, msg)
}
