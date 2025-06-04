grammar MyLang;

// Parser Rules for V language
program
    : moduleDeclaration? importDeclaration* statement* EOF
    ;

moduleDeclaration
    : 'module' moduleName
    ;

moduleName
    : IDENTIFIER ('.' IDENTIFIER)*
    ;

importDeclaration
    : 'import' importPath
    ;

importPath
    : IDENTIFIER ('.' IDENTIFIER)* 
    ;

statement
    : variableDeclaration
    | constDeclaration
    | assignmentStatement
    | expressionStatement
    | printStatement
    | ifStatement
    | matchStatement
    | forStatement
    | forInStatement
    | forCStatement
    | breakStatement
    | continueStatement
    | functionDeclaration
    | structDeclaration
    | enumDeclaration
    | interfaceDeclaration
    | returnStatement
    | deferStatement
    | block
    ;


block
    : '{' statement* '}'
    ;

variableDeclaration
    : ('mut')? IDENTIFIER (':' type)? ('=' expression)? ';'
    ;

constDeclaration
    : 'const' (
        '(' (IDENTIFIER '=' expression ';')* ')'
        | IDENTIFIER '=' expression ';'
    )
    ;

type
    : simpleType
    | arrayType
    | mapType
    | optionType
    | structType
    ;

simpleType
    : IDENTIFIER
    ;

arrayType
    : '[]' type
    ;

mapType
    : 'map[' type ']' type
    ;

optionType
    : '?' type
    ;
    
structType
    : 'struct' '{' structField* '}'
    ;
    
structField
    : IDENTIFIER type ';'
    ;

assignmentStatement
    : IDENTIFIER ('.' IDENTIFIER)* '=' expression ';'                        // Standard assignment
    | IDENTIFIER '[' expression ']' '=' expression ';'                       // Array assignment
    | IDENTIFIER ('.' IDENTIFIER)* ':=' expression ';'                       // Declaration assignment
    | IDENTIFIER ('.' IDENTIFIER)* ('+=' | '-=' | '*=' | '/=') expression ';' // Compound assignment
    ;

expressionStatement
    : expression ';'
    ;

printStatement
    : 'println' '(' expression ')' ';'
    | 'print' '(' expression ')' ';'
    | 'eprintln' '(' expression ')' ';'
    | 'eprint' '(' expression ')' ';'
    ;

ifStatement
    : 'if' expression statement ('else' ifStatement | 'else' statement)?
    ;

matchStatement
    : 'match' expression '{' matchBranch* '}'
    ;

matchBranch
    : matchPattern '{' statement* '}'
    ;

matchPattern
    : expression
    | expression ',' matchPattern
    | 'else'
    ;

forStatement
    : 'for' '{' statement* '}' // Infinite loop
    ;

forInStatement
    : 'for' IDENTIFIER (',' IDENTIFIER)* 'in' expression '{' statement* '}'
    ;

forCStatement
    : 'for' expression '{' statement* '}' // Condition-only loop (while loop)
    ;

deferStatement
    : 'defer' (block | expression ';')
    ;

breakStatement
    : 'break' ';'
    ;

continueStatement
    : 'continue' ';'
    ;

functionDeclaration
    : ('pub')? 'fn' IDENTIFIER '(' parameterList? ')' returnType? block
    ;

parameterList
    : parameter (',' parameter)*
    ;

parameter
    : ('mut')? IDENTIFIER type
    ;

returnType
    : type
    ;

structDeclaration
    : ('pub')? 'struct' IDENTIFIER '{' structField* '}'
    ;

enumDeclaration
    : ('pub')? 'enum' IDENTIFIER '{' enumValue (',' enumValue)* '}'
    ;

enumValue
    : IDENTIFIER
    | IDENTIFIER '=' expression
    ;

interfaceDeclaration
    : ('pub')? 'interface' IDENTIFIER '{' interfaceMethod* '}'
    ;

interfaceMethod
    : IDENTIFIER '(' parameterList? ')' returnType? ';'
    ;

returnStatement
    : 'return' expression? ';'
    ;

expression
    : assignment
    ;

assignment
    : logicalOr
    | IDENTIFIER (':=' | '=' | '+=' | '-=' | '*=' | '/=' | '%=') expression
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
    | methodCall
    | arrayAccess
    | arrayLiteral
    | mapLiteral
    | selectorExpression
    | optionCheck
    | NUMBER
    | STRING
    | 'true'
    | 'false'
    | 'none'
    | IDENTIFIER
    ;
    
selectorExpression
    : primary '.' IDENTIFIER
    ;
    
methodCall
    : primary '.' IDENTIFIER '(' argumentList? ')'
    ;
    
optionCheck
    : primary '?' // exists check
    | primary 'or' '{' statement* '}' // unwrap with alternative
    | primary 'or' expression // unwrap with alternative expression
    ;

functionCall
    : IDENTIFIER '(' argumentList? ')'
    ;

argumentList
    : expression (',' expression)*
    | namedArgument (',' namedArgument)*
    ;

namedArgument
    : IDENTIFIER ':' expression
    ;

arrayAccess
    : primary '[' expression ']'
    ;

arrayLiteral
    : '[' (expression (',' expression)*)? ']'
    | '[' type ']{' (expression (',' expression)*)? '}'
    ;

mapLiteral
    : '{' (mapEntry (',' mapEntry)*)? '}'
    | 'map[' type ']' type '{' (mapEntry (',' mapEntry)*)? '}'
    ;

mapEntry
    : expression ':' expression
    ;

// Lexer Rules
NUMBER
    : INT ('.' [0-9]+)?
    | INT ('.' [0-9]+)? [eE] [+-]? INT  // Scientific notation
    | '0x' [0-9a-fA-F]+                  // Hex notation
    | '0o' [0-7]+                        // Octal notation
    | '0b' [01]+                         // Binary notation
    ;

INT
    : '0' | [1-9][0-9]* | [0-9]+ '_' [0-9]+ // Support for numbers with underscores like 1_000_000
    ;

STRING
    : '"' (~["\\] | '\\' .)* '"'     // Regular string
    | '`' (~[`])* '`'                    // Raw string
    | "'" (~[']) "'"                     // Character/rune
    | '@"' (~["\\] | '\\' .)* '"'    // V-style identifier interpolation
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
