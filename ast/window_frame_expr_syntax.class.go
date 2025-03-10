package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type WindowFrameExprSyntax struct {
	*M_Syntax
	*M_WindowFrameExtentSyntax
	*M_WindowFrameStartEndSyntax
	Expr I_ExprSyntax
	Type enum.WindowFrameStartEndType
}

func (this *WindowFrameExprSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitWindowFrameExprSyntax(this)
}

func (this *WindowFrameExprSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Expr)
	builder.writeStr(" PRECEDING")
}

func NewWindowFrameExprSyntax() *WindowFrameExprSyntax {
	s := &WindowFrameExprSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_WindowFrameExtentSyntax = ExtendWindowFrameExtentSyntax(s)
	s.M_WindowFrameStartEndSyntax = ExtendWindowFrameStartEndSyntax(s)
	return s
}
