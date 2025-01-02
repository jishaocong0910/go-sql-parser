package ast

type DecimalNumberSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Sql string
}

func (this *DecimalNumberSyntax) accept(I_Visitor) {}

func (this *DecimalNumberSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewDecimalNumberSyntax() *DecimalNumberSyntax {
	s := &DecimalNumberSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
