package ast

type WindowFrameExtentSyntax_ interface {
	WindowFrameExtentSyntax_() *WindowFrameExtentSyntax__
	Syntax_
}

type WindowFrameExtentSyntax__ struct {
	I WindowFrameExtentSyntax_
}

func (this *WindowFrameExtentSyntax__) WindowFrameExtentSyntax_() *WindowFrameExtentSyntax__ {
	return this
}

func ExtendWindowFrameExtentSyntax(i WindowFrameExtentSyntax_) *WindowFrameExtentSyntax__ {
	return &WindowFrameExtentSyntax__{I: i}
}
