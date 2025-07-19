package ast

type PartitionBySyntax struct {
	*Syntax__
	Expr ExprSyntax_
}

func (this *PartitionBySyntax) accept(v_ Visitor_) {
	v_.visitor_().visitPartitionBySyntax(this)
}

func (this *PartitionBySyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("PARTITION BY ")
	builder.writeSyntax(this.Expr)
}

func NewPartitionBySyntax() *PartitionBySyntax {
	s := &PartitionBySyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
