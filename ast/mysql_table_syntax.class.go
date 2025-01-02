package ast

// MySqlTableSyntax
//
//	@Description: https://dev.mysql.com/doc/refman/8.0/en/table.html
type MySqlTableSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	TableNameItem *TableNameItemSyntax
	OrderBy       *OrderBySyntax
	Limit         *MySqlLimitSyntax
}

func (this *MySqlTableSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlTableSyntax(this)
}

func (this *MySqlTableSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("TABLE ")
	builder.writeSyntax(this.TableNameItem)
	if this.OrderBy != nil {
		builder.writeSpace()
		builder.writeSyntaxWithFormat(this.OrderBy, false)
	}
	if this.Limit != nil {
		builder.writeSpace()
		builder.writeSyntaxWithFormat(this.Limit, false)
	}
}

func NewMySqlTableSyntax() *MySqlTableSyntax {
	s := &MySqlTableSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
