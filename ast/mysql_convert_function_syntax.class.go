package ast

type MySqlConvertFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
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
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.Name = "CONVERT"
	return s
}
