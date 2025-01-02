package ast

type HexadecimalNumberSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Sql string
}

func (this *HexadecimalNumberSyntax) accept(I_Visitor) {}

func (this *HexadecimalNumberSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewHexadecimalNumberSyntax() *HexadecimalNumberSyntax {
	s := &HexadecimalNumberSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
