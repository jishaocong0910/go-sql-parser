package ast

// 标识符
type IdentifierSyntax_ interface {
	IdentifierSyntax_() *IdentifierSyntax__
	ExprSyntax_
	AliasSyntax_
	ColumnItemSyntax_
	PropertyValueSyntax_
	OverWindowSyntax_
	Sql() string
}

type IdentifierSyntax__ struct {
	I    IdentifierSyntax_
	Name string
}

func (this *IdentifierSyntax__) IdentifierSyntax_() *IdentifierSyntax__ {
	return this
}

func (this *IdentifierSyntax__) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.I.Sql())
}

func (this *IdentifierSyntax__) AliasName() string {
	return this.Name
}

func (this *IdentifierSyntax__) TableAlias() string {
	return ""
}

func (this *IdentifierSyntax__) Column() string {
	return this.Name
}

func (this *IdentifierSyntax__) FullColumn() string {
	return this.Name
}

func (this *IdentifierSyntax__) Value() string {
	return this.Name
}

func ExtendIdentifierSyntax(i IdentifierSyntax_) *IdentifierSyntax__ {
	return &IdentifierSyntax__{I: i}
}
