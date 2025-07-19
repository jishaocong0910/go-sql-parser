package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlGetFormatFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	Type       enum.MySqlGetFormatType
	DateFormat *MySqlStringSyntax
}

func (this *MySqlGetFormatFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	builder.writeStr(this.Type.Sql)
	builder.writeStr(", ")
	builder.writeSyntax(this.DateFormat)
	builder.writeStr(")")
}

func NewMySqlGetFormatFunctionSyntax() *MySqlGetFormatFunctionSyntax {
	s := &MySqlGetFormatFunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.Name = "GET_FORMAT"
	return s
}
