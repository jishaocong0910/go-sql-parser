package ast

type HexadecimalNumberSyntax struct {
	*Syntax__
	*ExprSyntax__
	Sql string
}

func (this *HexadecimalNumberSyntax) accept(Visitor_) {}

func (this *HexadecimalNumberSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewHexadecimalNumberSyntax() *HexadecimalNumberSyntax {
	s := &HexadecimalNumberSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
