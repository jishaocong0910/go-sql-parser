package ast

import "strings"

type MySqlStringSyntax struct {
	*Syntax__
	*ExprSyntax__
	*AliasSyntax__
	*StringSyntax__
}

func (this *MySqlStringSyntax) AliasName() string {
	sql := this.Sql()
	return sql[1 : len(sql)-1]
}

func (this *MySqlStringSyntax) accept(Visitor_) {}

func (this *MySqlStringSyntax) sqlToValue(str string) string {
	if str == "" {
		return str
	}
	var builder strings.Builder
	// 跳过两端单引号或双引号
	chars := []rune(str[1 : len(str)-1])
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c != '\\' {
			builder.WriteRune(c)
			continue
		}
		if i+1 == len(chars) {
			panic("illegal string")
		}
		i++
		nextC := chars[i]
		switch nextC {
		case '0':
			builder.WriteRune(0)
		case '\'':
			builder.WriteRune('\'')
		case '"':
			builder.WriteRune('"')
		case 'b':
			builder.WriteRune('\b')
		case 'n':
			builder.WriteRune('\n')
		case 'r':
			builder.WriteRune('\r')
		case 't':
			builder.WriteRune('\t')
		case 'Z':
			builder.WriteRune(26)
		case '\\':
			builder.WriteRune('\\')
		case '%':
			builder.WriteRune('%')
		case '_':
			builder.WriteRune('_')
		default:
			panic("unknown escape sequence \\" + string(nextC))
		}
	}

	return builder.String()
}

func (this *MySqlStringSyntax) valueToSql(str string) string {
	if str == "" {
		return str
	}
	var builder strings.Builder
	builder.WriteRune('\'')
	for _, c := range str {
		switch c {
		case 0:
			builder.WriteString(`\0`)
		case '\'':
			builder.WriteString(`\'`)
		case '"':
			builder.WriteString(`\"`)
		case '\b':
			builder.WriteString(`\b`)
		case '\n':
			builder.WriteString(`\n`)
		case '\r':
			builder.WriteString(`\r`)
		case '\t':
			builder.WriteString(`\t`)
		case 26:
			builder.WriteString(`\Z`)
		case '\\':
			builder.WriteString(`\\`)
		case '%':
			builder.WriteString(`\%`)
		case '_':
			builder.WriteString(`\_`)
		default:
			builder.WriteRune(c)
		}
	}
	builder.WriteRune('\'')
	return builder.String()
}

func NewMySqlStringSyntax() *MySqlStringSyntax {
	s := &MySqlStringSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.AliasSyntax__ = ExtendAliasSyntax(s)
	s.StringSyntax__ = ExtendStringSyntax(s)
	return s
}
