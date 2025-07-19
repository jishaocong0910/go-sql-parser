package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlUnarySyntax struct {
	*Syntax__
	*ExprSyntax__
	Expr          ExprSyntax_
	UnaryOperator enum.UnaryOperator
}

func (this *MySqlUnarySyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlUnarySyntax(this)
}

func (this *MySqlUnarySyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.UnaryOperator.Symbol)
	if enum.SymbolType_.Is(this.UnaryOperator.SymbolType, enum.SymbolType_.IDENTIFIER) {
		builder.writeSpace()
	}
	builder.writeSyntax(this.Expr)
}

func NewMySqlUnarySyntax() *MySqlUnarySyntax {
	s := &MySqlUnarySyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
