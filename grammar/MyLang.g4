grammar MyLang;

// Parser Rules
program
    : moduleDecl? (topLevelDecl | statement)* EOF
    ;

moduleDecl
    : MODULE name=IDENTIFIER
    ;

topLevelDecl
    : functionDecl
    | structDecl
    | interfaceDecl
    | enumDecl
    | constDecl
    | importDecl
    ;

importDecl
    : IMPORT importSpec
    ;

importSpec
    : name=IDENTIFIER
    | name=IDENTIFIER DOT subname=IDENTIFIER
    | LBRACE IDENTIFIER (COMMA IDENTIFIER)* RBRACE
    ;

functionDecl
    : PUB? FN name=IDENTIFIER genericParams? LPAREN params? RPAREN returnType? block
    ;

genericParams
    : LT IDENTIFIER (COMMA IDENTIFIER)* GT
    ;

params
    : param (COMMA param)*
    ;

param
    : MUT? name=IDENTIFIER type
    ;

structDecl
    : PUB? STRUCT name=IDENTIFIER LBRACE structField* RBRACE
    ;

structField
    : PUB? MUT? name=IDENTIFIER type
    ;

interfaceDecl
    : PUB? INTERFACE name=IDENTIFIER LBRACE interfaceMethod* RBRACE
    ;

interfaceMethod
    : name=IDENTIFIER LPAREN params? RPAREN returnType?
    ;

enumDecl
    : PUB? ENUM name=IDENTIFIER LBRACE enumField (COMMA enumField)* RBRACE
    ;

enumField
    : name=IDENTIFIER (ASSIGN expr)?
    ;

constDecl
    : PUB? CONST (
        LPAREN (constField SEMI)* RPAREN
        | constField
    )
    ;

constField
    : name=IDENTIFIER ASSIGN expr
    ;

type
    : name=IDENTIFIER
    | arrayType
    | mapType
    | optionType
    | resultType
    ;

arrayType
    : LBRACK RBRACK type
    ;

mapType
    : MAP LBRACK type RBRACK type
    ;

optionType
    : QUESTION type
    ;

resultType
    : EXCLAMATION type
    ;

block
    : LBRACE statement* RBRACE
    ;

statement
    : varDecl
    | assignment
    | exprStmt
    | ifStmt
    | forStmt
    | matchStmt
    | returnStmt
    | deferStmt
    | block
    ;

varDecl
    : MUT? name=IDENTIFIER (DECL_ASSIGN | COLON type ASSIGN) expr
    ;

assignment
    : expr assignOp expr
    ;

assignOp
    : ASSIGN | PLUS_ASSIGN | MINUS_ASSIGN | MULT_ASSIGN | DIV_ASSIGN | MOD_ASSIGN
    | AND_ASSIGN | OR_ASSIGN | XOR_ASSIGN | LSHIFT_ASSIGN | RSHIFT_ASSIGN | URSHIFT_ASSIGN
    | LOGICAL_OR_ASSIGN | LOGICAL_AND_ASSIGN
    ;

exprStmt
    : expr
    ;

ifStmt
    : IF expr block (ELSE (ifStmt | block))?
    ;

forStmt
    : FOR (
        block                                  // infinite loop
        | expr block                           // while loop
        | (expr SEMI expr SEMI expr) block     // C-style for
        | (IDENTIFIER IN expr) block           // for-in loop
    )
    ;

matchStmt
    : MATCH expr LBRACE matchBranch* RBRACE
    ;

matchBranch
    : (expr | ELSE) ARROW statement
    ;

returnStmt
    : RETURN expr?
    ;

deferStmt
    : DEFER block
    ;

expr
    : primary
    | expr DOT IDENTIFIER                          // field access
    | expr DOT IDENTIFIER LPAREN exprList? RPAREN  // method call
    | IDENTIFIER LPAREN exprList? RPAREN           // function call
    | expr LBRACK expr RBRACK                      // array/map access
    | (NOT | MINUS | TILDE) expr                   // unary operators
    | expr (MULT | DIV | MOD) expr                 // multiplicative
    | expr (PLUS | MINUS) expr                     // additive
    | expr (LSHIFT | RSHIFT | URSHIFT) expr        // shift
    | expr (LT | GT | LE | GE) expr                // relational
    | expr (EQ | NE) expr                          // equality
    | expr AND expr                                // bitwise AND
    | expr XOR expr                                // bitwise XOR
    | expr OR expr                                 // bitwise OR
    | expr LOGICAL_AND expr                        // logical AND
    | expr LOGICAL_OR expr                         // logical OR
    | expr QUESTION expr COLON expr                // ternary
    | <assoc=right> expr ASSIGN expr               // assignment
    ;

