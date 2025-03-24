package ast

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	. "github.com/jishaocong0910/go-sql-parser/enum"

	. "github.com/jishaocong0910/go-object"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type I_Visitor interface {
	m_Visitor_() *m_Visitor
	newSubqueryVisitor(I_StatementSyntax, Option) *m_Visitor

	SqlOperationType() SqlOperationType
	SingleTableSql() bool
	TableOfSingleTableSql() string
	Tables() *StrSet
	TableColumns() *StrKeyMap[*StrSet]
	SelectColumns() *StrKeyMap[*StrSet]
	WhereColumns() *StrKeyMap[*StrSet]
	UpdateTables() *StrSet
	HintContent() string

	TablesRaw() []string
	TableColumnsRaw() map[string][]string
	SelectColumnsRaw() map[string][]string
	WhereColumnsRaw() map[string][]string
	UpdateTablesRaw() []string

	Option() Option
	Warning() string
}

type m_Visitor struct {
	I I_Visitor
	// SQL语句
	sql string
	// 当前访问的完整SQL语法对象
	is I_StatementSyntax
	// 选项
	option Option
	// 访问中信息
	visitingInfo
	// 访问后信息
	visitedInfo
	// 查询信息缓存
	queryCache
}

func (this *m_Visitor) m_Visitor_() *m_Visitor {
	return this
}

func (this *m_Visitor) visit(is I_Syntax) {
	if !IsNull(is) {
		e := this.tracks.PushFront(is)
		is.accept(this.I)
		this.tracks.Remove(e)
	}
}

func (this *m_Visitor) visitSubquery(iq I_QuerySyntax) {
	sv := this.I.newSubqueryVisitor(iq, this.option)
	for k, v := range this.inheritedTableAliases {
		sv.inheritedTableAliases[k] = v
	}
	for k, v := range this.tableAliases {
		sv.inheritedTableAliases[k] = v
	}
	for k, v := range this.inheritedSubVisitors {
		sv.inheritedSubVisitors[k] = v
	}
	for k, v := range this.subVisitors {
		sv.inheritedSubVisitors[k] = v
	}
	sv.visit(iq)
	k := uuid.New().String()
	if this.tracks.Len() > 1 {
		if d, ok := this.traceSyntax(1).(*DerivedTableReferenceSyntax); ok {
			k = d.Alias.M_IdentifierSyntax_().Name
		}
	}
	this.subVisitors[k] = sv
}

func (this *m_Visitor) traceSyntax(depth int) I_Syntax {
	e := this.tracks.Front()
	for i := 0; i < depth; i++ {
		e = e.Next()
	}
	return e.Value.(I_Syntax)
}

func (this *m_Visitor) addLocalTableReference(a, n string) {
	if this.tableAliases[a] != "" {
		this.panic("table alias '%s' is duplicate", a)
	}
	this.tableAliases[a] = n
}

func (this *m_Visitor) addAllColumnItem() {
	for _, t := range this.tableAliases {
		this.selectAllColumnTables = append(this.selectAllColumnTables, t)
	}
	for _, sv := range this.subVisitors {
		this.selectAllColumnSubVisitors = append(this.selectAllColumnSubVisitors, sv)
	}
}

func (this *m_Visitor) addColumnItem(ic I_ColumnItemSyntax) {
	if this.inWindowSpecSyntax {
		this.addOtherColumnItems(ic)
	} else if this.inSelectItemListSyntax {
		this.addSelectColumnItems(ic)
	} else if this.inWhereSyntax {
		this.addWhereColumnItems(ic)
	} else if this.inInsertColumnListSyntax || this.isAssignmentSyntaxColumn {
		this.addAssignmentItems(ic)
	} else {
		this.addOtherColumnItems(ic)
	}
}

func (this *m_Visitor) addSelectColumnItems(ic I_ColumnItemSyntax) {
	cis := this.determineColumnItemDefault(ic)
	if len(cis) > 0 {
		var k string
		if sc, ok := this.traceSyntax(1).(*SelectColumnSyntax); ok && sc.Alias != nil {
			k = sc.Alias.AliasName()
		} else {
			k = ic.Column()
		}
		this.selectColumnItems[k] = append(this.selectColumnItems[k], cis)
	}
	return
}

