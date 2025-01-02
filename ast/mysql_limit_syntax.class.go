package ast

type MySqlLimitSyntax struct {
	*M_Syntax
	Offset   *DecimalNumberSyntax
	RowCount *DecimalNumberSyntax
}

func (this *MySqlLimitSyntax) accept(I_Visitor) {}

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
	s.M_Syntax = ExtendSyntax(s)
	return s
}
