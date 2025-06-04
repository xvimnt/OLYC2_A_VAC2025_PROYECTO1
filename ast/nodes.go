package ast

// Node represents a node in the abstract syntax tree
type Node interface {
	// Accept accepts a visitor to visit this node
	Accept(visitor Visitor) interface{}
}

// Statement represents a statement in the language
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression in the language
type Expression interface {
	Node
	expressionNode()
}

// Program represents the top-level program node
type Program struct {
	Statements []Statement
}

// Accept implements the Node interface
func (p *Program) Accept(visitor Visitor) interface{} {
	return visitor.VisitProgram(p)
}

// Block represents a block of statements
type Block struct {
	Statements []Statement
}

func (b *Block) statementNode() {}

// Accept implements the Node interface
func (b *Block) Accept(visitor Visitor) interface{} {
	return visitor.VisitBlock(b)
}

// VariableDeclaration represents a variable declaration
type VariableDeclaration struct {
	Name        string
	Initializer Expression
}

func (v *VariableDeclaration) statementNode() {}

// Accept implements the Node interface
func (v *VariableDeclaration) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariableDeclaration(v)
}

// AssignmentStatement represents an assignment statement
type AssignmentStatement struct {
	Name      string
	Index     Expression // For array assignments, nil for regular assignments
	Value     Expression
}

func (a *AssignmentStatement) statementNode() {}

// Accept implements the Node interface
func (a *AssignmentStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitAssignmentStatement(a)
}

// ExpressionStatement represents an expression used as a statement
type ExpressionStatement struct {
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}

// Accept implements the Node interface
func (e *ExpressionStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitExpressionStatement(e)
}

// PrintStatement represents a print statement
type PrintStatement struct {
	Expression Expression
}

func (p *PrintStatement) statementNode() {}

// Accept implements the Node interface
func (p *PrintStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitPrintStatement(p)
}

// IfStatement represents an if statement
type IfStatement struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (i *IfStatement) statementNode() {}

// Accept implements the Node interface
func (i *IfStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitIfStatement(i)
}

// WhileStatement represents a while loop
type WhileStatement struct {
	Condition Expression
	Body      Statement
}

func (w *WhileStatement) statementNode() {}

// Accept implements the Node interface
func (w *WhileStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitWhileStatement(w)
}

// DoWhileStatement represents a do-while loop
type DoWhileStatement struct {
	Body      Statement
	Condition Expression
}

func (d *DoWhileStatement) statementNode() {}

// Accept implements the Node interface
func (d *DoWhileStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitDoWhileStatement(d)
}

// ForStatement represents a for loop
type ForStatement struct {
	Initializer Statement
	Condition   Expression
	Increment   Expression
	Body        Statement
}

func (f *ForStatement) statementNode() {}

// Accept implements the Node interface
func (f *ForStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitForStatement(f)
}

// SwitchStatement represents a switch statement
type SwitchStatement struct {
	Value Expression
	Cases []*SwitchCase
}

func (s *SwitchStatement) statementNode() {}

// Accept implements the Node interface
func (s *SwitchStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitSwitchStatement(s)
}

// SwitchCase represents a case in a switch statement
type SwitchCase struct {
	Value      Expression // nil for default case
	Statements []Statement
	IsDefault  bool
}

// BreakStatement represents a break statement
type BreakStatement struct{}

func (b *BreakStatement) statementNode() {}

// Accept implements the Node interface
func (b *BreakStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitBreakStatement(b)
}

// ContinueStatement represents a continue statement
type ContinueStatement struct{}

func (c *ContinueStatement) statementNode() {}

// Accept implements the Node interface
func (c *ContinueStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitContinueStatement(c)
}

// FunctionDeclaration represents a function declaration
type FunctionDeclaration struct {
	Name       string
	Parameters []string
	Body       *Block
}

func (f *FunctionDeclaration) statementNode() {}

// Accept implements the Node interface
func (f *FunctionDeclaration) Accept(visitor Visitor) interface{} {
	return visitor.VisitFunctionDeclaration(f)
}

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Value Expression
}

func (r *ReturnStatement) statementNode() {}

