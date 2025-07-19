package ast

type PlaceholderSyntax struct {
	*Syntax__
	*ExprSyntax__
	Sql string
	// 参数占位符索引，从1开始
	Index int
}

func (this *PlaceholderSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitPlaceholderSyntax(this)
}

func (this *PlaceholderSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql)
}

func NewPlaceholderSyntax() *PlaceholderSyntax {
	s := &PlaceholderSyntax{Sql: "?"}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
