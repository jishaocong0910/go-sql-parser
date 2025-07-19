package ast

type MySqlCastDataTypeSyntax struct {
	*Syntax__
	Name       string
	Parameters *MySqlCastDataTypeParamListSyntax
	// for CHAR type
	CharsetName string
}

func (this *MySqlCastDataTypeSyntax) accept(Visitor_) {}

func (this *MySqlCastDataTypeSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	if this.Parameters != nil {
		builder.writeSyntaxWithFormat(this.Parameters, false)
	}
	if this.CharsetName != "" {
		builder.writeStr(" CHARACTER SET ")
		builder.writeStr(this.CharsetName)
	}
}

func NewMySqlCastDataTypeSyntax() *MySqlCastDataTypeSyntax {
	s := &MySqlCastDataTypeSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
