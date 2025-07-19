package ast

import (
	. "github.com/jishaocong0910/go-sql-parser/enum"
)

type mySqlVisitor struct {
	*visitor__
}

func (this *mySqlVisitor) newSubqueryVisitor(s_ StatementSyntax_, opt Option) *visitor__ {
	return newMySqlVisitor(s_, opt).visitor__
}

func (this *mySqlVisitor) visitMySqlDeleteSyntax(s *MySqlDeleteSyntax) {
	this.sqlOperationType = SqlOperationType_.DELETE
	if s.Hint != nil {
		this.hintContent = s.Hint.Content
	}
	if s.TableReference != nil {
		if i, ok := s.TableReference.(NameTableReferenceSyntax_); ok {
			t := i.NameTableReferenceSyntax_()
			tn := t.TableNameItem.FullTableName()
			this.singleTableSql = true
			this.tableOfSingleTableSql = tn
			this.addDeleteItem(tn)
		}
		this.visit(s.TableReference)
	}
	this.visit(s.MultiDeleteTableAliasList)
	this.visit(s.Where)
	this.visit(s.OrderBy)
}

func (this *mySqlVisitor) visitMySqlIdentifierSyntax(s *MySqlIdentifierSyntax) {
	this.addColumnItem(s)
}

func (this *mySqlVisitor) visitMySqlInsertSyntax(s *MySqlInsertSyntax) {
	this.sqlOperationType = SqlOperationType_.INSERT
	if s.Hint != nil {
		this.hintContent = s.Hint.Content
	}
	this.singleTableSql = true
	this.tableOfSingleTableSql = s.NameTableReference.NameTableReferenceSyntax_().TableNameItem.FullTableName()
	this.visit(s.NameTableReference)
	switch s.AssignmentType() {
	case MySqlAssignmentType_.VALUES_LIST:
		this.inInsertColumnListSyntax = true
		this.visit(s.InsertColumnList)
		this.inInsertColumnListSyntax = false
		this.visit(s.ValueListList)
	case MySqlAssignmentType_.ASSIGNMENT_LIST:
		this.visit(s.AssignmentList)
	}
	this.visit(s.OnDuplicateKeyUpdateAssignmentList)
}

func (this *mySqlVisitor) visitMySqlIntervalSyntax(s *MySqlIntervalSyntax) {
	this.visit(s.Expr)
}

func (this *mySqlVisitor) visitMySqlMultiDeleteTableAliasSyntax(s *MySqlMultiDeleteTableAliasSyntax) {
	if table, ok := this.tableAliases[s.Alias.Name]; ok {
		this.addDeleteItem(table)
	} else {
		this.panic("unknown table of alias '%s' in MULTI DELETE", s.Alias.Name)
	}
}

func (this *mySqlVisitor) visitMySqlSelectSyntax(s *MySqlSelectSyntax) {
	if !this.sqlOperationType.Undefined() {
		this.singleTableSql = false
		this.visitSubquery(s)
		return
	}
	this.sqlOperationType = SqlOperationType_.SELECT
	if s.Hint != nil {
		this.hintContent = s.Hint.Content
	}
	if s.TableReference != nil {
		if i, ok := s.TableReference.(NameTableReferenceSyntax_); ok {
			t := i.NameTableReferenceSyntax_()
			this.singleTableSql = true
			this.tableOfSingleTableSql = t.TableNameItem.FullTableName()
		}
		this.visit(s.TableReference)
	}
	this.visit(s.SelectItemList)
	this.visit(s.Where)
	this.visit(s.GroupBy)
	this.visit(s.Having)
	this.visit(s.OrderBy)
	this.visit(s.NamedWindowList)
}

func (this *mySqlVisitor) visitMySqlTrimFunctionSyntax(s *MySqlTrimFunctionSyntax) {
	this.visit(s.Str)
	this.visit(s.RemStr)
}

func (this *mySqlVisitor) visitMySqlUnarySyntax(s *MySqlUnarySyntax) {
	this.visit(s.Expr)
}

func (this *mySqlVisitor) visitMySqlUpdateSyntax(s *MySqlUpdateSyntax) {
	this.sqlOperationType = SqlOperationType_.UPDATE
	if s.Hint != nil {
		this.hintContent = s.Hint.Content
	}
	if i, ok := s.TableReference.(NameTableReferenceSyntax_); ok {
		this.singleTableSql = true
		this.tableOfSingleTableSql = i.NameTableReferenceSyntax_().TableNameItem.FullTableName()
	}
	this.visit(s.Hint)
	this.visit(s.TableReference)
	this.visit(s.AssignmentList)
	this.visit(s.Where)
}

func (this *mySqlVisitor) visitMySqlTableSyntax(s *MySqlTableSyntax) {
	n := s.TableNameItem.FullTableName()
	this.addLocalTableReference(n, n)
	if b, ok := this.traceSyntax(1).(*MySqlBinaryOperationSyntax); ok && !b.ComparisonMode.Undefined() {
		// 此场景Table语法中的表肯定只有一个字段，参考一下文档。
		// https://dev.mysql.com/doc/refman/8.0/en/any-in-some-subqueries.html
		// https://dev.mysql.com/doc/refman/8.0/en/all-subqueries.html
		// 由于从字面解析SQL无法得知字段名，所以记录该表查询了所有字段
		this.selectAllColumnTables = append(this.selectAllColumnTables, n)
	}
	this.visit(s.OrderBy)
}

func newMySqlVisitor(s_ StatementSyntax_, opt Option) *mySqlVisitor {
	v := &mySqlVisitor{}
	v.visitor__ = extendVisitor(v, s_, opt)
	return v
}
