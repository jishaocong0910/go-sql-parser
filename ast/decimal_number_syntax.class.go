package ast

type DecimalNumberSyntax struct {
	*Syntax__
	*ExprSyntax__
	Sql string
}

func (this *DecimalNumberSyntax) accept(Visitor_) {}

func (this *DecimalNumberSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewDecimalNumberSyntax() *DecimalNumberSyntax {
	s := &DecimalNumberSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
