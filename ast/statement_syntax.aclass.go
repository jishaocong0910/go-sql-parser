package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// 完整的可执行SQL语句
type StatementSyntax_ interface {
	StatementSyntax_() *StatementSyntax__
	Syntax_
	Dialect() enum.Dialect
}

type StatementSyntax__ struct {
	I StatementSyntax_
	// 原始SQL语句（由解析器填充）
	Sql string
}

func (this *StatementSyntax__) StatementSyntax_() *StatementSyntax__ {
	return this
}

func ExtendStatementSyntax(i StatementSyntax_) *StatementSyntax__ {
	return &StatementSyntax__{I: i}
}
