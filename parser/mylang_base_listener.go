// Code generated from ./grammar/MyLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MyLang
import "github.com/antlr4-go/antlr/v4"

// BaseMyLangListener is a complete listener for a parse tree produced by MyLangParser.
type BaseMyLangListener struct{}

var _ MyLangListener = &BaseMyLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMyLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMyLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMyLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMyLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMyLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMyLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterModuleDecl is called when production moduleDecl is entered.
func (s *BaseMyLangListener) EnterModuleDecl(ctx *ModuleDeclContext) {}

// ExitModuleDecl is called when production moduleDecl is exited.
func (s *BaseMyLangListener) ExitModuleDecl(ctx *ModuleDeclContext) {}

// EnterTopLevelDecl is called when production topLevelDecl is entered.
func (s *BaseMyLangListener) EnterTopLevelDecl(ctx *TopLevelDeclContext) {}

// ExitTopLevelDecl is called when production topLevelDecl is exited.
func (s *BaseMyLangListener) ExitTopLevelDecl(ctx *TopLevelDeclContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseMyLangListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseMyLangListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseMyLangListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseMyLangListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseMyLangListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseMyLangListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterGenericParams is called when production genericParams is entered.
func (s *BaseMyLangListener) EnterGenericParams(ctx *GenericParamsContext) {}

// ExitGenericParams is called when production genericParams is exited.
func (s *BaseMyLangListener) ExitGenericParams(ctx *GenericParamsContext) {}

// EnterParams is called when production params is entered.
func (s *BaseMyLangListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseMyLangListener) ExitParams(ctx *ParamsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseMyLangListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseMyLangListener) ExitParam(ctx *ParamContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseMyLangListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseMyLangListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseMyLangListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseMyLangListener) ExitStructField(ctx *StructFieldContext) {}

// EnterInterfaceDecl is called when production interfaceDecl is entered.
func (s *BaseMyLangListener) EnterInterfaceDecl(ctx *InterfaceDeclContext) {}

// ExitInterfaceDecl is called when production interfaceDecl is exited.
func (s *BaseMyLangListener) ExitInterfaceDecl(ctx *InterfaceDeclContext) {}

// EnterInterfaceMethod is called when production interfaceMethod is entered.
func (s *BaseMyLangListener) EnterInterfaceMethod(ctx *InterfaceMethodContext) {}

// ExitInterfaceMethod is called when production interfaceMethod is exited.
func (s *BaseMyLangListener) ExitInterfaceMethod(ctx *InterfaceMethodContext) {}

// EnterEnumDecl is called when production enumDecl is entered.
func (s *BaseMyLangListener) EnterEnumDecl(ctx *EnumDeclContext) {}

// ExitEnumDecl is called when production enumDecl is exited.
func (s *BaseMyLangListener) ExitEnumDecl(ctx *EnumDeclContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseMyLangListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseMyLangListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *BaseMyLangListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *BaseMyLangListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterConstField is called when production constField is entered.
func (s *BaseMyLangListener) EnterConstField(ctx *ConstFieldContext) {}

// ExitConstField is called when production constField is exited.
func (s *BaseMyLangListener) ExitConstField(ctx *ConstFieldContext) {}

// EnterType is called when production type is entered.
func (s *BaseMyLangListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseMyLangListener) ExitType(ctx *TypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseMyLangListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseMyLangListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseMyLangListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseMyLangListener) ExitMapType(ctx *MapTypeContext) {}

// EnterOptionType is called when production optionType is entered.
func (s *BaseMyLangListener) EnterOptionType(ctx *OptionTypeContext) {}

// ExitOptionType is called when production optionType is exited.
func (s *BaseMyLangListener) ExitOptionType(ctx *OptionTypeContext) {}

// EnterResultType is called when production resultType is entered.
func (s *BaseMyLangListener) EnterResultType(ctx *ResultTypeContext) {}

// ExitResultType is called when production resultType is exited.
func (s *BaseMyLangListener) ExitResultType(ctx *ResultTypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMyLangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMyLangListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMyLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMyLangListener) ExitStatement(ctx *StatementContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseMyLangListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseMyLangListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseMyLangListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseMyLangListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignOp is called when production assignOp is entered.
func (s *BaseMyLangListener) EnterAssignOp(ctx *AssignOpContext) {}

// ExitAssignOp is called when production assignOp is exited.
func (s *BaseMyLangListener) ExitAssignOp(ctx *AssignOpContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseMyLangListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseMyLangListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseMyLangListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseMyLangListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseMyLangListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseMyLangListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterMatchStmt is called when production matchStmt is entered.
func (s *BaseMyLangListener) EnterMatchStmt(ctx *MatchStmtContext) {}

// ExitMatchStmt is called when production matchStmt is exited.
func (s *BaseMyLangListener) ExitMatchStmt(ctx *MatchStmtContext) {}

// EnterMatchBranch is called when production matchBranch is entered.
func (s *BaseMyLangListener) EnterMatchBranch(ctx *MatchBranchContext) {}

// ExitMatchBranch is called when production matchBranch is exited.
func (s *BaseMyLangListener) ExitMatchBranch(ctx *MatchBranchContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseMyLangListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseMyLangListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *BaseMyLangListener) EnterDeferStmt(ctx *DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *BaseMyLangListener) ExitDeferStmt(ctx *DeferStmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMyLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMyLangListener) ExitExpr(ctx *ExprContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseMyLangListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseMyLangListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseMyLangListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseMyLangListener) ExitLiteral(ctx *LiteralContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseMyLangListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseMyLangListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseMyLangListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseMyLangListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterMapElement is called when production mapElement is entered.
func (s *BaseMyLangListener) EnterMapElement(ctx *MapElementContext) {}

// ExitMapElement is called when production mapElement is exited.
func (s *BaseMyLangListener) ExitMapElement(ctx *MapElementContext) {}

// EnterStructLiteral is called when production structLiteral is entered.
func (s *BaseMyLangListener) EnterStructLiteral(ctx *StructLiteralContext) {}

// ExitStructLiteral is called when production structLiteral is exited.
func (s *BaseMyLangListener) ExitStructLiteral(ctx *StructLiteralContext) {}

// EnterStructElement is called when production structElement is entered.
func (s *BaseMyLangListener) EnterStructElement(ctx *StructElementContext) {}

// ExitStructElement is called when production structElement is exited.
func (s *BaseMyLangListener) ExitStructElement(ctx *StructElementContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseMyLangListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseMyLangListener) ExitExprList(ctx *ExprListContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseMyLangListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseMyLangListener) ExitReturnType(ctx *ReturnTypeContext) {}
