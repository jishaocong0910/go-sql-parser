package ast

// 属性值
type I_PropertyValueSyntax interface {
	I_Syntax
	M_PropertyValueSyntax_() *M_PropertyValueSyntax
	Value() string
}

type M_PropertyValueSyntax struct {
	I I_PropertyValueSyntax
}

func (this *M_PropertyValueSyntax) M_PropertyValueSyntax_() *M_PropertyValueSyntax {
	return this
}

func ExtendPropertyValueSyntax(i I_PropertyValueSyntax) *M_PropertyValueSyntax {
	return &M_PropertyValueSyntax{I: i}
}
