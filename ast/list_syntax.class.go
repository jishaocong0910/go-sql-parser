package ast

type ListSyntax[T I_Syntax] struct {
	*M_Syntax
	*M_ListSyntax[T]
}

func NewListSyntax[T I_Syntax]() *ListSyntax[T] {
	s := &ListSyntax[T]{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[T](s)
	return s
}
