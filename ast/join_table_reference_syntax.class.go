package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type JoinTableReferenceSyntax struct {
	*Syntax__
	*TableReferenceSyntax__
	Left          TableReferenceSyntax_
	Right         TableReferenceSyntax_
	Natural       bool
	JoinType      enum.JoinType
	JoinCondition JoinConditionSyntax_
}

func (this *JoinTableReferenceSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitJoinTableReferenceSyntax(this)
}

func (this *JoinTableReferenceSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Left)
	if enum.JoinType_.Not(this.JoinType, enum.JoinType_.COMMA) && !this.Natural {
		builder.writeSpaceOrLf(this, false)
	} else {
		builder.writeSpace()
	}
	if this.Natural {
		builder.writeStr("NATURAL ")
	}
	builder.writeStr(this.JoinType.Sql)
	builder.writeSpace()
	builder.writeSyntax(this.Right)
	if this.JoinCondition != nil {
		builder.writeSpaceOrLf(this, true)
		builder.writeSyntax(this.JoinCondition)
	}
}

func NewJoinTableReferenceSyntax() *JoinTableReferenceSyntax {
	s := &JoinTableReferenceSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.TableReferenceSyntax__ = ExtendTableReferenceSyntax(s)
	return s
}
