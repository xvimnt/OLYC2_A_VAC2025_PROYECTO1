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
