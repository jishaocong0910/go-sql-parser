package ast

import "go-sql-parser/enum"

type MySqlGetFormatFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.Name = "GET_FORMAT"
	return s
}
