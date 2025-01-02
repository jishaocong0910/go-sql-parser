package ast

type MySqlGroupBySyntax struct {
	*M_Syntax
	*M_GroupBySyntax
	WithRollup bool
}

func (this *MySqlGroupBySyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("GROUP BY")
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntax(this.OrderingItemList)
	if this.WithRollup {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("WITH ROLLUP")
	}
}

func NewMySqlGroupBySyntax() *MySqlGroupBySyntax {
	s := &MySqlGroupBySyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_GroupBySyntax = ExtendGroupBySyntax(s)
	return s
}
