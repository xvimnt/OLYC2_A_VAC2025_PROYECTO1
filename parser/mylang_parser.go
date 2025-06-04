// Code generated from ./grammar/MyLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MyLang
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type MyLangParser struct {
	*antlr.BaseParser
}

var MyLangParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func mylangParserInit() {
  staticData := &MyLangParserStaticData
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
    "program", "moduleDecl", "topLevelDecl", "importDecl", "importSpec", 
    "functionDecl", "genericParams", "params", "param", "structDecl", "structField", 
    "interfaceDecl", "interfaceMethod", "enumDecl", "enumField", "constDecl", 
    "constField", "type", "arrayType", "mapType", "optionType", "resultType", 
    "block", "statement", "varDecl", "assignment", "assignOp", "exprStmt", 
    "ifStmt", "forStmt", "matchStmt", "matchBranch", "returnStmt", "deferStmt", 
    "expr", "primary", "literal", "arrayLiteral", "mapLiteral", "mapElement", 
    "structLiteral", "structElement", "exprList", "returnType",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 78, 528, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26, 
	7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7, 
	31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36, 
	2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2, 
	42, 7, 42, 2, 43, 7, 43, 1, 0, 3, 0, 90, 8, 0, 1, 0, 1, 0, 5, 0, 94, 8, 
	0, 10, 0, 12, 0, 97, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 
	2, 1, 2, 1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 
	4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 123, 8, 4, 10, 4, 12, 4, 126, 9, 
	4, 1, 4, 3, 4, 129, 8, 4, 1, 5, 3, 5, 132, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 
	137, 8, 5, 1, 5, 1, 5, 3, 5, 141, 8, 5, 1, 5, 1, 5, 3, 5, 145, 8, 5, 1, 
	5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 153, 8, 6, 10, 6, 12, 6, 156, 9, 
	6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 163, 8, 7, 10, 7, 12, 7, 166, 9, 
	7, 1, 8, 3, 8, 169, 8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 3, 9, 175, 8, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 5, 9, 181, 8, 9, 10, 9, 12, 9, 184, 9, 9, 1, 9, 1, 9, 
	1, 10, 3, 10, 189, 8, 10, 1, 10, 3, 10, 192, 8, 10, 1, 10, 1, 10, 1, 10, 
	1, 11, 3, 11, 198, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 204, 8, 11, 
	10, 11, 12, 11, 207, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 214, 
	8, 12, 1, 12, 1, 12, 3, 12, 218, 8, 12, 1, 13, 3, 13, 221, 8, 13, 1, 13, 
	1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 229, 8, 13, 10, 13, 12, 13, 232, 
	9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 239, 8, 14, 1, 15, 3, 
	15, 242, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 249, 8, 15, 10, 
	15, 12, 15, 252, 9, 15, 1, 15, 1, 15, 3, 15, 256, 8, 15, 1, 16, 1, 16, 
	1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 267, 8, 17, 1, 
	18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 
	1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22, 287, 8, 22, 10, 
	22, 12, 22, 290, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 
	1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 303, 8, 23, 1, 24, 3, 24, 306, 8, 24, 
	1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 314, 8, 24, 1, 24, 1, 
	24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 
	1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 332, 8, 28, 3, 28, 334, 8, 28, 1, 29, 
	1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 
	29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 355, 8, 29, 
	1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 361, 8, 30, 10, 30, 12, 30, 364, 9, 
	30, 1, 30, 1, 30, 1, 31, 1, 31, 3, 31, 370, 8, 31, 1, 31, 1, 31, 1, 31, 
	1, 32, 1, 32, 3, 32, 377, 8, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 
	34, 1, 34, 1, 34, 3, 34, 387, 8, 34, 1, 34, 1, 34, 1, 34, 3, 34, 392, 8, 
	34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 
	1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 
	34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 
	1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 
	34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 441, 8, 34, 1, 34, 1, 34, 
	1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 449, 8, 34, 10, 34, 12, 34, 452, 9, 
	34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 
	463, 8, 35, 1, 36, 1, 36, 1, 37, 1, 37, 3, 37, 469, 8, 37, 1, 37, 1, 37, 
	1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 477, 8, 37, 1, 37, 1, 37, 3, 37, 481, 
	8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 487, 8, 38, 10, 38, 12, 38, 490, 
	9, 38, 3, 38, 492, 8, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 
	40, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 505, 8, 40, 10, 40, 12, 40, 508, 
	9, 40, 3, 40, 510, 8, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 
	42, 1, 42, 1, 42, 5, 42, 521, 8, 42, 10, 42, 12, 42, 524, 9, 42, 1, 43, 
	1, 43, 1, 43, 0, 1, 68, 44, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 0, 8, 2, 0, 44, 
	44, 55, 67, 2, 0, 37, 38, 48, 48, 1, 0, 49, 51, 1, 0, 47, 48, 1, 0, 52, 
	54, 1, 0, 28, 31, 1, 0, 32, 33, 2, 0, 18, 20, 72, 75, 570, 0, 89, 1, 0, 
	0, 0, 2, 100, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 111, 1, 0, 0, 0, 8, 128, 
	1, 0, 0, 0, 10, 131, 1, 0, 0, 0, 12, 148, 1, 0, 0, 0, 14, 159, 1, 0, 0, 
	0, 16, 168, 1, 0, 0, 0, 18, 174, 1, 0, 0, 0, 20, 188, 1, 0, 0, 0, 22, 197, 
	1, 0, 0, 0, 24, 210, 1, 0, 0, 0, 26, 220, 1, 0, 0, 0, 28, 235, 1, 0, 0, 
	0, 30, 241, 1, 0, 0, 0, 32, 257, 1, 0, 0, 0, 34, 266, 1, 0, 0, 0, 36, 268, 
	1, 0, 0, 0, 38, 272, 1, 0, 0, 0, 40, 278, 1, 0, 0, 0, 42, 281, 1, 0, 0, 
	0, 44, 284, 1, 0, 0, 0, 46, 302, 1, 0, 0, 0, 48, 305, 1, 0, 0, 0, 50, 317, 
	1, 0, 0, 0, 52, 321, 1, 0, 0, 0, 54, 323, 1, 0, 0, 0, 56, 325, 1, 0, 0, 
	0, 58, 335, 1, 0, 0, 0, 60, 356, 1, 0, 0, 0, 62, 369, 1, 0, 0, 0, 64, 374, 
	1, 0, 0, 0, 66, 378, 1, 0, 0, 0, 68, 391, 1, 0, 0, 0, 70, 462, 1, 0, 0, 
	0, 72, 464, 1, 0, 0, 0, 74, 480, 1, 0, 0, 0, 76, 482, 1, 0, 0, 0, 78, 495, 
	1, 0, 0, 0, 80, 499, 1, 0, 0, 0, 82, 513, 1, 0, 0, 0, 84, 517, 1, 0, 0, 
	0, 86, 525, 1, 0, 0, 0, 88, 90, 3, 2, 1, 0, 89, 88, 1, 0, 0, 0, 89, 90, 
	1, 0, 0, 0, 90, 95, 1, 0, 0, 0, 91, 94, 3, 4, 2, 0, 92, 94, 3, 46, 23, 
	0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 
	1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 
	98, 99, 5, 0, 0, 1, 99, 1, 1, 0, 0, 0, 100, 101, 5, 1, 0, 0, 101, 102, 
	5, 71, 0, 0, 102, 3, 1, 0, 0, 0, 103, 110, 3, 10, 5, 0, 104, 110, 3, 18, 
	9, 0, 105, 110, 3, 22, 11, 0, 106, 110, 3, 26, 13, 0, 107, 110, 3, 30, 
	15, 0, 108, 110, 3, 6, 3, 0, 109, 103, 1, 0, 0, 0, 109, 104, 1, 0, 0, 0, 
	109, 105, 1, 0, 0, 0, 109, 106, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 
	108, 1, 0, 0, 0, 110, 5, 1, 0, 0, 0, 111, 112, 5, 2, 0, 0, 112, 113, 3, 
	8, 4, 0, 113, 7, 1, 0, 0, 0, 114, 129, 5, 71, 0, 0, 115, 116, 5, 71, 0, 
	0, 116, 117, 5, 43, 0, 0, 117, 129, 5, 71, 0, 0, 118, 119, 5, 24, 0, 0, 
	119, 124, 5, 71, 0, 0, 120, 121, 5, 42, 0, 0, 121, 123, 5, 71, 0, 0, 122, 
	120, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 
	1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 129, 5, 25, 
	0, 0, 128, 114, 1, 0, 0, 0, 128, 115, 1, 0, 0, 0, 128, 118, 1, 0, 0, 0, 
	129, 9, 1, 0, 0, 0, 130, 132, 5, 3, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 
	1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 5, 4, 0, 0, 134, 136, 5, 71, 
	0, 0, 135, 137, 3, 12, 6, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 
	137, 138, 1, 0, 0, 0, 138, 140, 5, 22, 0, 0, 139, 141, 3, 14, 7, 0, 140, 
	139, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 
	5, 23, 0, 0, 143, 145, 3, 86, 43, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 
	0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 3, 44, 22, 0, 147, 11, 1, 0, 0, 
	0, 148, 149, 5, 28, 0, 0, 149, 154, 5, 71, 0, 0, 150, 151, 5, 42, 0, 0, 
	151, 153, 5, 71, 0, 0, 152, 150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 
	152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 
	1, 0, 0, 0, 157, 158, 5, 29, 0, 0, 158, 13, 1, 0, 0, 0, 159, 164, 3, 16, 
	8, 0, 160, 161, 5, 42, 0, 0, 161, 163, 3, 16, 8, 0, 162, 160, 1, 0, 0, 
	0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 
	15, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 169, 5, 9, 0, 0, 168, 167, 1, 
	0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 5, 71, 0, 
	0, 171, 172, 3, 34, 17, 0, 172, 17, 1, 0, 0, 0, 173, 175, 5, 3, 0, 0, 174, 
	173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 
	5, 5, 0, 0, 177, 178, 5, 71, 0, 0, 178, 182, 5, 24, 0, 0, 179, 181, 3, 
	20, 10, 0, 180, 179, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 
	0, 0, 182, 183, 1, 0, 0, 0, 183, 185, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 
	185, 186, 5, 25, 0, 0, 186, 19, 1, 0, 0, 0, 187, 189, 5, 3, 0, 0, 188, 
	187, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0, 190, 192, 
	5, 9, 0, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 
	0, 0, 193, 194, 5, 71, 0, 0, 194, 195, 3, 34, 17, 0, 195, 21, 1, 0, 0, 
	0, 196, 198, 5, 3, 0, 0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 
	199, 1, 0, 0, 0, 199, 200, 5, 6, 0, 0, 200, 201, 5, 71, 0, 0, 201, 205, 
	5, 24, 0, 0, 202, 204, 3, 24, 12, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 
	0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208, 1, 0, 0, 
	0, 207, 205, 1, 0, 0, 0, 208, 209, 5, 25, 0, 0, 209, 23, 1, 0, 0, 0, 210, 
	211, 5, 71, 0, 0, 211, 213, 5, 22, 0, 0, 212, 214, 3, 14, 7, 0, 213, 212, 
	1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 5, 23, 
	0, 0, 216, 218, 3, 86, 43, 0, 217, 216, 1, 0, 0, 0, 217, 218, 1, 0, 0, 
	0, 218, 25, 1, 0, 0, 0, 219, 221, 5, 3, 0, 0, 220, 219, 1, 0, 0, 0, 220, 
	221, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 5, 7, 0, 0, 223, 224, 
	5, 71, 0, 0, 224, 225, 5, 24, 0, 0, 225, 230, 3, 28, 14, 0, 226, 227, 5, 
	42, 0, 0, 227, 229, 3, 28, 14, 0, 228, 226, 1, 0, 0, 0, 229, 232, 1, 0, 
	0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 
	232, 230, 1, 0, 0, 0, 233, 234, 5, 25, 0, 0, 234, 27, 1, 0, 0, 0, 235, 
	238, 5, 71, 0, 0, 236, 237, 5, 44, 0, 0, 237, 239, 3, 68, 34, 0, 238, 236, 
	1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 29, 1, 0, 0, 0, 240, 242, 5, 3, 
	0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 
	243, 255, 5, 8, 0, 0, 244, 250, 5, 22, 0, 0, 245, 246, 3, 32, 16, 0, 246, 
	247, 5, 41, 0, 0, 247, 249, 1, 0, 0, 0, 248, 245, 1, 0, 0, 0, 249, 252, 
	1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 
	0, 0, 252, 250, 1, 0, 0, 0, 253, 256, 5, 23, 0, 0, 254, 256, 3, 32, 16, 
	0, 255, 244, 1, 0, 0, 0, 255, 254, 1, 0, 0, 0, 256, 31, 1, 0, 0, 0, 257, 
	258, 5, 71, 0, 0, 258, 259, 5, 44, 0, 0, 259, 260, 3, 68, 34, 0, 260, 33, 
	1, 0, 0, 0, 261, 267, 5, 71, 0, 0, 262, 267, 3, 36, 18, 0, 263, 267, 3, 
	38, 19, 0, 264, 267, 3, 40, 20, 0, 265, 267, 3, 42, 21, 0, 266, 261, 1, 
	0, 0, 0, 266, 262, 1, 0, 0, 0, 266, 263, 1, 0, 0, 0, 266, 264, 1, 0, 0, 
	0, 266, 265, 1, 0, 0, 0, 267, 35, 1, 0, 0, 0, 268, 269, 5, 26, 0, 0, 269, 
	270, 5, 27, 0, 0, 270, 271, 3, 34, 17, 0, 271, 37, 1, 0, 0, 0, 272, 273, 
	5, 21, 0, 0, 273, 274, 5, 26, 0, 0, 274, 275, 3, 34, 17, 0, 275, 276, 5, 
	27, 0, 0, 276, 277, 3, 34, 17, 0, 277, 39, 1, 0, 0, 0, 278, 279, 5, 39, 
	0, 0, 279, 280, 3, 34, 17, 0, 280, 41, 1, 0, 0, 0, 281, 282, 5, 46, 0, 
	0, 282, 283, 3, 34, 17, 0, 283, 43, 1, 0, 0, 0, 284, 288, 5, 24, 0, 0, 
	285, 287, 3, 46, 23, 0, 286, 285, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 
	286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1, 0, 0, 0, 290, 288, 
	1, 0, 0, 0, 291, 292, 5, 25, 0, 0, 292, 45, 1, 0, 0, 0, 293, 303, 3, 48, 
	24, 0, 294, 303, 3, 50, 25, 0, 295, 303, 3, 54, 27, 0, 296, 303, 3, 56, 
	28, 0, 297, 303, 3, 58, 29, 0, 298, 303, 3, 60, 30, 0, 299, 303, 3, 64, 
	32, 0, 300, 303, 3, 66, 33, 0, 301, 303, 3, 44, 22, 0, 302, 293, 1, 0, 
	0, 0, 302, 294, 1, 0, 0, 0, 302, 295, 1, 0, 0, 0, 302, 296, 1, 0, 0, 0, 
	302, 297, 1, 0, 0, 0, 302, 298, 1, 0, 0, 0, 302, 299, 1, 0, 0, 0, 302, 
	300, 1, 0, 0, 0, 302, 301, 1, 0, 0, 0, 303, 47, 1, 0, 0, 0, 304, 306, 5, 
	9, 0, 0, 305, 304, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 1, 0, 0, 
	0, 307, 313, 5, 71, 0, 0, 308, 314, 5, 70, 0, 0, 309, 310, 5, 40, 0, 0, 
	310, 311, 3, 34, 17, 0, 311, 312, 5, 44, 0, 0, 312, 314, 1, 0, 0, 0, 313, 
	308, 1, 0, 0, 0, 313, 309, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 
	3, 68, 34, 0, 316, 49, 1, 0, 0, 0, 317, 318, 3, 68, 34, 0, 318, 319, 3, 
	52, 26, 0, 319, 320, 3, 68, 34, 0, 320, 51, 1, 0, 0, 0, 321, 322, 7, 0, 
	0, 0, 322, 53, 1, 0, 0, 0, 323, 324, 3, 68, 34, 0, 324, 55, 1, 0, 0, 0, 
	325, 326, 5, 11, 0, 0, 326, 327, 3, 68, 34, 0, 327, 333, 3, 44, 22, 0, 
	328, 331, 5, 12, 0, 0, 329, 332, 3, 56, 28, 0, 330, 332, 3, 44, 22, 0, 
	331, 329, 1, 0, 0, 0, 331, 330, 1, 0, 0, 0, 332, 334, 1, 0, 0, 0, 333, 
	328, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 57, 1, 0, 0, 0, 335, 354, 5, 
	13, 0, 0, 336, 355, 3, 44, 22, 0, 337, 338, 3, 68, 34, 0, 338, 339, 3, 
	44, 22, 0, 339, 355, 1, 0, 0, 0, 340, 341, 3, 68, 34, 0, 341, 342, 5, 41, 
	0, 0, 342, 343, 3, 68, 34, 0, 343, 344, 5, 41, 0, 0, 344, 345, 3, 68, 34, 
	0, 345, 346, 1, 0, 0, 0, 346, 347, 3, 44, 22, 0, 347, 355, 1, 0, 0, 0, 
	348, 349, 5, 71, 0, 0, 349, 350, 5, 14, 0, 0, 350, 351, 3, 68, 34, 0, 351, 
	352, 1, 0, 0, 0, 352, 353, 3, 44, 22, 0, 353, 355, 1, 0, 0, 0, 354, 336, 
	1, 0, 0, 0, 354, 337, 1, 0, 0, 0, 354, 340, 1, 0, 0, 0, 354, 348, 1, 0, 
	0, 0, 355, 59, 1, 0, 0, 0, 356, 357, 5, 15, 0, 0, 357, 358, 3, 68, 34, 
	0, 358, 362, 5, 24, 0, 0, 359, 361, 3, 62, 31, 0, 360, 359, 1, 0, 0, 0, 
	361, 364, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 
	365, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0, 365, 366, 5, 25, 0, 0, 366, 61, 
	1, 0, 0, 0, 367, 370, 3, 68, 34, 0, 368, 370, 5, 12, 0, 0, 369, 367, 1, 
	0, 0, 0, 369, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 372, 5, 45, 0, 
	0, 372, 373, 3, 46, 23, 0, 373, 63, 1, 0, 0, 0, 374, 376, 5, 16, 0, 0, 
	375, 377, 3, 68, 34, 0, 376, 375, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 
	65, 1, 0, 0, 0, 378, 379, 5, 17, 0, 0, 379, 380, 3, 44, 22, 0, 380, 67, 
	1, 0, 0, 0, 381, 382, 6, 34, -1, 0, 382, 392, 3, 70, 35, 0, 383, 384, 5, 
	71, 0, 0, 384, 386, 5, 22, 0, 0, 385, 387, 3, 84, 42, 0, 386, 385, 1, 0, 
	0, 0, 386, 387, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 392, 5, 23, 0, 0, 
	389, 390, 7, 1, 0, 0, 390, 392, 3, 68, 34, 13, 391, 381, 1, 0, 0, 0, 391, 
	383, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0, 392, 450, 1, 0, 0, 0, 393, 394, 
	10, 12, 0, 0, 394, 395, 7, 2, 0, 0, 395, 449, 3, 68, 34, 13, 396, 397, 
	10, 11, 0, 0, 397, 398, 7, 3, 0, 0, 398, 449, 3, 68, 34, 12, 399, 400, 
	10, 10, 0, 0, 400, 401, 7, 4, 0, 0, 401, 449, 3, 68, 34, 11, 402, 403, 
	10, 9, 0, 0, 403, 404, 7, 5, 0, 0, 404, 449, 3, 68, 34, 10, 405, 406, 10, 
	8, 0, 0, 406, 407, 7, 6, 0, 0, 407, 449, 3, 68, 34, 9, 408, 409, 10, 7, 
	0, 0, 409, 410, 5, 34, 0, 0, 410, 449, 3, 68, 34, 8, 411, 412, 10, 6, 0, 
	0, 412, 413, 5, 36, 0, 0, 413, 449, 3, 68, 34, 7, 414, 415, 10, 5, 0, 0, 
	415, 416, 5, 35, 0, 0, 416, 449, 3, 68, 34, 6, 417, 418, 10, 4, 0, 0, 418, 
	419, 5, 69, 0, 0, 419, 449, 3, 68, 34, 5, 420, 421, 10, 3, 0, 0, 421, 422, 
	5, 68, 0, 0, 422, 449, 3, 68, 34, 4, 423, 424, 10, 2, 0, 0, 424, 425, 5, 
	39, 0, 0, 425, 426, 3, 68, 34, 0, 426, 427, 5, 40, 0, 0, 427, 428, 3, 68, 
	34, 3, 428, 449, 1, 0, 0, 0, 429, 430, 10, 1, 0, 0, 430, 431, 5, 44, 0, 
	0, 431, 449, 3, 68, 34, 1, 432, 433, 10, 17, 0, 0, 433, 434, 5, 43, 0, 
	0, 434, 449, 5, 71, 0, 0, 435, 436, 10, 16, 0, 0, 436, 437, 5, 43, 0, 0, 
	437, 438, 5, 71, 0, 0, 438, 440, 5, 22, 0, 0, 439, 441, 3, 84, 42, 0, 440, 
	439, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 449, 
	5, 23, 0, 0, 443, 444, 10, 14, 0, 0, 444, 445, 5, 26, 0, 0, 445, 446, 3, 
	68, 34, 0, 446, 447, 5, 27, 0, 0, 447, 449, 1, 0, 0, 0, 448, 393, 1, 0, 
	0, 0, 448, 396, 1, 0, 0, 0, 448, 399, 1, 0, 0, 0, 448, 402, 1, 0, 0, 0, 
	448, 405, 1, 0, 0, 0, 448, 408, 1, 0, 0, 0, 448, 411, 1, 0, 0, 0, 448, 
	414, 1, 0, 0, 0, 448, 417, 1, 0, 0, 0, 448, 420, 1, 0, 0, 0, 448, 423, 
	1, 0, 0, 0, 448, 429, 1, 0, 0, 0, 448, 432, 1, 0, 0, 0, 448, 435, 1, 0, 
	0, 0, 448, 443, 1, 0, 0, 0, 449, 452, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 
	450, 451, 1, 0, 0, 0, 451, 69, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 453, 463, 
	5, 71, 0, 0, 454, 463, 3, 72, 36, 0, 455, 463, 3, 74, 37, 0, 456, 463, 
	3, 76, 38, 0, 457, 463, 3, 80, 40, 0, 458, 459, 5, 22, 0, 0, 459, 460, 
	3, 68, 34, 0, 460, 461, 5, 23, 0, 0, 461, 463, 1, 0, 0, 0, 462, 453, 1, 
	0, 0, 0, 462, 454, 1, 0, 0, 0, 462, 455, 1, 0, 0, 0, 462, 456, 1, 0, 0, 
	0, 462, 457, 1, 0, 0, 0, 462, 458, 1, 0, 0, 0, 463, 71, 1, 0, 0, 0, 464, 
	465, 7, 7, 0, 0, 465, 73, 1, 0, 0, 0, 466, 468, 5, 26, 0, 0, 467, 469, 
	3, 84, 42, 0, 468, 467, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 470, 1, 
	0, 0, 0, 470, 481, 5, 27, 0, 0, 471, 472, 5, 26, 0, 0, 472, 473, 3, 34, 
	17, 0, 473, 474, 5, 27, 0, 0, 474, 476, 5, 24, 0, 0, 475, 477, 3, 84, 42, 
	0, 476, 475, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 
	479, 5, 25, 0, 0, 479, 481, 1, 0, 0, 0, 480, 466, 1, 0, 0, 0, 480, 471, 
	1, 0, 0, 0, 481, 75, 1, 0, 0, 0, 482, 491, 5, 24, 0, 0, 483, 488, 3, 78, 
	39, 0, 484, 485, 5, 42, 0, 0, 485, 487, 3, 78, 39, 0, 486, 484, 1, 0, 0, 
	0, 487, 490, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 
	492, 1, 0, 0, 0, 490, 488, 1, 0, 0, 0, 491, 483, 1, 0, 0, 0, 491, 492, 
	1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 494, 5, 25, 0, 0, 494, 77, 1, 0, 
	0, 0, 495, 496, 3, 68, 34, 0, 496, 497, 5, 40, 0, 0, 497, 498, 3, 68, 34, 
	0, 498, 79, 1, 0, 0, 0, 499, 500, 3, 34, 17, 0, 500, 509, 5, 24, 0, 0, 
	501, 506, 3, 82, 41, 0, 502, 503, 5, 42, 0, 0, 503, 505, 3, 82, 41, 0, 
	504, 502, 1, 0, 0, 0, 505, 508, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 506, 
	507, 1, 0, 0, 0, 507, 510, 1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 509, 501, 
	1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 511, 1, 0, 0, 0, 511, 512, 5, 25, 
	0, 0, 512, 81, 1, 0, 0, 0, 513, 514, 5, 71, 0, 0, 514, 515, 5, 40, 0, 0, 
	515, 516, 3, 68, 34, 0, 516, 83, 1, 0, 0, 0, 517, 522, 3, 68, 34, 0, 518, 
	519, 5, 42, 0, 0, 519, 521, 3, 68, 34, 0, 520, 518, 1, 0, 0, 0, 521, 524, 
	1, 0, 0, 0, 522, 520, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 85, 1, 0, 
	0, 0, 524, 522, 1, 0, 0, 0, 525, 526, 3, 34, 17, 0, 526, 87, 1, 0, 0, 0, 
	52, 89, 93, 95, 109, 124, 128, 131, 136, 140, 144, 154, 164, 168, 174, 
	182, 188, 191, 197, 205, 213, 217, 220, 230, 238, 241, 250, 255, 266, 288, 
	302, 305, 313, 331, 333, 354, 362, 369, 376, 386, 391, 440, 448, 450, 462, 
	468, 476, 480, 488, 491, 506, 509, 522,
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

// MyLangParserInit initializes any static state used to implement MyLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMyLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyLangParserInit() {
  staticData := &MyLangParserStaticData
  staticData.once.Do(mylangParserInit)
}

// NewMyLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMyLangParser(input antlr.TokenStream) *MyLangParser {
	MyLangParserInit()
	this := new(MyLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &MyLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MyLang.g4"

	return this
}


// MyLangParser tokens.
const (
	MyLangParserEOF = antlr.TokenEOF
	MyLangParserMODULE = 1
	MyLangParserIMPORT = 2
	MyLangParserPUB = 3
	MyLangParserFN = 4
	MyLangParserSTRUCT = 5
	MyLangParserINTERFACE = 6
	MyLangParserENUM = 7
	MyLangParserCONST = 8
	MyLangParserMUT = 9
	MyLangParserTYPE = 10
	MyLangParserIF = 11
	MyLangParserELSE = 12
	MyLangParserFOR = 13
	MyLangParserIN = 14
	MyLangParserMATCH = 15
	MyLangParserRETURN = 16
	MyLangParserDEFER = 17
	MyLangParserTRUE = 18
	MyLangParserFALSE = 19
	MyLangParserNONE = 20
	MyLangParserMAP = 21
	MyLangParserLPAREN = 22
	MyLangParserRPAREN = 23
	MyLangParserLBRACE = 24
	MyLangParserRBRACE = 25
	MyLangParserLBRACK = 26
	MyLangParserRBRACK = 27
	MyLangParserLT = 28
	MyLangParserGT = 29
	MyLangParserLE = 30
	MyLangParserGE = 31
	MyLangParserEQ = 32
	MyLangParserNE = 33
	MyLangParserAND = 34
	MyLangParserOR = 35
	MyLangParserXOR = 36
	MyLangParserNOT = 37
	MyLangParserTILDE = 38
	MyLangParserQUESTION = 39
	MyLangParserCOLON = 40
	MyLangParserSEMI = 41
	MyLangParserCOMMA = 42
	MyLangParserDOT = 43
	MyLangParserASSIGN = 44
	MyLangParserARROW = 45
	MyLangParserEXCLAMATION = 46
	MyLangParserPLUS = 47
	MyLangParserMINUS = 48
	MyLangParserMULT = 49
	MyLangParserDIV = 50
	MyLangParserMOD = 51
	MyLangParserLSHIFT = 52
	MyLangParserRSHIFT = 53
	MyLangParserURSHIFT = 54
	MyLangParserPLUS_ASSIGN = 55
	MyLangParserMINUS_ASSIGN = 56
	MyLangParserMULT_ASSIGN = 57
	MyLangParserDIV_ASSIGN = 58
	MyLangParserMOD_ASSIGN = 59
	MyLangParserAND_ASSIGN = 60
	MyLangParserOR_ASSIGN = 61
	MyLangParserXOR_ASSIGN = 62
	MyLangParserLSHIFT_ASSIGN = 63
	MyLangParserRSHIFT_ASSIGN = 64
	MyLangParserURSHIFT_ASSIGN = 65
	MyLangParserLOGICAL_OR_ASSIGN = 66
	MyLangParserLOGICAL_AND_ASSIGN = 67
	MyLangParserLOGICAL_OR = 68
	MyLangParserLOGICAL_AND = 69
	MyLangParserDECL_ASSIGN = 70
	MyLangParserIDENTIFIER = 71
	MyLangParserINT = 72
	MyLangParserFLOAT = 73
	MyLangParserSTRING = 74
	MyLangParserCHAR = 75
	MyLangParserWHITESPACE = 76
	MyLangParserSINGLE_LINE_COMMENT = 77
	MyLangParserMULTI_LINE_COMMENT = 78
)