func (this *m_Visitor) addWhereColumnItems(ic I_ColumnItemSyntax) {
	cis := this.determineColumnItemDefault(ic)
	for _, ci := range cis {
		this.whereColumnItems = append(this.whereColumnItems, ci)
	}
	return
}

func (this *m_Visitor) addAssignmentItems(ic I_ColumnItemSyntax) {
	cis := this.determineColumnItemDefault(ic)
	if len(cis) > 0 {
		this.assignmentItems = append(this.assignmentItems, &assignmentItem{column: cis[0]})
	}
	return
}

func (this *m_Visitor) addOtherColumnItems(ic I_ColumnItemSyntax) {
	var cis []*columnItem
	if this.inJoinUsingSyntax {
		cis = this.determineColumnItemInJoinUsingSyntax(ic)
	} else if _, ok := this.is.(I_MultisetSyntax); ok {
		ta := ic.TableAlias()
		if ta != "" {
			this.panicBySyntax(ic, "cannot be used table alias in global clause of multiset syntax")
		}
		for _, sv := range this.visitedInfo.subVisitors {
			subCis := this.determineTableOfSubQuery(ic, sv)
			cis = append(cis, subCis...)
		}
	} else {
		cis = this.determineColumnItemDefault(ic)
	}
	for _, ci := range cis {
		this.otherColumnItems = append(this.otherColumnItems, ci)
	}
	return
}

func (this *m_Visitor) determineColumnItemDefault(ic I_ColumnItemSyntax) (cis []*columnItem) {
	ta := ic.TableAlias()
	cc := ic.Column()
	var t string
	var sv *m_Visitor
	if ta == "" {
		tableCount := len(this.tableAliases)
		subqueryCount := len(this.subVisitors)
		totalCount := tableCount + subqueryCount
		if totalCount > 1 {
			this.addWarning(fmt.Sprintf("column '%s' is ambiguous", ic.Column()))
		} else if totalCount == 0 {
			this.panic("unknown column '%s'", ic.Column())
		} else {
			if tableCount == 1 {
				for _, t = range this.tableAliases {
				}
			} else {
				for _, sv = range this.subVisitors {
				}
			}
		}
	} else {
		if t = this.tableAliases[ta]; t == "" {
			if sv = this.subVisitors[ta]; sv == nil {
				if t = this.inheritedTableAliases[ta]; t == "" {
					sv = this.inheritedSubVisitors[ta]
				}
			}
		}
	}
	if t != "" {
		if cc == "*" {
			this.selectAllColumnTables = append(this.selectAllColumnTables, t)
		} else {
			cis = append(cis, &columnItem{t, cc})
		}
	} else if sv != nil {
		if cc == "*" {
			this.selectAllColumnSubVisitors = append(this.selectAllColumnSubVisitors, sv)
		} else {
			cis = this.determineTableOfSubQuery(ic, sv)
		}
	}
	return
}

func (this *m_Visitor) determineColumnItemInJoinUsingSyntax(ic I_ColumnItemSyntax) (cis []*columnItem) {
	j := this.traceSyntax(3).(*JoinTableReferenceSyntax)
	cis = this.determineTableOfTableReference(ic, j)
	return
}

func (this *m_Visitor) determineTableOfSubQuery(ic I_ColumnItemSyntax, sv *m_Visitor) (cis []*columnItem) {
	tableCount := len(sv.selectAllColumnTables)
	subqueryCount := len(sv.selectAllColumnSubVisitors)
	totalCount := tableCount + subqueryCount
	if totalCount == 1 {
		if tableCount == 1 {
			cis = append(cis, &columnItem{table: sv.selectAllColumnTables[0], column: ic.Column()})
			return
		} else {
			return this.determineTableOfSubQuery(ic, sv.selectAllColumnSubVisitors[0])
		}
	} else if totalCount > 1 {
		this.addWarning(fmt.Sprintf("column '%s' is ambiguous", ic.Column()))
		return
	}
	ciss := sv.visitedInfo.selectColumnItems[ic.Column()]
	if ciss != nil {
		if len(ciss) > 1 {
			this.panicBySyntax(ic, "column '%s' is ambiguous", ic.FullColumn())
		}
		cis = ciss[0]
	}
	if cis == nil {
		this.panic("unknown column '%s'", ic.FullColumn())
	}
	return
}

