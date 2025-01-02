package ast

// GROUP BY语法
type I_GroupBySyntax interface {
	I_Syntax
	M_7E13FD01759C() *M_GroupBySyntax
}

type M_GroupBySyntax struct {
	I                I_GroupBySyntax
	OrderingItemList *OrderingItemListSyntax
}

func (this *M_GroupBySyntax) M_7E13FD01759C() *M_GroupBySyntax {
	return this
}

func (this *M_GroupBySyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitGroupBySyntax(this)
}

func ExtendGroupBySyntax(i I_GroupBySyntax) *M_GroupBySyntax {
	return &M_GroupBySyntax{I: i}
}
