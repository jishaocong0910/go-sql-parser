package ast

type MySqlConvertFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
	UsingTranscoding bool
	TranscodingName  string
	DataType         *MySqlCastDataTypeSyntax
	Collate          string
}

func (this *MySqlConvertFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	builder.writeSyntax(this.Parameters.elements[0])
	if this.UsingTranscoding {
		builder.writeStr(" USING ")
		builder.writeStr(this.TranscodingName)
	} else {
		builder.writeStr(", ")
		builder.writeSyntax(this.DataType)
	}
	builder.writeStr(")")
	if this.Collate != "" {
		builder.writeStr(" COLLATE ")
		builder.writeStr(this.Collate)
	}
}

func NewMySqlConvertFunctionSyntax() *MySqlConvertFunctionSyntax {
	s := &MySqlConvertFunctionSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.Name = "CONVERT"
	return s
}
