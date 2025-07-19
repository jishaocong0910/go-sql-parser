package ast

type JoinOnSyntax struct {
	*Syntax__
	*JoinConditionSyntax__
	Condition ExprSyntax_
}

func (this *JoinOnSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitJoinOnSyntax(this)
}

func (this *JoinOnSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("ON ")
	builder.writeSyntax(this.Condition)
}

func NewJoinOnSyntax() *JoinOnSyntax {
	s := &JoinOnSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.JoinConditionSyntax__ = ExtendJoinConditionSyntax(s)
	return s
}
