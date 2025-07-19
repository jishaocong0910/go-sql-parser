package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlTrimFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	TrimMode enum.MySqlTrimMode
	RemStr   ExprSyntax_
	Str      ExprSyntax_
}

func (this *MySqlTrimFunctionSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlTrimFunctionSyntax(this)
}

func (this *MySqlTrimFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	if !this.TrimMode.Undefined() {
		builder.writeStr(this.TrimMode.Sql)
		builder.writeSpace()
	}
	if this.RemStr != nil {
		builder.writeSyntax(this.RemStr)
		builder.writeStr(" FROM ")
	}
	builder.writeSyntax(this.Str)
	builder.writeStr(")")
}

func NewMySqlTrimFunctionSyntax() *MySqlTrimFunctionSyntax {
	s := &MySqlTrimFunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.Name = "TRIM"
	return s
}
