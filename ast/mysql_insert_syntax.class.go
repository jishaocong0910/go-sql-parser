package ast

import "go-sql-parser/enum"

type MySqlInsertSyntax struct {
	*M_Syntax
	*M_StatementSyntax
	*M_InsertSyntax
	LowPriority                        bool
	Delayed                            bool
	HighPriority                       bool
	Ignore                             bool
	RowAlias                           *MySqlIdentifierSyntax
	ColumnAliasList                    I_IdentifierListSyntax
	AssignmentList                     *AssignmentListSyntax
	OnDuplicateKeyUpdateAssignmentList *AssignmentListSyntax
}

func (this *MySqlInsertSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlInsertSyntax(this)
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
	return enum.Dialects.MYSQL
}

func (this *MySqlInsertSyntax) AssignmentType() (t enum.MySqlAssignmentType) {
	if this.AssignmentList != nil {
		return enum.MySqlAssignmentTypes.ASSIGNMENT_LIST
	} else if this.ValueListList != nil {
		return enum.MySqlAssignmentTypes.VALUES_LIST
	}
	return
}

func NewMySqlInsertSyntax() *MySqlInsertSyntax {
	s := &MySqlInsertSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_StatementSyntax = ExtendStatementSyntax(s)
	s.M_InsertSyntax = ExtendInsertSyntax(s)
	return s
}
