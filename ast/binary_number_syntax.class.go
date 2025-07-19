package ast

type BinaryNumberSyntax struct {
	*Syntax__
	*ExprSyntax__
	Sql string
}

func (this *BinaryNumberSyntax) accept(Visitor_) {}

func (this *BinaryNumberSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewBinaryNumberSyntax() *BinaryNumberSyntax {
	s := &BinaryNumberSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
