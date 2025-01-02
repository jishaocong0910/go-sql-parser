package ast

type JoinUsingSyntax struct {
	*M_Syntax
	*M_JoinConditionSyntax
	ColumnList I_ListSyntax[I_IdentifierSyntax]
}

func (this *JoinUsingSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitJoinUsingSyntax(this)
}

func (this *JoinUsingSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("USING")
	builder.writeSyntaxWithFormat(this.ColumnList, false)
}

func NewJoinUsingSyntax() *JoinUsingSyntax {
	s := &JoinUsingSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_JoinConditionSyntax = ExtendJoinConditionSyntax(s)
	return s
}
