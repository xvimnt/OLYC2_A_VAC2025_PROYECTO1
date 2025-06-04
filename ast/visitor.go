package ast

// Visitor defines the visitor pattern interface for traversing the AST
type Visitor interface {
	// Program nodes
	VisitProgram(program *Program) interface{}

	// Statement nodes
	VisitBlock(block *Block) interface{}
	VisitVariableDeclaration(declaration *VariableDeclaration) interface{}
	VisitAssignmentStatement(assignment *AssignmentStatement) interface{}
	VisitExpressionStatement(statement *ExpressionStatement) interface{}
	VisitPrintStatement(statement *PrintStatement) interface{}
	VisitIfStatement(statement *IfStatement) interface{}
	VisitWhileStatement(statement *WhileStatement) interface{}
	VisitDoWhileStatement(statement *DoWhileStatement) interface{}
	VisitForStatement(statement *ForStatement) interface{}
	VisitSwitchStatement(statement *SwitchStatement) interface{}
	VisitBreakStatement(statement *BreakStatement) interface{}
	VisitContinueStatement(statement *ContinueStatement) interface{}
	VisitFunctionDeclaration(declaration *FunctionDeclaration) interface{}
	VisitReturnStatement(statement *ReturnStatement) interface{}

	// Expression nodes
	VisitBinaryExpression(expression *BinaryExpression) interface{}
	VisitUnaryExpression(expression *UnaryExpression) interface{}
	VisitLiteralExpression(expression *LiteralExpression) interface{}
	VisitVariableExpression(expression *VariableExpression) interface{}
	VisitArrayAccessExpression(expression *ArrayAccessExpression) interface{}
	VisitArrayLiteralExpression(expression *ArrayLiteralExpression) interface{}
	VisitFunctionCallExpression(expression *FunctionCallExpression) interface{}
}

// BaseVisitor provides a default implementation of Visitor
// that can be embedded in concrete visitors. All methods return nil by default.
type BaseVisitor struct{}

// Program nodes
func (v *BaseVisitor) VisitProgram(program *Program) interface{} { return nil }

// Statement nodes
func (v *BaseVisitor) VisitBlock(block *Block) interface{} { return nil }
func (v *BaseVisitor) VisitVariableDeclaration(declaration *VariableDeclaration) interface{} { return nil }
func (v *BaseVisitor) VisitAssignmentStatement(assignment *AssignmentStatement) interface{} { return nil }
func (v *BaseVisitor) VisitExpressionStatement(statement *ExpressionStatement) interface{} { return nil }
func (v *BaseVisitor) VisitPrintStatement(statement *PrintStatement) interface{} { return nil }
func (v *BaseVisitor) VisitIfStatement(statement *IfStatement) interface{} { return nil }
func (v *BaseVisitor) VisitWhileStatement(statement *WhileStatement) interface{} { return nil }
func (v *BaseVisitor) VisitDoWhileStatement(statement *DoWhileStatement) interface{} { return nil }
func (v *BaseVisitor) VisitForStatement(statement *ForStatement) interface{} { return nil }
func (v *BaseVisitor) VisitSwitchStatement(statement *SwitchStatement) interface{} { return nil }
func (v *BaseVisitor) VisitBreakStatement(statement *BreakStatement) interface{} { return nil }
func (v *BaseVisitor) VisitContinueStatement(statement *ContinueStatement) interface{} { return nil }
func (v *BaseVisitor) VisitFunctionDeclaration(declaration *FunctionDeclaration) interface{} { return nil }
func (v *BaseVisitor) VisitReturnStatement(statement *ReturnStatement) interface{} { return nil }

// Expression nodes
func (v *BaseVisitor) VisitBinaryExpression(expression *BinaryExpression) interface{} { return nil }
func (v *BaseVisitor) VisitUnaryExpression(expression *UnaryExpression) interface{} { return nil }
func (v *BaseVisitor) VisitLiteralExpression(expression *LiteralExpression) interface{} { return nil }
func (v *BaseVisitor) VisitVariableExpression(expression *VariableExpression) interface{} { return nil }
func (v *BaseVisitor) VisitArrayAccessExpression(expression *ArrayAccessExpression) interface{} { return nil }
func (v *BaseVisitor) VisitArrayLiteralExpression(expression *ArrayLiteralExpression) interface{} { return nil }
func (v *BaseVisitor) VisitFunctionCallExpression(expression *FunctionCallExpression) interface{} { return nil }

