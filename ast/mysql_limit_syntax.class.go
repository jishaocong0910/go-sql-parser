package ast

type MySqlLimitSyntax struct {
	*Syntax__
	Offset   *DecimalNumberSyntax
	RowCount *DecimalNumberSyntax
}

func (this *MySqlLimitSyntax) accept(Visitor_) {}

func (this *MySqlLimitSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("LIMIT ")
	if this.Offset != nil {
		builder.writeSyntax(this.Offset)
		builder.writeStr(",")
	}
	builder.writeSyntax(this.RowCount)
}

func NewMySqlLimitSyntax() *MySqlLimitSyntax {
	s := &MySqlLimitSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