func (this *m_Visitor) determineTableOfTableReference(ic I_ColumnItemSyntax, it I_TableReferenceSyntax) (cis []*columnItem) {
	if i, ok := it.(I_NameTableReferenceSyntax); ok {
		tn := i.M_NameTableReferenceSyntax_()
		cis = append(cis, &columnItem{tn.TableNameItem.FullTableName(), ic.Column()})
	} else if d, ok := it.(*DerivedTableReferenceSyntax); ok {
		sv := this.subVisitors[d.Alias.M_IdentifierSyntax_().Name]
		cis = this.determineTableOfSubQuery(ic, sv)
	} else if j, ok := it.(*JoinTableReferenceSyntax); ok {
		cis = append(cis, this.determineTableOfTableReference(ic, j.Left)...)
		cis = append(cis, this.determineTableOfTableReference(ic, j.Right)...)
	}
	return
}

func (this *m_Visitor) addDeleteItem(table string) {
	this.deleteItems = append(this.deleteItems, &deleteItem{table})
}

func (this *m_Visitor) addTableColumns(ci *columnItem) {
	columns := this.tableColumns.Get(ci.table)
	if columns == nil {
		columns = NewStrSet(this.option.ColumnCaseSensitive)
		this.tableColumns.Put(ci.table, columns)
	}
	columns.Add(ci.column)
}

func (this *m_Visitor) addSelectColumns(ci *columnItem) {
	columns := this.selectColumns.Get(ci.table)
	if columns == nil {
		columns = NewStrSet(this.option.ColumnCaseSensitive)
		this.selectColumns.Put(ci.table, columns)
	}
	columns.Add(ci.column)
}

func (this *m_Visitor) addWhereColumns(ci *columnItem) {
	columns := this.whereColumns.Get(ci.table)
	if columns == nil {
		columns = NewStrSet(this.option.ColumnCaseSensitive)
		this.whereColumns.Put(ci.table, columns)
	}
	columns.Add(ci.column)
}

func (this *m_Visitor) panic(msg string, a ...any) {
	this.panicBySyntax(this.traceSyntax(0), msg, a...)
}

func (this *m_Visitor) panicBySyntax(is I_Syntax, msg string, a ...any) {
	panic(parseError(this.buildErrorMsg(is, msg, a...)))
}

func (this *m_Visitor) addWarning(msg string, a ...any) {
	this.warnings = append(this.warnings, this.buildErrorMsg(this.traceSyntax(0), msg, a...))
}

func (this *m_Visitor) buildErrorMsg(is I_Syntax, msg string, a ...any) string {
	var builder strings.Builder
	if msg != "" {
		msg = fmt.Sprintf(msg, a...)
		builder.WriteString(msg)
	}
	builder.WriteString("\n")
	chars := []rune(this.sql)
	for i := range chars {
		c := chars[i]
		if i == is.M_Syntax_().BeginPos {
			builder.WriteString("↪")
		}
		builder.WriteRune(c)
		if i == is.M_Syntax_().EndPos-1 {
			builder.WriteString("↩")
		}
	}
	return builder.String()
}

func (this *m_Visitor) visitAllColumnSyntax(*AllColumnSyntax) {
	if _, ok := this.traceSyntax(1).(*SelectItemListSyntax); ok {
		this.addAllColumnItem()
	}
}

func (this *m_Visitor) visitAssignmentSyntax(s *AssignmentSyntax) {
	this.isAssignmentSyntaxColumn = true
	this.visit(s.Column)
	this.isAssignmentSyntaxColumn = false
	this.visit(s.Value)
}

func (this *m_Visitor) visitAggregateFunctionSyntax(s I_AggregateFunctionSyntax) {
	a := s.M_AggregateFunctionSyntax_()
	if !a.AllColumnParameter {
		a.I.M_FunctionSyntax_().accept(this.I)
	}
	this.visit(a.Over)
}

func (this *m_Visitor) visitCaseSyntax(s *CaseSyntax) {
	this.visit(s.ValueExpr)
	this.visit(s.WhenItemList)
	this.visit(s.ElseExr)
}

func (this *m_Visitor) visitCaseWhenItemSyntax(s *CaseWhenItemSyntax) {
	this.visit(s.Condition)
	this.visit(s.Result)
}

func (this *m_Visitor) visitSelectItemListSyntax(s *SelectItemListSyntax) {
	this.inSelectItemListSyntax = true
	for _, item := range s.elements {
		this.visit(item)
	}
	this.inSelectItemListSyntax = false
}