// MyLangParser rules.
const (
	MyLangParserRULE_program = 0
	MyLangParserRULE_moduleDecl = 1
	MyLangParserRULE_topLevelDecl = 2
	MyLangParserRULE_importDecl = 3
	MyLangParserRULE_importSpec = 4
	MyLangParserRULE_functionDecl = 5
	MyLangParserRULE_genericParams = 6
	MyLangParserRULE_params = 7
	MyLangParserRULE_param = 8
	MyLangParserRULE_structDecl = 9
	MyLangParserRULE_structField = 10
	MyLangParserRULE_interfaceDecl = 11
	MyLangParserRULE_interfaceMethod = 12
	MyLangParserRULE_enumDecl = 13
	MyLangParserRULE_enumField = 14
	MyLangParserRULE_constDecl = 15
	MyLangParserRULE_constField = 16
	MyLangParserRULE_type = 17
	MyLangParserRULE_arrayType = 18
	MyLangParserRULE_mapType = 19
	MyLangParserRULE_optionType = 20
	MyLangParserRULE_resultType = 21
	MyLangParserRULE_block = 22
	MyLangParserRULE_statement = 23
	MyLangParserRULE_varDecl = 24
	MyLangParserRULE_assignment = 25
	MyLangParserRULE_assignOp = 26
	MyLangParserRULE_exprStmt = 27
	MyLangParserRULE_ifStmt = 28
	MyLangParserRULE_forStmt = 29
	MyLangParserRULE_matchStmt = 30
	MyLangParserRULE_matchBranch = 31
	MyLangParserRULE_returnStmt = 32
	MyLangParserRULE_deferStmt = 33
	MyLangParserRULE_expr = 34
	MyLangParserRULE_primary = 35
	MyLangParserRULE_literal = 36
	MyLangParserRULE_arrayLiteral = 37
	MyLangParserRULE_mapLiteral = 38
	MyLangParserRULE_mapElement = 39
	MyLangParserRULE_structLiteral = 40
	MyLangParserRULE_structElement = 41
	MyLangParserRULE_exprList = 42
	MyLangParserRULE_returnType = 43
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	ModuleDecl() IModuleDeclContext
	AllTopLevelDecl() []ITopLevelDeclContext
	TopLevelDecl(i int) ITopLevelDeclContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MyLangParserEOF, 0)
}

