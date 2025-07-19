package ast

type TableNameItemSyntax struct {
	*Syntax__
	Catalog   IdentifierSyntax_
	TableName IdentifierSyntax_
}

func (this *TableNameItemSyntax) accept(Visitor_) {}

func (this *TableNameItemSyntax) writeSql(builder *sqlBuilder) {
	if this.Catalog != nil {
		builder.writeSyntax(this.Catalog)
		builder.writeStr(".")
	}
	builder.writeSyntax(this.TableName)
}

func (this *TableNameItemSyntax) FullTableName() string {
	if this.Catalog != nil {
		return this.Catalog.IdentifierSyntax_().Name + "." + this.TableName.IdentifierSyntax_().Name
	}
	return this.TableName.IdentifierSyntax_().Name
}

func NewTableNameItemSyntax() *TableNameItemSyntax {
	s := &TableNameItemSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
