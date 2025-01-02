package ast

type JoinOnSyntax struct {
	*M_Syntax
	*M_JoinConditionSyntax
	Condition I_ExprSyntax
}

func (this *JoinOnSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitJoinOnSyntax(this)
}

func (this *JoinOnSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("ON ")
	builder.writeSyntax(this.Condition)
}

func NewJoinOnSyntax() *JoinOnSyntax {
	s := &JoinOnSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_JoinConditionSyntax = ExtendJoinConditionSyntax(s)
	return s
}
