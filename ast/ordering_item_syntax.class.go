package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type OrderingItemSyntax struct {
	*Syntax__
	Column           ColumnItemSyntax_
	OrderingSequence enum.OrderingSequence
}

func (this *OrderingItemSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitOrderingItemSyntax(this)
}

func (this *OrderingItemSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Column)
	if !this.OrderingSequence.Undefined() {
		builder.writeSpace()
		builder.writeStr(this.OrderingSequence.Sql)
	}
}

func NewOrderingItemSyntax() *OrderingItemSyntax {
	s := &OrderingItemSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
