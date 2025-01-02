package ast

type BinaryNumberSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Sql string
}

func (this *BinaryNumberSyntax) accept(I_Visitor) {}

func (this *BinaryNumberSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewBinaryNumberSyntax() *BinaryNumberSyntax {
	s := &BinaryNumberSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
