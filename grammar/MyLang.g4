grammar MyLang;

// Parser Rules
program
    : statement* EOF
    ;

statement
    : variableDeclaration
    | assignmentStatement
    | expressionStatement
    | printStatement
    | ifStatement
    | whileStatement
    | doWhileStatement
    | forStatement
    | switchStatement
    | breakStatement
    | continueStatement
    | functionDeclaration
    | returnStatement
    | block
    ;

block
    : '{' statement* '}'
    ;

variableDeclaration
    : 'var' IDENTIFIER ('=' expression)? ';'
    ;

assignmentStatement
    : IDENTIFIER '=' expression ';'
    | IDENTIFIER '[' expression ']' '=' expression ';'  // Array assignment
    ;

expressionStatement
    : expression ';'
    ;

printStatement
    : 'print' '(' expression ')' ';'
    ;

ifStatement
    : 'if' '(' expression ')' statement ('else' statement)?
    ;

whileStatement
    : 'while' '(' expression ')' statement
    ;

doWhileStatement
    : 'do' statement 'while' '(' expression ')' ';'
    ;

forStatement
    : 'for' '(' (variableDeclaration | expressionStatement | ';') expression? ';' expression? ')' statement
    ;

switchStatement
    : 'switch' '(' expression ')' '{' switchCase* defaultCase? '}'
    ;

switchCase
    : 'case' expression ':' statement*
    ;

defaultCase
    : 'default' ':' statement*
    ;

breakStatement
    : 'break' ';'
    ;

continueStatement
    : 'continue' ';'
    ;

functionDeclaration
    : 'function' IDENTIFIER '(' parameterList? ')' block
    ;

parameterList
    : IDENTIFIER (',' IDENTIFIER)*
    ;

returnStatement
    : 'return' expression? ';'
    ;

expression
    : assignment
    ;

assignment
    : logicalOr
    ;

logicalOr
    : logicalAnd ('||' logicalAnd)*
    ;

logicalAnd
    : equality ('&&' equality)*
    ;

equality
    : comparison (('==' | '!=') comparison)*
    ;

comparison
    : addition (('<' | '>' | '<=' | '>=') addition)*
    ;

addition
    : multiplication (('+' | '-') multiplication)*
    ;

multiplication
    : unary (('*' | '/' | '%') unary)*
    ;

unary
    : ('!' | '-') unary
    | primary
    ;

primary
    : '(' expression ')'
    | functionCall
    | arrayAccess
    | arrayLiteral
    | NUMBER
    | STRING
    | 'true'
    | 'false'
    | IDENTIFIER
    ;

functionCall
    : IDENTIFIER '(' argumentList? ')'
    ;

argumentList
    : expression (',' expression)*
    ;

arrayAccess
    : IDENTIFIER '[' expression ']'
    ;

arrayLiteral
    : '[' (expression (',' expression)*)? ']'
    ;

// Lexer Rules
NUMBER
    : INT ('.' [0-9]+)?
    ;

INT
    : '0' | [1-9][0-9]*
    ;

STRING
    : '"' (~["\\] | '\\' .)* '"'
    ;

IDENTIFIER
    : [a-zA-Z_][a-zA-Z_0-9]*
    ;

COMMENT
    : '//' ~[\r\n]* -> skip
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> skip
    ;

WHITESPACE
    : [ \t\r\n]+ -> skip
    ;
