package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

// Test parsing V language example files
func TestVLangExamples(t *testing.T) {
	examples := []struct {
		name     string
		filename string
	}{
		{"Basic", "basic.v"},
		{"Structs", "structs.v"},
		{"Advanced", "advanced.v"},
	}

	for _, tt := range examples {
		t.Run(tt.name, func(t *testing.T) {
			// Read example file
			content, err := os.ReadFile(filepath.Join("..", "examples", tt.filename))
			assert.NoError(t, err)

			// Create input stream
			input := antlr.NewInputStream(string(content))

			// Create lexer
			lexer := parser.NewMyLangLexer(input)
			assert.NotNil(t, lexer)

			// Create token stream
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			// Create parser
			p := parser.NewMyLangParser(stream)
			assert.NotNil(t, p)

			// Create error listener
			errorListener := &ErrorListener{t: t}
			p.RemoveErrorListeners()
			p.AddErrorListener(errorListener)

			// Parse program
			program := p.Program()
			assert.NotNil(t, program)
			assert.False(t, errorListener.hasErrors, "Parsing failed for %s", tt.filename)
		})
	}
}

// ErrorListener implementation
type ErrorListener struct {
	*antlr.DefaultErrorListener
	t         *testing.T
	hasErrors bool
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.hasErrors = true
	l.t.Errorf("Syntax error at %d:%d - %s", line, column, msg)
}
