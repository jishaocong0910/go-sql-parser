package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlDateAndTimeLiteralSyntax struct {
	*Syntax__
	*ExprSyntax__
	Type        enum.MySqlDatetimeLiteralType
	DateAndTime *MySqlStringSyntax
}

func (this *MySqlDateAndTimeLiteralSyntax) accept(Visitor_) {}

func (this *MySqlDateAndTimeLiteralSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Type.Sql)
	builder.writeSyntax(this.DateAndTime)
}

func NewMySqlDateAndTimeLiteralSyntax() *MySqlDateAndTimeLiteralSyntax {
	s := &MySqlDateAndTimeLiteralSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
