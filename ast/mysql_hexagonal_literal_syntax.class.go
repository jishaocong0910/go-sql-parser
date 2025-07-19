package ast

type MySqlHexagonalLiteralSyntax struct {
	*Syntax__
	*ExprSyntax__
	// e.g. x'4D7953514C'
	sql string
	// e.g. 4D7953514C
	hexStr string
	// 0-都无值，1-只有sql有值，2-只有hexStr有值，3-都有值
	flag int
}

func (this *MySqlHexagonalLiteralSyntax) accept(Visitor_) {}

func (this *MySqlHexagonalLiteralSyntax) Sql() string {
	if this.flag == 2 {
		this.sql = "x'" + this.hexStr + "'"
		this.flag = 3
	}
	return this.sql
}

func (this *MySqlHexagonalLiteralSyntax) SetSql(sql string) {
	this.sql = sql
	this.flag = 1
}

func (this *MySqlHexagonalLiteralSyntax) HexStr() string {
	if this.flag == 1 {
		this.hexStr = this.sql[2 : len(this.sql)-1]
		this.flag = 3
	}
	return this.hexStr
}

func (this *MySqlHexagonalLiteralSyntax) SetHexStr(hexStr string) {
	this.hexStr = hexStr
	this.flag = 2
}

func (this *MySqlHexagonalLiteralSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Sql())
}

func NewMySqlHexagonalLiteralSyntax() *MySqlHexagonalLiteralSyntax {
	s := &MySqlHexagonalLiteralSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	return s
}
