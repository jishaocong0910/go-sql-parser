package ast

// 属性值
type PropertyValueSyntax_ interface {
	PropertyValueSyntax() *PropertyValueSyntax__
	Syntax_

	Value() string
}

type PropertyValueSyntax__ struct {
	I PropertyValueSyntax_
}

func (this *PropertyValueSyntax__) PropertyValueSyntax() *PropertyValueSyntax__ {
	return this
}

func ExtendPropertyValueSyntax(i PropertyValueSyntax_) *PropertyValueSyntax__ {
	return &PropertyValueSyntax__{I: i}
}