func (s *ProgramContext) ModuleDecl() IModuleDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleDeclContext)
}

func (s *ProgramContext) AllTopLevelDecl() []ITopLevelDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelDeclContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelDeclContext); ok {
			tst[i] = t.(ITopLevelDeclContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) TopLevelDecl(i int) ITopLevelDeclContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelDeclContext)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MyLangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserMODULE {
		{
			p.SetState(88)
			p.ModuleDecl()
		}

	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 352805885815804) != 0) || ((int64((_la - 71)) & ^0x3f) == 0 && ((int64(1) << (_la - 71)) & 31) != 0) {
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MyLangParserIMPORT, MyLangParserPUB, MyLangParserFN, MyLangParserSTRUCT, MyLangParserINTERFACE, MyLangParserENUM, MyLangParserCONST:
			{
				p.SetState(91)
				p.TopLevelDecl()
			}


		case MyLangParserMUT, MyLangParserIF, MyLangParserFOR, MyLangParserMATCH, MyLangParserRETURN, MyLangParserDEFER, MyLangParserTRUE, MyLangParserFALSE, MyLangParserNONE, MyLangParserMAP, MyLangParserLPAREN, MyLangParserLBRACE, MyLangParserLBRACK, MyLangParserNOT, MyLangParserTILDE, MyLangParserQUESTION, MyLangParserEXCLAMATION, MyLangParserMINUS, MyLangParserIDENTIFIER, MyLangParserINT, MyLangParserFLOAT, MyLangParserSTRING, MyLangParserCHAR:
			{
				p.SetState(92)
				p.Statement()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Match(MyLangParserEOF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IModuleDeclContext is an interface to support dynamic dispatch.
type IModuleDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	MODULE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsModuleDeclContext differentiates from other interfaces.
	IsModuleDeclContext()
}

type ModuleDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyModuleDeclContext() *ModuleDeclContext {
	var p = new(ModuleDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_moduleDecl
	return p
}

func InitEmptyModuleDeclContext(p *ModuleDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_moduleDecl
}

func (*ModuleDeclContext) IsModuleDeclContext() {}

func NewModuleDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDeclContext {
	var p = new(ModuleDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_moduleDecl

	return p
}

func (s *ModuleDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDeclContext) GetName() antlr.Token { return s.name }


func (s *ModuleDeclContext) SetName(v antlr.Token) { s.name = v }


func (s *ModuleDeclContext) MODULE() antlr.TerminalNode {
	return s.GetToken(MyLangParserMODULE, 0)
}

func (s *ModuleDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *ModuleDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModuleDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterModuleDecl(s)
	}
}

func (s *ModuleDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitModuleDecl(s)
	}
}

func (s *ModuleDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitModuleDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ModuleDecl() (localctx IModuleDeclContext) {
	localctx = NewModuleDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MyLangParserRULE_moduleDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(MyLangParserMODULE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(101)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*ModuleDeclContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITopLevelDeclContext is an interface to support dynamic dispatch.
type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDecl() IFunctionDeclContext
	StructDecl() IStructDeclContext
	InterfaceDecl() IInterfaceDeclContext
	EnumDecl() IEnumDeclContext
	ConstDecl() IConstDeclContext
	ImportDecl() IImportDeclContext

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}

type TopLevelDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_topLevelDecl
	return p
}

func InitEmptyTopLevelDeclContext(p *TopLevelDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_topLevelDecl
}

func (*TopLevelDeclContext) IsTopLevelDeclContext() {}

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext {
	var p = new(TopLevelDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_topLevelDecl

	return p
}

func (s *TopLevelDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDeclContext) FunctionDecl() IFunctionDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclContext)
}

func (s *TopLevelDeclContext) StructDecl() IStructDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *TopLevelDeclContext) InterfaceDecl() IInterfaceDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclContext)
}

func (s *TopLevelDeclContext) EnumDecl() IEnumDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDeclContext)
}

func (s *TopLevelDeclContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *TopLevelDeclContext) ImportDecl() IImportDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TopLevelDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitTopLevelDecl(s)
	}
}

func (s *TopLevelDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitTopLevelDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) TopLevelDecl() (localctx ITopLevelDeclContext) {
	localctx = NewTopLevelDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MyLangParserRULE_topLevelDecl)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.FunctionDecl()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.StructDecl()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.InterfaceDecl()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.EnumDecl()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.ConstDecl()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.ImportDecl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	ImportSpec() IImportSpecContext

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_importDecl
	return p
}

func InitEmptyImportDeclContext(p *ImportDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_importDecl
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(MyLangParserIMPORT, 0)
}

func (s *ImportDeclContext) ImportSpec() IImportSpecContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportSpecContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MyLangParserRULE_importDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(MyLangParserIMPORT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(112)
		p.ImportSpec()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 

	// GetSubname returns the subname token.
	GetSubname() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 

	// SetSubname sets the subname token.
	SetSubname(antlr.Token) 


	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
	subname antlr.Token
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_importSpec
	return p
}

func InitEmptyImportSpecContext(p *ImportSpecContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_importSpec
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) GetName() antlr.Token { return s.name }

func (s *ImportSpecContext) GetSubname() antlr.Token { return s.subname }


func (s *ImportSpecContext) SetName(v antlr.Token) { s.name = v }

func (s *ImportSpecContext) SetSubname(v antlr.Token) { s.subname = v }


func (s *ImportSpecContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserIDENTIFIER)
}

func (s *ImportSpecContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, i)
}

func (s *ImportSpecContext) DOT() antlr.TerminalNode {
	return s.GetToken(MyLangParserDOT, 0)
}

func (s *ImportSpecContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *ImportSpecContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *ImportSpecContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserCOMMA)
}

