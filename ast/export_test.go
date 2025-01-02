package ast

import (
	"go-sql-parser/enum"
)

type NotSupportedDialectStatement struct {
	*M_Syntax
	*M_StatementSyntax
}

func (n *NotSupportedDialectStatement) accept(I_Visitor) {}

func (n *NotSupportedDialectStatement) writeSql(*sqlBuilder) {}

func (n *NotSupportedDialectStatement) Dialect() enum.Dialect {
	return enum.Dialects.SQLSERVER
}

func NewNotSupportedDialectStatement() *NotSupportedDialectStatement {
	s := &NotSupportedDialectStatement{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_StatementSyntax = ExtendStatementSyntax(s)
	return s
}
