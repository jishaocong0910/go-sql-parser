package ast

type MySqlCharFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	CharsetName string
}

func (this *MySqlCharFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	separator := ""
	for i := 0; i < this.Parameters.Len(); i++ {
		builder.writeStr(separator)
		builder.writeSyntax(this.Parameters.elements[i])
		separator = ", "
	}
	if this.CharsetName != "" {
		builder.writeStr(" USING ")
		builder.writeStr(this.CharsetName)
	}
	builder.writeStr(")")
}

func NewMySqlCharFunctionSyntax() *MySqlCharFunctionSyntax {
	s := &MySqlCharFunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.Name = "CHAR"
	return s
}
