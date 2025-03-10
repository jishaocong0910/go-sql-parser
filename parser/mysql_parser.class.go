package parser

import (
	"strconv"
	"strings"

	. "github.com/jishaocong0910/go-sql-parser/ast"
	. "github.com/jishaocong0910/go-sql-parser/enum"
)

type mySqlParser struct {
	*m_Parser
}

func (this *mySqlParser) hasQualifier() bool {
	return this.lexer.(*mySqlLexer).hasQualifier()
}

func (this *mySqlParser) parseIStatementSyntax() (is I_StatementSyntax) {
	is = this.parseIStatementSyntaxInner()
	if q, ok := is.(I_QuerySyntax); ok {
		is = this.parseIQuerySyntaxRest(q)
	}
	for {
		if Tokens.Not(this.token(), Tokens.SEMI) {
			break
		}
		this.nextToken()
	}
	if Tokens.Not(this.token(), Tokens.EOI) {
		this.panicByUnexpectedToken()
	}
	s := is.M_A88DB0CC837F()
	s.Sql = this.sql()
	return
}

func (this *mySqlParser) parseIStatementSyntaxInner() (is I_StatementSyntax) {
	switch this.token() {
	case Tokens.SELECT:
		is = this.parseIQuerySyntax(QuerySyntaxLevels.NORMAL)
	case Tokens.UPDATE:
		is = this.parseMySqlUpdateSyntax()
	case Tokens.INSERT:
		is = this.parseMySqlInsertSyntax()
	case Tokens.DELETE:
		is = this.parseMySqlDeleteSyntax()
	case Tokens.L_PAREN:
		beginPos := this.tokenBeginPos()
		this.nextToken()
		is = this.parseIStatementSyntaxInner()
		this.setBeginPos(is, beginPos)
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(is)
		this.parenthesizingSyntax(is)
	default:
		this.panicByUnexpectedToken()
	}
	return
}

func (this *mySqlParser) parseIQuerySyntax(l QuerySyntaxLevel) (iq I_QuerySyntax) {
	switch this.token() {
	case Tokens.SELECT:
		iq = this.parseMySqlSelectSyntax(l)
	case Tokens.L_PAREN:
		beginPos := this.tokenBeginPos()
		this.nextToken()
		iq = this.parseIQuerySyntax(l)
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setBeginPos(iq, beginPos)
		this.setEndPosDefault(iq)
		this.parenthesizingSyntax(iq)
	default:
		this.panicByUnexpectedToken()
	}
	if QuerySyntaxLevels.Not(l, QuerySyntaxLevels.QUERY_OPERAND) {
		iq = this.parseIQuerySyntaxRest(iq)
	}
	return
}

func (this *mySqlParser) parseIQuerySyntaxRest(before I_QuerySyntax) (last I_QuerySyntax) {
	last = before
	for {
		var mo MultisetOperator
		switch this.token() {
		case Tokens.UNION:
			mo = MultisetOperators.UNION
		case Tokens.EXCEPT:
			mo = MultisetOperators.EXCEPT
		case Tokens.INTERSECT:
			mo = MultisetOperators.INTERSECT
		}
		if mo.Undefined() {
			break
		}

		this.nextToken()
		if s, ok := last.(*MySqlSelectSyntax); ok {
			if ParenthesizeTypes.Not(s.ParenthesizeType, ParenthesizeTypes.TRUE) {
				if s.OrderBy != nil {
					this.panicBySyntax(s, "incorrect usage of ORDER BY")
				}
				if s.LimitSyntax != nil {
					this.panicBySyntax(s, "incorrect usage of LIMIT")
				}
			}
		}

		u := NewMySqlMultisetSyntax()
		this.setBeginPos(u, last.M_5CF6320E8474().BeginPos)
		u.LeftQuery = last
		u.MultisetOperator = mo
		u.AggregateOption = this.parseAggregateOption()

		var nextLevel QuerySyntaxLevel
		if Tokens.Is(this.token(), Tokens.SELECT) {
			nextLevel = QuerySyntaxLevels.QUERY_OPERAND
		} else {
			nextLevel = QuerySyntaxLevels.NORMAL
		}

		u.RightQuery = this.parseIQuerySyntax(nextLevel)
		this.setEndPosDefault(u)
		this.acceptEqualOperandCount(u.LeftQuery, u.RightQuery, false)
		last = u
	}

	if u, ok := last.(*MySqlMultisetSyntax); ok {
		u.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevels.IDENTIFIER)
		u.Limit = this.parseMySqlLimitSyntax()
	}

	return
}

func (this *mySqlParser) parseMySqlSelectSyntax(l QuerySyntaxLevel) (m *MySqlSelectSyntax) {
	m = NewMySqlSelectSyntax()
	this.setBeginPosDefault(m)
	for {
		isOption := false
		switch this.nextToken() {
		case Tokens.IDENTIFIER:
			switch this.tokenValUpper() {
			case "SQL_BUFFER_RESULT":
				m.SqlBufferResult = true
				isOption = true
			case "SQL_CACHE":
				m.SqlCache = true
				isOption = true
			case "SQL_NO_CACHE":
				m.SqlNoCache = true
				isOption = true
			}
		case Tokens.DISTINCT, Tokens.DISTINCTROW:
			m.AggregateOption = AggregateOptions.DISTINCT
			isOption = true
		case Tokens.ALL:
			m.AggregateOption = AggregateOptions.ALL
			isOption = true
		case Tokens.HIGHP_RIORITY:
			m.HighPriority = true
			isOption = true
		case Tokens.STRAIGHT_JOIN:
			m.StraightJoin = true
			isOption = true
		case Tokens.SQL_SMALL_RESULT:
			m.SqlSmallResult = true
			isOption = true
		case Tokens.SQL_BIG_RESULT:
			m.SqlBigResult = true
			isOption = true
		case Tokens.SQL_CALC_FOUND_ROWS:
			m.SqlCalcFoundRows = true
			isOption = true
		}
		if !isOption {
			break
		}
	}

	m.SelectItemList = this.parseSelectItemListSyntax()

	if Tokens.Is(this.token(), Tokens.FROM) {
		this.nextToken()
		m.TableReference = this.parseITableReferenceSyntax(TableReferenceSyntaxLevels.JOIN)
		m.Where = this.parseWhereSyntax()
		if QuerySyntaxLevels.Not(l, QuerySyntaxLevels.QUERY_OPERAND) {
			if g := this.parseMySqlGroupBySyntax(); g != nil {
				m.GroupBy = g
			}
		}
		m.Having = this.parseHavingSyntax()
		m.NamedWindowList = this.parseNamedWindowListSyntax()
		if QuerySyntaxLevels.Not(l, QuerySyntaxLevels.QUERY_OPERAND) {
			m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevels.NORNAL)
		}
		if QuerySyntaxLevels.Not(l, QuerySyntaxLevels.QUERY_OPERAND) {
			m.LimitSyntax = this.parseMySqlLimitSyntax()
		}

		lr := this.parseMySqlLockReadSyntax(false)
		if lr != nil {
			var lrs []*MySqlLockingReadSyntax
			lrs = append(lrs, lr)
			if lr.OfTableName != nil {
				for {
					if Tokens.Not(this.token(), Tokens.FOR) {
						break
					}
					lrs = append(lrs, this.parseMySqlLockReadSyntax(true))
				}
			}
			m.LockingReads = lrs
		}
	}
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlLockReadSyntax(mustOfTable bool) (m *MySqlLockingReadSyntax) {
	switch this.token() {
	case Tokens.FOR:
		m = NewMySqlLockingReadSyntax()
		this.setBeginPosDefault(m)
		if Tokens.Is(this.nextToken(), Tokens.UPDATE) {
			m.LockingRead = MySqlLockingReads.FOR_UPDATE
		} else if this.tokenValUpper() == "SHARE" {
			m.LockingRead = MySqlLockingReads.FOR_SHARE
		} else {
			this.panicByUnexpectedToken()
		}
		this.nextToken()

		if mustOfTable {
			this.acceptAnyToken(Tokens.OF)
		}
		if Tokens.Is(this.token(), Tokens.OF) {
			this.nextToken()
			m.OfTableName = this.parseMySqlIdentifierSyntax(true)
		}

		switch this.tokenValUpper() {
		case "NOWAIT":
			m.LockingReadConcurrency = MySqlLockingReadConcurrencys.NO_WAIT
			this.nextToken()
		case "SKIP":
			this.nextToken()
			this.acceptAnyTokenVal("LOCKED")
			m.LockingReadConcurrency = MySqlLockingReadConcurrencys.SKIP_LOCKED
			this.nextToken()
		}
		this.setEndPosDefault(m)
	case Tokens.LOCK:
		m = NewMySqlLockingReadSyntax()
		this.setBeginPosDefault(m)
		this.nextToken()
		this.acceptAnyToken(Tokens.IN)
		this.nextToken()
		this.acceptAnyTokenVal("SHARE")
		this.nextToken()
		this.acceptAnyTokenVal("MODE")
		this.nextToken()
		m.LockingRead = MySqlLockingReads.LOCK_IN_SHARE_MODE
		this.setEndPosDefault(m)
	}
	return
}