func (s *ImportSpecContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserCOMMA, i)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MyLangParserRULE_importSpec)
	var _la int

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)

			var _m = p.Match(MyLangParserIDENTIFIER)

			localctx.(*ImportSpecContext).name = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)

			var _m = p.Match(MyLangParserIDENTIFIER)

			localctx.(*ImportSpecContext).name = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(MyLangParserDOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(117)

			var _m = p.Match(MyLangParserIDENTIFIER)

			localctx.(*ImportSpecContext).subname = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(MyLangParserLBRACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(MyLangParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == MyLangParserCOMMA {
			{
				p.SetState(120)
				p.Match(MyLangParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(121)
				p.Match(MyLangParserIDENTIFIER)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(127)
			p.Match(MyLangParserRBRACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFunctionDeclContext is an interface to support dynamic dispatch.
type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	FN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	IDENTIFIER() antlr.TerminalNode
	PUB() antlr.TerminalNode
	GenericParams() IGenericParamsContext
	Params() IParamsContext
	ReturnType() IReturnTypeContext

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}

type FunctionDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext {
	var p = new(FunctionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_functionDecl
	return p
}

func InitEmptyFunctionDeclContext(p *FunctionDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_functionDecl
}

func (*FunctionDeclContext) IsFunctionDeclContext() {}

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_functionDecl

	return p
}

func (s *FunctionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclContext) GetName() antlr.Token { return s.name }


func (s *FunctionDeclContext) SetName(v antlr.Token) { s.name = v }


func (s *FunctionDeclContext) FN() antlr.TerminalNode {
	return s.GetToken(MyLangParserFN, 0)
}

func (s *FunctionDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLPAREN, 0)
}

func (s *FunctionDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserRPAREN, 0)
}

func (s *FunctionDeclContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *FunctionDeclContext) PUB() antlr.TerminalNode {
	return s.GetToken(MyLangParserPUB, 0)
}

func (s *FunctionDeclContext) GenericParams() IGenericParamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericParamsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericParamsContext)
}

func (s *FunctionDeclContext) Params() IParamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunctionDeclContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitFunctionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) FunctionDecl() (localctx IFunctionDeclContext) {
	localctx = NewFunctionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MyLangParserRULE_functionDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserPUB {
		{
			p.SetState(130)
			p.Match(MyLangParserPUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(133)
		p.Match(MyLangParserFN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(134)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*FunctionDeclContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserLT {
		{
			p.SetState(135)
			p.GenericParams()
		}

	}
	{
		p.SetState(138)
		p.Match(MyLangParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserMUT || _la == MyLangParserIDENTIFIER {
		{
			p.SetState(139)
			p.Params()
		}

	}
	{
		p.SetState(142)
		p.Match(MyLangParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64((_la - 21)) & ^0x3f) == 0 && ((int64(1) << (_la - 21)) & 1125899940659233) != 0) {
		{
			p.SetState(143)
			p.ReturnType()
		}

	}
	{
		p.SetState(146)
		p.Block()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IGenericParamsContext is an interface to support dynamic dispatch.
type IGenericParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	GT() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsGenericParamsContext differentiates from other interfaces.
	IsGenericParamsContext()
}

type GenericParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericParamsContext() *GenericParamsContext {
	var p = new(GenericParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_genericParams
	return p
}

func InitEmptyGenericParamsContext(p *GenericParamsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_genericParams
}

func (*GenericParamsContext) IsGenericParamsContext() {}

func NewGenericParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericParamsContext {
	var p = new(GenericParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_genericParams

	return p
}

func (s *GenericParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericParamsContext) LT() antlr.TerminalNode {
	return s.GetToken(MyLangParserLT, 0)
}

func (s *GenericParamsContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserIDENTIFIER)
}

func (s *GenericParamsContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, i)
}

func (s *GenericParamsContext) GT() antlr.TerminalNode {
	return s.GetToken(MyLangParserGT, 0)
}

func (s *GenericParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserCOMMA)
}

func (s *GenericParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserCOMMA, i)
}

func (s *GenericParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GenericParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterGenericParams(s)
	}
}

func (s *GenericParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitGenericParams(s)
	}
}

func (s *GenericParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitGenericParams(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) GenericParams() (localctx IGenericParamsContext) {
	localctx = NewGenericParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MyLangParserRULE_genericParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(MyLangParserLT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(MyLangParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MyLangParserCOMMA {
		{
			p.SetState(150)
			p.Match(MyLangParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(MyLangParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(MyLangParserGT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserCOMMA, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MyLangParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Param()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MyLangParserCOMMA {
		{
			p.SetState(160)
			p.Match(MyLangParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Param()
		}


		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	Type_() ITypeContext
	IDENTIFIER() antlr.TerminalNode
	MUT() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) GetName() antlr.Token { return s.name }


func (s *ParamContext) SetName(v antlr.Token) { s.name = v }


func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *ParamContext) MUT() antlr.TerminalNode {
	return s.GetToken(MyLangParserMUT, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MyLangParserRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserMUT {
		{
			p.SetState(167)
			p.Match(MyLangParserMUT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(170)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*ParamContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Type_()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	PUB() antlr.TerminalNode
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structDecl
	return p
}

func InitEmptyStructDeclContext(p *StructDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structDecl
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) GetName() antlr.Token { return s.name }


func (s *StructDeclContext) SetName(v antlr.Token) { s.name = v }


func (s *StructDeclContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(MyLangParserSTRUCT, 0)
}

func (s *StructDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *StructDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *StructDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *StructDeclContext) PUB() antlr.TerminalNode {
	return s.GetToken(MyLangParserPUB, 0)
}

func (s *StructDeclContext) AllStructField() []IStructFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldContext); ok {
			tst[i] = t.(IStructFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclContext) StructField(i int) IStructFieldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MyLangParserRULE_structDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserPUB {
		{
			p.SetState(173)
			p.Match(MyLangParserPUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(176)
		p.Match(MyLangParserSTRUCT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(177)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*StructDeclContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(MyLangParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MyLangParserPUB || _la == MyLangParserMUT || _la == MyLangParserIDENTIFIER {
		{
			p.SetState(179)
			p.StructField()
		}


		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(MyLangParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStructFieldContext is an interface to support dynamic dispatch.
type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	Type_() ITypeContext
	IDENTIFIER() antlr.TerminalNode
	PUB() antlr.TerminalNode
	MUT() antlr.TerminalNode

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyStructFieldContext() *StructFieldContext {
	var p = new(StructFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structField
	return p
}

func InitEmptyStructFieldContext(p *StructFieldContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structField
}

func (*StructFieldContext) IsStructFieldContext() {}

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext {
	var p = new(StructFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_structField

	return p
}

func (s *StructFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldContext) GetName() antlr.Token { return s.name }


func (s *StructFieldContext) SetName(v antlr.Token) { s.name = v }


func (s *StructFieldContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *StructFieldContext) PUB() antlr.TerminalNode {
	return s.GetToken(MyLangParserPUB, 0)
}

func (s *StructFieldContext) MUT() antlr.TerminalNode {
	return s.GetToken(MyLangParserMUT, 0)
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitStructField(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) StructField() (localctx IStructFieldContext) {
	localctx = NewStructFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MyLangParserRULE_structField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserPUB {
		{
			p.SetState(187)
			p.Match(MyLangParserPUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserMUT {
		{
			p.SetState(190)
			p.Match(MyLangParserMUT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(193)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*StructFieldContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Type_()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IInterfaceDeclContext is an interface to support dynamic dispatch.
type IInterfaceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	INTERFACE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	PUB() antlr.TerminalNode
	AllInterfaceMethod() []IInterfaceMethodContext
	InterfaceMethod(i int) IInterfaceMethodContext

	// IsInterfaceDeclContext differentiates from other interfaces.
	IsInterfaceDeclContext()
}

type InterfaceDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyInterfaceDeclContext() *InterfaceDeclContext {
	var p = new(InterfaceDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_interfaceDecl
	return p
}

func InitEmptyInterfaceDeclContext(p *InterfaceDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_interfaceDecl
}

func (*InterfaceDeclContext) IsInterfaceDeclContext() {}

func NewInterfaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceDeclContext {
	var p = new(InterfaceDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_interfaceDecl

	return p
}

func (s *InterfaceDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceDeclContext) GetName() antlr.Token { return s.name }


func (s *InterfaceDeclContext) SetName(v antlr.Token) { s.name = v }


func (s *InterfaceDeclContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserINTERFACE, 0)
}

func (s *InterfaceDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *InterfaceDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *InterfaceDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *InterfaceDeclContext) PUB() antlr.TerminalNode {
	return s.GetToken(MyLangParserPUB, 0)
}

func (s *InterfaceDeclContext) AllInterfaceMethod() []IInterfaceMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInterfaceMethodContext); ok {
			len++
		}
	}

	tst := make([]IInterfaceMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInterfaceMethodContext); ok {
			tst[i] = t.(IInterfaceMethodContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceDeclContext) InterfaceMethod(i int) IInterfaceMethodContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceMethodContext)
}

func (s *InterfaceDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InterfaceDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterInterfaceDecl(s)
	}
}

func (s *InterfaceDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitInterfaceDecl(s)
	}
}

func (s *InterfaceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitInterfaceDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) InterfaceDecl() (localctx IInterfaceDeclContext) {
	localctx = NewInterfaceDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MyLangParserRULE_interfaceDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserPUB {
		{
			p.SetState(196)
			p.Match(MyLangParserPUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(199)
		p.Match(MyLangParserINTERFACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(200)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*InterfaceDeclContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(MyLangParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MyLangParserIDENTIFIER {
		{
			p.SetState(202)
			p.InterfaceMethod()
		}


		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
		p.Match(MyLangParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IInterfaceMethodContext is an interface to support dynamic dispatch.
type IInterfaceMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Params() IParamsContext
	ReturnType() IReturnTypeContext

	// IsInterfaceMethodContext differentiates from other interfaces.
	IsInterfaceMethodContext()
}

type InterfaceMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyInterfaceMethodContext() *InterfaceMethodContext {
	var p = new(InterfaceMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_interfaceMethod
	return p
}

func InitEmptyInterfaceMethodContext(p *InterfaceMethodContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_interfaceMethod
}

func (*InterfaceMethodContext) IsInterfaceMethodContext() {}

func NewInterfaceMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceMethodContext {
	var p = new(InterfaceMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_interfaceMethod

	return p
}

func (s *InterfaceMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceMethodContext) GetName() antlr.Token { return s.name }


func (s *InterfaceMethodContext) SetName(v antlr.Token) { s.name = v }


func (s *InterfaceMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLPAREN, 0)
}

func (s *InterfaceMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserRPAREN, 0)
}

func (s *InterfaceMethodContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *InterfaceMethodContext) Params() IParamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *InterfaceMethodContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *InterfaceMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InterfaceMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterInterfaceMethod(s)
	}
}

func (s *InterfaceMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitInterfaceMethod(s)
	}
}

func (s *InterfaceMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitInterfaceMethod(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) InterfaceMethod() (localctx IInterfaceMethodContext) {
	localctx = NewInterfaceMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MyLangParserRULE_interfaceMethod)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*InterfaceMethodContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(MyLangParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserMUT || _la == MyLangParserIDENTIFIER {
		{
			p.SetState(212)
			p.Params()
		}

	}
	{
		p.SetState(215)
		p.Match(MyLangParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(216)
			p.ReturnType()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEnumDeclContext is an interface to support dynamic dispatch.
type IEnumDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	ENUM() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AllEnumField() []IEnumFieldContext
	EnumField(i int) IEnumFieldContext
	RBRACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	PUB() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsEnumDeclContext differentiates from other interfaces.
	IsEnumDeclContext()
}

type EnumDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyEnumDeclContext() *EnumDeclContext {
	var p = new(EnumDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_enumDecl
	return p
}

func InitEmptyEnumDeclContext(p *EnumDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_enumDecl
}

func (*EnumDeclContext) IsEnumDeclContext() {}

func NewEnumDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclContext {
	var p = new(EnumDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_enumDecl

	return p
}

func (s *EnumDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclContext) GetName() antlr.Token { return s.name }


func (s *EnumDeclContext) SetName(v antlr.Token) { s.name = v }


func (s *EnumDeclContext) ENUM() antlr.TerminalNode {
	return s.GetToken(MyLangParserENUM, 0)
}

func (s *EnumDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *EnumDeclContext) AllEnumField() []IEnumFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumFieldContext); ok {
			len++
		}
	}

	tst := make([]IEnumFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumFieldContext); ok {
			tst[i] = t.(IEnumFieldContext)
			i++
		}
	}

	return tst
}

func (s *EnumDeclContext) EnumField(i int) IEnumFieldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *EnumDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *EnumDeclContext) PUB() antlr.TerminalNode {
	return s.GetToken(MyLangParserPUB, 0)
}

func (s *EnumDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserCOMMA)
}

func (s *EnumDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserCOMMA, i)
}

func (s *EnumDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EnumDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterEnumDecl(s)
	}
}

func (s *EnumDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitEnumDecl(s)
	}
}

func (s *EnumDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitEnumDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) EnumDecl() (localctx IEnumDeclContext) {
	localctx = NewEnumDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MyLangParserRULE_enumDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserPUB {
		{
			p.SetState(219)
			p.Match(MyLangParserPUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(222)
		p.Match(MyLangParserENUM)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(223)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*EnumDeclContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(MyLangParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(225)
		p.EnumField()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MyLangParserCOMMA {
		{
			p.SetState(226)
			p.Match(MyLangParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(227)
			p.EnumField()
		}


		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(MyLangParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) GetName() antlr.Token { return s.name }


func (s *EnumFieldContext) SetName(v antlr.Token) { s.name = v }


func (s *EnumFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *EnumFieldContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserASSIGN, 0)
}

func (s *EnumFieldContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (s *EnumFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MyLangParserRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*EnumFieldContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserASSIGN {
		{
			p.SetState(236)
			p.Match(MyLangParserASSIGN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(237)
			p.expr(0)
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IConstDeclContext is an interface to support dynamic dispatch.
type IConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllConstField() []IConstFieldContext
	ConstField(i int) IConstFieldContext
	PUB() antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsConstDeclContext differentiates from other interfaces.
	IsConstDeclContext()
}

type ConstDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclContext() *ConstDeclContext {
	var p = new(ConstDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_constDecl
	return p
}

func InitEmptyConstDeclContext(p *ConstDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_constDecl
}

func (*ConstDeclContext) IsConstDeclContext() {}

func NewConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclContext {
	var p = new(ConstDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_constDecl

	return p
}

func (s *ConstDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclContext) CONST() antlr.TerminalNode {
	return s.GetToken(MyLangParserCONST, 0)
}

func (s *ConstDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLPAREN, 0)
}

func (s *ConstDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserRPAREN, 0)
}

func (s *ConstDeclContext) AllConstField() []IConstFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstFieldContext); ok {
			len++
		}
	}

	tst := make([]IConstFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstFieldContext); ok {
			tst[i] = t.(IConstFieldContext)
			i++
		}
	}

	return tst
}

func (s *ConstDeclContext) ConstField(i int) IConstFieldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstFieldContext)
}

func (s *ConstDeclContext) PUB() antlr.TerminalNode {
	return s.GetToken(MyLangParserPUB, 0)
}

func (s *ConstDeclContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserSEMI)
}

func (s *ConstDeclContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserSEMI, i)
}

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterConstDecl(s)
	}
}

func (s *ConstDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitConstDecl(s)
	}
}

func (s *ConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitConstDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ConstDecl() (localctx IConstDeclContext) {
	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MyLangParserRULE_constDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserPUB {
		{
			p.SetState(240)
			p.Match(MyLangParserPUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(243)
		p.Match(MyLangParserCONST)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyLangParserLPAREN:
		{
			p.SetState(244)
			p.Match(MyLangParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == MyLangParserIDENTIFIER {
			{
				p.SetState(245)
				p.ConstField()
			}
			{
				p.SetState(246)
				p.Match(MyLangParserSEMI)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(253)
			p.Match(MyLangParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MyLangParserIDENTIFIER:
		{
			p.SetState(254)
			p.ConstField()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IConstFieldContext is an interface to support dynamic dispatch.
type IConstFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	IDENTIFIER() antlr.TerminalNode

	// IsConstFieldContext differentiates from other interfaces.
	IsConstFieldContext()
}

type ConstFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyConstFieldContext() *ConstFieldContext {
	var p = new(ConstFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_constField
	return p
}

func InitEmptyConstFieldContext(p *ConstFieldContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_constField
}

func (*ConstFieldContext) IsConstFieldContext() {}

func NewConstFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstFieldContext {
	var p = new(ConstFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_constField

	return p
}

func (s *ConstFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstFieldContext) GetName() antlr.Token { return s.name }


func (s *ConstFieldContext) SetName(v antlr.Token) { s.name = v }


func (s *ConstFieldContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserASSIGN, 0)
}

func (s *ConstFieldContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *ConstFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterConstField(s)
	}
}

func (s *ConstFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitConstField(s)
	}
}

func (s *ConstFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitConstField(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ConstField() (localctx IConstFieldContext) {
	localctx = NewConstFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MyLangParserRULE_constField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*ConstFieldContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(MyLangParserASSIGN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(259)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ArrayType() IArrayTypeContext
	MapType() IMapTypeContext
	OptionType() IOptionTypeContext
	ResultType() IResultTypeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) GetName() antlr.Token { return s.name }


func (s *TypeContext) SetName(v antlr.Token) { s.name = v }


func (s *TypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *TypeContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *TypeContext) MapType() IMapTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *TypeContext) OptionType() IOptionTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionTypeContext)
}

func (s *TypeContext) ResultType() IResultTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultTypeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MyLangParserRULE_type)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)

			var _m = p.Match(MyLangParserIDENTIFIER)

			localctx.(*TypeContext).name = _m
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MyLangParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.ArrayType()
		}


	case MyLangParserMAP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(263)
			p.MapType()
		}


	case MyLangParserQUESTION:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(264)
			p.OptionType()
		}


	case MyLangParserEXCLAMATION:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(265)
			p.ResultType()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Type_() ITypeContext

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACK, 0)
}

func (s *ArrayTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACK, 0)
}

func (s *ArrayTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MyLangParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(MyLangParserLBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(MyLangParserRBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Type_()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	RBRACK() antlr.TerminalNode

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_mapType
	return p
}

func InitEmptyMapTypeContext(p *MapTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_mapType
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) MAP() antlr.TerminalNode {
	return s.GetToken(MyLangParserMAP, 0)
}

func (s *MapTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACK, 0)
}

func (s *MapTypeContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *MapTypeContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MapTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACK, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MyLangParserRULE_mapType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(MyLangParserMAP)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(MyLangParserLBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Type_()
	}
	{
		p.SetState(275)
		p.Match(MyLangParserRBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Type_()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IOptionTypeContext is an interface to support dynamic dispatch.
type IOptionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUESTION() antlr.TerminalNode
	Type_() ITypeContext

	// IsOptionTypeContext differentiates from other interfaces.
	IsOptionTypeContext()
}

type OptionTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionTypeContext() *OptionTypeContext {
	var p = new(OptionTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_optionType
	return p
}

func InitEmptyOptionTypeContext(p *OptionTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_optionType
}

func (*OptionTypeContext) IsOptionTypeContext() {}

func NewOptionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionTypeContext {
	var p = new(OptionTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_optionType

	return p
}

func (s *OptionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionTypeContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(MyLangParserQUESTION, 0)
}

func (s *OptionTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *OptionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterOptionType(s)
	}
}

func (s *OptionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitOptionType(s)
	}
}

func (s *OptionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitOptionType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) OptionType() (localctx IOptionTypeContext) {
	localctx = NewOptionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MyLangParserRULE_optionType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(MyLangParserQUESTION)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Type_()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IResultTypeContext is an interface to support dynamic dispatch.
type IResultTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXCLAMATION() antlr.TerminalNode
	Type_() ITypeContext

	// IsResultTypeContext differentiates from other interfaces.
	IsResultTypeContext()
}

type ResultTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultTypeContext() *ResultTypeContext {
	var p = new(ResultTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_resultType
	return p
}

func InitEmptyResultTypeContext(p *ResultTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_resultType
}

func (*ResultTypeContext) IsResultTypeContext() {}

func NewResultTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultTypeContext {
	var p = new(ResultTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_resultType

	return p
}

func (s *ResultTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultTypeContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(MyLangParserEXCLAMATION, 0)
}

func (s *ResultTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ResultTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ResultTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterResultType(s)
	}
}

func (s *ResultTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitResultType(s)
	}
}

func (s *ResultTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitResultType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ResultType() (localctx IResultTypeContext) {
	localctx = NewResultTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MyLangParserRULE_resultType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(MyLangParserEXCLAMATION)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Type_()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MyLangParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(MyLangParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 352805885815296) != 0) || ((int64((_la - 71)) & ^0x3f) == 0 && ((int64(1) << (_la - 71)) & 31) != 0) {
		{
			p.SetState(285)
			p.Statement()
		}


		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
		p.Match(MyLangParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	Assignment() IAssignmentContext
	ExprStmt() IExprStmtContext
	IfStmt() IIfStmtContext
	ForStmt() IForStmtContext
	MatchStmt() IMatchStmtContext
	ReturnStmt() IReturnStmtContext
	DeferStmt() IDeferStmtContext
	Block() IBlockContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ExprStmt() IExprStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) MatchStmt() IMatchStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchStmtContext)
}

func (s *StatementContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementContext) DeferStmt() IDeferStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeferStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeferStmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MyLangParserRULE_statement)
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.VarDecl()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Assignment()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.ExprStmt()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(296)
			p.IfStmt()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(297)
			p.ForStmt()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(298)
			p.MatchStmt()
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(299)
			p.ReturnStmt()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(300)
			p.DeferStmt()
		}


	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(301)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 


	// Getter signatures
	Expr() IExprContext
	IDENTIFIER() antlr.TerminalNode
	DECL_ASSIGN() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	ASSIGN() antlr.TerminalNode
	MUT() antlr.TerminalNode

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) GetName() antlr.Token { return s.name }


func (s *VarDeclContext) SetName(v antlr.Token) { s.name = v }


func (s *VarDeclContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *VarDeclContext) DECL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserDECL_ASSIGN, 0)
}

func (s *VarDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(MyLangParserCOLON, 0)
}

func (s *VarDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserASSIGN, 0)
}

func (s *VarDeclContext) MUT() antlr.TerminalNode {
	return s.GetToken(MyLangParserMUT, 0)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MyLangParserRULE_varDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserMUT {
		{
			p.SetState(304)
			p.Match(MyLangParserMUT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(307)

		var _m = p.Match(MyLangParserIDENTIFIER)

		localctx.(*VarDeclContext).name = _m
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyLangParserDECL_ASSIGN:
		{
			p.SetState(308)
			p.Match(MyLangParserDECL_ASSIGN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case MyLangParserCOLON:
		{
			p.SetState(309)
			p.Match(MyLangParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Type_()
		}
		{
			p.SetState(311)
			p.Match(MyLangParserASSIGN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(315)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AssignOp() IAssignOpContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllExpr() []IExprContext {
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

func (s *AssignmentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *AssignmentContext) AssignOp() IAssignOpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignOpContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignOpContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MyLangParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.expr(0)
	}
	{
		p.SetState(318)
		p.AssignOp()
	}
	{
		p.SetState(319)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAssignOpContext is an interface to support dynamic dispatch.
type IAssignOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	PLUS_ASSIGN() antlr.TerminalNode
	MINUS_ASSIGN() antlr.TerminalNode
	MULT_ASSIGN() antlr.TerminalNode
	DIV_ASSIGN() antlr.TerminalNode
	MOD_ASSIGN() antlr.TerminalNode
	AND_ASSIGN() antlr.TerminalNode
	OR_ASSIGN() antlr.TerminalNode
	XOR_ASSIGN() antlr.TerminalNode
	LSHIFT_ASSIGN() antlr.TerminalNode
	RSHIFT_ASSIGN() antlr.TerminalNode
	URSHIFT_ASSIGN() antlr.TerminalNode
	LOGICAL_OR_ASSIGN() antlr.TerminalNode
	LOGICAL_AND_ASSIGN() antlr.TerminalNode

	// IsAssignOpContext differentiates from other interfaces.
	IsAssignOpContext()
}

type AssignOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignOpContext() *AssignOpContext {
	var p = new(AssignOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_assignOp
	return p
}

func InitEmptyAssignOpContext(p *AssignOpContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_assignOp
}

func (*AssignOpContext) IsAssignOpContext() {}

func NewAssignOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignOpContext {
	var p = new(AssignOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_assignOp

	return p
}

func (s *AssignOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignOpContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserASSIGN, 0)
}

func (s *AssignOpContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserPLUS_ASSIGN, 0)
}

func (s *AssignOpContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserMINUS_ASSIGN, 0)
}

func (s *AssignOpContext) MULT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserMULT_ASSIGN, 0)
}

func (s *AssignOpContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserDIV_ASSIGN, 0)
}

func (s *AssignOpContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserMOD_ASSIGN, 0)
}

func (s *AssignOpContext) AND_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserAND_ASSIGN, 0)
}

func (s *AssignOpContext) OR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserOR_ASSIGN, 0)
}

func (s *AssignOpContext) XOR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserXOR_ASSIGN, 0)
}

func (s *AssignOpContext) LSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLSHIFT_ASSIGN, 0)
}

func (s *AssignOpContext) RSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserRSHIFT_ASSIGN, 0)
}

func (s *AssignOpContext) URSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserURSHIFT_ASSIGN, 0)
}

func (s *AssignOpContext) LOGICAL_OR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLOGICAL_OR_ASSIGN, 0)
}

func (s *AssignOpContext) LOGICAL_AND_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLOGICAL_AND_ASSIGN, 0)
}

func (s *AssignOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssignOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterAssignOp(s)
	}
}

func (s *AssignOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitAssignOp(s)
	}
}

func (s *AssignOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitAssignOp(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) AssignOp() (localctx IAssignOpContext) {
	localctx = NewAssignOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MyLangParserRULE_assignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 44)) & ^0x3f) == 0 && ((int64(1) << (_la - 44)) & 16775169) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_exprStmt
	return p
}

func InitEmptyExprStmtContext(p *ExprStmtContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_exprStmt
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MyLangParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	IfStmt() IIfStmtContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(MyLangParserIF, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MyLangParserELSE, 0)
}

func (s *IfStmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MyLangParserRULE_ifStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(MyLangParserIF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(326)
		p.expr(0)
	}
	{
		p.SetState(327)
		p.Block()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(328)
			p.Match(MyLangParserELSE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MyLangParserIF:
			{
				p.SetState(329)
				p.IfStmt()
			}


		case MyLangParserLBRACE:
			{
				p.SetState(330)
				p.Block()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Block() IBlockContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	IN() antlr.TerminalNode

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(MyLangParserFOR, 0)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) AllExpr() []IExprContext {
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

func (s *ForStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *ForStmtContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserSEMI)
}

func (s *ForStmtContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserSEMI, i)
}

func (s *ForStmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *ForStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(MyLangParserIN, 0)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MyLangParserRULE_forStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(MyLangParserFOR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(336)
			p.Block()
		}


	case 2:
		{
			p.SetState(337)
			p.expr(0)
		}
		{
			p.SetState(338)
			p.Block()
		}


	case 3:
		{
			p.SetState(340)
			p.expr(0)
		}
		{
			p.SetState(341)
			p.Match(MyLangParserSEMI)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(342)
			p.expr(0)
		}
		{
			p.SetState(343)
			p.Match(MyLangParserSEMI)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(344)
			p.expr(0)
		}

		{
			p.SetState(346)
			p.Block()
		}


	case 4:
		{
			p.SetState(348)
			p.Match(MyLangParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Match(MyLangParserIN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(350)
			p.expr(0)
		}

		{
			p.SetState(352)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMatchStmtContext is an interface to support dynamic dispatch.
type IMatchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MATCH() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllMatchBranch() []IMatchBranchContext
	MatchBranch(i int) IMatchBranchContext

	// IsMatchStmtContext differentiates from other interfaces.
	IsMatchStmtContext()
}

type MatchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchStmtContext() *MatchStmtContext {
	var p = new(MatchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_matchStmt
	return p
}

func InitEmptyMatchStmtContext(p *MatchStmtContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_matchStmt
}

func (*MatchStmtContext) IsMatchStmtContext() {}

func NewMatchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchStmtContext {
	var p = new(MatchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_matchStmt

	return p
}

func (s *MatchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchStmtContext) MATCH() antlr.TerminalNode {
	return s.GetToken(MyLangParserMATCH, 0)
}

func (s *MatchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *MatchStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *MatchStmtContext) AllMatchBranch() []IMatchBranchContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatchBranchContext); ok {
			len++
		}
	}

	tst := make([]IMatchBranchContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatchBranchContext); ok {
			tst[i] = t.(IMatchBranchContext)
			i++
		}
	}

	return tst
}

func (s *MatchStmtContext) MatchBranch(i int) IMatchBranchContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchBranchContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchBranchContext)
}

func (s *MatchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MatchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterMatchStmt(s)
	}
}

func (s *MatchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitMatchStmt(s)
	}
}

func (s *MatchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitMatchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) MatchStmt() (localctx IMatchStmtContext) {
	localctx = NewMatchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MyLangParserRULE_matchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(MyLangParserMATCH)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(357)
		p.expr(0)
	}
	{
		p.SetState(358)
		p.Match(MyLangParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64((_la - 12)) & ^0x3f) == 0 && ((int64(1) << (_la - 12)) & -576460666169174079) != 0) {
		{
			p.SetState(359)
			p.MatchBranch()
		}


		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(365)
		p.Match(MyLangParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMatchBranchContext is an interface to support dynamic dispatch.
type IMatchBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARROW() antlr.TerminalNode
	Statement() IStatementContext
	Expr() IExprContext
	ELSE() antlr.TerminalNode

	// IsMatchBranchContext differentiates from other interfaces.
	IsMatchBranchContext()
}

type MatchBranchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchBranchContext() *MatchBranchContext {
	var p = new(MatchBranchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_matchBranch
	return p
}

func InitEmptyMatchBranchContext(p *MatchBranchContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_matchBranch
}

func (*MatchBranchContext) IsMatchBranchContext() {}

func NewMatchBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchBranchContext {
	var p = new(MatchBranchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_matchBranch

	return p
}

func (s *MatchBranchContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchBranchContext) ARROW() antlr.TerminalNode {
	return s.GetToken(MyLangParserARROW, 0)
}

func (s *MatchBranchContext) Statement() IStatementContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *MatchBranchContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchBranchContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MyLangParserELSE, 0)
}

func (s *MatchBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchBranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MatchBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterMatchBranch(s)
	}
}

func (s *MatchBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitMatchBranch(s)
	}
}

func (s *MatchBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitMatchBranch(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) MatchBranch() (localctx IMatchBranchContext) {
	localctx = NewMatchBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MyLangParserRULE_matchBranch)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MyLangParserTRUE, MyLangParserFALSE, MyLangParserNONE, MyLangParserMAP, MyLangParserLPAREN, MyLangParserLBRACE, MyLangParserLBRACK, MyLangParserNOT, MyLangParserTILDE, MyLangParserQUESTION, MyLangParserEXCLAMATION, MyLangParserMINUS, MyLangParserIDENTIFIER, MyLangParserINT, MyLangParserFLOAT, MyLangParserSTRING, MyLangParserCHAR:
		{
			p.SetState(367)
			p.expr(0)
		}


	case MyLangParserELSE:
		{
			p.SetState(368)
			p.Match(MyLangParserELSE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(371)
		p.Match(MyLangParserARROW)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(372)
		p.Statement()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MyLangParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MyLangParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(MyLangParserRETURN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(375)
			p.expr(0)
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDeferStmtContext is an interface to support dynamic dispatch.
type IDeferStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFER() antlr.TerminalNode
	Block() IBlockContext

	// IsDeferStmtContext differentiates from other interfaces.
	IsDeferStmtContext()
}

type DeferStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeferStmtContext() *DeferStmtContext {
	var p = new(DeferStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_deferStmt
	return p
}

func InitEmptyDeferStmtContext(p *DeferStmtContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_deferStmt
}

func (*DeferStmtContext) IsDeferStmtContext() {}

func NewDeferStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeferStmtContext {
	var p = new(DeferStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_deferStmt

	return p
}

func (s *DeferStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeferStmtContext) DEFER() antlr.TerminalNode {
	return s.GetToken(MyLangParserDEFER, 0)
}

func (s *DeferStmtContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DeferStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeferStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DeferStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterDeferStmt(s)
	}
}

func (s *DeferStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitDeferStmt(s)
	}
}

func (s *DeferStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitDeferStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) DeferStmt() (localctx IDeferStmtContext) {
	localctx = NewDeferStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MyLangParserRULE_deferStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(MyLangParserDEFER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Block()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ExprList() IExprListContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	NOT() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	TILDE() antlr.TerminalNode
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode
	URSHIFT() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	AND() antlr.TerminalNode
	XOR() antlr.TerminalNode
	OR() antlr.TerminalNode
	LOGICAL_AND() antlr.TerminalNode
	LOGICAL_OR() antlr.TerminalNode
	QUESTION() antlr.TerminalNode
	COLON() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	DOT() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Primary() IPrimaryContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserRPAREN, 0)
}

func (s *ExprContext) ExprList() IExprListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(MyLangParserNOT, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MyLangParserMINUS, 0)
}

func (s *ExprContext) TILDE() antlr.TerminalNode {
	return s.GetToken(MyLangParserTILDE, 0)
}

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(MyLangParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(MyLangParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(MyLangParserMOD, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MyLangParserPLUS, 0)
}

func (s *ExprContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(MyLangParserLSHIFT, 0)
}

func (s *ExprContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(MyLangParserRSHIFT, 0)
}

func (s *ExprContext) URSHIFT() antlr.TerminalNode {
	return s.GetToken(MyLangParserURSHIFT, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(MyLangParserLT, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(MyLangParserGT, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLE, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(MyLangParserGE, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(MyLangParserEQ, 0)
}

func (s *ExprContext) NE() antlr.TerminalNode {
	return s.GetToken(MyLangParserNE, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(MyLangParserAND, 0)
}

func (s *ExprContext) XOR() antlr.TerminalNode {
	return s.GetToken(MyLangParserXOR, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(MyLangParserOR, 0)
}

func (s *ExprContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(MyLangParserLOGICAL_AND, 0)
}

func (s *ExprContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(MyLangParserLOGICAL_OR, 0)
}

func (s *ExprContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(MyLangParserQUESTION, 0)
}

func (s *ExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(MyLangParserCOLON, 0)
}

func (s *ExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MyLangParserASSIGN, 0)
}

func (s *ExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(MyLangParserDOT, 0)
}

func (s *ExprContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACK, 0)
}

func (s *ExprContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACK, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *MyLangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MyLangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, MyLangParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(382)
			p.Primary()
		}


	case 2:
		{
			p.SetState(383)
			p.Match(MyLangParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(384)
			p.Match(MyLangParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64((_la - 18)) & ^0x3f) == 0 && ((int64(1) << (_la - 18)) & 279223178242818399) != 0) {
			{
				p.SetState(385)
				p.ExprList()
			}

		}
		{
			p.SetState(388)
			p.Match(MyLangParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		{
			p.SetState(389)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 281887293571072) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(390)
			p.expr(13)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(448)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(393)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(394)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3940649673949184) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(395)
					p.expr(13)
				}


			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(396)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(397)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MyLangParserPLUS || _la == MyLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(398)
					p.expr(12)
				}


			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(399)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(400)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 31525197391593472) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(401)
					p.expr(11)
				}


			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(402)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(403)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 4026531840) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(404)
					p.expr(10)
				}


			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(406)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MyLangParserEQ || _la == MyLangParserNE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(407)
					p.expr(9)
				}


			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(409)
					p.Match(MyLangParserAND)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(410)
					p.expr(8)
				}


			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(412)
					p.Match(MyLangParserXOR)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(413)
					p.expr(7)
				}


			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(415)
					p.Match(MyLangParserOR)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(416)
					p.expr(6)
				}


			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(418)
					p.Match(MyLangParserLOGICAL_AND)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(419)
					p.expr(5)
				}


			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(421)
					p.Match(MyLangParserLOGICAL_OR)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(422)
					p.expr(4)
				}


			case 11:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(423)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(424)
					p.Match(MyLangParserQUESTION)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(425)
					p.expr(0)
				}
				{
					p.SetState(426)
					p.Match(MyLangParserCOLON)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(427)
					p.expr(3)
				}


			case 12:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(429)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(430)
					p.Match(MyLangParserASSIGN)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(431)
					p.expr(1)
				}


			case 13:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(432)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(433)
					p.Match(MyLangParserDOT)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(434)
					p.Match(MyLangParserIDENTIFIER)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 14:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(435)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(436)
					p.Match(MyLangParserDOT)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(437)
					p.Match(MyLangParserIDENTIFIER)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(438)
					p.Match(MyLangParserLPAREN)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				p.SetState(440)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)


				if ((int64((_la - 18)) & ^0x3f) == 0 && ((int64(1) << (_la - 18)) & 279223178242818399) != 0) {
					{
						p.SetState(439)
						p.ExprList()
					}

				}
				{
					p.SetState(442)
					p.Match(MyLangParserRPAREN)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 15:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MyLangParserRULE_expr)
				p.SetState(443)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(444)
					p.Match(MyLangParserLBRACK)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(445)
					p.expr(0)
				}
				{
					p.SetState(446)
					p.Match(MyLangParserRBRACK)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



	errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Literal() ILiteralContext
	ArrayLiteral() IArrayLiteralContext
	MapLiteral() IMapLiteralContext
	StructLiteral() IStructLiteralContext
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *PrimaryContext) MapLiteral() IMapLiteralContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapLiteralContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapLiteralContext)
}

