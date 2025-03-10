package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlExtractFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.Name = "EXTRACT"
	return s
}
