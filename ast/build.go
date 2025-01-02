package ast

import (
	. "go-sql-parser/enum"
	"strings"
	"unicode/utf8"

	. "github.com/jishaocong0910/go-object"
)

// BuildSql
//
//	@Description: 将完整的SQL语法结构构建为SQL字符串
//	@param is 完整SQL语法
//	@param Format 是否格式化
//	@return string SQL字符串
func BuildSql(is I_StatementSyntax, format bool) (sql string) {
	if !IsNull(is) {
		b := &sqlBuilder{defaultFormat: format}
		b.builder.Grow(len(is.M_A88DB0CC837F().Sql))
		b.writeSyntax(is)
		sql = b.toSql()
	}
	return
}

// sqlBuilder
//
//	@Description: SQL构建器
type sqlBuilder struct {
	builder strings.Builder
	// 当前语法缩进起始位置，随着SQL的拼接会变化。当一个Syntax对象开始拼接时，会将此值设置到Syntax对象中，作为该对象固定的缩进其实位置。
	currentIndentBeginIndex int
	// 默认是否格式化
	defaultFormat bool
}

// writeSyntax
//
//	@Description: 将语法结构转换为SQL进行拼接，使用SQL构建器的格式化默认值
//	@receiver b SQL构建器
//	@param is 语法
func (b *sqlBuilder) writeSyntax(is I_Syntax) {
	b.writeSyntaxWithFormat(is, b.defaultFormat)
}

// writeSyntaxWithFormat
//
//	@Description: 将语法结构转换为SQL进行拼接，并指定该语法结构是否格式化
//	@receiver b SQL构建器
//	@param is 语法
//	@param format 是否格式化
func (b *sqlBuilder) writeSyntaxWithFormat(is I_Syntax, format bool) {
	if s := is.M_5CF6320E8474(); s != nil {
		s.Format = format
		if ParenthesizeTypes.Is(s.ParenthesizeType, ParenthesizeTypes.TRUE) {
			b.writeStr("(")
			s.IndentBeginIndex = b.currentIndentBeginIndex
			s.I.writeSql(b)
			b.writeStr(")")
		} else {
			s.IndentBeginIndex = b.currentIndentBeginIndex
			s.I.writeSql(b)
		}
	}
}

// writeSpaceOrLf 拼接空格或者换行符。若语法结构不格式化则拼接空格，否则拼接换行符，换行后输入光标所在列与当前语法结构的缩进起始位
// 置一致。参数indentAfterLf可指定输入光标是否缩进
func (b *sqlBuilder) writeSpaceOrLf(is I_Syntax, indentAfterLf bool) {
	if s := is.M_5CF6320E8474(); s.Format {
		b.writeLf()
		b.currentIndentBeginIndex = 0
		for i := 0; i < s.IndentBeginIndex; i++ {
			b.writeStr(" ")
		}
		if indentAfterLf {
			b.writeStr("  ")
		}
	} else {
		b.writeStr(" ")
	}
}

// writeSpaceOrLfIndent 拼接空格或者换行符。若语法结构不格式化则拼接空格，否则拼接换行符，换行后自动缩进。参数caretPrefixes可指定在
// 光标前方拼接的字符，如果拼接的字符长度大于缩进的长度，则
func (b *sqlBuilder) writeSpaceOrLfIndent(is I_Syntax, caretPrefixes ...string) {
	if s := is.M_5CF6320E8474(); s.Format {
		b.writeLf()
		b.currentIndentBeginIndex = 0
		spaceNum := s.IndentBeginIndex
		for i := range caretPrefixes {
			spaceNum -= utf8.RuneCountInString(caretPrefixes[i])
		}
		for i := 0; i < spaceNum; i++ {
			b.writeStr(" ")
		}
	} else {
		b.writeStr(" ")
	}
	for i := range caretPrefixes {
		b.writeStr(caretPrefixes[i])
	}
}

// writeStr 拼接字符串
func (b *sqlBuilder) writeStr(s string) {
	b.builder.WriteString(s)
	b.currentIndentBeginIndex += utf8.RuneCountInString(s)
}

// writeSpace 拼接空格
func (b *sqlBuilder) writeSpace() {
	b.writeStr(" ")
}

// writeLf 拼接换行符
func (b *sqlBuilder) writeLf() {
	b.builder.WriteString("\n")
}

// toSql 将SQL构建器当前拼接的结果输出为SQL字符串
func (b *sqlBuilder) toSql() string {
	return b.builder.String()
}
