package ast

import (
	"strings"
	"unicode/utf8"

	. "github.com/jishaocong0910/go-sql-parser/enum"

	. "github.com/jishaocong0910/go-object-util"
)

// BuildSql 将语法对象构建为SQL字符串
func BuildSql(s_ Syntax_, format bool) (sql string) {
	if !IsNull(s_) {
		b := &sqlBuilder{defaultFormat: format}
		b.writeSyntax(s_)
		sql = b.toSql()
	}
	return
}

// sqlBuilder SQL构建器
type sqlBuilder struct {
	builder strings.Builder
	// 当前语法缩进起始位置，随着SQL的拼接会变化。当一个Syntax对象开始拼接时，会将此值设置到Syntax对象中，作为该对象固定的缩进其实位置。
	currentIndentBeginIndex int
	// 默认是否格式化
	defaultFormat bool
}

// writeSyntax 将语法结构转换为SQL进行拼接，使用SQL构建器的格式化默认值
func (b *sqlBuilder) writeSyntax(s_ Syntax_) {
	b.writeSyntaxWithFormat(s_, b.defaultFormat)
}

// writeSyntaxWithFormat 将语法结构转换为SQL进行拼接，并指定该语法结构是否格式化
func (b *sqlBuilder) writeSyntaxWithFormat(s_ Syntax_, format bool) {
	if s := s_.Syntax_(); s != nil {
		s.Format = format
		if ParenthesizeType_.Is(s.ParenthesizeType, ParenthesizeType_.TRUE) {
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
func (b *sqlBuilder) writeSpaceOrLf(s_ Syntax_, indentAfterLf bool) {
	if s := s_.Syntax_(); s.Format {
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
func (b *sqlBuilder) writeSpaceOrLfIndent(s_ Syntax_, caretPrefixes ...string) {
	if s := s_.Syntax_(); s.Format {
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