func (this *m_Visitor) visitDerivedTableTableReferenceSyntax(s *DerivedTableReferenceSyntax) {
	this.visit(s.Query)
}

func (this *m_Visitor) visitHavingSyntax(s *HavingSyntax) {
	this.visit(s.Condition)
}

func (this *m_Visitor) visitInsertColumnListSyntax(s *InsertColumnListSyntax) {
	this.inInsertColumnListSyntax = true
	for _, item := range s.elements {
		this.visit(item)
	}
	this.inInsertColumnListSyntax = false
}

func (this *m_Visitor) visitJoinOnSyntax(s *JoinOnSyntax) {
	this.visit(s.Condition)
}

func (this *m_Visitor) visitJoinTableReferenceSyntax(s *JoinTableReferenceSyntax) {
	this.visit(s.Left)
	this.visit(s.Right)
	this.visit(s.JoinCondition)
}

func (this *m_Visitor) visitJoinUsingSyntax(s *JoinUsingSyntax) {
	this.inJoinUsingSyntax = true
	this.visit(s.ColumnList)
	this.inJoinUsingSyntax = false
}

func (this *m_Visitor) visitOrderBySyntax(s *OrderBySyntax) {
	this.visit(s.OrderByItemList)
}

func (this *m_Visitor) visitOrderingItemSyntax(s *OrderingItemSyntax) {
	this.visit(s.Column)
}

func (this *m_Visitor) visitPartitionBySyntax(s *PartitionBySyntax) {
	this.visit(s.Expr)
}

func (this *m_Visitor) visitPropertySyntax(s *PropertySyntax) {
	this.addColumnItem(s)
}

func (this *m_Visitor) visitSelectColumnSyntax(s *SelectColumnSyntax) {
	this.visit(s.Expr)
}

func (this *m_Visitor) visitValueListSyntax(s *ValueListSyntax) {
	for i := 0; i < len(s.elements); i++ {
		this.assignmentItems[i].values = append(this.assignmentItems[i].values, s.elements[i].(I_ExprSyntax))
	}
}

func (this *m_Visitor) visitWhereSyntax(s *WhereSyntax) {
	this.inWhereSyntax = true
	this.visit(s.Condition)
	this.inWhereSyntax = false
}

func (this *m_Visitor) visitOverSyntax(s *OverSyntax) {
	if _, ok := s.Window.(*WindowSpecSyntax); ok {
		this.visit(s.Window)
	}
}

func (this *m_Visitor) visitWindowSpecSyntax(s *WindowSpecSyntax) {
	this.inWindowSpecSyntax = true
	this.visit(s.PartitionBy)
	this.visit(s.OrderBy)
	this.visit(s.Frame)
	this.inWindowSpecSyntax = false
}

func (this *m_Visitor) visitWindowFrameExprSyntax(s *WindowFrameExprSyntax) {
	this.visit(s.Expr)
}

func (this *m_Visitor) visitWindowFunctionSyntax(s *WindowFunctionSyntax) {
	s.M_FunctionSyntax.accept(this.I)
	this.visitOverSyntax(s.Over)
}

func (this *m_Visitor) visitExistsSyntax(s *ExistsSyntax) {
	this.visit(s.Query)
}

func (this *m_Visitor) visitNameTableReferenceSyntax(s *M_NameTableReferenceSyntax) {
	tn := s.TableNameItem.FullTableName()
	a := tn
	if s.Alias != nil {
		a = s.Alias.M_IdentifierSyntax_().Name
	}
	this.addLocalTableReference(a, tn)
}

func (this *m_Visitor) visitFunctionSyntax(s *M_FunctionSyntax) {
	this.visit(s.Parameters)
}

func (this *m_Visitor) visitGroupBySyntax(s *M_GroupBySyntax) {
	this.visit(s.OrderingItemList)
}

func (this *m_Visitor) visitMultisetSyntax(s *M_MultisetSyntax) {
	this.sqlOperationType = SqlOperationTypes.SELECT
	this.visitSubquery(s.LeftQuery)
	this.visitSubquery(s.RightQuery)
	this.visit(s.OrderBy)
}

func (this *m_Visitor) visitWindowFrameSyntax(s *WindowFrameSyntax) {
	this.visit(s.Extent)
}

