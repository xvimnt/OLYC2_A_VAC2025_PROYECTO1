// Code generated from ./grammar/MyLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // MyLang
import "github.com/antlr4-go/antlr/v4"


type BaseMyLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMyLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitModuleDecl(ctx *ModuleDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitGenericParams(ctx *GenericParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitInterfaceDecl(ctx *InterfaceDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitInterfaceMethod(ctx *InterfaceMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitEnumDecl(ctx *EnumDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitEnumField(ctx *EnumFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitConstField(ctx *ConstFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitOptionType(ctx *OptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitResultType(ctx *ResultTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitAssignOp(ctx *AssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitMatchStmt(ctx *MatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitMatchBranch(ctx *MatchBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitMapElement(ctx *MapElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitStructElement(ctx *StructElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMyLangVisitor) VisitReturnType(ctx *ReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