// VVisitor extends the Visitor interface with V language specific node visitors
type VVisitor interface {
	// Override the base Visitor interface
	Visitor
	
	// V-specific statement nodes
	VisitModuleDeclaration(module *ModuleDeclaration) interface{}
	VisitImportDeclaration(imp *ImportDeclaration) interface{}
	VisitConstDeclaration(constant *ConstDeclaration) interface{}
	VisitTypedVariable(variable *TypedVariable) interface{}
	VisitStructDeclaration(structDecl *StructDeclaration) interface{}
	VisitEnumDeclaration(enumDecl *EnumDeclaration) interface{}
	VisitInterfaceDeclaration(interfaceDecl *InterfaceDeclaration) interface{}
	VisitVFunctionDeclaration(function *VFunctionDeclaration) interface{}
	VisitMatchStatement(match *MatchStatement) interface{}
	VisitForInStatement(forIn *ForInStatement) interface{}
	VisitForCStatement(forC *ForCStatement) interface{}
	VisitDeferStatement(defer *DeferStatement) interface{}
	
	// V-specific expression nodes
	VisitSelectorExpression(selector *SelectorExpression) interface{}
	VisitMethodCallExpression(methodCall *MethodCallExpression) interface{}
	VisitMapLiteralExpression(mapLiteral *MapLiteralExpression) interface{}
	VisitOptionalExpression(optional *OptionalExpression) interface{}
	VisitOrBlockExpression(orBlock *OrBlockExpression) interface{}
}

// VBaseVisitor provides a default implementation of VVisitor
type VBaseVisitor struct {
	BaseVisitor
}

// V-specific statement nodes
func (v *VBaseVisitor) VisitModuleDeclaration(module *ModuleDeclaration) interface{} { return nil }
func (v *VBaseVisitor) VisitImportDeclaration(imp *ImportDeclaration) interface{} { return nil }
func (v *VBaseVisitor) VisitConstDeclaration(constant *ConstDeclaration) interface{} { return nil }
func (v *VBaseVisitor) VisitTypedVariable(variable *TypedVariable) interface{} { return nil }
func (v *VBaseVisitor) VisitStructDeclaration(structDecl *StructDeclaration) interface{} { return nil }
func (v *VBaseVisitor) VisitEnumDeclaration(enumDecl *EnumDeclaration) interface{} { return nil }
func (v *VBaseVisitor) VisitInterfaceDeclaration(interfaceDecl *InterfaceDeclaration) interface{} { return nil }
func (v *VBaseVisitor) VisitVFunctionDeclaration(function *VFunctionDeclaration) interface{} { return nil }
func (v *VBaseVisitor) VisitMatchStatement(match *MatchStatement) interface{} { return nil }
func (v *VBaseVisitor) VisitForInStatement(forIn *ForInStatement) interface{} { return nil }
func (v *VBaseVisitor) VisitForCStatement(forC *ForCStatement) interface{} { return nil }
func (v *VBaseVisitor) VisitDeferStatement(defer *DeferStatement) interface{} { return nil }

// V-specific expression nodes
func (v *VBaseVisitor) VisitSelectorExpression(selector *SelectorExpression) interface{} { return nil }
func (v *VBaseVisitor) VisitMethodCallExpression(methodCall *MethodCallExpression) interface{} { return nil }
func (v *VBaseVisitor) VisitMapLiteralExpression(mapLiteral *MapLiteralExpression) interface{} { return nil }
func (v *VBaseVisitor) VisitOptionalExpression(optional *OptionalExpression) interface{} { return nil }
func (v *VBaseVisitor) VisitOrBlockExpression(orBlock *OrBlockExpression) interface{} { return nil }
