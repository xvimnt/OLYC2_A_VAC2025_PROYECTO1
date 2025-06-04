// Code generated from ./grammar/MyLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MyLang
import "github.com/antlr4-go/antlr/v4"


// MyLangListener is a complete listener for a parse tree produced by MyLangParser.
type MyLangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterModuleDecl is called when entering the moduleDecl production.
	EnterModuleDecl(c *ModuleDeclContext)

	// EnterTopLevelDecl is called when entering the topLevelDecl production.
	EnterTopLevelDecl(c *TopLevelDeclContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterImportSpec is called when entering the importSpec production.
	EnterImportSpec(c *ImportSpecContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterGenericParams is called when entering the genericParams production.
	EnterGenericParams(c *GenericParamsContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterInterfaceDecl is called when entering the interfaceDecl production.
	EnterInterfaceDecl(c *InterfaceDeclContext)

	// EnterInterfaceMethod is called when entering the interfaceMethod production.
	EnterInterfaceMethod(c *InterfaceMethodContext)

	// EnterEnumDecl is called when entering the enumDecl production.
	EnterEnumDecl(c *EnumDeclContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterConstDecl is called when entering the constDecl production.
	EnterConstDecl(c *ConstDeclContext)

	// EnterConstField is called when entering the constField production.
	EnterConstField(c *ConstFieldContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterOptionType is called when entering the optionType production.
	EnterOptionType(c *OptionTypeContext)

	// EnterResultType is called when entering the resultType production.
	EnterResultType(c *ResultTypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAssignOp is called when entering the assignOp production.
	EnterAssignOp(c *AssignOpContext)

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterMatchStmt is called when entering the matchStmt production.
	EnterMatchStmt(c *MatchStmtContext)

	// EnterMatchBranch is called when entering the matchBranch production.
	EnterMatchBranch(c *MatchBranchContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterDeferStmt is called when entering the deferStmt production.
	EnterDeferStmt(c *DeferStmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterMapElement is called when entering the mapElement production.
	EnterMapElement(c *MapElementContext)

	// EnterStructLiteral is called when entering the structLiteral production.
	EnterStructLiteral(c *StructLiteralContext)

	// EnterStructElement is called when entering the structElement production.
	EnterStructElement(c *StructElementContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitModuleDecl is called when exiting the moduleDecl production.
	ExitModuleDecl(c *ModuleDeclContext)

	// ExitTopLevelDecl is called when exiting the topLevelDecl production.
	ExitTopLevelDecl(c *TopLevelDeclContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitImportSpec is called when exiting the importSpec production.
	ExitImportSpec(c *ImportSpecContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitGenericParams is called when exiting the genericParams production.
	ExitGenericParams(c *GenericParamsContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitInterfaceDecl is called when exiting the interfaceDecl production.
	ExitInterfaceDecl(c *InterfaceDeclContext)

	// ExitInterfaceMethod is called when exiting the interfaceMethod production.
	ExitInterfaceMethod(c *InterfaceMethodContext)

	// ExitEnumDecl is called when exiting the enumDecl production.
	ExitEnumDecl(c *EnumDeclContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitConstDecl is called when exiting the constDecl production.
	ExitConstDecl(c *ConstDeclContext)

	// ExitConstField is called when exiting the constField production.
	ExitConstField(c *ConstFieldContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitOptionType is called when exiting the optionType production.
	ExitOptionType(c *OptionTypeContext)

	// ExitResultType is called when exiting the resultType production.
	ExitResultType(c *ResultTypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAssignOp is called when exiting the assignOp production.
	ExitAssignOp(c *AssignOpContext)

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitMatchStmt is called when exiting the matchStmt production.
	ExitMatchStmt(c *MatchStmtContext)

	// ExitMatchBranch is called when exiting the matchBranch production.
	ExitMatchBranch(c *MatchBranchContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitDeferStmt is called when exiting the deferStmt production.
	ExitDeferStmt(c *DeferStmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitMapElement is called when exiting the mapElement production.
	ExitMapElement(c *MapElementContext)

	// ExitStructLiteral is called when exiting the structLiteral production.
	ExitStructLiteral(c *StructLiteralContext)

	// ExitStructElement is called when exiting the structElement production.
	ExitStructElement(c *StructElementContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)
}
