package ast

type ListSyntax[T Syntax_] struct {
	*Syntax__
	*ListSyntax__[T]
}

func NewListSyntax[T Syntax_]() *ListSyntax[T] {
	s := &ListSyntax[T]{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[T](s)
	return s
}
