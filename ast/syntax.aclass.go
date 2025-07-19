package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"
)

// 语法基类
type Syntax_ interface {
	Syntax_() *Syntax__
	accept(v_ Visitor_)
	writeSql(builder *sqlBuilder)
}

type Syntax__ struct {
	I Syntax_
	// 语法在SQL中的起始位置（由解析器填充）
	BeginPos int
	// 语法在SQL中的结束位置，不包含（由解析器填充）
	EndPos int
	// 括号类型
	ParenthesizeType enum.ParenthesizeType
	// 缩进起始位置（构建SQL使用）
	IndentBeginIndex int
	// 是否格式化（构建SQL使用）
	Format bool
}

func (this *Syntax__) Syntax_() *Syntax__ {
	return this
}

func ExtendSyntax(i Syntax_) *Syntax__ {
	return &Syntax__{I: i}
}
