package ast

type MySqlTranscodingStringSyntax struct {
	*Syntax__
	*ExprSyntax__
	CharsetName string
	Str         *MySqlStringSyntax
	Collate     string
}

func (this *MySqlTranscodingStringSyntax) accept(Visitor_) {}

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
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
