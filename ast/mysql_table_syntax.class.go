package ast

// MySqlTableSyntax
//
//	@Description: https://dev.mysql.com/doc/refman/8.0/en/table.html
type MySqlTableSyntax struct {
	*Syntax__
	*ExprSyntax__
	TableNameItem *TableNameItemSyntax
	OrderBy       *OrderBySyntax
	Limit         *MySqlLimitSyntax
}

func (this *MySqlTableSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlTableSyntax(this)
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
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
