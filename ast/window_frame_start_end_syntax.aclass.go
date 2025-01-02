package ast

type I_WindowFrameStartEndSyntax interface {
	I_WindowFrameExtentSyntax
	M_A233E334438C() *M_WindowFrameStartEndSyntax
}

type M_WindowFrameStartEndSyntax struct {
	I I_WindowFrameStartEndSyntax
}

func (this *M_WindowFrameStartEndSyntax) M_A233E334438C() *M_WindowFrameStartEndSyntax {
	return this
}

func ExtendWindowFrameStartEndSyntax(i I_WindowFrameStartEndSyntax) *M_WindowFrameStartEndSyntax {
	return &M_WindowFrameStartEndSyntax{I: i}
}
