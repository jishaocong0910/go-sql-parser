package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"
)

// 语法基类
type I_Syntax interface {
	M_Syntax_() *M_Syntax
	accept(iv I_Visitor)
	writeSql(builder *sqlBuilder)
}

type M_Syntax struct {
	I I_Syntax
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

func (this *M_Syntax) M_Syntax_() *M_Syntax {
	return this
}

func ExtendSyntax(i I_Syntax) *M_Syntax {
	return &M_Syntax{I: i}
}