// Accept implements the Node interface
func (r *ReturnStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitReturnStatement(r)
}

// BinaryExpression represents a binary expression
type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryExpression) expressionNode() {}

// Accept implements the Node interface
func (b *BinaryExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryExpression(b)
}

// UnaryExpression represents a unary expression
type UnaryExpression struct {
	Operator string
	Right    Expression
}

func (u *UnaryExpression) expressionNode() {}

// Accept implements the Node interface
func (u *UnaryExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnaryExpression(u)
}

// LiteralExpression represents a literal value
type LiteralExpression struct {
	Value interface{}
}

func (l *LiteralExpression) expressionNode() {}

// Accept implements the Node interface
func (l *LiteralExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteralExpression(l)
}

// VariableExpression represents a variable reference
type VariableExpression struct {
	Name string
}

func (v *VariableExpression) expressionNode() {}

// Accept implements the Node interface
func (v *VariableExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariableExpression(v)
}

// ArrayAccessExpression represents an array access expression
type ArrayAccessExpression struct {
	Name  string
	Index Expression
}

func (a *ArrayAccessExpression) expressionNode() {}

// Accept implements the Node interface
func (a *ArrayAccessExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitArrayAccessExpression(a)
}

// ArrayLiteralExpression represents an array literal
type ArrayLiteralExpression struct {
	Elements []Expression
}

func (a *ArrayLiteralExpression) expressionNode() {}

// Accept implements the Node interface
func (a *ArrayLiteralExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitArrayLiteralExpression(a)
}

// FunctionCallExpression represents a function call
type FunctionCallExpression struct {
	Callee    string
	Arguments []Expression
}

func (f *FunctionCallExpression) expressionNode() {}

// Accept implements the Node interface
func (f *FunctionCallExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitFunctionCallExpression(f)
}

// V Language Specific AST Nodes

// ModuleDeclaration represents a V module declaration
type ModuleDeclaration struct {
	Name string
}

func (m *ModuleDeclaration) statementNode() {}

// Accept implements the Node interface
func (m *ModuleDeclaration) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitModuleDeclaration(m)
	}
	return nil
}

// ImportDeclaration represents a V import declaration
type ImportDeclaration struct {
	Path string
}

func (i *ImportDeclaration) statementNode() {}

// Accept implements the Node interface
func (i *ImportDeclaration) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitImportDeclaration(i)
	}
	return nil
}

// ConstDeclaration represents a V const declaration
type ConstDeclaration struct {
	Name  string
	Value Expression
}

func (c *ConstDeclaration) statementNode() {}

// Accept implements the Node interface
func (c *ConstDeclaration) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitConstDeclaration(c)
	}
	return nil
}

// TypedVariable represents a variable with a type in V
type TypedVariable struct {
	Name        string
	TypeName    string
	IsMutable   bool
	Initializer Expression
}

func (v *TypedVariable) statementNode() {}

// Accept implements the Node interface
func (v *TypedVariable) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitTypedVariable(v)
	}
	return nil
}

// StructDeclaration represents a V struct declaration
type StructDeclaration struct {
	Name     string
	IsPublic bool
	Fields   []StructField
}

func (s *StructDeclaration) statementNode() {}

// Accept implements the Node interface
func (s *StructDeclaration) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitStructDeclaration(s)
	}
	return nil
}

// StructField represents a field in a struct
type StructField struct {
	Name     string
	Type     string
	IsMutable bool
}

// EnumDeclaration represents a V enum declaration
type EnumDeclaration struct {
	Name     string
	IsPublic bool
	Values   []EnumValue
}

func (e *EnumDeclaration) statementNode() {}

// Accept implements the Node interface
func (e *EnumDeclaration) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitEnumDeclaration(e)
	}
	return nil
}

// EnumValue represents a value in an enum
type EnumValue struct {
	Name  string
	Value Expression // Optional explicit value
}

// InterfaceDeclaration represents a V interface declaration
type InterfaceDeclaration struct {
	Name     string
	IsPublic bool
	Methods  []InterfaceMethod
}

func (i *InterfaceDeclaration) statementNode() {}

