package ast

type MySqlNameTableReferenceSyntax struct {
	*Syntax__
	*TableReferenceSyntax__
	*NameTableReferenceSyntax__
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
	s.Syntax__ = ExtendSyntax(s)
	s.TableReferenceSyntax__ = ExtendTableReferenceSyntax(s)
	s.NameTableReferenceSyntax__ = ExtendNameTableReferenceSyntax(s)
	return s
}
