package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// 完整的可执行SQL语句
type I_StatementSyntax interface {
	I_Syntax
	M_A88DB0CC837F() *M_StatementSyntax
	Dialect() enum.Dialect
}

type M_StatementSyntax struct {
	I I_StatementSyntax
	// 原始SQL语句（由解析器填充）
	Sql string
}

func (this *M_StatementSyntax) M_A88DB0CC837F() *M_StatementSyntax {
	return this
}

func ExtendStatementSyntax(i I_StatementSyntax) *M_StatementSyntax {
	return &M_StatementSyntax{I: i}
}
