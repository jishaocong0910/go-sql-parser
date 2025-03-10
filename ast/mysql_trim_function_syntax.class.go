package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlTrimFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
	TrimMode enum.MySqlTrimMode
	RemStr   I_ExprSyntax
	Str      I_ExprSyntax
}

func (this *MySqlTrimFunctionSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlTrimFunctionSyntax(this)
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.Name = "TRIM"
	return s
}
