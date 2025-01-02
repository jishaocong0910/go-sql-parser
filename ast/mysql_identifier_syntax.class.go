package ast

type MySqlIdentifierSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_AliasSyntax
	*M_ColumnItemSyntax
	*M_PropertyValueSyntax
	*M_OverWindowSyntax
	*M_IdentifierSyntax
	Qualifier bool
	sql       string
}

func (this *MySqlIdentifierSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlIdentifierSyntax(this)
}

func (this *MySqlIdentifierSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql())
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_AliasSyntax = ExtendAliasSyntax(s)
	s.M_ColumnItemSyntax = ExtendColumnItemSyntax(s)
	s.M_PropertyValueSyntax = ExtendPropertyValueSyntax(s)
	s.M_OverWindowSyntax = ExtendOverWindowSyntax(s)
	s.M_IdentifierSyntax = ExtendIdentifierSyntax(s)
	return s
}