primary
    : IDENTIFIER
    | literal
    | arrayLiteral
    | mapLiteral
    | structLiteral
    | LPAREN expr RPAREN
    ;

literal
    : INT
    | FLOAT
    | STRING
    | CHAR
    | TRUE
    | FALSE
    | NONE
    ;

arrayLiteral
    : LBRACK exprList? RBRACK
    | LBRACK type RBRACK LBRACE exprList? RBRACE
    ;

mapLiteral
    : LBRACE (mapElement (COMMA mapElement)*)? RBRACE
    ;

mapElement
    : expr COLON expr
    ;

structLiteral
    : type LBRACE (structElement (COMMA structElement)*)? RBRACE
    ;

structElement
    : IDENTIFIER COLON expr
    ;

exprList
    : expr (COMMA expr)*
    ;

returnType
    : type
    ;

// Lexer Rules
MODULE: 'module';
IMPORT: 'import';
PUB: 'pub';
FN: 'fn';
STRUCT: 'struct';
INTERFACE: 'interface';
ENUM: 'enum';
CONST: 'const';
MUT: 'mut';
TYPE: 'type';
IF: 'if';
ELSE: 'else';
FOR: 'for';
IN: 'in';
MATCH: 'match';
RETURN: 'return';
DEFER: 'defer';
TRUE: 'true';
FALSE: 'false';
NONE: 'none';
MAP: 'map';

// Operators and punctuation
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACK: '[';
RBRACK: ']';
LT: '<';
GT: '>';
LE: '<=';
GE: '>=';
EQ: '==';
NE: '!=';
AND: '&';
OR: '|';
XOR: '^';
NOT: '!';
TILDE: '~';
QUESTION: '?';
COLON: ':';
SEMI: ';';
COMMA: ',';
DOT: '.';
ASSIGN: '=';
ARROW: '=>';
EXCLAMATION: '!';

// Arithmetic operators
PLUS: '+';
MINUS: '-';
MULT: '*';
DIV: '/';
MOD: '%';
LSHIFT: '<<';
RSHIFT: '>>';
URSHIFT: '>>>';

// Assignment operators
PLUS_ASSIGN: '+=';
MINUS_ASSIGN: '-=';
MULT_ASSIGN: '*=';
DIV_ASSIGN: '/=';
MOD_ASSIGN: '%=';
AND_ASSIGN: '&=';
OR_ASSIGN: '|=';
XOR_ASSIGN: '^=';
LSHIFT_ASSIGN: '<<=';
RSHIFT_ASSIGN: '>>=';
URSHIFT_ASSIGN: '>>>=';
LOGICAL_OR_ASSIGN: '||=';
LOGICAL_AND_ASSIGN: '&&=';

// Logical operators
LOGICAL_OR: '||';
LOGICAL_AND: '&&';
DECL_ASSIGN: ':=';

// Identifiers and literals
IDENTIFIER
    : [a-zA-Z_][a-zA-Z0-9_]*
    ;

INT
    : [0-9]+
    | '0x' [0-9a-fA-F]+
    | '0o' [0-7]+
    | '0b' [0-1]+
    | [0-9]+ '_' [0-9]+
    ;

FLOAT
    : [0-9]+ '.' [0-9]*
    | '.' [0-9]+
    | [0-9]+ ('.' [0-9]*)? [eE] [+-]? [0-9]+
    ;

STRING
    : '"' (~["\\] | '\\' .)* '"'
    | '`' (~[`])* '`'
    ;

CHAR
    : '\'' (~[']) '\''
    ;

WHITESPACE
    : [ \t\r\n]+ -> skip
    ;

SINGLE_LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;

MULTI_LINE_COMMENT
    : '/*' .*? '*/' -> skip
    ;
