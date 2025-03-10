package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlIntervalSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Expr I_ExprSyntax
	Unit enum.MySqlTemporalInterval
}

func (this *MySqlIntervalSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlIntervalSyntax(this)
}

func (this *MySqlIntervalSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("INTERVAL ")
	builder.writeSyntax(this.Expr)
	builder.writeSpace()
	builder.writeStr(this.Unit.Sql)
}

func NewMySqlIntervalSyntax() *MySqlIntervalSyntax {
	s := &MySqlIntervalSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.ParenthesizeType = enum.ParenthesizeTypes.NOT_SUPPORT
	return s
}
