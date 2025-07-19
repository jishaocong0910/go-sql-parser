package ast

type JoinUsingSyntax struct {
	*Syntax__
	*JoinConditionSyntax__
	ColumnList ListSyntax_[IdentifierSyntax_]
}

func (this *JoinUsingSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitJoinUsingSyntax(this)
}

func (this *JoinUsingSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("USING")
	builder.writeSyntaxWithFormat(this.ColumnList, false)
}

func NewJoinUsingSyntax() *JoinUsingSyntax {
	s := &JoinUsingSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.JoinConditionSyntax__ = ExtendJoinConditionSyntax(s)
	return s
}