func (s *PrimaryContext) StructLiteral() IStructLiteralContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructLiteralContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructLiteralContext)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserLPAREN, 0)
}

func (s *PrimaryContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MyLangParserRPAREN, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MyLangParserRULE_primary)
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(453)
			p.Match(MyLangParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(454)
			p.Literal()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(455)
			p.ArrayLiteral()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(456)
			p.MapLiteral()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(457)
			p.StructLiteral()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(458)
			p.Match(MyLangParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(459)
			p.expr(0)
		}
		{
			p.SetState(460)
			p.Match(MyLangParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NONE() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(MyLangParserINT, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(MyLangParserFLOAT, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(MyLangParserSTRING, 0)
}

func (s *LiteralContext) CHAR() antlr.TerminalNode {
	return s.GetToken(MyLangParserCHAR, 0)
}

func (s *LiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MyLangParserTRUE, 0)
}

func (s *LiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MyLangParserFALSE, 0)
}

func (s *LiteralContext) NONE() antlr.TerminalNode {
	return s.GetToken(MyLangParserNONE, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MyLangParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 18)) & ^0x3f) == 0 && ((int64(1) << (_la - 18)) & 270215977642229767) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	ExprList() IExprListContext
	Type_() ITypeContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACK, 0)
}

