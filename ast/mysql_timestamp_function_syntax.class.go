package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// TIMESTAMPADD、TIMESTAMPDIFF函数的统一结构
type MySqlTimestampFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
	Unit enum.MySqlTemporalInterval
}

func (this *MySqlTimestampFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	builder.writeStr(this.Unit.Sql)
	for i := 0; i < this.Parameters.Len(); i++ {
		builder.writeStr(", ")
		builder.writeSyntax(this.Parameters.elements[i])

	}
	builder.writeStr(")")
}

func NewMySqlTimestampFunctionSyntax() *MySqlTimestampFunctionSyntax {
	s := &MySqlTimestampFunctionSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	return s
}