func (this *m_Visitor) visitWindowFrameBetweenSyntax(s *WindowFrameBetweenSyntax) {
	this.visit(s.Start)
	this.visit(s.End)
}

func (this *m_Visitor) visitNamedWindowsSyntax(s *NamedWindowsSyntax) {
	this.visit(s.WindowSpec)
}

func (this *m_Visitor) SqlOperationType() SqlOperationType {
	return this.sqlOperationType
}

func (this *m_Visitor) SingleTableSql() bool {
	return this.singleTableSql
}

func (this *m_Visitor) TableOfSingleTableSql() string {
	return this.tableOfSingleTableSql
}

// Tables 获取SQL中所有涉及的表名
func (this *m_Visitor) Tables() *StrSet {
	if this.tables == nil {
		this.tables = NewStrSet(this.option.TableCaseSensitive)
		for _, tn := range this.tableAliases {
			this.tables.Add(tn)
		}
		for _, sv := range this.subVisitors {
			this.tables.AddSet(sv.Tables())
		}
	}
	return this.tables
}

func (this *m_Visitor) TableColumns() *StrKeyMap[*StrSet] {
	if this.tableColumns == nil {
		this.tableColumns = NewStrKeyMap[*StrSet](this.option.TableCaseSensitive)
		for _, t := range this.selectAllColumnTables {
			this.addTableColumns(&columnItem{t, "*"})
		}
		for _, item := range this.selectColumnItems {
			for _, item := range item {
				for _, item := range item {
					this.addTableColumns(item)
				}
			}
		}
		for _, item := range this.whereColumnItems {
			this.addTableColumns(item)
		}
		for _, item := range this.assignmentItems {
			this.addTableColumns(item.column)
		}
		for _, item := range this.otherColumnItems {
			this.addTableColumns(item)
		}
		for _, sv := range this.subVisitors {
			sub := sv.TableColumns()
			for _, table := range sub.Keys() {
				columns := this.tableColumns.Get(table)
				if columns == nil {
					this.tableColumns.Put(table, sub.Get(table))
				} else {
					columns.AddSet(sub.Get(table))
				}
			}
		}
	}
	return this.tableColumns
}

func (this *m_Visitor) SelectColumns() *StrKeyMap[*StrSet] {
	if this.selectColumns == nil {
		this.selectColumns = NewStrKeyMap[*StrSet](this.option.TableCaseSensitive)
		for _, t := range this.selectAllColumnTables {
			this.addSelectColumns(&columnItem{t, "*"})
		}
		for _, item := range this.selectColumnItems {
			for _, item := range item {
				for _, item := range item {
					this.addSelectColumns(item)
				}
			}
		}
		for _, sv := range this.subVisitors {
			sub := sv.SelectColumns()
			for _, table := range sub.Keys() {
				columns := this.selectColumns.Get(table)
				if columns == nil {
					this.selectColumns.Put(table, sub.Get(table))
				} else {
					columns.AddSet(sub.Get(table))
				}
			}
		}
	}
	return this.selectColumns
}

func (this *m_Visitor) WhereColumns() *StrKeyMap[*StrSet] {
	if this.whereColumns == nil {
		this.whereColumns = NewStrKeyMap[*StrSet](this.option.ColumnCaseSensitive)
		for _, item := range this.whereColumnItems {
			this.addWhereColumns(item)
		}
		for _, sv := range this.subVisitors {
			sub := sv.WhereColumns()
			for _, table := range sub.Keys() {
				columns := this.whereColumns.Get(table)
				if columns == nil {
					this.whereColumns.Put(table, sub.Get(table))
				} else {
					columns.AddSet(sub.Get(table))
				}
			}
		}
	}
	return this.whereColumns
}

// UpdateTables DML语句中有做更新的表
//
// e.g.
//
//		 UPDATE tab1 t1
//			  INNER JOIN tab2 t2
//			     ON t1.id = t2.pid
//			  INNER JOIN tab3 t3
//			     ON t1.id = t3.pid
//		 SET
//			  t1.nickname = 'jack',
//			  t2.email = '123@a.com'
//		 WHERE t1.id = 1001
//			  AND t3.type = 1;
//	  以上SQL语句中只有表tab1和tab2有更新数据，tab3无更新，所以此方法结果只有tab1和tab2。
func (this *m_Visitor) UpdateTables() *StrSet {
	if this.updateTables == nil {
		this.updateTables = NewStrSet(this.option.TableCaseSensitive)
		switch this.sqlOperationType {
		case SqlOperationTypes.UPDATE:
			for _, item := range this.assignmentItems {
				this.updateTables.Add(item.column.table)
			}
		case SqlOperationTypes.INSERT:
			this.updateTables.Add(this.tableOfSingleTableSql)
		case SqlOperationTypes.DELETE:
			for _, item := range this.deleteItems {
				this.updateTables.Add(item.table)
			}
		}
	}
	return this.updateTables
}

