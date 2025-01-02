package ast

import (
	. "go-sql-parser/enum"
)

type mySqlVisitor struct {
	*m_Visitor
}

func (v *mySqlVisitor) newSubqueryVisitor(is I_StatementSyntax, opt Option) *m_Visitor {
	return newMySqlVisitor(is, opt).m_E61B18189B57()
}

func (v *mySqlVisitor) visitMySqlBinaryOperationSyntax(s *MySqlBinaryOperationSyntax) {
	v.visit(s.LeftOperand)
	v.visit(s.RightOperand)
	v.visit(s.BetweenThirdOperand)
}

func (v *mySqlVisitor) visitMySqlDeleteSyntax(s *MySqlDeleteSyntax) {
	v.sqlOperationType = SqlOperationTypes.DELETE
	if s.Hint != nil {
		v.hintContent = s.Hint.Content
	}
	if s.TableReference != nil {
		if i, ok := s.TableReference.(I_NameTableReferenceSyntax); ok {
			t := i.M_0E797D96D386()
			tn := t.TableNameItem.FullTableName()
			v.singleTableSql = true
			v.tableOfSingleTableSql = tn
			v.addDeleteItem(tn)
		}
		v.visit(s.TableReference)
	}
	v.visit(s.MultiDeleteTableAliasList)
	v.visit(s.Where)
	v.visit(s.OrderBy)
}

func (v *mySqlVisitor) visitMySqlIdentifierSyntax(s *MySqlIdentifierSyntax) {
	v.addColumnItem(s)
}

func (v *mySqlVisitor) visitMySqlInsertSyntax(s *MySqlInsertSyntax) {
	v.sqlOperationType = SqlOperationTypes.INSERT
	if s.Hint != nil {
		v.hintContent = s.Hint.Content
	}
	v.singleTableSql = true
	v.tableOfSingleTableSql = s.NameTableReference.M_0E797D96D386().TableNameItem.FullTableName()
	v.visit(s.NameTableReference)
	switch s.AssignmentType() {
	case MySqlAssignmentTypes.VALUES_LIST:
		v.inInsertColumnListSyntax = true
		v.visit(s.InsertColumnList)
		v.inInsertColumnListSyntax = false
		v.visit(s.ValueListList)
	case MySqlAssignmentTypes.ASSIGNMENT_LIST:
		v.visit(s.AssignmentList)
	}
	v.visit(s.OnDuplicateKeyUpdateAssignmentList)
}

func (v *mySqlVisitor) visitMySqlIntervalSyntax(s *MySqlIntervalSyntax) {
	v.visit(s.Expr)
}

func (v *mySqlVisitor) visitMySqlMultiDeleteTableAliasSyntax(s *MySqlMultiDeleteTableAliasSyntax) {
	if table, ok := v.tableAliases[s.Alias.Name]; ok {
		v.addDeleteItem(table)
	} else {
		v.panic("unknown table of alias '%s' in MULTI DELETE", s.Alias.Name)
	}
}

func (v *mySqlVisitor) visitMySqlSelectSyntax(s *MySqlSelectSyntax) {
	if !v.sqlOperationType.Undefined() {
		v.singleTableSql = false
		v.visitSubquery(s)
		return
	}
	v.sqlOperationType = SqlOperationTypes.SELECT
	if s.Hint != nil {
		v.hintContent = s.Hint.Content
	}
	if s.TableReference != nil {
		if i, ok := s.TableReference.(I_NameTableReferenceSyntax); ok {
			t := i.M_0E797D96D386()
			v.singleTableSql = true
			v.tableOfSingleTableSql = t.TableNameItem.FullTableName()
		}
		v.visit(s.TableReference)
	}
	v.visit(s.SelectItemList)
	v.visit(s.Where)
	v.visit(s.GroupBy)
	v.visit(s.Having)
	v.visit(s.OrderBy)
	v.visit(s.NamedWindowList)
}

func (v *mySqlVisitor) visitMySqlTrimFunctionSyntax(s *MySqlTrimFunctionSyntax) {
	v.visit(s.Str)
	v.visit(s.RemStr)
}

func (v *mySqlVisitor) visitMySqlUnarySyntax(s *MySqlUnarySyntax) {
	v.visit(s.Expr)
}

func (v *mySqlVisitor) visitMySqlUpdateSyntax(s *MySqlUpdateSyntax) {
	v.sqlOperationType = SqlOperationTypes.UPDATE
	if s.Hint != nil {
		v.hintContent = s.Hint.Content
	}
	if i, ok := s.TableReference.(I_NameTableReferenceSyntax); ok {
		v.singleTableSql = true
		v.tableOfSingleTableSql = i.M_0E797D96D386().TableNameItem.FullTableName()
	}
	v.visit(s.Hint)
	v.visit(s.TableReference)
	v.visit(s.AssignmentList)
	v.visit(s.Where)
}

func (v *mySqlVisitor) visitMySqlTableSyntax(s *MySqlTableSyntax) {
	n := s.TableNameItem.FullTableName()
	v.addLocalTableReference(n, n)
	if b, ok := v.traceSyntax(1).(*MySqlBinaryOperationSyntax); ok && !b.ComparisonMode.Undefined() {
		// 此场景Table语法中的表肯定只有一个字段，参考一下文档。
		// https://dev.mysql.com/doc/refman/8.0/en/any-in-some-subqueries.html
		// https://dev.mysql.com/doc/refman/8.0/en/all-subqueries.html
		// 由于从字面解析SQL无法得知字段名，所以记录该表查询了所有字段
		v.selectAllColumnTables = append(v.selectAllColumnTables, n)
	}
	v.visit(s.OrderBy)
}

func newMySqlVisitor(is I_StatementSyntax, opt Option) *mySqlVisitor {
	v := &mySqlVisitor{}
	v.m_Visitor = extendVisitor(v, is, opt)
	return v
}
