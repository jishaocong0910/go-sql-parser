package ast

type AllColumnSyntax struct {
	*M_Syntax
	*M_SelectItemSyntax
	*M_PropertyValueSyntax
}

func (this *AllColumnSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitAllColumnSyntax(this)
}

func (this *AllColumnSyntax) Value() string {
	return "*"
}

func (this *AllColumnSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("*")
}

func NewAllColumnSyntax() *AllColumnSyntax {
	s := &AllColumnSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_SelectItemSyntax = ExtendSelectItemSyntax(s)
	s.M_PropertyValueSyntax = ExtendPropertyValueSyntax(s)
	return s
}
