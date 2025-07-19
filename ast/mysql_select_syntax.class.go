package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlSelectSyntax struct {
	*Syntax__
	*ExprSyntax__
	*StatementSyntax__
	*QuerySyntax__
	*HaveWhereSyntax__
	*SelectSyntax__
	HighPriority     bool
	StraightJoin     bool
	SqlSmallResult   bool
	SqlBigResult     bool
	SqlBufferResult  bool
	SqlCache         bool // 8.0已删除
	SqlNoCache       bool
	SqlCalcFoundRows bool
	LimitSyntax      *MySqlLimitSyntax
	LockingReads     []*MySqlLockingReadSyntax
}

func (this *MySqlSelectSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlSelectSyntax(this)
}

func (this *MySqlSelectSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("SELECT")
	if this.Hint != nil {
		builder.writeSpace()
		builder.writeSyntax(this.Hint)
	}
	if !this.AggregateOption.Undefined() {
		builder.writeSpace()
		builder.writeStr(this.AggregateOption.Sql)
	}
	if this.HighPriority {
		builder.writeStr(" HIGH_PRIORITY")
	}
	if this.StraightJoin {
		builder.writeStr(" STRAIGHT_JOIN")
	}
	if this.SqlSmallResult {
		builder.writeStr(" SQL_SMALL_RESULT")
	}
	if this.SqlBigResult {
		builder.writeStr(" SQL_BIG_RESULT")
	}
	if this.SqlBufferResult {
		builder.writeStr(" SQL_BUFFER_RESULT")
	}
	if this.SqlCache {
		builder.writeStr(" SQL_CACHE")
	}
	if this.SqlNoCache {
		builder.writeStr(" SQL_NO_CACHE")
	}
	if this.SqlCalcFoundRows {
		builder.writeStr(" SQL_CALC_FOUND_ROWS")
	}
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntax(this.SelectItemList)
	if this.TableReference != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("FROM")
		builder.writeSpaceOrLf(this, true)
		builder.writeSyntax(this.TableReference)
	}
	if this.Where != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.Where)
	}
	if this.GroupBy != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.GroupBy)
	}
	if this.Having != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.Having)
	}
	if this.NamedWindowList != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("WINDOW ")
		builder.writeSyntax(this.NamedWindowList)
	}
	if this.OrderBy != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.OrderBy)
	}
	if this.LimitSyntax != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.LimitSyntax)
	}
	for _, l := range this.LockingReads {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(l)
	}
}

func (this *MySqlSelectSyntax) OperandCount() int {
	return this.SelectSyntax__.OperandCount()
}

func (this *MySqlSelectSyntax) Dialect() enum.Dialect {
	return enum.Dialect_.MYSQL
}

func NewMySqlSelectSyntax() *MySqlSelectSyntax {
	s := &MySqlSelectSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.StatementSyntax__ = ExtendStatementSyntax(s)
	s.QuerySyntax__ = ExtendQuerySyntax(s)
	s.HaveWhereSyntax__ = ExtendHaveWhereSyntax(s)
	s.SelectSyntax__ = ExtendSelectSyntax(s)
	return s
}
