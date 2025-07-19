package ast

type MySqlCastDataTypeParamListSyntax struct {
	*Syntax__
	*ListSyntax__[*DecimalNumberSyntax]
}

func NewMySqlCastDataTypeParamListSyntax() *MySqlCastDataTypeParamListSyntax {
	s := &MySqlCastDataTypeParamListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*DecimalNumberSyntax](s)
	return s
}
