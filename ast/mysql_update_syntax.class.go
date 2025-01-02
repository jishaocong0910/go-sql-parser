package ast

import "go-sql-parser/enum"

type MySqlUpdateSyntax struct {
	*M_Syntax
	*M_StatementSyntax
	*M_UpdateSyntax
	LowPriority bool
	Ignore      bool
	OrderBy     *OrderBySyntax
	Limit       *MySqlLimitSyntax
}

func (this *MySqlUpdateSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlUpdateSyntax(this)
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
	return enum.Dialects.MYSQL
}

func NewMySqlUpdateSyntax() *MySqlUpdateSyntax {
	s := &MySqlUpdateSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_StatementSyntax = ExtendStatementSyntax(s)
	s.M_UpdateSyntax = ExtendUpdateSyntax(s)
	return s
}
