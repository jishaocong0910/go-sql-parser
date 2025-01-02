package ast

type PartitionBySyntax struct {
	*M_Syntax
	Expr I_ExprSyntax
}

func (this *PartitionBySyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitPartitionBySyntax(this)
}

func (this *PartitionBySyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("PARTITION BY ")
	builder.writeSyntax(this.Expr)
}

func NewPartitionBySyntax() *PartitionBySyntax {
	s := &PartitionBySyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
