package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type JoinTableReferenceSyntax struct {
	*M_Syntax
	*M_TableReferenceSyntax
	Left          I_TableReferenceSyntax
	Right         I_TableReferenceSyntax
	Natural       bool
	JoinType      enum.JoinType
	JoinCondition I_JoinConditionSyntax
}

func (this *JoinTableReferenceSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitJoinTableReferenceSyntax(this)
}

func (this *JoinTableReferenceSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Left)
	if enum.JoinTypes.Not(this.JoinType, enum.JoinTypes.COMMA) && !this.Natural {
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_TableReferenceSyntax = ExtendTableReferenceSyntax(s)
	return s
}
