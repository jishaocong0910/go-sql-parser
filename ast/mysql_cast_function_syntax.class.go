package ast

type MySqlCastFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	DataType   *MySqlCastDataTypeSyntax
	AtTimeZone *MySqlStringSyntax
	Collate    string
}

func (this *MySqlCastFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	builder.writeSyntax(this.Parameters.elements[0])
	if this.AtTimeZone != nil {
		builder.writeStr(" AT TIME ZONE ")
		builder.writeSyntax(this.AtTimeZone)
	}
	builder.writeStr(" AS ")
	builder.writeSyntax(this.DataType)
	builder.writeStr(")")
	if this.Collate != "" {
		builder.writeStr(" COLLATE ")
		builder.writeStr(this.Collate)
	}
}

func NewMySqlCastFunctionSyntax() *MySqlCastFunctionSyntax {
	s := &MySqlCastFunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.Name = "CAST"
	return s
}
