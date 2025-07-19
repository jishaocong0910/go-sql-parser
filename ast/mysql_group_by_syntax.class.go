package ast

type MySqlGroupBySyntax struct {
	*Syntax__
	*GroupBySyntax__
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
	s.Syntax__ = ExtendSyntax(s)
	s.GroupBySyntax__ = ExtendGroupBySyntax(s)
	return s
}
