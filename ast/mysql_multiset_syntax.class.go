package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlMultisetSyntax struct {
	*Syntax__
	*ExprSyntax__
	*StatementSyntax__
	*QuerySyntax__
	*MultisetSyntax__
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
	return this.MultisetSyntax__.OperandCount()
}

func (this *MySqlMultisetSyntax) Dialect() enum.Dialect {
	return enum.Dialect_.MYSQL
}

func NewMySqlMultisetSyntax() *MySqlMultisetSyntax {
	s := &MySqlMultisetSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.StatementSyntax__ = ExtendStatementSyntax(s)
	s.QuerySyntax__ = ExtendQuerySyntax(s)
	s.MultisetSyntax__ = ExtendMultisetSyntax(s)
	return s
}
