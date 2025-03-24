package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type OrderingItemSyntax struct {
	*M_Syntax
	Column           I_ColumnItemSyntax
	OrderingSequence enum.OrderingSequence
}

func (this *OrderingItemSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitOrderingItemSyntax(this)
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
	s.M_Syntax = ExtendSyntax(s)
	return s
}
