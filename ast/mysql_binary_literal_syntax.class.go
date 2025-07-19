package ast

type MySqlBinaryLiteralSyntax struct {
	*Syntax__
	*ExprSyntax__
	// e.g. b'01100001'
	sql string
	// e.g. 01100001
	binStr string
	// 0-都无值，1-只有sql有值，2-只有binStr有值，3-都有值
	flag int
}

func (this *MySqlBinaryLiteralSyntax) accept(Visitor_) {}

func (this *MySqlBinaryLiteralSyntax) Sql() string {
	if this.flag == 2 {
		this.sql = "b'" + this.binStr + "'"
		this.flag = 3
	}
	return this.sql
}

func (this *MySqlBinaryLiteralSyntax) SetSql(sql string) {
	this.sql = sql
	this.flag = 1
}

func (this *MySqlBinaryLiteralSyntax) BinStr() string {
	if this.flag == 1 {
		this.binStr = this.sql[2 : len(this.sql)-1]
	}
	return this.binStr
}

func (this *MySqlBinaryLiteralSyntax) SetBinStr(binStr string) {
	this.binStr = binStr
	this.flag = 2
}

func (this *MySqlBinaryLiteralSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql())
}

func NewMySqlBinaryLiteralSyntax() *MySqlBinaryLiteralSyntax {
	s := &MySqlBinaryLiteralSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
