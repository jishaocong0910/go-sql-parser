package ast

import "go-sql-parser/enum"

type MySqlIndexHintSyntax struct {
	*M_Syntax
	IndexHintMode enum.MySqlIndexHintMode
	IndexHintFor  enum.MySqlIndexHintFor
	IndexList     I_IdentifierListSyntax
}

func (this *MySqlIndexHintSyntax) accept(I_Visitor) {}

func (this *MySqlIndexHintSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.IndexHintMode.Sql)
	builder.writeStr(" INDEX")
	if !this.IndexHintFor.Undefined() {
		builder.writeStr(" FOR ")
		builder.writeStr(this.IndexHintFor.Sql)
	}
	if this.IndexList != nil {
		builder.writeSpace()
		builder.writeSyntax(this.IndexList)
	}
}

func NewMySqlIndexHintSyntax() *MySqlIndexHintSyntax {
	s := &MySqlIndexHintSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
