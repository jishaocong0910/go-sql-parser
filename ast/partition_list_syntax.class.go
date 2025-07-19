package ast

// 查询语句的分区列表
type PartitionListSyntax struct {
	*Syntax__
	PartitionList IdentifierListSyntax_
}

func (this *PartitionListSyntax) accept(Visitor_) {}

func (this *PartitionListSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("PARTITION")
	builder.writeSyntaxWithFormat(this.PartitionList, false)
}

func NewPartitionListSyntax() *PartitionListSyntax {
	s := &PartitionListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