// Accept implements the Node interface
func (i *InterfaceDeclaration) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitInterfaceDeclaration(i)
	}
	return nil
}

// InterfaceMethod represents a method in an interface
type InterfaceMethod struct {
	Name       string
	Params     []Parameter
	ReturnType string
}

// Parameter represents a function parameter
type Parameter struct {
	Name     string
	Type     string
	IsMutable bool
}

// VFunctionDeclaration represents a V function declaration with a receiver for methods
type VFunctionDeclaration struct {
	Name       string
	IsPublic   bool
	Receiver   *Parameter // Optional receiver for methods
	Params     []Parameter
	ReturnType string
	Body       *Block
}

func (f *VFunctionDeclaration) statementNode() {}

// Accept implements the Node interface
func (f *VFunctionDeclaration) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitVFunctionDeclaration(f)
	}
	return nil
}

// MatchStatement represents a V match statement (similar to switch)
type MatchStatement struct {
	Subject  Expression
	Branches []*MatchBranch
}

func (m *MatchStatement) statementNode() {}

// Accept implements the Node interface
func (m *MatchStatement) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitMatchStatement(m)
	}
	return nil
}

// MatchBranch represents a branch in a match statement
type MatchBranch struct {
	Patterns []Expression // Can be multiple patterns separated by comma
	Body     []Statement
	IsElse   bool
}

// ForInStatement represents a V for-in loop (for item in collection)
type ForInStatement struct {
	IndexVariable string // Optional index variable name
	ValueVariable string // Value variable name
	Iterable      Expression
	Body          []Statement
}

func (f *ForInStatement) statementNode() {}

// Accept implements the Node interface
func (f *ForInStatement) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitForInStatement(f)
	}
	return nil
}

// ForCStatement represents a V for loop with only a condition (while loop)
type ForCStatement struct {
	Condition Expression
	Body      []Statement
}

func (f *ForCStatement) statementNode() {}

// Accept implements the Node interface
func (f *ForCStatement) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitForCStatement(f)
	}
	return nil
}

// DeferStatement represents a V defer statement
type DeferStatement struct {
	Statement Statement
}

func (d *DeferStatement) statementNode() {}

// Accept implements the Node interface
func (d *DeferStatement) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitDeferStatement(d)
	}
	return nil
}

// SelectorExpression represents a V field access expression (x.y)
type SelectorExpression struct {
	Object Expression
	Field  string
}

func (s *SelectorExpression) expressionNode() {}

// Accept implements the Node interface
func (s *SelectorExpression) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitSelectorExpression(s)
	}
	return nil
}

// MethodCallExpression represents a V method call expression (x.method())
type MethodCallExpression struct {
	Object    Expression
	Method    string
	Arguments []Expression
}

func (m *MethodCallExpression) expressionNode() {}

// Accept implements the Node interface
func (m *MethodCallExpression) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitMethodCallExpression(m)
	}
	return nil
}

// MapLiteralExpression represents a V map literal
type MapLiteralExpression struct {
	KeyType   string
	ValueType string
	Entries   []MapEntry
}

func (m *MapLiteralExpression) expressionNode() {}

// Accept implements the Node interface
func (m *MapLiteralExpression) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitMapLiteralExpression(m)
	}
	return nil
}

// MapEntry represents a key-value pair in a map
type MapEntry struct {
	Key   Expression
	Value Expression
}

// OptionalExpression represents a V option type expression
type OptionalExpression struct {
	Expression Expression
}

func (o *OptionalExpression) expressionNode() {}

// Accept implements the Node interface
func (o *OptionalExpression) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitOptionalExpression(o)
	}
	return nil
}

// OrBlockExpression represents the block in an "or" expression for handling None
type OrBlockExpression struct {
	Expression Expression
	OrBlock    []Statement
}

func (o *OrBlockExpression) expressionNode() {}

// Accept implements the Node interface
func (o *OrBlockExpression) Accept(visitor Visitor) interface{} {
	if vVisitor, ok := visitor.(VVisitor); ok {
		return vVisitor.VisitOrBlockExpression(o)
	}
	return nil
}