func (this *mySqlParser) parseMySqlUpdateSyntax() (m *MySqlUpdateSyntax) {
	m = NewMySqlUpdateSyntax()
	this.setBeginPosDefault(m)
	this.nextTokenIncludeComment()

	m.Hint = this.parseHintSyntax()

	for {
		isOption := false
		switch this.token() {
		case Tokens.LOW_PRIORITY:
			m.LowPriority = true
			isOption = true
		case Tokens.IGNORE:
			m.Ignore = true
			isOption = true
		}

		if !isOption {
			break
		}
		this.nextToken()
	}

	m.TableReference = this.parseITableReferenceSyntax(TableReferenceSyntaxLevels.JOIN)
	this.acceptAnyToken(Tokens.SET)
	this.nextToken()

	cl := NewAssignmentListSyntax()
	this.setBeginPosDefault(cl)
	for {
		c := this.parseAssignmentSyntax()
		cl.Add(c)
		if Tokens.Not(this.token(), Tokens.COMMA) {
			break
		}
		this.nextToken()
	}
	this.setEndPosDefault(cl)

	m.AssignmentList = cl
	m.Where = this.parseWhereSyntax()
	m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevels.NORNAL)
	m.Limit = this.parseMySqlLimitSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseHintSyntax() (h *HintSyntax) {
	if Tokens.Not(this.token(), Tokens.COMMENT) {
		return
	}
	comment := this.tokenVal()
	if !strings.HasPrefix(comment, "/*+") {
		return
	}
	h = NewHintSyntax()
	h.CommentType = CommentTypes.MULTI_LINE
	h.Content = comment[3 : len(comment)-2]
	this.setBeginPosDefault(h)
	this.setEndPosDefault(h)
	this.nextToken()
	return
}

