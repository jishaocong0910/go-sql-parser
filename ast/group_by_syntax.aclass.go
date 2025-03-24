package ast

// GROUP BY语法
type I_GroupBySyntax interface {
	I_Syntax
	M_GroupBySyntax_() *M_GroupBySyntax
}

type M_GroupBySyntax struct {
	I                I_GroupBySyntax
	OrderingItemList *OrderingItemListSyntax
}

func (this *M_GroupBySyntax) M_GroupBySyntax_() *M_GroupBySyntax {
	return this
}

func (this *M_GroupBySyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitGroupBySyntax(this)
}

func ExtendGroupBySyntax(i I_GroupBySyntax) *M_GroupBySyntax {
	return &M_GroupBySyntax{I: i}
}