func (this *m_Visitor) HintContent() string {
	return this.hintContent
}

func (this *m_Visitor) TablesRaw() []string {
	if this.tables == nil {
		this.Tables()
	}
	return this.tables.Raw()
}

func (this *m_Visitor) TableColumnsRaw() map[string][]string {
	if this.tableColumns == nil {
		this.TableColumns()
	}
	m := make(map[string][]string, this.tableColumns.Len())
	for t, cs := range this.tableColumns.Raw() {
		m[t] = cs.Raw()
	}
	return m
}

func (this *m_Visitor) SelectColumnsRaw() map[string][]string {
	if this.selectColumns == nil {
		this.SelectColumns()
	}
	m := make(map[string][]string, this.selectColumns.Len())
	for t, cs := range this.selectColumns.Raw() {
		m[t] = cs.Raw()
	}
	return m
}

func (this *m_Visitor) WhereColumnsRaw() map[string][]string {
	if this.whereColumns == nil {
		this.WhereColumns()
	}
	m := make(map[string][]string, this.whereColumns.Len())
	for t, cs := range this.whereColumns.Raw() {
		m[t] = cs.Raw()
	}
	return m
}

func (this *m_Visitor) UpdateTablesRaw() []string {
	if this.updateTables == nil {
		this.UpdateTables()
	}
	return this.updateTables.Raw()
}

func (this *m_Visitor) Option() Option {
	return this.option
}

func (this *m_Visitor) Warning() string {
	var b strings.Builder
	for i, warning := range this.warnings {
		b.WriteRune('[')
		b.WriteString(strconv.Itoa(i))
		b.WriteRune(']')
		b.WriteString(warning)
	}
	return b.String()
}

func extendVisitor(i I_Visitor, is I_StatementSyntax, opt Option) *m_Visitor {
	return &m_Visitor{
		I:      i,
		sql:    is.M_StatementSyntax_().Sql,
		is:     is,
		option: opt,
		visitingInfo: visitingInfo{
			tracks: list.New(),
		},
		visitedInfo: visitedInfo{
			selectColumnItems:     make(map[string][][]*columnItem),
			inheritedTableAliases: make(map[string]string),
			tableAliases:          make(map[string]string),
			inheritedSubVisitors:  make(map[string]*m_Visitor),
			subVisitors:           make(map[string]*m_Visitor),
		},
	}
}

type Option struct {
	// 表名是否大小写敏感
	TableCaseSensitive bool
	// 字段名是否大小写敏感
	ColumnCaseSensitive bool
}

type visitingInfo struct {
	// 当前访问的语法树深度
	tracks *list.List
	// 当前处于SELECT语句的查询列表中
	inSelectItemListSyntax bool
	// 当前处于INSERT语句的字段列表中
	inInsertColumnListSyntax bool
	// 当前处于WHERE条件中
	inWhereSyntax bool
	// 当前处于JOIN...USING...语法
	inJoinUsingSyntax bool
	// 当前处于Window函数声明中
	inWindowSpecSyntax bool
	// 当前处于赋值语法中的字段
	isAssignmentSyntaxColumn bool
}

