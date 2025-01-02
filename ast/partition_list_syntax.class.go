package ast

// 查询语句的分区列表
type PartitionListSyntax struct {
	*M_Syntax
	PartitionList I_IdentifierListSyntax
}

func (this *PartitionListSyntax) accept(I_Visitor) {}

func (this *PartitionListSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("PARTITION")
	builder.writeSyntaxWithFormat(this.PartitionList, false)
}

func NewPartitionListSyntax() *PartitionListSyntax {
	s := &PartitionListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
