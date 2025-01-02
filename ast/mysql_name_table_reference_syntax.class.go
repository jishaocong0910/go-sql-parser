package ast

type MySqlNameTableReferenceSyntax struct {
	*M_Syntax
	*M_TableReferenceSyntax
	*M_NameTableReferenceSyntax
	PartitionList *PartitionListSyntax
	IndexHintList *MySqlIndexHintListSyntax
}

func (this *MySqlNameTableReferenceSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.TableNameItem)
	if this.Alias != nil {
		builder.writeSpace()
		builder.writeSyntax(this.Alias)
	}
	if this.PartitionList != nil {
		builder.writeSpace()
		builder.writeSyntax(this.PartitionList)
	}
	if this.IndexHintList != nil {
		builder.writeSpace()
		builder.writeSyntax(this.IndexHintList)
	}
}

func NewMySqlNameTableReferenceSyntax() *MySqlNameTableReferenceSyntax {
	s := &MySqlNameTableReferenceSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_TableReferenceSyntax = ExtendTableReferenceSyntax(s)
	s.M_NameTableReferenceSyntax = ExtendNameTableReferenceSyntax(s)
	return s
}
