package ast

type MySqlMultiDeleteTableAliasSyntax struct {
	*M_Syntax
	Alias   *MySqlIdentifierSyntax
	HasStar bool
}

func (this *MySqlMultiDeleteTableAliasSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlMultiDeleteTableAliasSyntax(this)
}

func (this *MySqlMultiDeleteTableAliasSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Alias)
	if this.HasStar {
		builder.writeStr(".*")
	}
}

func NewMySqlMultiDeleteTableAliasSyntax() *MySqlMultiDeleteTableAliasSyntax {
	s := &MySqlMultiDeleteTableAliasSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