func (s *ArrayLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACK, 0)
}

func (s *ArrayLiteralContext) ExprList() IExprListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ArrayLiteralContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArrayLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *ArrayLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MyLangParserRULE_arrayLiteral)
	var _la int

	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(466)
			p.Match(MyLangParserLBRACK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64((_la - 18)) & ^0x3f) == 0 && ((int64(1) << (_la - 18)) & 279223178242818399) != 0) {
			{
				p.SetState(467)
				p.ExprList()
			}

		}
		{
			p.SetState(470)
			p.Match(MyLangParserRBRACK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(471)
			p.Match(MyLangParserLBRACK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(472)
			p.Type_()
		}
		{
			p.SetState(473)
			p.Match(MyLangParserRBRACK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(474)
			p.Match(MyLangParserLBRACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64((_la - 18)) & ^0x3f) == 0 && ((int64(1) << (_la - 18)) & 279223178242818399) != 0) {
			{
				p.SetState(475)
				p.ExprList()
			}

		}
		{
			p.SetState(478)
			p.Match(MyLangParserRBRACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMapLiteralContext is an interface to support dynamic dispatch.
type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllMapElement() []IMapElementContext
	MapElement(i int) IMapElementContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}

type MapLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapLiteralContext() *MapLiteralContext {
	var p = new(MapLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_mapLiteral
	return p
}

func InitEmptyMapLiteralContext(p *MapLiteralContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_mapLiteral
}

func (*MapLiteralContext) IsMapLiteralContext() {}

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext {
	var p = new(MapLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_mapLiteral

	return p
}

func (s *MapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *MapLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *MapLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *MapLiteralContext) AllMapElement() []IMapElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapElementContext); ok {
			len++
		}
	}

	tst := make([]IMapElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapElementContext); ok {
			tst[i] = t.(IMapElementContext)
			i++
		}
	}

	return tst
}

func (s *MapLiteralContext) MapElement(i int) IMapElementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapElementContext)
}

