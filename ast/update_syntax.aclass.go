package ast

// UPDATE语法
type I_UpdateSyntax interface {
	I_StatementSyntax
	M_2A4708829A9C() *M_UpdateSyntax
}

type M_UpdateSyntax struct {
	I              I_UpdateSyntax
	AssignmentList *AssignmentListSyntax
	TableReference I_TableReferenceSyntax
	Where          *WhereSyntax
	Hint           *HintSyntax
}

func (this *M_UpdateSyntax) M_2A4708829A9C() *M_UpdateSyntax {
	return this
}

func ExtendUpdateSyntax(i I_UpdateSyntax) *M_UpdateSyntax {
	return &M_UpdateSyntax{I: i}
}
