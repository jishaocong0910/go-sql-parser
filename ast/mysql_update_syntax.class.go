package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlUpdateSyntax struct {
	*Syntax__
	*StatementSyntax__
	*HaveWhereSyntax__
	*UpdateSyntax__
	LowPriority bool
	Ignore      bool
	OrderBy     *OrderBySyntax
	Limit       *MySqlLimitSyntax
}

func (this *MySqlUpdateSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlUpdateSyntax(this)
}

func (this *MySqlUpdateSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("UPDATE")
	if this.Hint != nil {
		builder.writeSpace()
		builder.writeSyntax(this.Hint)
	}
	if this.LowPriority {
		builder.writeStr(" LOW_PRIORITY")
	}
	if this.Ignore {
		builder.writeStr(" IGNORE")
	}
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntax(this.TableReference)
	builder.writeSpaceOrLf(this, false)
	builder.writeStr("SET")
	builder.writeSpaceOrLf(this, true)
	builder.writeSyntax(this.AssignmentList)
	if this.Where != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.Where)
	}
	if this.OrderBy != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.OrderBy)
	}
	if this.Limit != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeSyntax(this.Limit)
	}
}

func (this *MySqlUpdateSyntax) Dialect() enum.Dialect {
	return enum.Dialect_.MYSQL
}

func NewMySqlUpdateSyntax() *MySqlUpdateSyntax {
	s := &MySqlUpdateSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.StatementSyntax__ = ExtendStatementSyntax(s)
	s.HaveWhereSyntax__ = ExtendHaveWhereSyntax(s)
	s.UpdateSyntax__ = ExtendUpdateSyntax(s)
	return s
}
