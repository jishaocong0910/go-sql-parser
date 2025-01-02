package ast

import "go-sql-parser/enum"

type MySqlDateAndTimeLiteralSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Type        enum.MySqlDatetimeLiteralType
	DateAndTime *MySqlStringSyntax
}

func (this *MySqlDateAndTimeLiteralSyntax) accept(I_Visitor) {}

func (this *MySqlDateAndTimeLiteralSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Type.Sql)
	builder.writeSyntax(this.DateAndTime)
}

func NewMySqlDateAndTimeLiteralSyntax() *MySqlDateAndTimeLiteralSyntax {
	s := &MySqlDateAndTimeLiteralSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
