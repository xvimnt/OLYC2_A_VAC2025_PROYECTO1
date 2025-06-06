// Code generated from ./grammar/MyLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser
import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type MyLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var MyLangLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func mylanglexerLexerInit() {
  staticData := &MyLangLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'module'", "'import'", "'pub'", "'fn'", "'struct'", "'interface'", 
    "'enum'", "'const'", "'mut'", "'type'", "'if'", "'else'", "'for'", "'in'", 
    "'match'", "'return'", "'defer'", "'true'", "'false'", "'none'", "'map'", 
    "'('", "')'", "'{'", "'}'", "'['", "']'", "'<'", "'>'", "'<='", "'>='", 
    "'=='", "'!='", "'&'", "'|'", "'^'", "", "'~'", "'?'", "':'", "';'", 
    "','", "'.'", "'='", "'=>'", "", "'+'", "'-'", "'*'", "'/'", "'%'", 
    "'<<'", "'>>'", "'>>>'", "'+='", "'-='", "'*='", "'/='", "'%='", "'&='", 
    "'|='", "'^='", "'<<='", "'>>='", "'>>>='", "'||='", "'&&='", "'||'", 
    "'&&'", "':='",
  }
  staticData.SymbolicNames = []string{
    "", "MODULE", "IMPORT", "PUB", "FN", "STRUCT", "INTERFACE", "ENUM", 
    "CONST", "MUT", "TYPE", "IF", "ELSE", "FOR", "IN", "MATCH", "RETURN", 
    "DEFER", "TRUE", "FALSE", "NONE", "MAP", "LPAREN", "RPAREN", "LBRACE", 
    "RBRACE", "LBRACK", "RBRACK", "LT", "GT", "LE", "GE", "EQ", "NE", "AND", 
    "OR", "XOR", "NOT", "TILDE", "QUESTION", "COLON", "SEMI", "COMMA", "DOT", 
    "ASSIGN", "ARROW", "EXCLAMATION", "PLUS", "MINUS", "MULT", "DIV", "MOD", 
    "LSHIFT", "RSHIFT", "URSHIFT", "PLUS_ASSIGN", "MINUS_ASSIGN", "MULT_ASSIGN", 
    "DIV_ASSIGN", "MOD_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN", 
    "LSHIFT_ASSIGN", "RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "LOGICAL_OR_ASSIGN", 
    "LOGICAL_AND_ASSIGN", "LOGICAL_OR", "LOGICAL_AND", "DECL_ASSIGN", "IDENTIFIER", 
    "INT", "FLOAT", "STRING", "CHAR", "WHITESPACE", "SINGLE_LINE_COMMENT", 
    "MULTI_LINE_COMMENT",
  }
  staticData.RuleNames = []string{
    "MODULE", "IMPORT", "PUB", "FN", "STRUCT", "INTERFACE", "ENUM", "CONST", 
    "MUT", "TYPE", "IF", "ELSE", "FOR", "IN", "MATCH", "RETURN", "DEFER", 
    "TRUE", "FALSE", "NONE", "MAP", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
    "LBRACK", "RBRACK", "LT", "GT", "LE", "GE", "EQ", "NE", "AND", "OR", 
    "XOR", "NOT", "TILDE", "QUESTION", "COLON", "SEMI", "COMMA", "DOT", 
    "ASSIGN", "ARROW", "EXCLAMATION", "PLUS", "MINUS", "MULT", "DIV", "MOD", 
    "LSHIFT", "RSHIFT", "URSHIFT", "PLUS_ASSIGN", "MINUS_ASSIGN", "MULT_ASSIGN", 
    "DIV_ASSIGN", "MOD_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN", 
    "LSHIFT_ASSIGN", "RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "LOGICAL_OR_ASSIGN", 
    "LOGICAL_AND_ASSIGN", "LOGICAL_OR", "LOGICAL_AND", "DECL_ASSIGN", "IDENTIFIER", 
    "INT", "FLOAT", "STRING", "CHAR", "WHITESPACE", "SINGLE_LINE_COMMENT", 
    "MULTI_LINE_COMMENT",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 78, 546, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 
	2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 
	31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 
	7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 
	41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 
	2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 
	52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 
	7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 
	62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 
	2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 
	73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 1, 0, 
	1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 
	1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 
	1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 
	10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 
	1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 
	15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 
	1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 
	18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 
	1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 
	26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 
	1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 
	35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 
	1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 
	45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 
	1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 
	53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 57, 
	1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 
	60, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 
	1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 
	66, 1, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 69, 
	1, 69, 1, 69, 1, 70, 1, 70, 5, 70, 401, 8, 70, 10, 70, 12, 70, 404, 9, 
	70, 1, 71, 4, 71, 407, 8, 71, 11, 71, 12, 71, 408, 1, 71, 1, 71, 1, 71, 
	1, 71, 4, 71, 415, 8, 71, 11, 71, 12, 71, 416, 1, 71, 1, 71, 1, 71, 1, 
	71, 4, 71, 423, 8, 71, 11, 71, 12, 71, 424, 1, 71, 1, 71, 1, 71, 1, 71, 
	4, 71, 431, 8, 71, 11, 71, 12, 71, 432, 1, 71, 4, 71, 436, 8, 71, 11, 71, 
	12, 71, 437, 1, 71, 1, 71, 4, 71, 442, 8, 71, 11, 71, 12, 71, 443, 3, 71, 
	446, 8, 71, 1, 72, 4, 72, 449, 8, 72, 11, 72, 12, 72, 450, 1, 72, 1, 72, 
	5, 72, 455, 8, 72, 10, 72, 12, 72, 458, 9, 72, 1, 72, 1, 72, 4, 72, 462, 
	8, 72, 11, 72, 12, 72, 463, 1, 72, 4, 72, 467, 8, 72, 11, 72, 12, 72, 468, 
	1, 72, 1, 72, 5, 72, 473, 8, 72, 10, 72, 12, 72, 476, 9, 72, 3, 72, 478, 
	8, 72, 1, 72, 1, 72, 3, 72, 482, 8, 72, 1, 72, 4, 72, 485, 8, 72, 11, 72, 
	12, 72, 486, 3, 72, 489, 8, 72, 1, 73, 1, 73, 1, 73, 1, 73, 5, 73, 495, 
	8, 73, 10, 73, 12, 73, 498, 9, 73, 1, 73, 1, 73, 1, 73, 5, 73, 503, 8, 
	73, 10, 73, 12, 73, 506, 9, 73, 1, 73, 3, 73, 509, 8, 73, 1, 74, 1, 74, 
	1, 74, 1, 74, 1, 75, 4, 75, 516, 8, 75, 11, 75, 12, 75, 517, 1, 75, 1, 
	75, 1, 76, 1, 76, 1, 76, 1, 76, 5, 76, 526, 8, 76, 10, 76, 12, 76, 529, 
	9, 76, 1, 76, 1, 76, 1, 77, 1, 77, 1, 77, 1, 77, 5, 77, 537, 8, 77, 10, 
	77, 12, 77, 540, 9, 77, 1, 77, 1, 77, 1, 77, 1, 77, 1, 77, 1, 538, 0, 78, 
	1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 
	23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 
	41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 
	59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 
	77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 
	95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 
	56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 
	64, 129, 65, 131, 66, 133, 67, 135, 68, 137, 69, 139, 70, 141, 71, 143, 
	72, 145, 73, 147, 74, 149, 75, 151, 76, 153, 77, 155, 78, 1, 0, 13, 3, 
	0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 
	48, 57, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 48, 55, 1, 0, 48, 49, 2, 0, 
	69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 34, 34, 92, 92, 1, 0, 96, 
	96, 1, 0, 39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 573, 
	0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 
	0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 
	0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 
	0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 
	0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 
	1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 
	47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 
	0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 
	0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 
	0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 
	0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 
	1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 
	93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 
	0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 
	0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 
	115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 
	0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 
	1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 
	0, 137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 
	0, 0, 0, 0, 145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 
	151, 1, 0, 0, 0, 0, 153, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 1, 157, 1, 0, 
	0, 0, 3, 164, 1, 0, 0, 0, 5, 171, 1, 0, 0, 0, 7, 175, 1, 0, 0, 0, 9, 178, 
	1, 0, 0, 0, 11, 185, 1, 0, 0, 0, 13, 195, 1, 0, 0, 0, 15, 200, 1, 0, 0, 
	0, 17, 206, 1, 0, 0, 0, 19, 210, 1, 0, 0, 0, 21, 215, 1, 0, 0, 0, 23, 218, 
	1, 0, 0, 0, 25, 223, 1, 0, 0, 0, 27, 227, 1, 0, 0, 0, 29, 230, 1, 0, 0, 
	0, 31, 236, 1, 0, 0, 0, 33, 243, 1, 0, 0, 0, 35, 249, 1, 0, 0, 0, 37, 254, 
	1, 0, 0, 0, 39, 260, 1, 0, 0, 0, 41, 265, 1, 0, 0, 0, 43, 269, 1, 0, 0, 
	0, 45, 271, 1, 0, 0, 0, 47, 273, 1, 0, 0, 0, 49, 275, 1, 0, 0, 0, 51, 277, 
	1, 0, 0, 0, 53, 279, 1, 0, 0, 0, 55, 281, 1, 0, 0, 0, 57, 283, 1, 0, 0, 
	0, 59, 285, 1, 0, 0, 0, 61, 288, 1, 0, 0, 0, 63, 291, 1, 0, 0, 0, 65, 294, 
	1, 0, 0, 0, 67, 297, 1, 0, 0, 0, 69, 299, 1, 0, 0, 0, 71, 301, 1, 0, 0, 
	0, 73, 303, 1, 0, 0, 0, 75, 305, 1, 0, 0, 0, 77, 307, 1, 0, 0, 0, 79, 309, 
	1, 0, 0, 0, 81, 311, 1, 0, 0, 0, 83, 313, 1, 0, 0, 0, 85, 315, 1, 0, 0, 
	0, 87, 317, 1, 0, 0, 0, 89, 319, 1, 0, 0, 0, 91, 322, 1, 0, 0, 0, 93, 324, 
	1, 0, 0, 0, 95, 326, 1, 0, 0, 0, 97, 328, 1, 0, 0, 0, 99, 330, 1, 0, 0, 
	0, 101, 332, 1, 0, 0, 0, 103, 334, 1, 0, 0, 0, 105, 337, 1, 0, 0, 0, 107, 
	340, 1, 0, 0, 0, 109, 344, 1, 0, 0, 0, 111, 347, 1, 0, 0, 0, 113, 350, 
	1, 0, 0, 0, 115, 353, 1, 0, 0, 0, 117, 356, 1, 0, 0, 0, 119, 359, 1, 0, 
	0, 0, 121, 362, 1, 0, 0, 0, 123, 365, 1, 0, 0, 0, 125, 368, 1, 0, 0, 0, 
	127, 372, 1, 0, 0, 0, 129, 376, 1, 0, 0, 0, 131, 381, 1, 0, 0, 0, 133, 
	385, 1, 0, 0, 0, 135, 389, 1, 0, 0, 0, 137, 392, 1, 0, 0, 0, 139, 395, 
	1, 0, 0, 0, 141, 398, 1, 0, 0, 0, 143, 445, 1, 0, 0, 0, 145, 488, 1, 0, 
	0, 0, 147, 508, 1, 0, 0, 0, 149, 510, 1, 0, 0, 0, 151, 515, 1, 0, 0, 0, 
	153, 521, 1, 0, 0, 0, 155, 532, 1, 0, 0, 0, 157, 158, 5, 109, 0, 0, 158, 
	159, 5, 111, 0, 0, 159, 160, 5, 100, 0, 0, 160, 161, 5, 117, 0, 0, 161, 
	162, 5, 108, 0, 0, 162, 163, 5, 101, 0, 0, 163, 2, 1, 0, 0, 0, 164, 165, 
	5, 105, 0, 0, 165, 166, 5, 109, 0, 0, 166, 167, 5, 112, 0, 0, 167, 168, 
	5, 111, 0, 0, 168, 169, 5, 114, 0, 0, 169, 170, 5, 116, 0, 0, 170, 4, 1, 
	0, 0, 0, 171, 172, 5, 112, 0, 0, 172, 173, 5, 117, 0, 0, 173, 174, 5, 98, 
	0, 0, 174, 6, 1, 0, 0, 0, 175, 176, 5, 102, 0, 0, 176, 177, 5, 110, 0, 
	0, 177, 8, 1, 0, 0, 0, 178, 179, 5, 115, 0, 0, 179, 180, 5, 116, 0, 0, 
	180, 181, 5, 114, 0, 0, 181, 182, 5, 117, 0, 0, 182, 183, 5, 99, 0, 0, 
	183, 184, 5, 116, 0, 0, 184, 10, 1, 0, 0, 0, 185, 186, 5, 105, 0, 0, 186, 
	187, 5, 110, 0, 0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 101, 0, 0, 189, 
	190, 5, 114, 0, 0, 190, 191, 5, 102, 0, 0, 191, 192, 5, 97, 0, 0, 192, 
	193, 5, 99, 0, 0, 193, 194, 5, 101, 0, 0, 194, 12, 1, 0, 0, 0, 195, 196, 
	5, 101, 0, 0, 196, 197, 5, 110, 0, 0, 197, 198, 5, 117, 0, 0, 198, 199, 
	5, 109, 0, 0, 199, 14, 1, 0, 0, 0, 200, 201, 5, 99, 0, 0, 201, 202, 5, 
	111, 0, 0, 202, 203, 5, 110, 0, 0, 203, 204, 5, 115, 0, 0, 204, 205, 5, 
	116, 0, 0, 205, 16, 1, 0, 0, 0, 206, 207, 5, 109, 0, 0, 207, 208, 5, 117, 
	0, 0, 208, 209, 5, 116, 0, 0, 209, 18, 1, 0, 0, 0, 210, 211, 5, 116, 0, 
	0, 211, 212, 5, 121, 0, 0, 212, 213, 5, 112, 0, 0, 213, 214, 5, 101, 0, 
	0, 214, 20, 1, 0, 0, 0, 215, 216, 5, 105, 0, 0, 216, 217, 5, 102, 0, 0, 
	217, 22, 1, 0, 0, 0, 218, 219, 5, 101, 0, 0, 219, 220, 5, 108, 0, 0, 220, 
	221, 5, 115, 0, 0, 221, 222, 5, 101, 0, 0, 222, 24, 1, 0, 0, 0, 223, 224, 
	5, 102, 0, 0, 224, 225, 5, 111, 0, 0, 225, 226, 5, 114, 0, 0, 226, 26, 
	1, 0, 0, 0, 227, 228, 5, 105, 0, 0, 228, 229, 5, 110, 0, 0, 229, 28, 1, 
	0, 0, 0, 230, 231, 5, 109, 0, 0, 231, 232, 5, 97, 0, 0, 232, 233, 5, 116, 
	0, 0, 233, 234, 5, 99, 0, 0, 234, 235, 5, 104, 0, 0, 235, 30, 1, 0, 0, 
	0, 236, 237, 5, 114, 0, 0, 237, 238, 5, 101, 0, 0, 238, 239, 5, 116, 0, 
	0, 239, 240, 5, 117, 0, 0, 240, 241, 5, 114, 0, 0, 241, 242, 5, 110, 0, 
	0, 242, 32, 1, 0, 0, 0, 243, 244, 5, 100, 0, 0, 244, 245, 5, 101, 0, 0, 
	245, 246, 5, 102, 0, 0, 246, 247, 5, 101, 0, 0, 247, 248, 5, 114, 0, 0, 
	248, 34, 1, 0, 0, 0, 249, 250, 5, 116, 0, 0, 250, 251, 5, 114, 0, 0, 251, 
	252, 5, 117, 0, 0, 252, 253, 5, 101, 0, 0, 253, 36, 1, 0, 0, 0, 254, 255, 
	5, 102, 0, 0, 255, 256, 5, 97, 0, 0, 256, 257, 5, 108, 0, 0, 257, 258, 
	5, 115, 0, 0, 258, 259, 5, 101, 0, 0, 259, 38, 1, 0, 0, 0, 260, 261, 5, 
	110, 0, 0, 261, 262, 5, 111, 0, 0, 262, 263, 5, 110, 0, 0, 263, 264, 5, 
	101, 0, 0, 264, 40, 1, 0, 0, 0, 265, 266, 5, 109, 0, 0, 266, 267, 5, 97, 
	0, 0, 267, 268, 5, 112, 0, 0, 268, 42, 1, 0, 0, 0, 269, 270, 5, 40, 0, 
	0, 270, 44, 1, 0, 0, 0, 271, 272, 5, 41, 0, 0, 272, 46, 1, 0, 0, 0, 273, 
	274, 5, 123, 0, 0, 274, 48, 1, 0, 0, 0, 275, 276, 5, 125, 0, 0, 276, 50, 
	1, 0, 0, 0, 277, 278, 5, 91, 0, 0, 278, 52, 1, 0, 0, 0, 279, 280, 5, 93, 
	0, 0, 280, 54, 1, 0, 0, 0, 281, 282, 5, 60, 0, 0, 282, 56, 1, 0, 0, 0, 
	283, 284, 5, 62, 0, 0, 284, 58, 1, 0, 0, 0, 285, 286, 5, 60, 0, 0, 286, 
	287, 5, 61, 0, 0, 287, 60, 1, 0, 0, 0, 288, 289, 5, 62, 0, 0, 289, 290, 
	5, 61, 0, 0, 290, 62, 1, 0, 0, 0, 291, 292, 5, 61, 0, 0, 292, 293, 5, 61, 
	0, 0, 293, 64, 1, 0, 0, 0, 294, 295, 5, 33, 0, 0, 295, 296, 5, 61, 0, 0, 
	296, 66, 1, 0, 0, 0, 297, 298, 5, 38, 0, 0, 298, 68, 1, 0, 0, 0, 299, 300, 
	5, 124, 0, 0, 300, 70, 1, 0, 0, 0, 301, 302, 5, 94, 0, 0, 302, 72, 1, 0, 
	0, 0, 303, 304, 5, 33, 0, 0, 304, 74, 1, 0, 0, 0, 305, 306, 5, 126, 0, 
	0, 306, 76, 1, 0, 0, 0, 307, 308, 5, 63, 0, 0, 308, 78, 1, 0, 0, 0, 309, 
	310, 5, 58, 0, 0, 310, 80, 1, 0, 0, 0, 311, 312, 5, 59, 0, 0, 312, 82, 
	1, 0, 0, 0, 313, 314, 5, 44, 0, 0, 314, 84, 1, 0, 0, 0, 315, 316, 5, 46, 
	0, 0, 316, 86, 1, 0, 0, 0, 317, 318, 5, 61, 0, 0, 318, 88, 1, 0, 0, 0, 
	319, 320, 5, 61, 0, 0, 320, 321, 5, 62, 0, 0, 321, 90, 1, 0, 0, 0, 322, 
	323, 5, 33, 0, 0, 323, 92, 1, 0, 0, 0, 324, 325, 5, 43, 0, 0, 325, 94, 
	1, 0, 0, 0, 326, 327, 5, 45, 0, 0, 327, 96, 1, 0, 0, 0, 328, 329, 5, 42, 
	0, 0, 329, 98, 1, 0, 0, 0, 330, 331, 5, 47, 0, 0, 331, 100, 1, 0, 0, 0, 
	332, 333, 5, 37, 0, 0, 333, 102, 1, 0, 0, 0, 334, 335, 5, 60, 0, 0, 335, 
	336, 5, 60, 0, 0, 336, 104, 1, 0, 0, 0, 337, 338, 5, 62, 0, 0, 338, 339, 
	5, 62, 0, 0, 339, 106, 1, 0, 0, 0, 340, 341, 5, 62, 0, 0, 341, 342, 5, 
	62, 0, 0, 342, 343, 5, 62, 0, 0, 343, 108, 1, 0, 0, 0, 344, 345, 5, 43, 
	0, 0, 345, 346, 5, 61, 0, 0, 346, 110, 1, 0, 0, 0, 347, 348, 5, 45, 0, 
	0, 348, 349, 5, 61, 0, 0, 349, 112, 1, 0, 0, 0, 350, 351, 5, 42, 0, 0, 
	351, 352, 5, 61, 0, 0, 352, 114, 1, 0, 0, 0, 353, 354, 5, 47, 0, 0, 354, 
	355, 5, 61, 0, 0, 355, 116, 1, 0, 0, 0, 356, 357, 5, 37, 0, 0, 357, 358, 
	5, 61, 0, 0, 358, 118, 1, 0, 0, 0, 359, 360, 5, 38, 0, 0, 360, 361, 5, 
	61, 0, 0, 361, 120, 1, 0, 0, 0, 362, 363, 5, 124, 0, 0, 363, 364, 5, 61, 
	0, 0, 364, 122, 1, 0, 0, 0, 365, 366, 5, 94, 0, 0, 366, 367, 5, 61, 0, 
	0, 367, 124, 1, 0, 0, 0, 368, 369, 5, 60, 0, 0, 369, 370, 5, 60, 0, 0, 
	370, 371, 5, 61, 0, 0, 371, 126, 1, 0, 0, 0, 372, 373, 5, 62, 0, 0, 373, 
	374, 5, 62, 0, 0, 374, 375, 5, 61, 0, 0, 375, 128, 1, 0, 0, 0, 376, 377, 
	5, 62, 0, 0, 377, 378, 5, 62, 0, 0, 378, 379, 5, 62, 0, 0, 379, 380, 5, 
	61, 0, 0, 380, 130, 1, 0, 0, 0, 381, 382, 5, 124, 0, 0, 382, 383, 5, 124, 
	0, 0, 383, 384, 5, 61, 0, 0, 384, 132, 1, 0, 0, 0, 385, 386, 5, 38, 0, 
	0, 386, 387, 5, 38, 0, 0, 387, 388, 5, 61, 0, 0, 388, 134, 1, 0, 0, 0, 
	389, 390, 5, 124, 0, 0, 390, 391, 5, 124, 0, 0, 391, 136, 1, 0, 0, 0, 392, 
	393, 5, 38, 0, 0, 393, 394, 5, 38, 0, 0, 394, 138, 1, 0, 0, 0, 395, 396, 
	5, 58, 0, 0, 396, 397, 5, 61, 0, 0, 397, 140, 1, 0, 0, 0, 398, 402, 7, 
	0, 0, 0, 399, 401, 7, 1, 0, 0, 400, 399, 1, 0, 0, 0, 401, 404, 1, 0, 0, 
	0, 402, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 142, 1, 0, 0, 0, 404, 
	402, 1, 0, 0, 0, 405, 407, 7, 2, 0, 0, 406, 405, 1, 0, 0, 0, 407, 408, 
	1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 446, 1, 0, 
	0, 0, 410, 411, 5, 48, 0, 0, 411, 412, 5, 120, 0, 0, 412, 414, 1, 0, 0, 
	0, 413, 415, 7, 3, 0, 0, 414, 413, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 
	414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 446, 1, 0, 0, 0, 418, 419, 
	5, 48, 0, 0, 419, 420, 5, 111, 0, 0, 420, 422, 1, 0, 0, 0, 421, 423, 7, 
	4, 0, 0, 422, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 422, 1, 0, 0, 
	0, 424, 425, 1, 0, 0, 0, 425, 446, 1, 0, 0, 0, 426, 427, 5, 48, 0, 0, 427, 
	428, 5, 98, 0, 0, 428, 430, 1, 0, 0, 0, 429, 431, 7, 5, 0, 0, 430, 429, 
	1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 432, 433, 1, 0, 
	0, 0, 433, 446, 1, 0, 0, 0, 434, 436, 7, 2, 0, 0, 435, 434, 1, 0, 0, 0, 
	436, 437, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 
	439, 1, 0, 0, 0, 439, 441, 5, 95, 0, 0, 440, 442, 7, 2, 0, 0, 441, 440, 
	1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 
	0, 0, 444, 446, 1, 0, 0, 0, 445, 406, 1, 0, 0, 0, 445, 410, 1, 0, 0, 0, 
	445, 418, 1, 0, 0, 0, 445, 426, 1, 0, 0, 0, 445, 435, 1, 0, 0, 0, 446, 
	144, 1, 0, 0, 0, 447, 449, 7, 2, 0, 0, 448, 447, 1, 0, 0, 0, 449, 450, 
	1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 452, 1, 0, 
	0, 0, 452, 456, 5, 46, 0, 0, 453, 455, 7, 2, 0, 0, 454, 453, 1, 0, 0, 0, 
	455, 458, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 
	489, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 459, 461, 5, 46, 0, 0, 460, 462, 
	7, 2, 0, 0, 461, 460, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 461, 1, 0, 
	0, 0, 463, 464, 1, 0, 0, 0, 464, 489, 1, 0, 0, 0, 465, 467, 7, 2, 0, 0, 
	466, 465, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0, 468, 
	469, 1, 0, 0, 0, 469, 477, 1, 0, 0, 0, 470, 474, 5, 46, 0, 0, 471, 473, 
	7, 2, 0, 0, 472, 471, 1, 0, 0, 0, 473, 476, 1, 0, 0, 0, 474, 472, 1, 0, 
	0, 0, 474, 475, 1, 0, 0, 0, 475, 478, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 
	477, 470, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479, 
	481, 7, 6, 0, 0, 480, 482, 7, 7, 0, 0, 481, 480, 1, 0, 0, 0, 481, 482, 
	1, 0, 0, 0, 482, 484, 1, 0, 0, 0, 483, 485, 7, 2, 0, 0, 484, 483, 1, 0, 
	0, 0, 485, 486, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 
	487, 489, 1, 0, 0, 0, 488, 448, 1, 0, 0, 0, 488, 459, 1, 0, 0, 0, 488, 
	466, 1, 0, 0, 0, 489, 146, 1, 0, 0, 0, 490, 496, 5, 34, 0, 0, 491, 495, 
	8, 8, 0, 0, 492, 493, 5, 92, 0, 0, 493, 495, 9, 0, 0, 0, 494, 491, 1, 0, 
	0, 0, 494, 492, 1, 0, 0, 0, 495, 498, 1, 0, 0, 0, 496, 494, 1, 0, 0, 0, 
	496, 497, 1, 0, 0, 0, 497, 499, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0, 499, 
	509, 5, 34, 0, 0, 500, 504, 5, 96, 0, 0, 501, 503, 8, 9, 0, 0, 502, 501, 
	1, 0, 0, 0, 503, 506, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 504, 505, 1, 0, 
	0, 0, 505, 507, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 507, 509, 5, 96, 0, 0, 
	508, 490, 1, 0, 0, 0, 508, 500, 1, 0, 0, 0, 509, 148, 1, 0, 0, 0, 510, 
	511, 5, 39, 0, 0, 511, 512, 8, 10, 0, 0, 512, 513, 5, 39, 0, 0, 513, 150, 
	1, 0, 0, 0, 514, 516, 7, 11, 0, 0, 515, 514, 1, 0, 0, 0, 516, 517, 1, 0, 
	0, 0, 517, 515, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 
	519, 520, 6, 75, 0, 0, 520, 152, 1, 0, 0, 0, 521, 522, 5, 47, 0, 0, 522, 
	523, 5, 47, 0, 0, 523, 527, 1, 0, 0, 0, 524, 526, 8, 12, 0, 0, 525, 524, 
	1, 0, 0, 0, 526, 529, 1, 0, 0, 0, 527, 525, 1, 0, 0, 0, 527, 528, 1, 0, 
	0, 0, 528, 530, 1, 0, 0, 0, 529, 527, 1, 0, 0, 0, 530, 531, 6, 76, 0, 0, 
	531, 154, 1, 0, 0, 0, 532, 533, 5, 47, 0, 0, 533, 534, 5, 42, 0, 0, 534, 
	538, 1, 0, 0, 0, 535, 537, 9, 0, 0, 0, 536, 535, 1, 0, 0, 0, 537, 540, 
	1, 0, 0, 0, 538, 539, 1, 0, 0, 0, 538, 536, 1, 0, 0, 0, 539, 541, 1, 0, 
	0, 0, 540, 538, 1, 0, 0, 0, 541, 542, 5, 42, 0, 0, 542, 543, 5, 47, 0, 
	0, 543, 544, 1, 0, 0, 0, 544, 545, 6, 77, 0, 0, 545, 156, 1, 0, 0, 0, 25, 
	0, 402, 408, 416, 424, 432, 437, 443, 445, 450, 456, 463, 468, 474, 477, 
	481, 486, 488, 494, 496, 504, 508, 517, 527, 538, 1, 6, 0, 0,
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

// MyLangLexerInit initializes any static state used to implement MyLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMyLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyLangLexerInit() {
  staticData := &MyLangLexerLexerStaticData
  staticData.once.Do(mylanglexerLexerInit)
}

// NewMyLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMyLangLexer(input antlr.CharStream) *MyLangLexer {
  MyLangLexerInit()
	l := new(MyLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &MyLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MyLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MyLangLexer tokens.
const (
	MyLangLexerMODULE = 1
	MyLangLexerIMPORT = 2
	MyLangLexerPUB = 3
	MyLangLexerFN = 4
	MyLangLexerSTRUCT = 5
	MyLangLexerINTERFACE = 6
	MyLangLexerENUM = 7
	MyLangLexerCONST = 8
	MyLangLexerMUT = 9
	MyLangLexerTYPE = 10
	MyLangLexerIF = 11
	MyLangLexerELSE = 12
	MyLangLexerFOR = 13
	MyLangLexerIN = 14
	MyLangLexerMATCH = 15
	MyLangLexerRETURN = 16
	MyLangLexerDEFER = 17
	MyLangLexerTRUE = 18
	MyLangLexerFALSE = 19
	MyLangLexerNONE = 20
	MyLangLexerMAP = 21
	MyLangLexerLPAREN = 22
	MyLangLexerRPAREN = 23
	MyLangLexerLBRACE = 24
	MyLangLexerRBRACE = 25
	MyLangLexerLBRACK = 26
	MyLangLexerRBRACK = 27
	MyLangLexerLT = 28
	MyLangLexerGT = 29
	MyLangLexerLE = 30
	MyLangLexerGE = 31
	MyLangLexerEQ = 32
	MyLangLexerNE = 33
	MyLangLexerAND = 34
	MyLangLexerOR = 35
	MyLangLexerXOR = 36
	MyLangLexerNOT = 37
	MyLangLexerTILDE = 38
	MyLangLexerQUESTION = 39
	MyLangLexerCOLON = 40
	MyLangLexerSEMI = 41
	MyLangLexerCOMMA = 42
	MyLangLexerDOT = 43
	MyLangLexerASSIGN = 44
	MyLangLexerARROW = 45
	MyLangLexerEXCLAMATION = 46
	MyLangLexerPLUS = 47
	MyLangLexerMINUS = 48
	MyLangLexerMULT = 49
	MyLangLexerDIV = 50
	MyLangLexerMOD = 51
	MyLangLexerLSHIFT = 52
	MyLangLexerRSHIFT = 53
	MyLangLexerURSHIFT = 54
	MyLangLexerPLUS_ASSIGN = 55
	MyLangLexerMINUS_ASSIGN = 56
	MyLangLexerMULT_ASSIGN = 57
	MyLangLexerDIV_ASSIGN = 58
	MyLangLexerMOD_ASSIGN = 59
	MyLangLexerAND_ASSIGN = 60
	MyLangLexerOR_ASSIGN = 61
	MyLangLexerXOR_ASSIGN = 62
	MyLangLexerLSHIFT_ASSIGN = 63
	MyLangLexerRSHIFT_ASSIGN = 64
	MyLangLexerURSHIFT_ASSIGN = 65
	MyLangLexerLOGICAL_OR_ASSIGN = 66
	MyLangLexerLOGICAL_AND_ASSIGN = 67
	MyLangLexerLOGICAL_OR = 68
	MyLangLexerLOGICAL_AND = 69
	MyLangLexerDECL_ASSIGN = 70
	MyLangLexerIDENTIFIER = 71
	MyLangLexerINT = 72
	MyLangLexerFLOAT = 73
	MyLangLexerSTRING = 74
	MyLangLexerCHAR = 75
	MyLangLexerWHITESPACE = 76
	MyLangLexerSINGLE_LINE_COMMENT = 77
	MyLangLexerMULTI_LINE_COMMENT = 78
)

