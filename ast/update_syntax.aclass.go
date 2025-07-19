package ast

// UPDATE语法
type UpdateSyntax_ interface {
	UpdateSyntax_() *UpdateSyntax__
	StatementSyntax_
	HaveWhereSyntax_
}

type UpdateSyntax__ struct {
	I              UpdateSyntax_
	AssignmentList *AssignmentListSyntax
	TableReference TableReferenceSyntax_
	Hint           *HintSyntax
}

func (this *UpdateSyntax__) UpdateSyntax_() *UpdateSyntax__ {
	return this
}

func ExtendUpdateSyntax(i UpdateSyntax_) *UpdateSyntax__ {
	return &UpdateSyntax__{I: i}
}
