package ast

type TableNameItemSyntax struct {
	*M_Syntax
	Catalog   I_IdentifierSyntax
	TableName I_IdentifierSyntax
}

func (this *TableNameItemSyntax) accept(I_Visitor) {}

func (this *TableNameItemSyntax) writeSql(builder *sqlBuilder) {
	if this.Catalog != nil {
		builder.writeSyntax(this.Catalog)
		builder.writeStr(".")
	}
	builder.writeSyntax(this.TableName)
}

func (this *TableNameItemSyntax) FullTableName() string {
	if this.Catalog != nil {
		return this.Catalog.M_A2CE003580A2().Name + "." + this.TableName.M_A2CE003580A2().Name
	}
	return this.TableName.M_A2CE003580A2().Name
}

func NewTableNameItemSyntax() *TableNameItemSyntax {
	s := &TableNameItemSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
