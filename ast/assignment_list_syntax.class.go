package ast

type AssignmentListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*AssignmentSyntax]
}

func NewAssignmentListSyntax() *AssignmentListSyntax {
	s := &AssignmentListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*AssignmentSyntax](s)
	return s
}
