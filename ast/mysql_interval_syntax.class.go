package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlIntervalSyntax struct {
	*Syntax__
	*ExprSyntax__
	Expr ExprSyntax_
	Unit enum.MySqlTemporalInterval
}

func (this *MySqlIntervalSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlIntervalSyntax(this)
}

func (this *MySqlIntervalSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("INTERVAL ")
	builder.writeSyntax(this.Expr)
	builder.writeSpace()
	builder.writeStr(this.Unit.Sql)
}

func NewMySqlIntervalSyntax() *MySqlIntervalSyntax {
	s := &MySqlIntervalSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.ParenthesizeType = enum.ParenthesizeType_.NOT_SUPPORT
	return s
}
