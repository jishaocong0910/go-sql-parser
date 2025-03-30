package ast

// UPDATE语法
type I_UpdateSyntax interface {
	I_StatementSyntax
	I_HaveWhereSyntax
	M_UpdateSyntax_() *M_UpdateSyntax
}

type M_UpdateSyntax struct {
	I              I_UpdateSyntax
	AssignmentList *AssignmentListSyntax
	TableReference I_TableReferenceSyntax
	Hint           *HintSyntax
}

func (this *M_UpdateSyntax) M_UpdateSyntax_() *M_UpdateSyntax {
	return this
}

func ExtendUpdateSyntax(i I_UpdateSyntax) *M_UpdateSyntax {
	return &M_UpdateSyntax{I: i}
}
