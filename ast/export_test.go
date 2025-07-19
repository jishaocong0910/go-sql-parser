package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"
)

type NotSupportedDialectStatement struct {
	*Syntax__
	*StatementSyntax__
}

func (n *NotSupportedDialectStatement) accept(Visitor_) {}

func (n *NotSupportedDialectStatement) writeSql(*sqlBuilder) {}

func (n *NotSupportedDialectStatement) Dialect() enum.Dialect {
	return enum.Dialect_.SQLSERVER
}

func NewNotSupportedDialectStatement() *NotSupportedDialectStatement {
	s := &NotSupportedDialectStatement{}
	s.Syntax__ = ExtendSyntax(s)
	s.StatementSyntax__ = ExtendStatementSyntax(s)
	return s
}
