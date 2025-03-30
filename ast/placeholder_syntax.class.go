package ast

type PlaceholderSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Sql string
	// 参数占位符索引，从1开始
	Index int
}

func (this *PlaceholderSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitPlaceholderSyntax(this)
}

func (this *PlaceholderSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewPlaceholderSyntax() *PlaceholderSyntax {
	s := &PlaceholderSyntax{Sql: "?"}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
