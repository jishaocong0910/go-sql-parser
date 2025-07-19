package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlExtractFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	Unit enum.MySqlTemporalInterval
}

func (this *MySqlExtractFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	builder.writeStr(this.Unit.Sql)
	builder.writeStr(" FROM ")
	builder.writeSyntax(this.Parameters.elements[0])
	builder.writeStr(")")
}

func NewMySqlExtractFunctionSyntax() *MySqlExtractFunctionSyntax {
	s := &MySqlExtractFunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.Name = "EXTRACT"
	return s
}
