package ast

// Node is the base interface for all AST nodes
type Node interface {
	Accept(visitor Visitor) interface{}
}

// Expression is the base interface for all expression nodes
type Expression interface {
	Node
	ExpressionNode()
}

// Statement is the base interface for all statement nodes
type Statement interface {
	Node
	StatementNode()
}

// Program represents a complete program
type Program struct {
	Statements []Statement
}

// Accept implements the Node interface
func (p *Program) Accept(visitor Visitor) interface{} {
	return visitor.VisitProgram(p)
}

// ExpressionStatement wraps an expression as a statement
type ExpressionStatement struct {
	Expression Expression
}

// Accept implements the Node interface
func (s *ExpressionStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitExpressionStatement(s)
}

// StatementNode implements the Statement interface
func (s *ExpressionStatement) StatementNode() {}

// VariableDeclaration represents a variable declaration
type VariableDeclaration struct {
	Name        string
	Initializer Expression
}

// Accept implements the Node interface
func (d *VariableDeclaration) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariableDeclaration(d)
}

// StatementNode implements the Statement interface
func (d *VariableDeclaration) StatementNode() {}

// AssignmentStatement represents an assignment statement
type AssignmentStatement struct {
	Name  string
	Value Expression
	Index Expression // Optional, for array assignments
}

// Accept implements the Node interface
func (s *AssignmentStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitAssignmentStatement(s)
}

// StatementNode implements the Statement interface
func (s *AssignmentStatement) StatementNode() {}

// ArrayAssignmentStatement represents an array assignment statement
type ArrayAssignmentStatement struct {
	Array Expression
	Index Expression
	Value Expression
}

// Accept implements the Node interface
func (s *ArrayAssignmentStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitArrayAssignmentStatement(s)
}

// StatementNode implements the Statement interface
func (s *ArrayAssignmentStatement) StatementNode() {}

// PrintStatement represents a print statement
type PrintStatement struct {
	Expression Expression
}

// Accept implements the Node interface
func (s *PrintStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitPrintStatement(s)
}

// StatementNode implements the Statement interface
func (s *PrintStatement) StatementNode() {}

// IfStatement represents an if statement
type IfStatement struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

// Accept implements the Node interface
func (s *IfStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitIfStatement(s)
}

// StatementNode implements the Statement interface
func (s *IfStatement) StatementNode() {}

// WhileStatement represents a while statement
type WhileStatement struct {
	Condition Expression
	Body      Statement
}

// Accept implements the Node interface
func (s *WhileStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitWhileStatement(s)
}

// StatementNode implements the Statement interface
func (s *WhileStatement) StatementNode() {}

// DoWhileStatement represents a do-while statement
type DoWhileStatement struct {
	Body      Statement
	Condition Expression
}

// Accept implements the Node interface
func (s *DoWhileStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitDoWhileStatement(s)
}

// StatementNode implements the Statement interface
func (s *DoWhileStatement) StatementNode() {}

// ForStatement represents a for statement
type ForStatement struct {
	Initializer Statement
	Condition   Expression
	Increment   Expression
	Body        Statement
}

// Accept implements the Node interface
func (s *ForStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitForStatement(s)
}

// StatementNode implements the Statement interface
func (s *ForStatement) StatementNode() {}

// SwitchStatement represents a switch statement
type SwitchStatement struct {
	Expression  Expression
	Cases       []*SwitchCase
	DefaultCase *Block
}

// Accept implements the Node interface
func (s *SwitchStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitSwitchStatement(s)
}

// StatementNode implements the Statement interface
func (s *SwitchStatement) StatementNode() {}

// SwitchCase represents a case in a switch statement
type SwitchCase struct {
	Value      Expression
	Statements *Block
}

// Accept implements the Node interface
func (c *SwitchCase) Accept(visitor Visitor) interface{} {
	return visitor.VisitSwitchCase(c)
}

// BreakStatement represents a break statement
type BreakStatement struct{}

// Accept implements the Node interface
func (s *BreakStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitBreakStatement(s)
}

// StatementNode implements the Statement interface
func (s *BreakStatement) StatementNode() {}

// ContinueStatement represents a continue statement
type ContinueStatement struct{}

// Accept implements the Node interface
func (s *ContinueStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitContinueStatement(s)
}

// StatementNode implements the Statement interface
func (s *ContinueStatement) StatementNode() {}

// Block represents a block of statements
type Block struct {
	Statements []Statement
}

// Accept implements the Node interface
func (b *Block) Accept(visitor Visitor) interface{} {
	return visitor.VisitBlock(b)
}

// StatementNode implements the Statement interface
func (b *Block) StatementNode() {}

// FunctionDeclaration represents a function declaration
type FunctionDeclaration struct {
	Name       string
	Parameters []string
	Body       *Block
}

// Accept implements the Node interface
func (d *FunctionDeclaration) Accept(visitor Visitor) interface{} {
	return visitor.VisitFunctionDeclaration(d)
}

// StatementNode implements the Statement interface
func (d *FunctionDeclaration) StatementNode() {}

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Value Expression
}

// Accept implements the Node interface
func (s *ReturnStatement) Accept(visitor Visitor) interface{} {
	return visitor.VisitReturnStatement(s)
}

// StatementNode implements the Statement interface
func (s *ReturnStatement) StatementNode() {}

// LiteralExpression represents a literal value
type LiteralExpression struct {
	Value interface{}
}

// Accept implements the Node interface
func (e *LiteralExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteralExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *LiteralExpression) ExpressionNode() {}

// VariableExpression represents a variable reference
type VariableExpression struct {
	Name string
}

// Accept implements the Node interface
func (e *VariableExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariableExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *VariableExpression) ExpressionNode() {}

// UnaryExpression represents a unary operation
type UnaryExpression struct {
	Operator string
	Right    Expression
}

// Accept implements the Node interface
func (e *UnaryExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnaryExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *UnaryExpression) ExpressionNode() {}

// BinaryExpression represents a binary operation
type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

// Accept implements the Node interface
func (e *BinaryExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *BinaryExpression) ExpressionNode() {}

// GroupingExpression represents a grouping expression (parentheses)
type GroupingExpression struct {
	Expression Expression
}

// Accept implements the Node interface
func (e *GroupingExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitGroupingExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *GroupingExpression) ExpressionNode() {}

// ArrayLiteralExpression represents an array literal
type ArrayLiteralExpression struct {
	Elements []Expression
}

// Accept implements the Node interface
func (e *ArrayLiteralExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitArrayLiteralExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *ArrayLiteralExpression) ExpressionNode() {}

// ArrayAccessExpression represents array access with an index
type ArrayAccessExpression struct {
	Array Expression
	Index Expression
}

// Accept implements the Node interface
func (e *ArrayAccessExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitArrayAccessExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *ArrayAccessExpression) ExpressionNode() {}

// FunctionCallExpression represents a function call
type FunctionCallExpression struct {
	Callee    string
	Arguments []Expression
}

// Accept implements the Node interface
func (e *FunctionCallExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitFunctionCallExpression(e)
}

// ExpressionNode implements the Expression interface
func (e *FunctionCallExpression) ExpressionNode() {}
