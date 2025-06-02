package ast

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/mylang/interpreter/parser"
)

// ASTBuilder converts ANTLR parse trees to AST nodes
type ASTBuilder struct {
	parser.BaseMyLangVisitor
}

// NewASTBuilder creates a new AST builder
func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{}
}

// BuildAST builds an AST from an ANTLR parse tree
func (b *ASTBuilder) BuildAST(tree antlr.ParseTree) *Program {
	program := tree.Accept(b).(*Program)
	return program
}

// VisitProgram visits a program parse tree
func (b *ASTBuilder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	statements := make([]Statement, 0)

	for _, stmtCtx := range ctx.AllStatement() {
		stmt := stmtCtx.Accept(b).(Statement)
		statements = append(statements, stmt)
	}

	return &Program{Statements: statements}
}

// VisitVariableDeclaration visits a variable declaration parse tree
func (b *ASTBuilder) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	var initializer Expression

	if ctx.Expression() != nil {
		initializer = ctx.Expression().Accept(b).(Expression)
	} else {
		// Default to null if no initializer
		initializer = &LiteralExpression{Value: nil}
	}

	return &VariableDeclaration{Name: name, Initializer: initializer}
}

// VisitAssignmentStatement visits an assignment statement parse tree
func (b *ASTBuilder) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	value := ctx.Expression().Accept(b).(Expression)

	return &AssignmentStatement{Name: name, Value: value}
}

// VisitArrayAssignmentStatement visits an array assignment statement parse tree
func (b *ASTBuilder) VisitArrayAssignmentStatement(ctx *parser.ArrayAssignmentStatementContext) interface{} {
	array := ctx.IDENTIFIER().Accept(b).(Expression)
	index := ctx.Expression(0).Accept(b).(Expression)
	value := ctx.Expression(1).Accept(b).(Expression)

	return &ArrayAssignmentStatement{Array: array, Index: index, Value: value}
}

// VisitPrintStatement visits a print statement parse tree
func (b *ASTBuilder) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	expression := ctx.Expression().Accept(b).(Expression)

	return &PrintStatement{Expression: expression}
}

// VisitIfStatement visits an if statement parse tree
func (b *ASTBuilder) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	condition := ctx.Expression().Accept(b).(Expression)
	thenBranch := ctx.Statement(0).Accept(b).(Statement)

	var elseBranch Statement
	if len(ctx.AllStatement()) > 1 {
		elseBranch = ctx.Statement(1).Accept(b).(Statement)
	}

	return &IfStatement{Condition: condition, ThenBranch: thenBranch, ElseBranch: elseBranch}
}

// VisitWhileStatement visits a while statement parse tree
func (b *ASTBuilder) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	condition := ctx.Expression().Accept(b).(Expression)
	body := ctx.Statement().Accept(b).(Statement)

	return &WhileStatement{Condition: condition, Body: body}
}

// VisitDoWhileStatement visits a do-while statement parse tree
func (b *ASTBuilder) VisitDoWhileStatement(ctx *parser.DoWhileStatementContext) interface{} {
	body := ctx.Statement().Accept(b).(Statement)
	condition := ctx.Expression().Accept(b).(Expression)

	return &DoWhileStatement{Body: body, Condition: condition}
}

// VisitForStatement visits a for statement parse tree
func (b *ASTBuilder) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	var initializer Statement
	if ctx.VarDeclaration() != nil {
		initializer = ctx.VarDeclaration().Accept(b).(Statement)
	} else if ctx.Assignment() != nil {
		initializer = ctx.Assignment().Accept(b).(Statement)
	}

	condition := ctx.Expression().Accept(b).(Expression)

	var increment Expression
	if ctx.Increment() != nil {
		increment = ctx.Increment().Accept(b).(Expression)
	}

	body := ctx.Statement().Accept(b).(Statement)

	return &ForStatement{Initializer: initializer, Condition: condition, Increment: increment, Body: body}
}

// VisitFunctionDeclaration visits a function declaration parse tree
func (b *ASTBuilder) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	parameters := make([]string, 0)

	if ctx.Parameters() != nil {
		for _, param := range ctx.Parameters().AllIDENTIFIER() {
			parameters = append(parameters, param.GetText())
		}
	}

	body := ctx.Block().Accept(b).(*Block)

	return &FunctionDeclaration{Name: name, Parameters: parameters, Body: body}
}

// VisitReturnStatement visits a return statement parse tree
func (b *ASTBuilder) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	var value Expression
	if ctx.Expression() != nil {
		value = ctx.Expression().Accept(b).(Expression)
	}

	return &ReturnStatement{Value: value}
}

// VisitBreakStatement visits a break statement parse tree
func (b *ASTBuilder) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	return &BreakStatement{}
}

// VisitContinueStatement visits a continue statement parse tree
func (b *ASTBuilder) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	return &ContinueStatement{}
}

// VisitSwitchStatement visits a switch statement parse tree
func (b *ASTBuilder) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	expression := ctx.Expression().Accept(b).(Expression)
	cases := make([]*SwitchCase, 0)

	for _, caseCtx := range ctx.AllSwitchCase() {
		case_ := caseCtx.Accept(b).(*SwitchCase)
		cases = append(cases, case_)
	}

	var defaultCase *Block
	if ctx.DefaultCase() != nil {
		defaultCase = ctx.DefaultCase().Block().Accept(b).(*Block)
	}

	return &SwitchStatement{Expression: expression, Cases: cases, DefaultCase: defaultCase}
}

