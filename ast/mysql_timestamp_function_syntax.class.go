package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// TIMESTAMPADD、TIMESTAMPDIFF函数的统一结构
type MySqlTimestampFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
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
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	return s
}
