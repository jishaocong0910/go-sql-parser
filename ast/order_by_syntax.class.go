package ast

type OrderBySyntax struct {
	*Syntax__
	OrderByItemList *OrderingItemListSyntax
}

func (this *OrderBySyntax) accept(v_ Visitor_) {
	v_.visitor_().visitOrderBySyntax(this)
}

func (this *OrderBySyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("ORDER BY")
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntaxWithFormat(this.OrderByItemList, this.Format)
}

func NewOrderBySyntax() *OrderBySyntax {
	s := &OrderBySyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
