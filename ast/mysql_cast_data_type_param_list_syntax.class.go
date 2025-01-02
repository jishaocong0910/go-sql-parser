package ast

type MySqlCastDataTypeParamListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*DecimalNumberSyntax]
}

func NewMySqlCastDataTypeParamListSyntax() *MySqlCastDataTypeParamListSyntax {
	s := &MySqlCastDataTypeParamListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*DecimalNumberSyntax](s)
	return s
}