// VisitSwitchCase visits a switch case parse tree
func (b *ASTBuilder) VisitSwitchCase(ctx *parser.SwitchCaseContext) interface{} {
	value := ctx.Expression().Accept(b).(Expression)
	statements := ctx.Block().Accept(b).(*Block)

	return &SwitchCase{Value: value, Statements: statements}
}

// VisitBlock visits a block parse tree
func (b *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	statements := make([]Statement, 0)

	for _, stmtCtx := range ctx.AllStatement() {
		stmt := stmtCtx.Accept(b).(Statement)
		statements = append(statements, stmt)
	}

	return &Block{Statements: statements}
}

// VisitExpressionStatement visits an expression statement parse tree
func (b *ASTBuilder) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	expression := ctx.Expression().Accept(b).(Expression)

	return &ExpressionStatement{Expression: expression}
}

// VisitLiteralExpression visits a literal expression parse tree
func (b *ASTBuilder) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) interface{} {
	if ctx.NUMBER() != nil {
		value, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
		return &LiteralExpression{Value: value}
	} else if ctx.STRING() != nil {
		// Remove the quotes from the string
		text := ctx.STRING().GetText()
		value := text[1 : len(text)-1] // Remove first and last character (quotes)
		return &LiteralExpression{Value: value}
	} else if ctx.TRUE() != nil {
		return &LiteralExpression{Value: true}
	} else if ctx.FALSE() != nil {
		return &LiteralExpression{Value: false}
	} else if ctx.NULL() != nil {
		return &LiteralExpression{Value: nil}
	}

	return &LiteralExpression{Value: nil}
}

// VisitArrayLiteralExpression visits an array literal expression parse tree
func (b *ASTBuilder) VisitArrayLiteralExpression(ctx *parser.ArrayLiteralExpressionContext) interface{} {
	elements := make([]Expression, 0)

	if ctx.ExpressionList() != nil {
		for _, exprCtx := range ctx.ExpressionList().AllExpression() {
			expr := exprCtx.Accept(b).(Expression)
			elements = append(elements, expr)
		}
	}

	return &ArrayLiteralExpression{Elements: elements}
}

// VisitIdentifierExpression visits an identifier expression parse tree
func (b *ASTBuilder) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {
	name := ctx.IDENTIFIER().GetText()

	return &VariableExpression{Name: name}
}

// VisitFunctionCallExpression visits a function call expression parse tree
func (b *ASTBuilder) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) interface{} {
	callee := ctx.IDENTIFIER().GetText()
	arguments := make([]Expression, 0)

	if ctx.ExpressionList() != nil {
		for _, exprCtx := range ctx.ExpressionList().AllExpression() {
			expr := exprCtx.Accept(b).(Expression)
			arguments = append(arguments, expr)
		}
	}

	return &FunctionCallExpression{Callee: callee, Arguments: arguments}
}

// VisitArrayAccessExpression visits an array access expression parse tree
func (b *ASTBuilder) VisitArrayAccessExpression(ctx *parser.ArrayAccessExpressionContext) interface{} {
	array := ctx.Expression(0).Accept(b).(Expression)
	index := ctx.Expression(1).Accept(b).(Expression)

	return &ArrayAccessExpression{Array: array, Index: index}
}

// VisitGroupingExpression visits a grouping expression parse tree
func (b *ASTBuilder) VisitGroupingExpression(ctx *parser.GroupingExpressionContext) interface{} {
	expression := ctx.Expression().Accept(b).(Expression)

	return &GroupingExpression{Expression: expression}
}

// VisitUnaryExpression visits a unary expression parse tree
func (b *ASTBuilder) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	operator := ctx.GetOp().GetText()
	right := ctx.Expression().Accept(b).(Expression)

	return &UnaryExpression{Operator: operator, Right: right}
}

// VisitBinaryExpression visits a binary expression parse tree
func (b *ASTBuilder) VisitBinaryExpression(ctx *parser.BinaryExpressionContext) interface{} {
	left := ctx.Expression(0).Accept(b).(Expression)
	operator := ctx.GetOp().GetText()
	right := ctx.Expression(1).Accept(b).(Expression)

	return &BinaryExpression{Left: left, Operator: operator, Right: right}
}

// Visit is the generic visit method for the visitor pattern
func (b *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(b)
}

// VisitChildren is the generic visit children method for the visitor pattern
func (b *ASTBuilder) VisitChildren(node antlr.RuleNode) interface{} {
	result := make([]interface{}, 0)

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)

		if childResult := child.Accept(b); childResult != nil {
			result = append(result, childResult)
		}
	}

	return result
}

// VisitTerminal is the generic visit terminal method for the visitor pattern
func (b *ASTBuilder) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

// VisitErrorNode is the generic visit error node method for the visitor pattern
func (b *ASTBuilder) VisitErrorNode(node antlr.ErrorNode) interface{} {
	fmt.Println("Error during parsing:", node.GetText())
	return nil
}
