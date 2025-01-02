package ast

import "go-sql-parser/enum"

type MySqlUnarySyntax struct {
	*M_Syntax
	*M_ExprSyntax
	Expr          I_ExprSyntax
	UnaryOperator enum.UnaryOperator
}

func (this *MySqlUnarySyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlUnarySyntax(this)
}

func (this *MySqlUnarySyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.UnaryOperator.Symbol)
	if enum.SymbolTypes.Is(this.UnaryOperator.SymbolType, enum.SymbolTypes.IDENTIFIER) {
		builder.writeSpace()
	}
	builder.writeSyntax(this.Expr)
}

func NewMySqlUnarySyntax() *MySqlUnarySyntax {
	s := &MySqlUnarySyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
