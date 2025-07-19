package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type MySqlDeleteSyntax struct {
	*Syntax__
	*StatementSyntax__
	*HaveWhereSyntax__
	*DeleteSyntax__
	LowPriority               bool
	Quick                     bool
	Ignore                    bool
	MultiDeleteMode           enum.MySqlMultiDeleteMode
	MultiDeleteTableAliasList *MySqlMultiDeleteTableAliasListSyntax
	OrderBy                   *OrderBySyntax
	Limit                     *MySqlLimitSyntax
}

func (this *MySqlDeleteSyntax) accept(v_ Visitor_) {
	v_.(*mySqlVisitor).visitMySqlDeleteSyntax(this)
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
		if enum.MySqlMultiDeleteMode_.Is(this.MultiDeleteMode, enum.MySqlMultiDeleteMode_.MODE1) {
			builder.writeSpaceOrLf(this, true)
			builder.writeSyntax(this.MultiDeleteTableAliasList)
			builder.writeSpaceOrLf(this, false)
			builder.writeStr("FROM")
			builder.writeSpaceOrLf(this, true)
		} else if enum.MySqlMultiDeleteMode_.Is(this.MultiDeleteMode, enum.MySqlMultiDeleteMode_.MODE2) {
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
	return enum.Dialect_.MYSQL
}

func NewMySqlDeleteSyntax() *MySqlDeleteSyntax {
	s := &MySqlDeleteSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.StatementSyntax__ = ExtendStatementSyntax(s)
	s.HaveWhereSyntax__ = ExtendHaveWhereSyntax(s)
	s.DeleteSyntax__ = ExtendDeleteSyntax(s)
	return s
}
