package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlInsertSyntax struct {
	*Syntax__
	*StatementSyntax__
	*InsertSyntax__
	LowPriority                        bool
	Delayed                            bool
	HighPriority                       bool
	Ignore                             bool
	RowAlias                           *MySqlIdentifierSyntax
	ColumnAliasList                    IdentifierListSyntax_
	AssignmentList                     *AssignmentListSyntax
	OnDuplicateKeyUpdateAssignmentList *AssignmentListSyntax
}

func (this *MySqlInsertSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlInsertSyntax(this)
}

func (this *MySqlInsertSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("INSERT")
	if this.Hint != nil {
		builder.writeSpace()
		builder.writeSyntax(this.Hint)
	}
	if this.LowPriority {
		builder.writeStr(" LOW_PRIORITY")
	}
	if this.HighPriority {
		builder.writeStr(" HIGH_PRIORITY")
	}
	if this.Delayed {
		builder.writeStr(" DELAYED")
	}
	if this.Ignore {
		builder.writeStr(" IGNORE")
	}
	builder.writeStr(" INTO")
	builder.writeSpaceOrLf(this, false)
	builder.writeSyntax(this.NameTableReference)
	if this.InsertColumnList != nil {
		builder.writeSpaceOrLf(this, true)
		builder.writeSyntaxWithFormat(this.InsertColumnList, false)
	}
	if this.ValueListList != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("VALUES")
		builder.writeSpaceOrLf(this, true)
		builder.writeSyntax(this.ValueListList)
	}
	if this.AssignmentList != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("SET")
		builder.writeSpaceOrLf(this, true)
		builder.writeSyntax(this.AssignmentList)
	}
	if this.RowAlias != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("AS")
		builder.writeSpace()
		builder.writeSyntax(this.RowAlias)
	}
	if this.ColumnAliasList != nil {
		builder.writeSyntaxWithFormat(this.ColumnAliasList, false)
	}
	if this.OnDuplicateKeyUpdateAssignmentList != nil {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("ON DUPLICATE KEY UPDATE")
		builder.writeSpaceOrLf(this, true)
		builder.writeSyntax(this.OnDuplicateKeyUpdateAssignmentList)
	}
}

func (this *MySqlInsertSyntax) Dialect() enum.Dialect {
	return enum.Dialect_.MYSQL
}

func (this *MySqlInsertSyntax) AssignmentType() (t enum.MySqlAssignmentType) {
	if this.AssignmentList != nil {
		return enum.MySqlAssignmentType_.ASSIGNMENT_LIST
	} else if this.ValueListList != nil {
		return enum.MySqlAssignmentType_.VALUES_LIST
	}
	return
}

func NewMySqlInsertSyntax() *MySqlInsertSyntax {
	s := &MySqlInsertSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.StatementSyntax__ = ExtendStatementSyntax(s)
	s.InsertSyntax__ = ExtendInsertSyntax(s)
	return s
}
