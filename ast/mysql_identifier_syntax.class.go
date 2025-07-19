package ast

type MySqlIdentifierSyntax struct {
	*Syntax__
	*ExprSyntax__
	*AliasSyntax__
	*ColumnItemSyntax__
	*PropertyValueSyntax__
	*OverWindowSyntax__
	*IdentifierSyntax__
	Qualifier bool
	sql       string
}

func (this *MySqlIdentifierSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlIdentifierSyntax(this)
}

func (this *MySqlIdentifierSyntax) Sql() string {
	if this.sql == "" {
		if this.Qualifier {
			this.sql = "`" + this.Name + "`"
		} else {
			this.sql = this.Name
		}
	}
	return this.sql
}

func NewMySqlIdentifierSyntax() *MySqlIdentifierSyntax {
	s := &MySqlIdentifierSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.AliasSyntax__ = ExtendAliasSyntax(s)
	s.ColumnItemSyntax__ = ExtendColumnItemSyntax(s)
	s.PropertyValueSyntax__ = ExtendPropertyValueSyntax(s)
	s.OverWindowSyntax__ = ExtendOverWindowSyntax(s)
	s.IdentifierSyntax__ = ExtendIdentifierSyntax(s)
	return s
}
