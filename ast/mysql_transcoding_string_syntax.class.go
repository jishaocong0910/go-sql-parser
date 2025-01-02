package ast

type MySqlTranscodingStringSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	CharsetName string
	Str         *MySqlStringSyntax
	Collate     string
}

func (this *MySqlTranscodingStringSyntax) accept(I_Visitor) {}

func (this *MySqlTranscodingStringSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("_")
	builder.writeStr(this.CharsetName)
	builder.writeSyntax(this.Str)
	if this.Collate != "" {
		builder.writeStr(" COLLATE ")
		builder.writeStr(this.Collate)
	}
}

func NewMySqlTranscodingStringSyntax() *MySqlTranscodingStringSyntax {
	s := &MySqlTranscodingStringSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
