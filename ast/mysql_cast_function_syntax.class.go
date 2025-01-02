package ast

type MySqlCastFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.Name = "CAST"
	return s
}
