package ast

import "go-sql-parser/enum"

type MySqlMultisetSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_StatementSyntax
	*M_QuerySyntax
	*M_MultisetSyntax
	Limit *MySqlLimitSyntax
}

func (this *MySqlMultisetSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.LeftQuery)
	builder.writeSpaceOrLf(this, false)
	builder.writeStr(this.MultisetOperator.Sql)
	if !this.AggregateOption.Undefined() {
		builder.writeSpace()
		builder.writeStr(this.AggregateOption.Sql)
	}
	builder.writeSpaceOrLf(this, false)
	builder.writeSyntax(this.RightQuery)
	if this.OrderBy != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.OrderBy)
	}
	if this.Limit != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.Limit)
	}
}

func (this *MySqlMultisetSyntax) OperandCount() int {
	return this.M_MultisetSyntax.OperandCount()
}

func (this *MySqlMultisetSyntax) Dialect() enum.Dialect {
	return enum.Dialects.MYSQL
}

func NewMySqlMultisetSyntax() *MySqlMultisetSyntax {
	s := &MySqlMultisetSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_StatementSyntax = ExtendStatementSyntax(s)
	s.M_QuerySyntax = ExtendQuerySyntax(s)
	s.M_MultisetSyntax = ExtendMultisetSyntax(s)
	return s
}
