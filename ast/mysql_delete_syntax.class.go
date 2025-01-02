package ast

import "go-sql-parser/enum"

type MySqlDeleteSyntax struct {
	*M_Syntax
	*M_StatementSyntax
	*M_DeleteSyntax
	LowPriority               bool
	Quick                     bool
	Ignore                    bool
	MultiDeleteMode           enum.MySqlMultiDeleteMode
	MultiDeleteTableAliasList *MySqlMultiDeleteTableAliasListSyntax
	OrderBy                   *OrderBySyntax
	Limit                     *MySqlLimitSyntax
}

func (this *MySqlDeleteSyntax) accept(iv I_Visitor) {
	iv.(*mySqlVisitor).visitMySqlDeleteSyntax(this)
}

func (this *MySqlDeleteSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("DELETE")
	if this.Hint != nil {
		builder.writeSpace()
		builder.writeSyntax(this.Hint)
	}
	if this.LowPriority {
		builder.writeStr(" LOW_PRIORITY")
	}
	if this.Quick {
		builder.writeStr(" QUICK")
	}
	if this.Ignore {
		builder.writeStr(" IGNORE")
	}

	if !this.MultiDeleteMode.Undefined() {
		if enum.MySqlMultiDeleteModes.Is(this.MultiDeleteMode, enum.MySqlMultiDeleteModes.MODE1) {
			builder.writeSpaceOrLf(this, true)
			builder.writeSyntax(this.MultiDeleteTableAliasList)
			builder.writeSpaceOrLf(this, false)
			builder.writeStr("FROM")
			builder.writeSpaceOrLf(this, true)
		} else if enum.MySqlMultiDeleteModes.Is(this.MultiDeleteMode, enum.MySqlMultiDeleteModes.MODE2) {
			builder.writeSpaceOrLf(this, false)
			builder.writeStr("FROM")
			builder.writeSpaceOrLf(this, true)
			builder.writeSyntax(this.MultiDeleteTableAliasList)
			builder.writeStr(" USING ")
		}
	} else {
		builder.writeSpaceOrLf(this, false)
		builder.writeStr("FROM")
		builder.writeSpaceOrLf(this, true)
	}

	builder.writeSyntax(this.TableReference)
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

func (this *MySqlDeleteSyntax) Dialect() enum.Dialect {
	return enum.Dialects.MYSQL
}

func NewMySqlDeleteSyntax() *MySqlDeleteSyntax {
	s := &MySqlDeleteSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_StatementSyntax = ExtendStatementSyntax(s)
	s.M_DeleteSyntax = ExtendDeleteSyntax(s)
	return s
}
