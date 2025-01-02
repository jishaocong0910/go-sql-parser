package ast

import "go-sql-parser/enum"

type MySqlVariableSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	sql  string
	name string
	// 0-都无值，1-只有sql有值，2-都有值
	flag         int
	VariableType enum.MySqlVariableType
}

func (this *MySqlVariableSyntax) accept(I_Visitor) {}

func (this *MySqlVariableSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.sql)
}

func (this *MySqlVariableSyntax) Sql() string {
	return this.sql
}

func (this *MySqlVariableSyntax) SetSql(sql string) {
	this.sql = sql
	this.flag = 1
}

func (this *MySqlVariableSyntax) Name() string {
	if this.flag == 1 {
		chars := []rune(this.sql)
		begin := 0
		end := len(chars)
		for i := range chars {
			c := chars[i]
			if c != '@' && c != '\'' && c != '"' {
				begin = i
				break
			}
		}
		for i := len(chars) - 1; i > 0; i-- {
			c := chars[i]
			if c == '\'' || c == '"' {
				end = i
				break
			}
		}
		this.name = string(chars[begin:end])
		this.flag = 2
	}
	return this.name
}

func NewMySqlVariableSyntax() *MySqlVariableSyntax {
	s := &MySqlVariableSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	return s
}
