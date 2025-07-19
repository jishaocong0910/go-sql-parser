package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlIndexHintSyntax struct {
	*Syntax__
	IndexHintMode enum.MySqlIndexHintMode
	IndexHintFor  enum.MySqlIndexHintFor
	IndexList     IdentifierListSyntax_
}

func (this *MySqlIndexHintSyntax) accept(Visitor_) {}

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
	s.Syntax__ = ExtendSyntax(s)
	return s
}
