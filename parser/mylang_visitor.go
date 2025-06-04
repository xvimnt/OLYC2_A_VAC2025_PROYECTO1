// Code generated from ./grammar/MyLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MyLang
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by MyLangParser.
type MyLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MyLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MyLangParser#moduleDecl.
	VisitModuleDecl(ctx *ModuleDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#topLevelDecl.
	VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by MyLangParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#genericParams.
	VisitGenericParams(ctx *GenericParamsContext) interface{}

	// Visit a parse tree produced by MyLangParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by MyLangParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by MyLangParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by MyLangParser#interfaceDecl.
	VisitInterfaceDecl(ctx *InterfaceDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#interfaceMethod.
	VisitInterfaceMethod(ctx *InterfaceMethodContext) interface{}

	// Visit a parse tree produced by MyLangParser#enumDecl.
	VisitEnumDecl(ctx *EnumDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#enumField.
	VisitEnumField(ctx *EnumFieldContext) interface{}

	// Visit a parse tree produced by MyLangParser#constDecl.
	VisitConstDecl(ctx *ConstDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#constField.
	VisitConstField(ctx *ConstFieldContext) interface{}

	// Visit a parse tree produced by MyLangParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by MyLangParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by MyLangParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by MyLangParser#optionType.
	VisitOptionType(ctx *OptionTypeContext) interface{}

	// Visit a parse tree produced by MyLangParser#resultType.
	VisitResultType(ctx *ResultTypeContext) interface{}

	// Visit a parse tree produced by MyLangParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MyLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MyLangParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by MyLangParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by MyLangParser#assignOp.
	VisitAssignOp(ctx *AssignOpContext) interface{}

	// Visit a parse tree produced by MyLangParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by MyLangParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by MyLangParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by MyLangParser#matchStmt.
	VisitMatchStmt(ctx *MatchStmtContext) interface{}

	// Visit a parse tree produced by MyLangParser#matchBranch.
	VisitMatchBranch(ctx *MatchBranchContext) interface{}

	// Visit a parse tree produced by MyLangParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by MyLangParser#deferStmt.
	VisitDeferStmt(ctx *DeferStmtContext) interface{}

	// Visit a parse tree produced by MyLangParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by MyLangParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by MyLangParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by MyLangParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by MyLangParser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by MyLangParser#mapElement.
	VisitMapElement(ctx *MapElementContext) interface{}

	// Visit a parse tree produced by MyLangParser#structLiteral.
	VisitStructLiteral(ctx *StructLiteralContext) interface{}

	// Visit a parse tree produced by MyLangParser#structElement.
	VisitStructElement(ctx *StructElementContext) interface{}

	// Visit a parse tree produced by MyLangParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by MyLangParser#returnType.
	VisitReturnType(ctx *ReturnTypeContext) interface{}

}