package ast

type MySqlMultiDeleteTableAliasSyntax struct {
	*Syntax__
	Alias   *MySqlIdentifierSyntax
	HasStar bool
}

func (this *MySqlMultiDeleteTableAliasSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlMultiDeleteTableAliasSyntax(this)
}

func (this *MySqlMultiDeleteTableAliasSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Alias)
	if this.HasStar {
		builder.writeStr(".*")
	}
}

func NewMySqlMultiDeleteTableAliasSyntax() *MySqlMultiDeleteTableAliasSyntax {
	s := &MySqlMultiDeleteTableAliasSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
