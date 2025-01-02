package ast

type MySqlCharFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.Name = "CHAR"
	return s
}
