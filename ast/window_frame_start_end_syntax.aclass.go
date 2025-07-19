package ast

type WindowFrameStartEndSyntax_ interface {
	WindowFrameStartEndSyntax_() *WindowFrameStartEndSyntax__
	WindowFrameExtentSyntax_
}

type WindowFrameStartEndSyntax__ struct {
	I WindowFrameStartEndSyntax_
}

func (this *WindowFrameStartEndSyntax__) WindowFrameStartEndSyntax_() *WindowFrameStartEndSyntax__ {
	return this
}

func ExtendWindowFrameStartEndSyntax(i WindowFrameStartEndSyntax_) *WindowFrameStartEndSyntax__ {
	return &WindowFrameStartEndSyntax__{I: i}
}
