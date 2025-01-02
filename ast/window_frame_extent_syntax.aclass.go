package ast

type I_WindowFrameExtentSyntax interface {
	I_Syntax
	M_57DD18D91B86() *M_WindowFrameExtentSyntax
}

type M_WindowFrameExtentSyntax struct {
	I I_WindowFrameExtentSyntax
}

func (this *M_WindowFrameExtentSyntax) M_57DD18D91B86() *M_WindowFrameExtentSyntax {
	return this
}

func ExtendWindowFrameExtentSyntax(i I_WindowFrameExtentSyntax) *M_WindowFrameExtentSyntax {
	return &M_WindowFrameExtentSyntax{I: i}
}
