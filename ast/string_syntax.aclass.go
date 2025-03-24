package ast

// 字符串
type I_StringSyntax interface {
	I_ExprSyntax
	I_AliasSyntax
	M_StringSyntax_() *M_StringSyntax
	// sqlToValue string with escape sequence convert to character string
	sqlToValue(string) string
	// valueToSql character string convert to string with escape sequence
	valueToSql(string) string
}

type M_StringSyntax struct {
	I I_StringSyntax
	// 在SQL中的字符表示
	sql string
	// 将成员变量sql去除掉单引号或双引号、转化转义字符后实际的值
	value string
	// 0-都无值，1-只有sql有值，2-只有value有值，3-都有值
	flag int
}

func (this *M_StringSyntax) M_StringSyntax_() *M_StringSyntax {
	return this
}

func (this *M_StringSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql())
}

func (this *M_StringSyntax) Sql() string {
	if this.flag == 2 {
		this.sql = this.I.valueToSql(this.value)
		this.flag = 3
	}
	return this.sql
}

func (this *M_StringSyntax) Value() string {
	if this.flag == 1 {
		this.value = this.I.sqlToValue(this.sql)
		this.flag = 3
	}
	return this.value
}

func (this *M_StringSyntax) SetSql(sql string) {
	this.sql = sql
	this.flag = 1
}

func (this *M_StringSyntax) SetValue(value string) {
	this.value = value
	this.flag = 2
}

func ExtendStringSyntax(i I_StringSyntax) *M_StringSyntax {
	return &M_StringSyntax{I: i}
}