func (this *mySqlParser) parseMySqlInsertSyntax() (m *MySqlInsertSyntax) {
	m = NewMySqlInsertSyntax()
	this.setBeginPosDefault(m)
	for {
		isOption := false
		switch this.nextToken() {
		case Tokens.LOW_PRIORITY:
			m.LowPriority = true
			isOption = true
		case Tokens.DELAYED:
			m.Delayed = true
			isOption = true
		case Tokens.HIGHP_RIORITY:
			m.HighPriority = true
			isOption = true
		case Tokens.IGNORE:
			m.Ignore = true
			isOption = true
		}

		if !isOption {
			break
		}
	}

	if Tokens.Is(this.token(), Tokens.INTO) {
		this.nextToken()
	}
	this.acceptAnyToken(Tokens.IDENTIFIER)

	t := NewMySqlNameTableReferenceSyntax()
	this.setBeginPosDefault(t)
	t.TableNameItem = this.parseTableNameItemSyntax()
	m.NameTableReference = t
	this.setEndPosDefault(m)

	if Tokens.Is(this.token(), Tokens.L_PAREN) {
		icl := NewInsertColumnListSyntax()
		this.setBeginPosDefault(icl)

		if Tokens.Not(this.nextToken(), Tokens.R_PAREN) {
			for {
				i := this.parseMySqlIdentifierSyntax(true)
				icl.Add(i)
				if Tokens.Not(this.token(), Tokens.COMMA) {
					break
				}
				this.nextToken()
			}
		}

		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(icl)
		m.InsertColumnList = icl
	}

	columnNum := -1
	if m.InsertColumnList != nil {
		columnNum = m.InsertColumnList.Len()
	}

	rowConstructorList := false
	if Tokens.Is(this.token(), Tokens.VALUES) {
		if Tokens.Is(this.nextToken(), Tokens.ROW) {
			rowConstructorList = true
		}
		m.ValueListList = this.parseMySqlValueListListSyntax(columnNum, rowConstructorList)
	} else if this.equalTokenVal("VALUE") {
		this.nextToken()
		m.ValueListList = this.parseMySqlValueListListSyntax(columnNum, rowConstructorList)
	} else if Tokens.Is(this.token(), Tokens.SET) {
		al := NewAssignmentListSyntax()
		this.setBeginPosDefault(al)
		this.nextToken()
		for {
			ia := this.parseAssignmentSyntax()
			al.Add(ia)
			if Tokens.Not(this.token(), Tokens.COMMA) {
				break
			}
			this.nextToken()
		}
		this.setEndPosDefault(al)
		m.AssignmentList = al
	} else {
		this.panicByUnexpectedToken()
	}

	if !rowConstructorList {
		if Tokens.Is(this.token(), Tokens.AS) {
			this.nextToken()
			m.RowAlias = this.parseMySqlIdentifierSyntax(true)
			if Tokens.Is(this.token(), Tokens.L_PAREN) {
				ml := this.parseMySqlIdentifierListSyntax()
				oc := ml.OperandCount()
				if columnNum != oc {
					this.panicBySyntax(ml, "column alias count doesn't match value count, column: "+strconv.Itoa(columnNum)+", alias: "+strconv.Itoa(oc))
				}
				m.ColumnAliasList = ml
			}
		}
	}

	if Tokens.Is(this.token(), Tokens.ON) {
		ial := NewAssignmentListSyntax()
		this.setBeginPosDefault(ial)
		this.nextToken()
		this.acceptAnyTokenVal("DUPLICATE")
		this.nextToken()
		this.acceptAnyToken(Tokens.KEY)
		this.nextToken()
		this.acceptAnyToken(Tokens.UPDATE)
		this.nextToken()
		for {
			ia := this.parseAssignmentSyntax()
			ial.Add(ia)
			if Tokens.Not(this.token(), Tokens.COMMA) {
				break
			}
			this.nextToken()
		}
		this.setEndPosDefault(ial)
		m.OnDuplicateKeyUpdateAssignmentList = ial
	}
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlValueListListSyntax(columnNum int, rowConstructorList bool) (vll *MySqlValueListListSyntax) {
	vll = NewMySqlValueListListSyntax()
	this.setBeginPosDefault(vll)
	vll.RowConstructorList = rowConstructorList
	for {
		if rowConstructorList {
			this.acceptAnyToken(Tokens.ROW)
			this.nextToken()
		}

		ol := this.parseSingleOperandExprListSyntax()
		oc := ol.OperandCount()
		if columnNum != -1 {
			if columnNum != oc {
				this.panicBySyntax(ol, "column count doesn't match value count, column: "+strconv.Itoa(columnNum)+", value: "+strconv.Itoa(oc))
			}
		} else {
			columnNum = oc
		}

		vl := NewValueListSyntax()
		this.setBeginPosDefault(vl)
		for i := 0; i < ol.Len(); i++ {
			vl.Add(ol.Get(i))
		}
		this.setEndPosDefault(vl)
		vll.Add(vl)

		if Tokens.Not(this.token(), Tokens.COMMA) {
			break
		}
		this.nextToken()
	}
	this.setEndPosDefault(vll)
	return
}

func (this *mySqlParser) parseMySqlDeleteSyntax() (m *MySqlDeleteSyntax) {
	m = NewMySqlDeleteSyntax()
	this.setBeginPosDefault(m)
	for {
		isOption := false
		switch this.nextToken() {
		case Tokens.LOW_PRIORITY:
			m.LowPriority = true
			isOption = true
		case Tokens.IGNORE:
			m.Ignore = true
			isOption = true
		}
		if this.equalTokenVal("QUICK") {
			m.Quick = true
			isOption = true
		}
		if !isOption {
			break
		}
	}
	if Tokens.Not(this.token(), Tokens.FROM) {
		m.MultiDeleteMode = MySqlMultiDeleteModes.MODE1
		ml := NewMySqlMultiDeleteTableAliasListSyntax()
		this.setBeginPosDefault(ml)
		m.MultiDeleteTableAliasList = ml
		for {
			md := NewMySqlMultiDeleteTableAliasSyntax()
			this.setBeginPosDefault(md)
			md.Alias = this.parseMySqlIdentifierSyntax(true)
			if Tokens.Is(this.token(), Tokens.DOT) {
				this.nextToken()
				this.acceptAnyToken(Tokens.STAR)
				this.nextToken()
				md.HasStar = true
			}
			this.setEndPosDefault(md)
			ml.Add(md)
			if Tokens.Not(this.token(), Tokens.COMMA) {
				break
			}
			this.nextToken()
		}
		this.setEndPosDefault(ml)
		this.acceptAnyToken(Tokens.FROM)
		this.nextToken()
		m.TableReference = this.parseITableReferenceSyntax(TableReferenceSyntaxLevels.JOIN)
	} else {
		this.nextToken()
		i := this.parseMySqlIdentifierSyntax(true)

		switch this.token() {
		case Tokens.DOT:
			switch this.nextToken() {
			case Tokens.STAR:
				this.nextToken()
				m.MultiDeleteMode = MySqlMultiDeleteModes.MODE2

				md := NewMySqlMultiDeleteTableAliasSyntax()
				this.setBeginPos(md, i.BeginPos)
				md.Alias = i
				md.HasStar = true
				this.setEndPos(md, i.EndPos)

				m.MultiDeleteTableAliasList = this.parseMySqlMultiDeleteTableAliasListSyntax(md)
				this.acceptAnyToken(Tokens.USING)
				this.nextToken()
				m.TableReference = this.parseITableReferenceSyntax(TableReferenceSyntaxLevels.JOIN)
			case Tokens.IDENTIFIER:
				t := NewTableNameItemSyntax()
				this.setBeginPos(t, i.BeginPos)
				t.Catalog = i
				t.TableName = this.parseMySqlIdentifierSyntax(false)
				this.setEndPosDefault(t)

				tnt := NewMySqlNameTableReferenceSyntax()
				this.setBeginPos(tnt, t.BeginPos)
				tnt.TableNameItem = t
				if alias := this.parseIAliasSyntax(AliasSyntaxLevels.IDENTIFIER); alias != nil {
					tnt.Alias = alias.(I_IdentifierSyntax)
				}
				this.setEndPos(tnt, t.EndPos)
				m.TableReference = tnt
			default:
				this.panicByUnexpectedToken()
			}
		case Tokens.COMMA:
			m.MultiDeleteMode = MySqlMultiDeleteModes.MODE2
			md := NewMySqlMultiDeleteTableAliasSyntax()
			this.setBeginPos(md, i.BeginPos)
			md.Alias = i
			this.setEndPos(md, i.EndPos)

			m.MultiDeleteTableAliasList = this.parseMySqlMultiDeleteTableAliasListSyntax(md)
			this.acceptAnyToken(Tokens.USING)
			this.nextToken()
			m.TableReference = this.parseITableReferenceSyntax(TableReferenceSyntaxLevels.JOIN)
		default:
			t := NewTableNameItemSyntax()
			this.setBeginPos(t, i.BeginPos)
			t.TableName = i
			this.setEndPos(t, i.EndPos)

			tnt := NewMySqlNameTableReferenceSyntax()
			this.setBeginPos(tnt, i.BeginPos)
			tnt.TableNameItem = t
			if alias := this.parseIAliasSyntax(AliasSyntaxLevels.IDENTIFIER); alias != nil {
				tnt.Alias = alias.(I_IdentifierSyntax)
			}
			this.setEndPos(tnt, t.EndPos)
			m.TableReference = tnt
		}
	}
	m.Where = this.parseWhereSyntax()
	m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevels.NORNAL)
	m.Limit = this.parseMySqlLimitSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlMultiDeleteTableAliasListSyntax(m *MySqlMultiDeleteTableAliasSyntax) (ml *MySqlMultiDeleteTableAliasListSyntax) {
	ml = NewMySqlMultiDeleteTableAliasListSyntax()
	this.setBeginPos(ml, m.BeginPos)
	ml.Add(m)

	for {
		if Tokens.Not(this.token(), Tokens.COMMA) {
			break
		}
		this.nextToken()
		ml.Add(this.parseMySqlMultiDeleteTableAliasSyntax())
	}

	this.setEndPosDefault(ml)
	return
}

func (this *mySqlParser) parseMySqlMultiDeleteTableAliasSyntax() (m *MySqlMultiDeleteTableAliasSyntax) {
	m = NewMySqlMultiDeleteTableAliasSyntax()
	this.setBeginPosDefault(m)
	m.Alias = this.parseMySqlIdentifierSyntax(true)
	if Tokens.Is(this.token(), Tokens.DOT) {
		this.nextToken()
		this.acceptAnyToken(Tokens.STAR)
		this.nextToken()
		m.HasStar = true
	}
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlIdentifierListSyntax() (ml *MySqlIdentifierListSyntax) {
	ml = NewMySqlIdentifierListSyntax()
	this.setBeginPosDefault(ml)
	this.acceptAnyToken(Tokens.L_PAREN)
	this.nextToken()
	if Tokens.Not(this.token(), Tokens.R_PAREN) {
		for {
			ml.Add(this.parseMySqlIdentifierSyntax(true))
			if Tokens.Not(this.token(), Tokens.COMMA) {
				break
			}
			this.nextToken()
		}
	}
	this.acceptAnyToken(Tokens.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(ml)
	this.parenthesizingSyntax(ml)
	return
}

func (this *mySqlParser) parseAggregateOption() (a AggregateOption) {
	switch this.token() {
	case Tokens.DISTINCT:
		a = AggregateOptions.DISTINCT
		this.nextToken()
	case Tokens.ALL:
		a = AggregateOptions.ALL
		this.nextToken()
	}
	return
}

func (this *mySqlParser) parseOrderBySyntax(l OrderingItemSyntaxLevel) (o *OrderBySyntax) {
	if Tokens.Not(this.token(), Tokens.ORDER) {
		return
	}
	o = NewOrderBySyntax()
	this.setBeginPosDefault(o)

	this.nextToken()
	this.acceptAnyToken(Tokens.BY)
	this.nextToken()

	o.OrderByItemList = this.parseOrderingItemListSyntax(l)
	this.setEndPosDefault(o)
	return
}

func (this *mySqlParser) parseOrderingItemListSyntax(l OrderingItemSyntaxLevel) (ol *OrderingItemListSyntax) {
	ol = NewOrderingItemListSyntax()
	this.setBeginPosDefault(ol)
	for {
		ol.Add(this.parseOrderingItemSyntax(l))
		if Tokens.Not(this.token(), Tokens.COMMA) {
			break
		}
		this.nextToken()
	}
	this.setEndPosDefault(ol)
	return
}

func (this *mySqlParser) parseOrderingItemSyntax(l OrderingItemSyntaxLevel) (o *OrderingItemSyntax) {
	o = NewOrderingItemSyntax()
	this.setBeginPosDefault(o)
	ic := this.parseIColumnItemSyntax()
	if OrderingItemSyntaxLevels.Is(l, OrderingItemSyntaxLevels.IDENTIFIER) {
		if _, ok := ic.(*PropertySyntax); ok {
			this.panicBySyntax(ic, "cannot be used table alias in global clause of multiset syntax")
		}
	}
	o.Column = ic
	if Tokens.Is(this.token(), Tokens.ASC) {
		o.OrderingSequence = OrderingSequences.ASC
		this.nextToken()
	} else if Tokens.Is(this.token(), Tokens.DESC) {
		o.OrderingSequence = OrderingSequences.DESC
		this.nextToken()
	}
	this.setEndPosDefault(o)
	return
}

func (this *mySqlParser) parseMySqlLimitSyntax() (m *MySqlLimitSyntax) {
	if Tokens.Not(this.token(), Tokens.LIMIT) {
		return
	}
	m = NewMySqlLimitSyntax()
	this.setBeginPosDefault(m)

	this.nextToken()
	this.acceptAnyToken(Tokens.DECIMAL_NUMBER)

	d := this.parseDecimalNumberSyntax()
	if Tokens.Is(this.token(), Tokens.COMMA) {
		m.Offset = d
		this.nextToken()
		this.acceptAnyToken(Tokens.DECIMAL_NUMBER)
		m.RowCount = this.parseDecimalNumberSyntax()
	} else if this.equalTokenVal("OFFSET") {
		m.RowCount = d
		this.nextToken()
		this.acceptAnyToken(Tokens.DECIMAL_NUMBER)
		m.Offset = this.parseDecimalNumberSyntax()
	} else {
		m.RowCount = d
	}
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseSelectItemListSyntax() (sil *SelectItemListSyntax) {
	sil = NewSelectItemListSyntax()
	this.setBeginPosDefault(sil)
	for {
		si := this.parseISelectItemSyntax(!sil.HasAllColumn)
		if _, ok := si.(*AllColumnSyntax); ok {
			sil.HasAllColumn = true
		}
		sil.Add(si)
		if Tokens.Not(this.token(), Tokens.COMMA) {
			break
		}
		this.nextToken()
	}
	this.setEndPosDefault(sil)
	return
}

func (this *mySqlParser) parseITableReferenceSyntax(l TableReferenceSyntaxLevel) (it I_TableReferenceSyntax) {
	if Tokens.Is(this.token(), Tokens.L_PAREN) {
		beginPos := this.tokenBeginPos()
		this.nextToken()
		it = this.parseITableReferenceSyntax(TableReferenceSyntaxLevels.JOIN)
		this.setBeginPos(it, beginPos)
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(it)

		if d, ok := it.(*DerivedTableReferenceSyntax); ok {
			this.parenthesizingSyntax(d.Query)
			if alias := this.parseIAliasSyntax(AliasSyntaxLevels.IDENTIFIER); alias == nil {
				this.panicBySyntax(d, "every derived table must have its own alias")
			} else {
				d.Alias = alias.(I_IdentifierSyntax)
			}
		} else {
			this.parenthesizingSyntax(it)
		}
	} else {
		it = this.parseITableReferenceSyntaxInner()
	}

	if TableReferenceSyntaxLevels.Is(l, TableReferenceSyntaxLevels.JOIN) {
		it = this.parseITableReferenceSyntaxRest(it)
	}
	return
}

func (this *mySqlParser) parseMySqlIndexHintSyntax() (m *MySqlIndexHintSyntax) {
	var mode MySqlIndexHintMode
	switch this.token() {
	case Tokens.USE:
		mode = MySqlIndexHintModes.USE
	case Tokens.IGNORE:
		mode = MySqlIndexHintModes.IGNORE
	case Tokens.FORCE:
		mode = MySqlIndexHintModes.FORCE
	default:
		return
	}
	m = NewMySqlIndexHintSyntax()
	this.setBeginPosDefault(m)
	m.IndexHintMode = mode
	this.nextToken()
	this.acceptAnyToken(Tokens.INDEX, Tokens.KEY)
	if Tokens.Is(this.nextToken(), Tokens.FOR) {
		switch this.nextToken() {
		case Tokens.JOIN:
			m.IndexHintFor = MySqlIndexHintFors.JOIN
			this.nextToken()
		case Tokens.GROUP:
			m.IndexHintFor = MySqlIndexHintFors.GROUP_BY
			this.nextToken()
			this.acceptAnyToken(Tokens.BY)
			this.nextToken()
		case Tokens.ORDER:
			m.IndexHintFor = MySqlIndexHintFors.ORDER_BY
			this.nextToken()
			this.acceptAnyToken(Tokens.BY)
			this.nextToken()
		default:
			this.panicByUnexpectedToken()
		}
	}

	m.IndexList = this.parseMySqlIdentifierListSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseITableReferenceSyntaxInner() (it I_TableReferenceSyntax) {
	switch this.token() {
	case Tokens.IDENTIFIER:
		tnt := NewMySqlNameTableReferenceSyntax()
		this.setBeginPosDefault(tnt)
		tnt.TableNameItem = this.parseTableNameItemSyntax()
		tnt.PartitionList = this.parsePartitionListSyntax()
		if alias := this.parseIAliasSyntax(AliasSyntaxLevels.IDENTIFIER); alias != nil {
			tnt.Alias = alias.(I_IdentifierSyntax)
		}
		tnt.IndexHintList = this.parseMySqlIndexHintListSyntax()
		this.setEndPosDefault(tnt)
		it = tnt
	case Tokens.SELECT:
		if Tokens.Not(this.prevToken(), Tokens.L_PAREN) {
			this.panicByToken("subquery expression must be parenthesized")
		}
		dtt := NewDerivedTableTableReferenceSyntax()
		dtt.Query = this.parseIQuerySyntax(QuerySyntaxLevels.NORMAL)
		it = dtt
	case Tokens.DUAL:
		d := NewDualTableReferenceSyntax()
		this.setBeginPosDefault(d)
		this.nextToken()
		this.setEndPosDefault(d)
		it = d
	default:
		this.panicByUnexpectedToken()
	}
	return
}

func (this *mySqlParser) parsePartitionListSyntax() (pl *PartitionListSyntax) {
	if Tokens.Not(this.token(), Tokens.PARTITION) {
		return
	}
	pl = NewPartitionListSyntax()
	this.setBeginPosDefault(pl)
	this.nextToken()
	pl.PartitionList = this.parseMySqlIdentifierListSyntax()
	this.setEndPosDefault(pl)
	return
}

func (this *mySqlParser) parseMySqlIndexHintListSyntax() (ihl *MySqlIndexHintListSyntax) {
	ih := this.parseMySqlIndexHintSyntax()
	if ih != nil {
		ihl = NewMySqlIndexHintListSyntax()
		ihl.Add(ih)
		for {
			ih := this.parseMySqlIndexHintSyntax()
			if ih == nil {
				break
			}
			ihl.Add(ih)
		}
	}
	return
}

func (this *mySqlParser) parseITableReferenceSyntaxRest(source I_TableReferenceSyntax) (last I_TableReferenceSyntax) {
	last = source
	natural := false
	if Tokens.Is(this.token(), Tokens.NATURAL) {
		natural = true
		this.nextToken()
	}

	var joinType JoinType
	switch this.token() {
	case Tokens.LEFT:
		if Tokens.Is(this.nextToken(), Tokens.OUTER) {
			this.nextToken()
		}
		this.acceptAnyToken(Tokens.JOIN)
		joinType = JoinTypes.LEFT_OUTER_JOIN
		this.nextToken()
	case Tokens.RIGHT:
		if Tokens.Is(this.nextToken(), Tokens.OUTER) {
			this.nextToken()
		}
		this.acceptAnyToken(Tokens.JOIN)
		joinType = JoinTypes.RIGHT_OUTER_JOIN
		this.nextToken()
	case Tokens.INNER:
		this.nextToken()
		this.acceptAnyToken(Tokens.JOIN)
		joinType = JoinTypes.INNER_JOIN
		this.nextToken()
	case Tokens.JOIN:
		joinType = JoinTypes.JOIN
		this.nextToken()
	case Tokens.COMMA:
		joinType = JoinTypes.COMMA
		this.nextToken()
	case Tokens.STRAIGHT_JOIN:
		joinType = JoinTypes.STRAIGHT_JOIN
		this.nextToken()
	case Tokens.CROSS:
		this.nextToken()
		this.acceptAnyToken(Tokens.JOIN)
		joinType = JoinTypes.CROSS_JOIN
		this.nextToken()
	}

	if !joinType.Undefined() {
		jt := NewJoinTableReferenceSyntax()
		jt.Left = last
		jt.Natural = natural
		jt.JoinType = joinType
		jt.Right = this.parseITableReferenceSyntax(TableReferenceSyntaxLevels.DERIVED)

		if !natural {
			switch this.token() {
			case Tokens.ON:
				this.nextToken()
				jt.JoinCondition = this.parseJoinOnSyntax()
			case Tokens.USING:
				this.nextToken()
				jt.JoinCondition = this.parseJoinUsingSyntax()
			default:
				if JoinTypes.Not(joinType, JoinTypes.COMMA, JoinTypes.INNER_JOIN, JoinTypes.CROSS_JOIN, JoinTypes.STRAIGHT_JOIN) {
					this.panicByUnexpectedToken()
				}
			}
		}
		last = this.parseITableReferenceSyntaxRest(jt)
	}
	return
}

func (this *mySqlParser) parseWhereSyntax() (w *WhereSyntax) {
	if Tokens.Not(this.token(), Tokens.WHERE) {
		return
	}
	w = NewWhereSyntax()
	this.setBeginPosDefault(w)
	this.nextToken()
	w.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	this.setEndPosDefault(w)
	return
}

func (this *mySqlParser) parseMySqlGroupBySyntax() (m *MySqlGroupBySyntax) {
	if Tokens.Not(this.token(), Tokens.GROUP) {
		return
	}
	m = NewMySqlGroupBySyntax()
	this.setBeginPosDefault(m)
	this.nextToken()
	this.acceptAnyToken(Tokens.BY)
	this.nextToken()
	// GROUP BY语法的ASC、DESC修饰在8.0中被删除，本解析器依然支持，兼容5.x
	m.OrderingItemList = this.parseOrderingItemListSyntax(OrderingItemSyntaxLevels.NORNAL)
	if Tokens.Is(this.token(), Tokens.WITH) {
		this.nextToken()
		this.acceptAnyTokenVal("ROLLUP")
		m.WithRollup = true
		this.nextToken()
	}
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseHavingSyntax() (h *HavingSyntax) {
	if Tokens.Not(this.token(), Tokens.HAVING) {
		return
	}
	h = NewHavingSyntax()
	this.setBeginPosDefault(h)
	this.nextToken()
	h.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	this.setEndPosDefault(h)
	return
}

func (this *mySqlParser) parseAssignmentSyntax() (a *AssignmentSyntax) {
	a = NewAssignmentSyntax()
	this.setBeginPosDefault(a)
	a.Column = this.parseIColumnItemSyntax()
	this.acceptAnyToken(Tokens.EQ, Tokens.COLON_EQ)

	if Tokens.Is(this.nextToken(), Tokens.DEFAULT) {
		a.Default = true
		this.nextToken()
	} else {
		a.Value = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.BOOLEAN_PREDICATE)
	}
	this.setEndPosDefault(a)
	return
}

func (this *mySqlParser) parseTableNameItemSyntax() (t *TableNameItemSyntax) {
	t = NewTableNameItemSyntax()
	this.setBeginPosDefault(t)
	i := this.parseMySqlIdentifierSyntax(true)
	if Tokens.Is(this.token(), Tokens.DOT) {
		this.nextToken()
		t.Catalog = i
		t.TableName = this.parseMySqlIdentifierSyntax(true)
	} else {
		t.TableName = i
	}
	this.setEndPosDefault(t)
	return
}

func (this *mySqlParser) parseSingleOperandExprListSyntax() (el *ExprListSyntax) {
	el = NewExprListSyntax()
	this.setBeginPosDefault(el)
	this.acceptAnyToken(Tokens.L_PAREN)
	this.nextToken()
	if Tokens.Not(this.token(), Tokens.R_PAREN) {
		for {
			el.Add(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
			if Tokens.Not(this.token(), Tokens.COMMA) {
				break
			}
			this.nextToken()
		}
	}
	this.acceptAnyToken(Tokens.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(el)
	this.parenthesizingSyntax(el)
	return
}

func (this *mySqlParser) parseMySqlIdentifierSyntax(check bool) (m *MySqlIdentifierSyntax) {
	if check {
		this.acceptAnyToken(Tokens.IDENTIFIER)
	}
	m = NewMySqlIdentifierSyntax()
	this.setBeginPosDefault(m)
	m.Name = this.tokenVal()
	m.Qualifier = this.hasQualifier()
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseSingleOperandIExprSyntax(l ExprSyntaxLevel) I_ExprSyntax {
	e := this.parseAnyOperandIExprSyntax(l, 0)
	this.acceptExpectedOperandCount(1, e)
	return e
}

// parseAnyOperandIExprSyntax 解析任意操作数的表达式
// @param operandCount 若表达式为一个列表时，每个元素的操作数
func (this *mySqlParser) parseAnyOperandIExprSyntax(l ExprSyntaxLevel, listElementOperandCount int) (e I_ExprSyntax) {
	if Tokens.Is(this.token(), Tokens.L_PAREN) {
		beginPos := this.tokenBeginPos()
		this.nextToken()
		e = this.parseAnyOperandIExprSyntax(ExprSyntaxLevels.EXPR, 0)
		switch this.token() {
		case Tokens.R_PAREN:
			this.nextToken()
			this.setBeginPos(e, beginPos)
			this.setEndPosDefault(e)
			this.parenthesizingSyntax(e)
		case Tokens.COMMA:
			this.nextToken()
			el := NewExprListSyntax()
			this.setBeginPos(el, beginPos)
			this.acceptExpectedOperandCount(listElementOperandCount, e)
			el.Add(e)
			for {
				e = this.parseAnyOperandIExprSyntax(ExprSyntaxLevels.EXPR, 0)
				this.acceptExpectedOperandCount(listElementOperandCount, e)
				el.Add(e)
				if Tokens.Not(this.token(), Tokens.COMMA) {
					break
				}
				this.nextToken()
			}
			this.acceptAnyToken(Tokens.R_PAREN)
			this.nextToken()
			this.setEndPosDefault(el)
			this.parenthesizingSyntax(el)
			e = el
		default:
			this.panicByUnexpectedToken()
		}

		if ExprSyntaxLevels.Not(l, ExprSyntaxLevels.SINGLE) {
			e = this.parseOperandSyntaxRest(l, e)
		}
	} else {
		e = this.parseIExprSyntax(l)
	}
	return
}

func (this *mySqlParser) parseIExprSyntax(l ExprSyntaxLevel) (e I_ExprSyntax) {
	switch this.token() {
	case Tokens.IDENTIFIER:
		i := this.parseMySqlIdentifierSyntax(false)
		switch this.token() {
		case Tokens.DOT:
			e = this.parsePropertySyntax(i)
		case Tokens.L_PAREN:
			e = this.parseIFunctionSyntax(i)
		case Tokens.STRING:
			if !i.Qualifier {
				if strings.EqualFold("x", i.Name) {
					e = this.parseMySqlHexagonalLiteralSyntax()
				} else if strings.EqualFold("b", i.Name) {
					e = this.parseMySqlBinaryLiteralSyntax()
				} else if strings.EqualFold("N", i.Name) {
					e = this.parseNStringSyntax()
				} else if i.Name[0] == '_' {
					e = this.parseMySqlTranscodingStringSyntax(i)
				} else if strings.EqualFold("DATE", i.Name) || strings.EqualFold("TIME", i.Name) || strings.EqualFold("TIMESTAMP", i.Name) {
					e = this.parseMySqlDateAndTimeLiteralSyntax(i)
				} else {
					e = i
				}
			}
		default:
			switch strings.ToUpper(i.Name) {
			case "CURRENT_DATE", "CURRENT_TIME", "CURRENT_TIMESTAMP", "LOCALTIME", "LOCALTIMESTAMP":
				e = this.parseIFunctionSyntax(i)
			default:
				e = i
			}
		}
	case Tokens.STRING:
		e = this.parseMySqlStringSyntax(false)
	case Tokens.DECIMAL_NUMBER:
		e = this.parseDecimalNumberSyntax()
	case Tokens.SELECT:
		if Tokens.Not(this.prevToken(), Tokens.L_PAREN) {
			this.panicByToken("subquery expression must be parenthesized")
		}
		e = this.parseIQuerySyntax(QuerySyntaxLevels.NORMAL)
	case Tokens.NULL:
		e = this.parseNullSyntax()
	case Tokens.CASE:
		e = this.parseCaseSyntax()
	case Tokens.EXISTS:
		e = this.parseExistsSyntax()
	case Tokens.BINARY, Tokens.SUB, Tokens.BANG, Tokens.TILDE, Tokens.PLUS, Tokens.NOT:
		e = this.parseMySqlUnarySyntax()
	case Tokens.INTERVAL:
		i := this.parseMySqlIdentifierSyntax(false)
		if Tokens.Is(this.token(), Tokens.L_PAREN) {
			o := this.parseAnyOperandIExprSyntax(ExprSyntaxLevels.EXPR, 1)
			if o.IsExprList() {
				// https://dev.mysql.com/doc/refman/8.0/en/comparison-operators.html#function_interval
				pf := NewFunctionSyntax()
				this.setBeginPos(pf, i.BeginPos)
				pf.Name = i.Name
				pf.Parameters = o.(*ExprListSyntax)
				this.setEndPosDefault(pf)
				e = pf
			} else {
				t := this.parseMySqlTemporalInterval()
				// 如果没有间隔单位，则完全无法确定究竟是Interval表达式，还是Interval函数（Interval函数需要两个以上参数）
				if t.Undefined() {
					this.panicBySyntax(o, "syntax error")
				}
				inr := NewMySqlIntervalSyntax()
				this.setBeginPos(inr, this.prevTokenBeginPos())
				inr.Expr = o
				inr.Unit = t
				this.setEndPosDefault(inr)
				e = inr
			}
		} else {
			// https://dev.mysql.com/doc/refman/8.0/en/expressions.html#temporal-intervals
			e = this.parseMySqlIntervalSyntax()
		}
	case Tokens.TRUE:
		e = this.parseMySqlTrueSyntax()
	case Tokens.FALSE:
		e = this.parseMySqlFalseSyntax()
	case Tokens.QUES:
		e = this.parseParameterSyntax()
	case Tokens.HEXADECIMAL_NUMBER:
		e = this.parseHexadecimalNumberSyntax()
	case Tokens.BINARY_NUMBER:
		e = this.parseBinaryNumberSyntax()
	case Tokens.VALUES, Tokens.CHAR:
		i := this.parseMySqlIdentifierSyntax(false)
		this.acceptAnyToken(Tokens.L_PAREN)
		e = this.parseIFunctionSyntax(i)
	case Tokens.AT, Tokens.AT_AT:
		e = this.parseMySqlVariableSyntax()
	default:
		this.panicByUnexpectedToken()
	}
	if ExprSyntaxLevels.Not(l, ExprSyntaxLevels.SINGLE) {
		e = this.parseOperandSyntaxRest(l, e)
	}
	return
}

func (this *mySqlParser) parseOperandSyntaxRest(l ExprSyntaxLevel, before I_ExprSyntax) (last I_ExprSyntax) {
	last = before
	for {
		bo := this.parseMySqlBinaryOperator(l, last)
		if bo.Undefined() {
			break
		}
		last = this.parseMySqlBinaryOperationSyntax(l, last, bo)
	}
	return
}

func (this *mySqlParser) parseIAliasSyntax(l AliasSyntaxLevel) (a I_AliasSyntax) {
	switch this.token() {
	case Tokens.AS:
		this.nextToken()
		if a = this.parseIAliasSyntax(l); a == nil {
			this.panicByUnexpectedToken()
		}
	case Tokens.STRING:
		if AliasSyntaxLevels.Is(l, AliasSyntaxLevels.STRING) {
			a = this.parseMySqlStringSyntax(false)
		}
	case Tokens.IDENTIFIER:
		a = this.parseMySqlIdentifierSyntax(false)
	}
	return
}

func (this *mySqlParser) parseDecimalNumberSyntax() (d *DecimalNumberSyntax) {
	d = NewDecimalNumberSyntax()
	this.setBeginPosDefault(d)
	d.Sql = this.tokenVal()
	this.nextToken()
	this.setEndPosDefault(d)
	return
}

func (this *mySqlParser) parseISelectItemSyntax(allowAllColumnSyntax bool) (s I_SelectItemSyntax) {
	if Tokens.Is(this.token(), Tokens.STAR) {
		if !allowAllColumnSyntax {
			this.panicByUnexpectedToken()
		}
		s = this.parseAllColumnSyntax()
	} else {
		s = this.parseSelectColumnSyntax()
	}
	return
}

func (this *mySqlParser) parseJoinOnSyntax() (j *JoinOnSyntax) {
	j = NewJoinOnSyntax()
	this.setBeginPosDefault(j)
	j.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	this.setEndPosDefault(j)
	return
}

func (this *mySqlParser) parseJoinUsingSyntax() (j *JoinUsingSyntax) {
	j = NewJoinUsingSyntax()
	this.setBeginPosDefault(j)
	l := this.parseMySqlIdentifierListSyntax()
	j.ColumnList = l
	this.setEndPosDefault(j)
	return
}

func (this *mySqlParser) parseIColumnItemSyntax() (ic I_ColumnItemSyntax) {
	i := this.parseMySqlIdentifierSyntax(true)
	if Tokens.Is(this.token(), Tokens.DOT) {
		ic = this.parsePropertySyntax(i)
	} else {
		ic = i
	}
	return
}

func (this *mySqlParser) parsePropertySyntax(i *MySqlIdentifierSyntax) (pt *PropertySyntax) {
	pt = NewPropertySyntax()
	this.setBeginPos(pt, i.BeginPos)
	pt.Owner = i
	pt.Value = this.parseIPropertyValueSyntax()
	this.setEndPosDefault(pt)
	return
}

func (this *mySqlParser) parseIPropertyValueSyntax() (pv I_PropertyValueSyntax) {
	switch this.nextToken() {
	case Tokens.STAR:
		pv = this.parseAllColumnSyntax()
	case Tokens.IDENTIFIER:
		pv = this.parseMySqlIdentifierSyntax(false)
	default:
		this.panicByUnexpectedToken()
	}
	return
}

func (this *mySqlParser) parseAllColumnSyntax() (a *AllColumnSyntax) {
	a = NewAllColumnSyntax()
	this.setBeginPosDefault(a)
	this.nextToken()
	this.setEndPosDefault(a)
	return
}

func (this *mySqlParser) parseSelectColumnSyntax() (s *SelectColumnSyntax) {
	s = NewSelectColumnSyntax()
	this.setBeginPosDefault(s)
	s.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	s.Alias = this.parseIAliasSyntax(AliasSyntaxLevels.STRING)
	this.setEndPosDefault(s)
	return
}

func (this *mySqlParser) parseIFunctionSyntax(functionName *MySqlIdentifierSyntax) (f I_FunctionSyntax) {
	upperFunctionName := strings.ToUpper(functionName.Name)
	switch upperFunctionName {
	case "CONVERT":
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		c := NewMySqlConvertFunctionSyntax()
		this.setBeginPos(c, functionName.BeginPos)
		c.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
		switch this.token() {
		case Tokens.USING:
			c.UsingTranscoding = true
			this.nextToken()
			c.TranscodingName = this.tokenVal()
			this.nextToken()
		case Tokens.COMMA:
			this.nextToken()
			c.DataType = this.parseMySqlCastDataTypeSyntax()
		default:
			this.panicByUnexpectedToken()
		}
		this.acceptAnyToken(Tokens.R_PAREN)
		if Tokens.Is(this.nextToken(), Tokens.COLLATE) {
			this.nextToken()
			c.Collate = this.parseMySqlIdentifierSyntax(true).Name
		}
		this.setEndPosDefault(c)
		f = c
	case "CAST":
		// https://dev.mysql.com/doc/refman/8.0/en/cast-functions.html#function_cast
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		c := NewMySqlCastFunctionSyntax()
		this.setBeginPos(c, functionName.BeginPos)
		c.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
		if this.tokenValUpper() == "AT" {
			this.nextToken()
			this.acceptAnyTokenVal("TIME")
			this.nextToken()
			this.acceptAnyTokenVal("ZONE")
			hasInterval := false
			if Tokens.Is(this.nextToken(), Tokens.INTERVAL) {
				hasInterval = true
				this.nextToken()
			}
			s := this.parseMySqlStringSyntax(true)
			if !(s.Value() == "+00:00" || (s.Value() == "UTC" && !hasInterval)) {
				this.panicBySyntax(s, "unknown or incorrect time zone: %s", s.Sql())
			}
			c.AtTimeZone = s
		}
		this.acceptAnyToken(Tokens.AS)
		this.nextToken()
		c.DataType = this.parseMySqlCastDataTypeSyntax()
		this.acceptAnyToken(Tokens.R_PAREN)
		if Tokens.Is(this.nextToken(), Tokens.COLLATE) {
			this.nextToken()
			c.Collate = this.parseMySqlIdentifierSyntax(true).Name
		}
		this.setEndPosDefault(c)
		f = c
	case "EXTRACT":
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		e := NewMySqlExtractFunctionSyntax()
		this.setBeginPos(e, functionName.BeginPos)
		t := this.parseMySqlTemporalInterval()
		if t.Undefined() {
			this.panicByUnexpectedToken()
		}
		e.Unit = t
		this.acceptAnyToken(Tokens.FROM)
		this.nextToken()
		e.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(e)
		f = e
	case "TIMESTAMPADD", "TIMESTAMPDIFF":
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		t := NewMySqlTimestampFunctionSyntax()
		this.setBeginPos(t, functionName.BeginPos)
		t.Name = upperFunctionName
		ti := this.parseMySqlTemporalInterval()
		if ti.Undefined() {
			this.panicByUnexpectedToken()
		}
		t.Unit = ti
		for i := 0; i < 2; i++ {
			this.acceptAnyToken(Tokens.COMMA)
			this.nextToken()
			t.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
		}
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(t)
		f = t
	case "TRIM":
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		t := NewMySqlTrimFunctionSyntax()
		this.setBeginPos(t, functionName.BeginPos)
		if Tokens.Is(this.token(), Tokens.BOTH) {
			t.TrimMode = MySqlTrimModes.BOTH
			this.nextToken()
		} else if Tokens.Is(this.token(), Tokens.LEADING) {
			t.TrimMode = MySqlTrimModes.LEADING
			this.nextToken()
		} else if this.tokenValUpper() == "TRAILING" {
			t.TrimMode = MySqlTrimModes.TRAILING
			this.nextToken()
		}

		tmpExpr := this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
		if Tokens.Is(this.token(), Tokens.FROM) {
			t.RemStr = tmpExpr
			this.nextToken()
			t.Str = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
		} else {
			t.Str = tmpExpr
		}

		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(t)
		f = t
	case "CHAR":
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		c := NewMySqlCharFunctionSyntax()
		this.setBeginPos(c, functionName.BeginPos)
		for {
			c.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
			if Tokens.Not(this.token(), Tokens.COMMA) {
				break
			}
			this.nextToken()
		}
		if Tokens.Is(this.token(), Tokens.USING) {
			this.nextToken()
			c.CharsetName = this.tokenVal()
			this.nextToken()
		}
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(c)
		f = c
	case "AVG", "BIT_AND", "BIT_OR", "BIT_XOR", "COUNT", "JSON_ARRAYAGG", "JSON_OBJECTAGG", "MAX", "MIN",
		"STD", "STDDEV", "STDDEV_POP", "STDDEV_SAMP", "SUM", "VAR_POP", "VAR_SAMP", "VARIANCE":
		// 通用的聚合函数解析，统一解析为：function_name([DISTINCT ](*|[<expr>[, <expr>]...]))[ <over_syntax>]
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		a := NewAggregateFunctionSyntax()
		this.setBeginPos(a, functionName.BeginPos)
		a.Name = upperFunctionName
		if Tokens.Not(this.token(), Tokens.R_PAREN) {
			a.AggregateOption = this.parseAggregateOption()
			if Tokens.Is(this.token(), Tokens.STAR) {
				a.AllColumnParameter = true
				this.nextToken()
			} else {
				for {
					a.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
					if Tokens.Not(this.token(), Tokens.COMMA) {
						break
					}
					this.nextToken()
				}
			}
		}
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		a.Over = this.parseOverSyntax()
		this.setEndPosDefault(a)
		f = a
	case "GROUP_CONCAT":
		// GROUP_CONCAT聚合函数特殊处理
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		g := NewMySqlGroupConcatFunctionSyntax()
		this.setBeginPos(g, functionName.BeginPos)
		g.AggregateOption = this.parseAggregateOption()
		for {
			g.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR))
			if Tokens.Not(this.token(), Tokens.COMMA) {
				break
			}
		}
		g.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevels.NORNAL)
		if Tokens.Is(this.token(), Tokens.SEPARATOR) {
			this.nextToken()
			g.Separator = this.parseMySqlStringSyntax(true)
		}
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(g)
		f = g
	case "GET_FORMAT":
		this.acceptAnyToken(Tokens.L_PAREN)
		this.nextToken()
		g := NewMySqlGetFormatFunctionSyntax()
		this.setBeginPos(g, functionName.BeginPos)
		d := MySqlGetFormatTypes.OfSql(strings.ToUpper(this.tokenVal()))
		if d.Undefined() {
			this.panicByUnexpectedToken()
		}
		g.Type = d
		this.nextToken()
		this.acceptAnyToken(Tokens.COMMA)
		this.nextToken()
		g.DateFormat = this.parseMySqlStringSyntax(true)
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(g)
		f = g
	case "CUME_DIST", "DENSE_RANK", "FIRST_VALUE", "LAG", "LAST_VALUE", "LEAD", "NTH_VALUE", "NTILE", "PERCENT_RANK", "RANK", "ROW_NUMBER":
		wf := NewWindowFunctionSyntax()
		this.setBeginPos(wf, functionName.BeginPos)
		wf.Name = upperFunctionName
		wf.Parameters = this.parseSingleOperandExprListSyntax()
		if Tokens.Is(this.token(), Tokens.IGNORE) {
			this.nextToken()
			this.acceptAnyToken(Tokens.NULL)
			wf.NullTreatment = NullTreatments.IGNORE_NULLS
			this.nextToken()
		}
		if this.tokenValUpper() == "RESPECT" {
			this.nextToken()
			this.acceptAnyToken(Tokens.NULL)
			wf.NullTreatment = NullTreatments.RESPECT_NULLS
			this.nextToken()
		}
		wf.Over = this.parseOverSyntax()
		this.setEndPosDefault(wf)
		f = wf
	case "CURRENT_DATE", "CURRENT_TIME", "CURRENT_TIMESTAMP", "LOCALTIME", "LOCALTIMESTAMP":
		pf := NewFunctionSyntax()
		this.setBeginPos(pf, functionName.BeginPos)
		pf.Name = upperFunctionName
		this.setEndPosDefault(pf)
		f = pf
	default:
		pf := NewFunctionSyntax()
		this.setBeginPos(pf, functionName.BeginPos)
		pf.Name = functionName.Name
		pf.Parameters = this.parseSingleOperandExprListSyntax()
		this.setEndPosDefault(pf)
		f = pf
	}
	return
}

func (this *mySqlParser) parseMySqlHexagonalLiteralSyntax() (m *MySqlHexagonalLiteralSyntax) {
	m = NewMySqlHexagonalLiteralSyntax()
	this.setBeginPosDefault(m)
	m.SetSql("x" + this.tokenVal())
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlBinaryLiteralSyntax() (m *MySqlBinaryLiteralSyntax) {
	m = NewMySqlBinaryLiteralSyntax()
	this.setBeginPosDefault(m)
	m.SetSql("b" + this.tokenVal())
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseNStringSyntax() (n *NStringSyntax) {
	n = NewNStringSyntax()
	this.setBeginPosDefault(n)
	n.Str = this.parseMySqlStringSyntax(true)
	this.setEndPosDefault(n)
	return
}

func (this *mySqlParser) parseMySqlTranscodingStringSyntax(i *MySqlIdentifierSyntax) (t *MySqlTranscodingStringSyntax) {
	t = NewMySqlTranscodingStringSyntax()
	this.setBeginPos(t, i.BeginPos)
	t.CharsetName = i.Name[1:]
	t.Str = this.parseMySqlStringSyntax(true)
	if Tokens.Is(this.token(), Tokens.COLLATE) {
		this.nextToken()
		t.Collate = this.parseMySqlIdentifierSyntax(true).Name
	}
	this.setEndPosDefault(t)
	return
}

func (this *mySqlParser) parseMySqlDateAndTimeLiteralSyntax(i *MySqlIdentifierSyntax) (d *MySqlDateAndTimeLiteralSyntax) {
	d = NewMySqlDateAndTimeLiteralSyntax()
	this.setBeginPos(d, i.BeginPos)
	d.Type = MySqlDatetimeLiteralTypes.OfSql(strings.ToUpper(i.Name))
	d.DateAndTime = this.parseMySqlStringSyntax(true)
	this.setEndPosDefault(d)
	return
}

func (this *mySqlParser) parseMySqlStringSyntax(check bool) (m *MySqlStringSyntax) {
	if check {
		this.acceptAnyToken(Tokens.STRING)
	}
	m = NewMySqlStringSyntax()
	this.setBeginPosDefault(m)
	m.SetSql(this.tokenVal())
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseNullSyntax() (n *NullSyntax) {
	n = NewNullSyntax()
	this.setBeginPosDefault(n)
	this.nextToken()
	this.setEndPosDefault(n)
	return
}

func (this *mySqlParser) parseCaseSyntax() (c *CaseSyntax) {
	c = NewCaseSyntax()
	this.setBeginPosDefault(c)
	if Tokens.Not(this.nextToken(), Tokens.WHEN) {
		c.ValueExpr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	}
	this.acceptAnyToken(Tokens.WHEN)
	c.WhenItemList = this.parseCaseWhenItemListSyntax()
	if Tokens.Is(this.token(), Tokens.ELSE) {
		this.nextToken()
		c.ElseExr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	}
	this.acceptAnyTokenVal("END")
	this.nextToken()
	this.setEndPosDefault(c)
	return
}
func (this *mySqlParser) parseCaseWhenItemListSyntax() (cl *CaseWhenItemListSyntax) {
	cl = NewCaseWhenItemListSyntax()
	this.setBeginPosDefault(cl)
	for {
		cl.Add(this.parseCaseWhenItemSyntax())
		if Tokens.Not(this.token(), Tokens.WHEN) {
			break
		}
	}
	this.setEndPosDefault(cl)
	return
}

func (this *mySqlParser) parseCaseWhenItemSyntax() (c *CaseWhenItemSyntax) {
	c = NewCaseWhenItem()
	this.setBeginPosDefault(c)
	this.acceptAnyToken(Tokens.WHEN)
	this.nextToken()
	c.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	this.acceptAnyToken(Tokens.THEN)
	this.nextToken()
	c.Result = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	this.setEndPosDefault(c)
	return
}

func (this *mySqlParser) parseExistsSyntax() (e *ExistsSyntax) {
	e = NewExistsSyntax()
	this.setBeginPosDefault(e)
	this.nextToken()
	this.acceptAnyToken(Tokens.L_PAREN)
	e.Query = this.parseIQuerySyntax(QuerySyntaxLevels.NORMAL)
	this.setEndPosDefault(e)
	return
}

func (this *mySqlParser) parseMySqlUnarySyntax() (m *MySqlUnarySyntax) {
	m = NewMySqlUnarySyntax()
	this.setBeginPosDefault(m)
	uo := this.parseMySqlUnaryOperator()
	m.UnaryOperator = uo
	m.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.SINGLE)
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlIntervalSyntax() (m *MySqlIntervalSyntax) {
	m = NewMySqlIntervalSyntax()
	this.setBeginPos(m, this.prevTokenBeginPos())
	m.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	t := this.parseMySqlTemporalInterval()
	if t.Undefined() {
		this.panicByUnexpectedToken()
	}
	m.Unit = t
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlTrueSyntax() (m *MySqlTrueSyntax) {
	m = NewMySqlTrueSyntax()
	this.setBeginPosDefault(m)
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlFalseSyntax() (m *MySqlFalseSyntax) {
	m = NewMySqlFalseSyntax()
	this.setBeginPosDefault(m)
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseParameterSyntax() (pa *ParameterSyntax) {
	pa = NewParameterSyntax()
	this.setBeginPosDefault(pa)
	pa.Index = this.nextParameterIndex()
	this.nextToken()
	this.setEndPosDefault(pa)
	return
}

func (this *mySqlParser) parseHexadecimalNumberSyntax() (h *HexadecimalNumberSyntax) {
	h = NewHexadecimalNumberSyntax()
	this.setBeginPosDefault(h)
	h.Sql = this.tokenVal()
	this.nextToken()
	this.setEndPosDefault(h)
	return
}

func (this *mySqlParser) parseBinaryNumberSyntax() (b *BinaryNumberSyntax) {
	b = NewBinaryNumberSyntax()
	this.setBeginPosDefault(b)
	b.Sql = this.tokenVal()
	this.nextToken()
	this.setEndPosDefault(b)
	return
}

func (this *mySqlParser) parseMySqlVariableSyntax() (m *MySqlVariableSyntax) {
	m = NewMySqlVariableSyntax()
	this.setBeginPosDefault(m)
	m.SetSql(this.tokenVal())
	if Tokens.Is(this.token(), Tokens.AT) {
		m.VariableType = MySqlVariableTypes.SESSION
	} else if Tokens.Is(this.token(), Tokens.AT_AT) {
		m.VariableType = MySqlVariableTypes.GLOBAL
	}
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlBinaryOperator(l ExprSyntaxLevel, leftOperand I_ExprSyntax) (bo BinaryOperator) {
	c := this.saveCursor()
	if Tokens.Is(this.token(), Tokens.IS) {
		if Tokens.Is(this.nextToken(), Tokens.NOT) {
			bo = MySqlBinaryOperators.IS_NOT
			this.nextToken()
		} else {
			bo = MySqlBinaryOperators.IS
		}
		this.acceptAnyTokenVal("NULL", "TRUE", "FALSE", "Undefined")
	} else if Tokens.Is(this.token(), Tokens.NOT) {
		switch this.nextToken() {
		case Tokens.LIKE:
			bo = MySqlBinaryOperators.NOT_LIKE
		case Tokens.BETWEEN:
			bo = MySqlBinaryOperators.NOT_BETWEEN
		case Tokens.REGEXP:
			bo = MySqlBinaryOperators.NOT_REGEXP
		case Tokens.R_LIKE:
			bo = MySqlBinaryOperators.NOT_RLIKE
		case Tokens.IN:
			bo = MySqlBinaryOperators.NOT_IN
		default:
			this.panicByUnexpectedToken()
		}
		this.nextToken()
	} else if this.tokenValUpper() == "SOUNDS" {
		if Tokens.Is(this.nextToken(), Tokens.LIKE) {
			bo = MySqlBinaryOperators.SOUNDS_LIKE
			this.nextToken()
		} else {
			this.panicByUnexpectedToken()
		}
	} else if this.tokenValUpper() == "MEMBER" {
		if Tokens.Is(this.nextToken(), Tokens.OF) {
			bo = MySqlBinaryOperators.MEMBER_OF
			this.nextToken()
		}
	} else {
		bo = MYSQL_TOKEN_TO_BINARY_OPERATORS[this.token()]
		if !bo.Undefined() {
			this.nextToken()
		}
	}
	if !MYSQL_EXPR_LEVEL_TO_BINARY_OPERATORS[l].Contains(bo) {
		bo = BinaryOperator{}
	}
	if !bo.Undefined() {
		if !bo.AllowMultipleOperand && leftOperand != nil {
			this.acceptExpectedOperandCount(1, leftOperand)
		}
	} else {
		this.rollback(c)
	}
	return
}

func (this *mySqlParser) parseMySqlBinaryOperationSyntax(l ExprSyntaxLevel, leftOperand I_ExprSyntax, bo BinaryOperator) (b *MySqlBinaryOperationSyntax) {
	b = NewMySqlBinaryOperationSyntax()
	this.setBeginPos(b, leftOperand.M_5CF6320E8474().BeginPos)
	b.LeftOperand = leftOperand
	b.BinaryOperator = bo
	if _, ok := leftOperand.(*MySqlIntervalSyntax); ok && MySqlBinaryOperators.Is(bo, MySqlBinaryOperators.SUBTRACT) {
		this.panicBySyntax(leftOperand, "for the - operator, INTERVAL expr unit is permitted only on the right side")
	}
	if MySqlBinaryOperators.Is(bo, MySqlBinaryOperators.BETWEEN, MySqlBinaryOperators.NOT_BETWEEN) {
		b.RightOperand = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.CALCULATION)
		this.acceptAnyToken(Tokens.AND)
		this.nextToken()
		b.BetweenThirdOperand = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.BOOLEAN_PREDICATE)
	} else {
		var (
			rightOperand I_ExprSyntax
			cm           MySqlComparisonMode
		)
		switch bo {
		case MySqlBinaryOperators.EQUAL_OR_ASSIGNMENT, MySqlBinaryOperators.GREATER_THAN, MySqlBinaryOperators.LESS_THAN, MySqlBinaryOperators.GREATER_THAN_OR_EQUAL, MySqlBinaryOperators.LESS_THAN_OR_EQUAL, MySqlBinaryOperators.LESS_THAN_OR_GREATER, MySqlBinaryOperators.NOT_EQUAL:
			if Tokens.Is(this.token(), Tokens.ALL) {
				cm = MySqlComparisonModes.ALL
				this.nextToken()
			} else {
				var (
					tmpCsm MySqlComparisonMode
					c      *cursor
				)
				switch this.tokenValUpper() {
				case "ANY":
					c = this.saveCursor()
					tmpCsm = MySqlComparisonModes.ANY
				case "SOME":
					c = this.saveCursor()
					tmpCsm = MySqlComparisonModes.SOME
				}
				// 若下一个记号非运算符，则为特殊语法
				if !tmpCsm.Undefined() {
					this.nextToken()
					if this.parseMySqlBinaryOperator(ExprSyntaxLevels.EXPR, nil).Undefined() {
						cm = tmpCsm
					} else {
						this.rollback(c)
					}
				}
			}
			if !cm.Undefined() {
				rightOperand = this.parseMySqlComparisonModeRightOperand()
			} else {
				rightOperand = this.parseAnyOperandIExprSyntax(ExprSyntaxLevels.SINGLE, 0)
			}
		case MySqlBinaryOperators.IN, MySqlBinaryOperators.NOT_IN:
			this.acceptAnyToken(Tokens.L_PAREN)
			c := this.saveCursor()
			if Tokens.Is(this.nextToken(), Tokens.SELECT) {
				this.rollback(c)
				rightOperand = this.parseIQuerySyntax(QuerySyntaxLevels.NORMAL)
			} else {
				this.rollback(c)
				rightOperand = this.parseAnyOperandIExprSyntax(ExprSyntaxLevels.SINGLE, 0)
			}
		case MySqlBinaryOperators.MEMBER_OF:
			this.acceptAnyToken(Tokens.L_PAREN)
			this.nextToken()
			rightOperand = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.SINGLE)
			this.acceptAnyToken(Tokens.R_PAREN)
			this.nextToken()
			this.parenthesizingSyntax(rightOperand)
		default:
			rightOperand = this.parseAnyOperandIExprSyntax(ExprSyntaxLevels.SINGLE, 0)
		}
		for {
			c := this.saveCursor()
			nextBo := this.parseMySqlBinaryOperator(l, nil)
			if bo.Compare(nextBo) < 0 {
				rightOperand = this.parseMySqlBinaryOperationSyntax(l, rightOperand, nextBo)
			} else {
				this.rollback(c)
				break
			}
		}
		this.acceptEqualOperandCount(leftOperand, rightOperand, MySqlBinaryOperators.Is(bo, MySqlBinaryOperators.IN, MySqlBinaryOperators.NOT_IN))
		b.ComparisonMode = cm
		b.RightOperand = rightOperand
		if MySqlBinaryOperators.Is(bo, MySqlBinaryOperators.LIKE, MySqlBinaryOperators.NOT_LIKE) {
			if this.equalTokenVal("ESCAPE") {
				this.nextToken()
				b.LikeEscape = this.parseMySqlStringSyntax(true)
			}
		}
	}
	this.setEndPosDefault(b)
	return
}

// parseMySqlComparisonModeRightOperand
// https://dev.mysql.com/doc/refman/8.0/en/any-in-some-subqueries.html
// https://dev.mysql.com/doc/refman/8.0/en/all-subqueries.html
func (this *mySqlParser) parseMySqlComparisonModeRightOperand() (ie I_ExprSyntax) {
	this.acceptAnyToken(Tokens.L_PAREN)
	switch this.nextToken() {
	case Tokens.SELECT:
		ie = this.parseIQuerySyntax(QuerySyntaxLevels.NORMAL)
	case Tokens.TABLE:
		ie = this.parseMySqlTablesSyntax()
	default:
		this.panicByUnexpectedToken()
	}
	this.acceptAnyToken(Tokens.R_PAREN)
	this.nextToken()
	this.parenthesizingSyntax(ie)
	return
}

func (this *mySqlParser) parseMySqlTablesSyntax() (m *MySqlTableSyntax) {
	this.nextToken()
	m = NewMySqlTableSyntax()
	this.setBeginPosDefault(m)
	m.TableNameItem = this.parseTableNameItemSyntax()
	m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevels.IDENTIFIER)
	m.Limit = this.parseMySqlLimitSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlTemporalInterval() (m MySqlTemporalInterval) {
	m = MySqlTemporalIntervals.OfSql(this.tokenVal())
	if !m.Undefined() {
		this.nextToken()
	}
	return
}

func (this *mySqlParser) parseMySqlUnaryOperator() (m UnaryOperator) {
	m = MYSQL_TOKEN_TO_UNARY_OPERATORS[this.token()]
	this.nextToken()
	return
}

func (this *mySqlParser) parseMySqlCastDataTypeSyntax() (m *MySqlCastDataTypeSyntax) {
	m = NewMySqlCastDataTypeSyntax()
	this.setBeginPosDefault(m)
	// 通用的处理方式，不对具体类型做检查。https://dev.mysql.com/doc/refman/8.1/en/cast-functions.html#function_cast
	if Tokens.Is(this.token(), Tokens.IDENTIFIER) && !this.hasQualifier() || this.reserved() {
		m.Name = this.tokenValUpper()
		if Tokens.Is(this.nextToken(), Tokens.L_PAREN) {
			m.Parameters = this.parseMySqlCastDataTypeParamListSyntax()
		}
		if Tokens.Is(this.token(), Tokens.CHARACTER) {
			this.nextToken()
			this.acceptAnyToken(Tokens.SET)
			this.nextToken()
			m.CharsetName = this.tokenVal()
			this.nextToken()
		}
	} else {
		this.panicByUnexpectedToken()
	}
	this.setBeginPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlCastDataTypeParamListSyntax() (m *MySqlCastDataTypeParamListSyntax) {
	m = NewMySqlCastDataTypeParamListSyntax()
	this.setBeginPosDefault(m)
	this.acceptAnyToken(Tokens.L_PAREN)
	this.nextToken()
	for {
		this.acceptAnyToken(Tokens.DECIMAL_NUMBER)
		m.Add(this.parseDecimalNumberSyntax())
		if Tokens.Not(this.token(), Tokens.COMMA) {
			break
		}
		this.nextToken()
	}
	this.acceptAnyToken(Tokens.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(m)
	this.parenthesizingSyntax(m)
	return
}

func (this *mySqlParser) parseOverSyntax() (o *OverSyntax) {
	if Tokens.Not(this.token(), Tokens.OVER) {
		return
	}
	o = NewOverSyntax()
	this.setBeginPosDefault(o)
	switch this.nextToken() {
	case Tokens.IDENTIFIER:
		o.Window = this.parseMySqlIdentifierSyntax(false)
	case Tokens.L_PAREN:
		this.nextToken()
		o.Window = this.parseWindowSpecSyntax()
		this.acceptAnyToken(Tokens.R_PAREN)
		this.nextToken()
	default:
		this.panicByUnexpectedToken()
	}
	this.setEndPosDefault(o)
	return
}

func (this *mySqlParser) parseWindowSpecSyntax() (w *WindowSpecSyntax) {
	w = NewWindowSpecSyntax()
	this.setBeginPosDefault(w)
	if Tokens.Is(this.token(), Tokens.IDENTIFIER) {
		w.Name = this.parseMySqlIdentifierSyntax(false)
	}

	w.PartitionBy = this.parsePartitionBySyntax()
	w.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevels.NORNAL)
	w.Frame = this.parseWindowFrameSyntax()
	this.setEndPosDefault(w)
	return
}

func (this *mySqlParser) parsePartitionBySyntax() (py *PartitionBySyntax) {
	if Tokens.Not(this.token(), Tokens.PARTITION) {
		return
	}
	py = NewPartitionBySyntax()
	this.setBeginPosDefault(py)
	this.nextToken()
	this.acceptAnyToken(Tokens.BY)
	this.nextToken()
	py.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
	this.setEndPosDefault(py)
	return
}

func (this *mySqlParser) parseWindowFrameSyntax() (w *WindowFrameSyntax) {
	var unit WindowFrameUnit
	switch this.token() {
	case Tokens.ROWS:
		unit = WindowFrameUnits.ROWS
	case Tokens.RANGE:
		unit = WindowFrameUnits.RANGE
	}
	if !unit.Undefined() {
		w = NewWindowFrameSyntax()
		this.setBeginPosDefault(w)
		w.Unit = unit
		if Tokens.Is(this.nextToken(), Tokens.BETWEEN) {
			this.nextToken()
			w.Extent = this.parseWindowFrameBetweenSyntax()
		} else {
			w.Extent = this.parseIWindowFrameStartSyntax()
		}
		this.setEndPosDefault(w)
	}
	return
}

func (this *mySqlParser) parseWindowFrameBetweenSyntax() (w *WindowFrameBetweenSyntax) {
	w = NewWindowFrameBetweenSyntax()
	this.setBeginPosDefault(w)
	w.Start = this.parseIWindowFrameStartSyntax()
	this.acceptAnyToken(Tokens.AND)
	this.nextToken()
	w.End = this.parseIWindowFrameStartSyntax()
	this.setEndPosDefault(w)
	return
}

func (this *mySqlParser) parseIWindowFrameStartSyntax() (w I_WindowFrameStartEndSyntax) {
	switch this.tokenValUpper() {
	case "CURRENT":
		wf := NewWindowFrameCurrentRowSyntax()
		this.setBeginPosDefault(wf)
		this.nextToken()
		this.acceptAnyToken(Tokens.ROW)
		this.nextToken()
		this.setEndPosDefault(wf)
		w = wf
	case "UNBOUNDED":
		wf := NewWindowFrameUnboundedSyntax()
		this.setBeginPosDefault(wf)
		this.nextToken()
		this.acceptAnyTokenVal("PRECEDING", "FOLLOWING")
		wf.Type = WindowFrameStartEndTypes.OfSql(this.tokenVal())
		this.nextToken()
		this.setEndPosDefault(wf)
		w = wf
	default:
		wf := NewWindowFrameExprSyntax()
		this.setBeginPosDefault(wf)
		wf.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevels.EXPR)
		this.acceptAnyTokenVal("PRECEDING", "FOLLOWING")
		wf.Type = WindowFrameStartEndTypes.OfSql(this.tokenVal())
		this.nextToken()
		this.setEndPosDefault(wf)
		w = wf
	}
	return
}

func (this *mySqlParser) parseNamedWindowListSyntax() (nwl *NamedWindowsListSyntax) {
	if Tokens.Not(this.token(), Tokens.WINDOW) {
		return
	}
	nwl = NewNamedWindowsListSyntax()
	this.setBeginPosDefault(nwl)
	this.nextToken()
	for {
		nwl.Add(this.parseNamedWindowsSyntax())
		if Tokens.Not(this.token(), Tokens.COMMA) {
			break
		}
		this.nextToken()
	}
	this.setEndPosDefault(nwl)
	return
}

func (this *mySqlParser) parseNamedWindowsSyntax() (n *NamedWindowsSyntax) {
	n = NewNamedWindowsSyntax()
	this.setBeginPosDefault(n)
	n.Name = this.parseMySqlIdentifierSyntax(true)
	this.acceptAnyToken(Tokens.AS)
	this.nextToken()
	this.acceptAnyToken(Tokens.L_PAREN)
	this.nextToken()
	n.WindowSpec = this.parseWindowSpecSyntax()
	this.acceptAnyToken(Tokens.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(n)
	return
}

func newMySqlParser(sql string) *mySqlParser {
	p := &mySqlParser{}
	p.m_Parser = extendParser(p, newMySqlLexer(sql))
	return p
}
