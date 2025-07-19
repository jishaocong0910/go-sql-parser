package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type WindowFrameExprSyntax struct {
	*Syntax__
	*WindowFrameExtentSyntax__
	*WindowFrameStartEndSyntax__
	Expr ExprSyntax_
	Type enum.WindowFrameStartEndType
}

func (this *WindowFrameExprSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitWindowFrameExprSyntax(this)
}

func (this *WindowFrameExprSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Expr)
	builder.writeStr(" PRECEDING")
}

func NewWindowFrameExprSyntax() *WindowFrameExprSyntax {
	s := &WindowFrameExprSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.WindowFrameExtentSyntax__ = ExtendWindowFrameExtentSyntax(s)
	s.WindowFrameStartEndSyntax__ = ExtendWindowFrameStartEndSyntax(s)
	return s
}