type visitedInfo struct {
	// 当前SQL层级查询的列项，k：列名称，v：列信息（可能有多个）。列名称优先为别名，无别名则为列名。列信息是一个二维数组，其中一维是查询列表的字
	// 段，二维是多结果集的同个字段。
	//
	// e.g.
	//  SELECT
	//    t1.col_1 AS name,
	//    t2.col_1 AS name,
	//    tt.name
	//  FROM
	//    tab_1 t1
	//    JOIN tab_2 t2
	//      ON t1.id = t2.id
	//    JOIN (SELECT id, name FROM tab_3
	//          UNION
	//          SELECT id, name FROM tab_4) tt
	//      ON t1.id = tt.id
	//
	//  此字段的内容为：
	//  {
	//    name: [
	//    	[{tab_1, name}], // 对应『t1.col_1 AS name』
	//    	[{tab_2, name}], // 对应『t2.col_1 AS name』
	//    	[{tab_3, name}, {tab_4, name}], // 对应『tt.name』，而『tt.name』是由tab_3和tab_4联合起来的相同字段
	//    ]
	//  }
	selectColumnItems map[string][][]*columnItem
	// 查询所有字段的表
	// e.g.
	//  SELECT
	//    t1.*, t2.*, t3.col1
	//  FROM
	//    tab1 t1
	//    JOIN tab2 t2
	//      ON t1.id = t2.pid
	//    JOIN tab3 t3
	//      ON t1.id = t3.pid
	//  tab1和tab2为查询了所有字段的表
	selectAllColumnTables []string
	// 在被查询所有字段的子查询
	// e.g.
	//  SELECT
	//    t1.id, t2.*
	//  FROM
	//    tab1 t1
	//    JOIN (SELECT id, pid, col1, col2 from tab2) t2
	//      ON t1.id = t2.pid
	//  由于『t2.*』，子查询『(SELECT id, pid, col1, col2 from tab2) t2』被查询了所有字段
	selectAllColumnSubVisitors []*m_Visitor
	// 当前SQL层级查询的列项
	whereColumnItems []*columnItem
	// 当前SQL层级其他列项
	otherColumnItems []*columnItem
	// 赋值项
	assignmentItems []*assignmentItem
	// 删除项
	deleteItems []*deleteItem
	// 继承的父语句的表别名，k：表别名，v：表名
	inheritedTableAliases map[string]string
	// 当前SQL层级的表别名（不包括子查询），如果表没有别名则别名为表名，k：表别名，v：表名
	tableAliases map[string]string
	// 继承的父语句的子查询别名
	inheritedSubVisitors map[string]*m_Visitor
	// 子查询别名和访问者，k：子查询别名，v：访问者
	subVisitors map[string]*m_Visitor
	// 语句类型
	sqlOperationType SqlOperationType
	// 是否单表SQL语句。SQL语句只操作一个表，不包单个表在派生表或括联合查询的SQL
	singleTableSql bool
	// 单表SQL语句中的表名。当SQL是INSERT语句时为插入的表的名称；当SQL为SELECT、UPDATE、DELETE时，只有FROM后面的表应用只出现一个表（非派生表）时为该表的名称
	tableOfSingleTableSql string
	// 暗示
	hintContent string
	// 警告
	warnings []string
}

type queryCache struct {
	// 所有涉及的表名（包括子查询）
	tables *StrSet
	// 所有涉及的表字段（包括子查询），k：表名，v：字段名
	tableColumns *StrKeyMap[*StrSet]
	// 所有涉及的表的查询字段（包括子查询），k：表名，v：查询字段信息
	selectColumns *StrKeyMap[*StrSet]
	// 所有涉及的表的条件字段（包括子查询），k：表名，v：查询字段信息
	whereColumns *StrKeyMap[*StrSet]
	// DML语句有做更改的表
	updateTables *StrSet
	// INSERT语句的字段列表索引，即字段的顺序，从0开始，k：字段名，v：字段索引
	columnsIndexOfInsertStatement *StrKeyMap[int]
	// 赋值信息，k：表名，v：赋值信息
	assignmentInfos *StrKeyMap[*AssignmentInfo]
}

type columnItem struct {
	table, column string
}

type assignmentItem struct {
	column *columnItem
	values []I_ExprSyntax
}

type deleteItem struct {
	table string
}

type AssignmentInfo struct {
}

type parseError string

func (r parseError) Error() string {
	return string(r)
}

func Visit(is I_StatementSyntax) (I_Visitor, error) {
	return VisitWithOption(is, Option{})
}

func VisitWithOption(is I_StatementSyntax, opt Option) (iv I_Visitor, err error) {
	switch is.Dialect() {
	case Dialects.MYSQL:
		iv = newMySqlVisitor(is, opt)
	default:
		return nil, fmt.Errorf("not supported dialect of '%s' yet", is.Dialect().Name)
	}
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(parseError); ok {
				err = e
			} else {
				err = errors.WithStack(fmt.Errorf("%v", r))
			}
		}
	}()
	iv.m_Visitor_().visit(is)
	return iv, nil
}
