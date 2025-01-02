package ast

// 标识符
type I_IdentifierSyntax interface {
	I_ExprSyntax
	I_AliasSyntax
	I_ColumnItemSyntax
	I_PropertyValueSyntax
	I_OverWindowSyntax
	M_A2CE003580A2() *M_IdentifierSyntax
}

type M_IdentifierSyntax struct {
	I    I_IdentifierSyntax
	Name string
}

func (this *M_IdentifierSyntax) M_A2CE003580A2() *M_IdentifierSyntax {
	return this
}

func (this *M_IdentifierSyntax) AliasName() string {
	return this.Name
}

func (this *M_IdentifierSyntax) TableAlias() string {
	return ""
}

func (this *M_IdentifierSyntax) Column() string {
	return this.Name
}

func (this *M_IdentifierSyntax) FullColumn() string {
	return this.Name
}

func (this *M_IdentifierSyntax) Value() string {
	return this.Name
}

func ExtendIdentifierSyntax(i I_IdentifierSyntax) *M_IdentifierSyntax {
	return &M_IdentifierSyntax{I: i}
}
