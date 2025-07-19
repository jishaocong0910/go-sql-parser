package ast

type AssignmentListSyntax struct {
	*Syntax__
	*ListSyntax__[*AssignmentSyntax]
}

func NewAssignmentListSyntax() *AssignmentListSyntax {
	s := &AssignmentListSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*AssignmentSyntax](s)
	return s
}