func (s *MapLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserCOMMA)
}

func (s *MapLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserCOMMA, i)
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterMapLiteral(s)
	}
}

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitMapLiteral(s)
	}
}

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitMapLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) MapLiteral() (localctx IMapLiteralContext) {
	localctx = NewMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MyLangParserRULE_mapLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(482)
		p.Match(MyLangParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64((_la - 18)) & ^0x3f) == 0 && ((int64(1) << (_la - 18)) & 279223178242818399) != 0) {
		{
			p.SetState(483)
			p.MapElement()
		}
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == MyLangParserCOMMA {
			{
				p.SetState(484)
				p.Match(MyLangParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(485)
				p.MapElement()
			}


			p.SetState(490)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(493)
		p.Match(MyLangParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMapElementContext is an interface to support dynamic dispatch.
type IMapElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COLON() antlr.TerminalNode

	// IsMapElementContext differentiates from other interfaces.
	IsMapElementContext()
}

type MapElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapElementContext() *MapElementContext {
	var p = new(MapElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_mapElement
	return p
}

func InitEmptyMapElementContext(p *MapElementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_mapElement
}

func (*MapElementContext) IsMapElementContext() {}

func NewMapElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapElementContext {
	var p = new(MapElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_mapElement

	return p
}

func (s *MapElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MapElementContext) AllExpr() []IExprContext {
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

func (s *MapElementContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *MapElementContext) COLON() antlr.TerminalNode {
	return s.GetToken(MyLangParserCOLON, 0)
}

func (s *MapElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MapElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterMapElement(s)
	}
}

func (s *MapElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitMapElement(s)
	}
}

func (s *MapElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitMapElement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) MapElement() (localctx IMapElementContext) {
	localctx = NewMapElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MyLangParserRULE_mapElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.expr(0)
	}
	{
		p.SetState(496)
		p.Match(MyLangParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(497)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStructLiteralContext is an interface to support dynamic dispatch.
type IStructLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStructElement() []IStructElementContext
	StructElement(i int) IStructElementContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructLiteralContext differentiates from other interfaces.
	IsStructLiteralContext()
}

type StructLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructLiteralContext() *StructLiteralContext {
	var p = new(StructLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structLiteral
	return p
}

func InitEmptyStructLiteralContext(p *StructLiteralContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structLiteral
}

func (*StructLiteralContext) IsStructLiteralContext() {}

func NewStructLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructLiteralContext {
	var p = new(StructLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_structLiteral

	return p
}

func (s *StructLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StructLiteralContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserLBRACE, 0)
}

func (s *StructLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MyLangParserRBRACE, 0)
}

func (s *StructLiteralContext) AllStructElement() []IStructElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructElementContext); ok {
			len++
		}
	}

	tst := make([]IStructElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructElementContext); ok {
			tst[i] = t.(IStructElementContext)
			i++
		}
	}

	return tst
}

func (s *StructLiteralContext) StructElement(i int) IStructElementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructElementContext)
}

func (s *StructLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserCOMMA)
}

func (s *StructLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserCOMMA, i)
}

func (s *StructLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterStructLiteral(s)
	}
}

func (s *StructLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitStructLiteral(s)
	}
}

func (s *StructLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitStructLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) StructLiteral() (localctx IStructLiteralContext) {
	localctx = NewStructLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, MyLangParserRULE_structLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		p.Type_()
	}
	{
		p.SetState(500)
		p.Match(MyLangParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MyLangParserIDENTIFIER {
		{
			p.SetState(501)
			p.StructElement()
		}
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == MyLangParserCOMMA {
			{
				p.SetState(502)
				p.Match(MyLangParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(503)
				p.StructElement()
			}


			p.SetState(508)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(511)
		p.Match(MyLangParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStructElementContext is an interface to support dynamic dispatch.
type IStructElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expr() IExprContext

	// IsStructElementContext differentiates from other interfaces.
	IsStructElementContext()
}

type StructElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructElementContext() *StructElementContext {
	var p = new(StructElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structElement
	return p
}

func InitEmptyStructElementContext(p *StructElementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_structElement
}

func (*StructElementContext) IsStructElementContext() {}

func NewStructElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructElementContext {
	var p = new(StructElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_structElement

	return p
}

func (s *StructElementContext) GetParser() antlr.Parser { return s.parser }

func (s *StructElementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MyLangParserIDENTIFIER, 0)
}

func (s *StructElementContext) COLON() antlr.TerminalNode {
	return s.GetToken(MyLangParserCOLON, 0)
}

func (s *StructElementContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterStructElement(s)
	}
}

func (s *StructElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitStructElement(s)
	}
}

func (s *StructElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitStructElement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) StructElement() (localctx IStructElementContext) {
	localctx = NewStructElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, MyLangParserRULE_structElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		p.Match(MyLangParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(514)
		p.Match(MyLangParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(515)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
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

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MyLangParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MyLangParserCOMMA, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, MyLangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.expr(0)
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MyLangParserCOMMA {
		{
			p.SetState(518)
			p.Match(MyLangParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(519)
			p.expr(0)
		}


		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_returnType
	return p
}

func InitEmptyReturnTypeContext(p *ReturnTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MyLangParserRULE_returnType
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MyLangParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MyLangListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (s *ReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MyLangVisitor:
		return t.VisitReturnType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MyLangParser) ReturnType() (localctx IReturnTypeContext) {
	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, MyLangParserRULE_returnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(525)
		p.Type_()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


func (p *MyLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 34:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MyLangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
			return p.Precpred(p.GetParserRuleContext(), 1)

	case 12:
			return p.Precpred(p.GetParserRuleContext(), 17)

	case 13:
			return p.Precpred(p.GetParserRuleContext(), 16)

	case 14:
			return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

