package parser

import (
	"strconv"
	"strings"

	o "github.com/jishaocong0910/go-object-util"
	. "github.com/jishaocong0910/go-sql-parser/ast"
	. "github.com/jishaocong0910/go-sql-parser/enum"
)

type mySqlParser struct {
	*parser__
}

// 比较操作符o1与o2优先级。
//
//	@return success true则比较成功，false则失败（存在未定义的操作符）
//	@return precedence 优先级，o1>o2则大于0，o1<o2则小于0，o1=o2则等于0
func (this *mySqlParser) compareBinaryOperator(o1, o2 MySqlBinaryOperator) (bool, int) {
	if o1.Undefined() || o2.Undefined() {
		return false, 0
	}
	return true, o2.Precedence - o1.Precedence
}

func (this *mySqlParser) hasQualifier() bool {
	return this.lexer.(*mySqlLexer).hasQualifier()
}

func (this *mySqlParser) parseStatementSyntax_() (s_ StatementSyntax_) {
	s_ = this.parseStatementSyntax_Inner()
	if q, ok := s_.(QuerySyntax_); ok {
		s_ = this.parseQuerySyntax_Rest(q)
	}
	for {
		if Token_.Not(this.token(), Token_.SEMI) {
			break
		}
		this.nextToken()
	}
	if Token_.Not(this.token(), Token_.EOI) {
		this.panicByUnexpectedToken()
	}
	s := s_.StatementSyntax_()
	s.Sql = this.sql()
	return
}

func (this *mySqlParser) parseStatementSyntax_Inner() (s_ StatementSyntax_) {
	switch this.token() {
	case Token_.SELECT:
		s_ = this.parseQuerySyntax_(QuerySyntaxLevel_.NORMAL)
	case Token_.UPDATE:
		s_ = this.parseMySqlUpdateSyntax()
	case Token_.INSERT:
		s_ = this.parseMySqlInsertSyntax()
	case Token_.DELETE:
		s_ = this.parseMySqlDeleteSyntax()
	case Token_.L_PAREN:
		beginPos := this.tokenBeginPos()
		this.nextToken()
		s_ = this.parseStatementSyntax_Inner()
		this.setBeginPos(s_, beginPos)
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(s_)
		this.parenthesizingSyntax(s_)
	default:
		this.panicByUnexpectedToken()
	}
	return
}

func (this *mySqlParser) parseQuerySyntax_(l QuerySyntaxLevel) (q_ QuerySyntax_) {
	switch this.token() {
	case Token_.SELECT:
		q_ = this.parseMySqlSelectSyntax(l)
	case Token_.L_PAREN:
		beginPos := this.tokenBeginPos()
		this.nextToken()
		q_ = this.parseQuerySyntax_(l)
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setBeginPos(q_, beginPos)
		this.setEndPosDefault(q_)
		this.parenthesizingSyntax(q_)
	default:
		this.panicByUnexpectedToken()
	}
	if QuerySyntaxLevel_.Not(l, QuerySyntaxLevel_.QUERY_OPERAND) {
		q_ = this.parseQuerySyntax_Rest(q_)
	}
	return
}

func (this *mySqlParser) parseQuerySyntax_Rest(before QuerySyntax_) (last QuerySyntax_) {
	last = before
	for {
		var mo MultisetOperator
		switch this.token() {
		case Token_.UNION:
			mo = MultisetOperator_.UNION
		case Token_.EXCEPT:
			mo = MultisetOperator_.EXCEPT
		case Token_.INTERSECT:
			mo = MultisetOperator_.INTERSECT
		}
		if mo.Undefined() {
			break
		}

		this.nextToken()
		if s, ok := last.(*MySqlSelectSyntax); ok {
			if ParenthesizeType_.Not(s.ParenthesizeType, ParenthesizeType_.TRUE) {
				if s.OrderBy != nil {
					this.panicBySyntax(s, "incorrect usage of ORDER BY")
				}
				if s.LimitSyntax != nil {
					this.panicBySyntax(s, "incorrect usage of LIMIT")
				}
			}
		}

		u := NewMySqlMultisetSyntax()
		this.setBeginPos(u, last.Syntax_().BeginPos)
		u.LeftQuery = last
		u.MultisetOperator = mo
		u.AggregateOption = this.parseAggregateOption()

		var nextLevel QuerySyntaxLevel
		if Token_.Is(this.token(), Token_.SELECT) {
			nextLevel = QuerySyntaxLevel_.QUERY_OPERAND
		} else {
			nextLevel = QuerySyntaxLevel_.NORMAL
		}

		u.RightQuery = this.parseQuerySyntax_(nextLevel)
		this.setEndPosDefault(u)
		this.acceptEqualOperandCount(u.LeftQuery, u.RightQuery, false)
		last = u
	}

	if u, ok := last.(*MySqlMultisetSyntax); ok {
		u.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevel_.IDENTIFIER)
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
		case Token_.IDENTIFIER:
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
		case Token_.DISTINCT, Token_.DISTINCTROW:
			m.AggregateOption = AggregateOption_.DISTINCT
			isOption = true
		case Token_.ALL:
			m.AggregateOption = AggregateOption_.ALL
			isOption = true
		case Token_.HIGHP_RIORITY:
			m.HighPriority = true
			isOption = true
		case Token_.STRAIGHT_JOIN:
			m.StraightJoin = true
			isOption = true
		case Token_.SQL_SMALL_RESULT:
			m.SqlSmallResult = true
			isOption = true
		case Token_.SQL_BIG_RESULT:
			m.SqlBigResult = true
			isOption = true
		case Token_.SQL_CALC_FOUND_ROWS:
			m.SqlCalcFoundRows = true
			isOption = true
		}
		if !isOption {
			break
		}
	}

	m.SelectItemList = this.parseSelectItemListSyntax()

	if Token_.Is(this.token(), Token_.FROM) {
		this.nextToken()
		m.TableReference = this.parseTableReferenceSyntax_(TableReferenceSyntaxLevel_.JOIN)
		m.Where = this.parseWhereSyntax()
		if QuerySyntaxLevel_.Not(l, QuerySyntaxLevel_.QUERY_OPERAND) {
			if g := this.parseMySqlGroupBySyntax(); g != nil {
				m.GroupBy = g
			}
		}
		m.Having = this.parseHavingSyntax()
		m.NamedWindowList = this.parseNamedWindowListSyntax()
		if QuerySyntaxLevel_.Not(l, QuerySyntaxLevel_.QUERY_OPERAND) {
			m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevel_.NORNAL)
		}
		if QuerySyntaxLevel_.Not(l, QuerySyntaxLevel_.QUERY_OPERAND) {
			m.LimitSyntax = this.parseMySqlLimitSyntax()
		}

		lr := this.parseMySqlLockReadSyntax(false)
		if lr != nil {
			var lrs []*MySqlLockingReadSyntax
			lrs = append(lrs, lr)
			if lr.OfTableName != nil {
				for {
					if Token_.Not(this.token(), Token_.FOR) {
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
	case Token_.FOR:
		m = NewMySqlLockingReadSyntax()
		this.setBeginPosDefault(m)
		if Token_.Is(this.nextToken(), Token_.UPDATE) {
			m.LockingRead = MySqlLockingRead_.FOR_UPDATE
		} else if this.tokenValUpper() == "SHARE" {
			m.LockingRead = MySqlLockingRead_.FOR_SHARE
		} else {
			this.panicByUnexpectedToken()
		}
		this.nextToken()

		if mustOfTable {
			this.acceptAnyToken(Token_.OF)
		}
		if Token_.Is(this.token(), Token_.OF) {
			this.nextToken()
			m.OfTableName = this.parseMySqlIdentifierSyntax(true)
		}

		switch this.tokenValUpper() {
		case "NOWAIT":
			m.LockingReadConcurrency = MySqlLockingReadConcurrency_.NO_WAIT
			this.nextToken()
		case "SKIP":
			this.nextToken()
			this.acceptAnyTokenVal("LOCKED")
			m.LockingReadConcurrency = MySqlLockingReadConcurrency_.SKIP_LOCKED
			this.nextToken()
		}
		this.setEndPosDefault(m)
	case Token_.LOCK:
		m = NewMySqlLockingReadSyntax()
		this.setBeginPosDefault(m)
		this.nextToken()
		this.acceptAnyToken(Token_.IN)
		this.nextToken()
		this.acceptAnyTokenVal("SHARE")
		this.nextToken()
		this.acceptAnyTokenVal("MODE")
		this.nextToken()
		m.LockingRead = MySqlLockingRead_.LOCK_IN_SHARE_MODE
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
		case Token_.LOW_PRIORITY:
			m.LowPriority = true
			isOption = true
		case Token_.IGNORE:
			m.Ignore = true
			isOption = true
		}

		if !isOption {
			break
		}
		this.nextToken()
	}

	m.TableReference = this.parseTableReferenceSyntax_(TableReferenceSyntaxLevel_.JOIN)
	this.acceptAnyToken(Token_.SET)
	this.nextToken()

	cl := NewAssignmentListSyntax()
	this.setBeginPosDefault(cl)
	for {
		c := this.parseAssignmentSyntax()
		cl.Add(c)
		if Token_.Not(this.token(), Token_.COMMA) {
			break
		}
		this.nextToken()
	}
	this.setEndPosDefault(cl)

	m.AssignmentList = cl
	m.Where = this.parseWhereSyntax()
	m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevel_.NORNAL)
	m.Limit = this.parseMySqlLimitSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseHintSyntax() (h *HintSyntax) {
	if Token_.Not(this.token(), Token_.COMMENT) {
		return
	}
	comment := this.tokenVal()
	if !strings.HasPrefix(comment, "/*+") {
		return
	}
	h = NewHintSyntax()
	h.CommentType = CommentType_.MULTI_LINE
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
		case Token_.LOW_PRIORITY:
			m.LowPriority = true
			isOption = true
		case Token_.DELAYED:
			m.Delayed = true
			isOption = true
		case Token_.HIGHP_RIORITY:
			m.HighPriority = true
			isOption = true
		case Token_.IGNORE:
			m.Ignore = true
			isOption = true
		}

		if !isOption {
			break
		}
	}

	if Token_.Is(this.token(), Token_.INTO) {
		this.nextToken()
	}
	this.acceptAnyToken(Token_.IDENTIFIER)

	t := NewMySqlNameTableReferenceSyntax()
	this.setBeginPosDefault(t)
	t.TableNameItem = this.parseTableNameItemSyntax()
	m.NameTableReference = t
	this.setEndPosDefault(m)

	if Token_.Is(this.token(), Token_.L_PAREN) {
		icl := NewInsertColumnListSyntax()
		this.setBeginPosDefault(icl)

		if Token_.Not(this.nextToken(), Token_.R_PAREN) {
			for {
				i := this.parseMySqlIdentifierSyntax(true)
				icl.Add(i)
				if Token_.Not(this.token(), Token_.COMMA) {
					break
				}
				this.nextToken()
			}
		}

		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(icl)
		m.InsertColumnList = icl
	}

	columnNum := -1
	if m.InsertColumnList != nil {
		columnNum = m.InsertColumnList.Len()
	}

	rowConstructorList := false
	if Token_.Is(this.token(), Token_.VALUES) {
		if Token_.Is(this.nextToken(), Token_.ROW) {
			rowConstructorList = true
		}
		m.ValueListList = this.parseMySqlValueListListSyntax(columnNum, rowConstructorList)
	} else if this.equalTokenVal("VALUE") {
		this.nextToken()
		m.ValueListList = this.parseMySqlValueListListSyntax(columnNum, rowConstructorList)
	} else if Token_.Is(this.token(), Token_.SET) {
		al := NewAssignmentListSyntax()
		this.setBeginPosDefault(al)
		this.nextToken()
		for {
			a := this.parseAssignmentSyntax()
			al.Add(a)
			if Token_.Not(this.token(), Token_.COMMA) {
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
		if Token_.Is(this.token(), Token_.AS) {
			this.nextToken()
			m.RowAlias = this.parseMySqlIdentifierSyntax(true)
			if Token_.Is(this.token(), Token_.L_PAREN) {
				ml := this.parseMySqlIdentifierListSyntax()
				oc := ml.OperandCount()
				if columnNum != oc {
					this.panicBySyntax(ml, "column alias count doesn't match value count, column: "+strconv.Itoa(columnNum)+", alias: "+strconv.Itoa(oc))
				}
				m.ColumnAliasList = ml
			}
		}
	}

	if Token_.Is(this.token(), Token_.ON) {
		al := NewAssignmentListSyntax()
		this.setBeginPosDefault(al)
		this.nextToken()
		this.acceptAnyTokenVal("DUPLICATE")
		this.nextToken()
		this.acceptAnyToken(Token_.KEY)
		this.nextToken()
		this.acceptAnyToken(Token_.UPDATE)
		this.nextToken()
		for {
			a := this.parseAssignmentSyntax()
			al.Add(a)
			if Token_.Not(this.token(), Token_.COMMA) {
				break
			}
			this.nextToken()
		}
		this.setEndPosDefault(al)
		m.OnDuplicateKeyUpdateAssignmentList = al
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
			this.acceptAnyToken(Token_.ROW)
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

		if Token_.Not(this.token(), Token_.COMMA) {
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
		case Token_.LOW_PRIORITY:
			m.LowPriority = true
			isOption = true
		case Token_.IGNORE:
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
	if Token_.Not(this.token(), Token_.FROM) {
		m.MultiDeleteMode = MySqlMultiDeleteMode_.MODE1
		ml := NewMySqlMultiDeleteTableAliasListSyntax()
		this.setBeginPosDefault(ml)
		m.MultiDeleteTableAliasList = ml
		for {
			md := NewMySqlMultiDeleteTableAliasSyntax()
			this.setBeginPosDefault(md)
			md.Alias = this.parseMySqlIdentifierSyntax(true)
			if Token_.Is(this.token(), Token_.DOT) {
				this.nextToken()
				this.acceptAnyToken(Token_.STAR)
				this.nextToken()
				md.HasStar = true
			}
			this.setEndPosDefault(md)
			ml.Add(md)
			if Token_.Not(this.token(), Token_.COMMA) {
				break
			}
			this.nextToken()
		}
		this.setEndPosDefault(ml)
		this.acceptAnyToken(Token_.FROM)
		this.nextToken()
		m.TableReference = this.parseTableReferenceSyntax_(TableReferenceSyntaxLevel_.JOIN)
	} else {
		this.nextToken()
		i := this.parseMySqlIdentifierSyntax(true)

		switch this.token() {
		case Token_.DOT:
			switch this.nextToken() {
			case Token_.STAR:
				this.nextToken()
				m.MultiDeleteMode = MySqlMultiDeleteMode_.MODE2

				md := NewMySqlMultiDeleteTableAliasSyntax()
				this.setBeginPos(md, i.BeginPos)
				md.Alias = i
				md.HasStar = true
				this.setEndPos(md, i.EndPos)

				m.MultiDeleteTableAliasList = this.parseMySqlMultiDeleteTableAliasListSyntax(md)
				this.acceptAnyToken(Token_.USING)
				this.nextToken()
				m.TableReference = this.parseTableReferenceSyntax_(TableReferenceSyntaxLevel_.JOIN)
			case Token_.IDENTIFIER:
				t := NewTableNameItemSyntax()
				this.setBeginPos(t, i.BeginPos)
				t.Catalog = i
				t.TableName = this.parseMySqlIdentifierSyntax(false)
				this.setEndPosDefault(t)

				tnt := NewMySqlNameTableReferenceSyntax()
				this.setBeginPos(tnt, t.BeginPos)
				tnt.TableNameItem = t
				if alias := this.parseAliasSyntax_(AliasSyntaxLevel_.IDENTIFIER); alias != nil {
					tnt.Alias = alias.(IdentifierSyntax_)
				}
				this.setEndPos(tnt, t.EndPos)
				m.TableReference = tnt
			default:
				this.panicByUnexpectedToken()
			}
		case Token_.COMMA:
			m.MultiDeleteMode = MySqlMultiDeleteMode_.MODE2
			md := NewMySqlMultiDeleteTableAliasSyntax()
			this.setBeginPos(md, i.BeginPos)
			md.Alias = i
			this.setEndPos(md, i.EndPos)

			m.MultiDeleteTableAliasList = this.parseMySqlMultiDeleteTableAliasListSyntax(md)
			this.acceptAnyToken(Token_.USING)
			this.nextToken()
			m.TableReference = this.parseTableReferenceSyntax_(TableReferenceSyntaxLevel_.JOIN)
		default:
			t := NewTableNameItemSyntax()
			this.setBeginPos(t, i.BeginPos)
			t.TableName = i
			this.setEndPos(t, i.EndPos)

			tnt := NewMySqlNameTableReferenceSyntax()
			this.setBeginPos(tnt, i.BeginPos)
			tnt.TableNameItem = t
			if alias := this.parseAliasSyntax_(AliasSyntaxLevel_.IDENTIFIER); alias != nil {
				tnt.Alias = alias.(IdentifierSyntax_)
			}
			this.setEndPos(tnt, t.EndPos)
			m.TableReference = tnt
		}
	}
	m.Where = this.parseWhereSyntax()
	m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevel_.NORNAL)
	m.Limit = this.parseMySqlLimitSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlMultiDeleteTableAliasListSyntax(m *MySqlMultiDeleteTableAliasSyntax) (ml *MySqlMultiDeleteTableAliasListSyntax) {
	ml = NewMySqlMultiDeleteTableAliasListSyntax()
	this.setBeginPos(ml, m.BeginPos)
	ml.Add(m)

	for {
		if Token_.Not(this.token(), Token_.COMMA) {
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
	if Token_.Is(this.token(), Token_.DOT) {
		this.nextToken()
		this.acceptAnyToken(Token_.STAR)
		this.nextToken()
		m.HasStar = true
	}
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlIdentifierListSyntax() (ml *MySqlIdentifierListSyntax) {
	ml = NewMySqlIdentifierListSyntax()
	this.setBeginPosDefault(ml)
	this.acceptAnyToken(Token_.L_PAREN)
	this.nextToken()
	if Token_.Not(this.token(), Token_.R_PAREN) {
		for {
			ml.Add(this.parseMySqlIdentifierSyntax(true))
			if Token_.Not(this.token(), Token_.COMMA) {
				break
			}
			this.nextToken()
		}
	}
	this.acceptAnyToken(Token_.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(ml)
	this.parenthesizingSyntax(ml)
	return
}

func (this *mySqlParser) parseAggregateOption() (a AggregateOption) {
	switch this.token() {
	case Token_.DISTINCT:
		a = AggregateOption_.DISTINCT
		this.nextToken()
	case Token_.ALL:
		a = AggregateOption_.ALL
		this.nextToken()
	}
	return
}

func (this *mySqlParser) parseOrderBySyntax(l OrderingItemSyntaxLevel) (o *OrderBySyntax) {
	if Token_.Not(this.token(), Token_.ORDER) {
		return
	}
	o = NewOrderBySyntax()
	this.setBeginPosDefault(o)

	this.nextToken()
	this.acceptAnyToken(Token_.BY)
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
		if Token_.Not(this.token(), Token_.COMMA) {
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
	ci_ := this.parseColumnItemSyntax_()
	if OrderingItemSyntaxLevel_.Is(l, OrderingItemSyntaxLevel_.IDENTIFIER) {
		if _, ok := ci_.(*PropertySyntax); ok {
			this.panicBySyntax(ci_, "cannot be used table alias in global clause of multiset syntax")
		}
	}
	o.Column = ci_
	if Token_.Is(this.token(), Token_.ASC) {
		o.OrderingSequence = OrderingSequence_.ASC
		this.nextToken()
	} else if Token_.Is(this.token(), Token_.DESC) {
		o.OrderingSequence = OrderingSequence_.DESC
		this.nextToken()
	}
	this.setEndPosDefault(o)
	return
}

func (this *mySqlParser) parseMySqlLimitSyntax() (m *MySqlLimitSyntax) {
	if Token_.Not(this.token(), Token_.LIMIT) {
		return
	}
	m = NewMySqlLimitSyntax()
	this.setBeginPosDefault(m)

	this.nextToken()
	this.acceptAnyToken(Token_.DECIMAL_NUMBER)

	d := this.parseDecimalNumberSyntax()
	if Token_.Is(this.token(), Token_.COMMA) {
		m.Offset = d
		this.nextToken()
		this.acceptAnyToken(Token_.DECIMAL_NUMBER)
		m.RowCount = this.parseDecimalNumberSyntax()
	} else if this.equalTokenVal("OFFSET") {
		m.RowCount = d
		this.nextToken()
		this.acceptAnyToken(Token_.DECIMAL_NUMBER)
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
		si := this.parseSelectItemSyntax_(!sil.HasAllColumn)
		if _, ok := si.(*AllColumnSyntax); ok {
			sil.HasAllColumn = true
		}
		sil.Add(si)
		if Token_.Not(this.token(), Token_.COMMA) {
			break
		}
		this.nextToken()
	}
	this.setEndPosDefault(sil)
	return
}

func (this *mySqlParser) parseTableReferenceSyntax_(l TableReferenceSyntaxLevel) (tr_ TableReferenceSyntax_) {
	if Token_.Is(this.token(), Token_.L_PAREN) {
		beginPos := this.tokenBeginPos()
		this.nextToken()
		tr_ = this.parseTableReferenceSyntax_(TableReferenceSyntaxLevel_.JOIN)
		this.setBeginPos(tr_, beginPos)
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(tr_)

		if d, ok := tr_.(*DerivedTableReferenceSyntax); ok {
			this.parenthesizingSyntax(d.Query)
			if alias := this.parseAliasSyntax_(AliasSyntaxLevel_.IDENTIFIER); alias == nil {
				this.panicBySyntax(d, "every derived table must have its own alias")
			} else {
				d.Alias = alias.(IdentifierSyntax_)
			}
		} else {
			this.parenthesizingSyntax(tr_)
		}
	} else {
		tr_ = this.parseTableReferenceSyntax_Inner()
	}

	if TableReferenceSyntaxLevel_.Is(l, TableReferenceSyntaxLevel_.JOIN) {
		tr_ = this.parseTableReferenceSyntax_Rest(tr_)
	}
	return
}

func (this *mySqlParser) parseMySqlIndexHintSyntax() (m *MySqlIndexHintSyntax) {
	var mode MySqlIndexHintMode
	switch this.token() {
	case Token_.USE:
		mode = MySqlIndexHintMode_.USE
	case Token_.IGNORE:
		mode = MySqlIndexHintMode_.IGNORE
	case Token_.FORCE:
		mode = MySqlIndexHintMode_.FORCE
	default:
		return
	}
	m = NewMySqlIndexHintSyntax()
	this.setBeginPosDefault(m)
	m.IndexHintMode = mode
	this.nextToken()
	this.acceptAnyToken(Token_.INDEX, Token_.KEY)
	if Token_.Is(this.nextToken(), Token_.FOR) {
		switch this.nextToken() {
		case Token_.JOIN:
			m.IndexHintFor = MySqlIndexHintFor_.JOIN
			this.nextToken()
		case Token_.GROUP:
			m.IndexHintFor = MySqlIndexHintFor_.GROUP_BY
			this.nextToken()
			this.acceptAnyToken(Token_.BY)
			this.nextToken()
		case Token_.ORDER:
			m.IndexHintFor = MySqlIndexHintFor_.ORDER_BY
			this.nextToken()
			this.acceptAnyToken(Token_.BY)
			this.nextToken()
		default:
			this.panicByUnexpectedToken()
		}
	}

	m.IndexList = this.parseMySqlIdentifierListSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseTableReferenceSyntax_Inner() (tr_ TableReferenceSyntax_) {
	switch this.token() {
	case Token_.IDENTIFIER:
		tnt := NewMySqlNameTableReferenceSyntax()
		this.setBeginPosDefault(tnt)
		tnt.TableNameItem = this.parseTableNameItemSyntax()
		tnt.PartitionList = this.parsePartitionListSyntax()
		if alias := this.parseAliasSyntax_(AliasSyntaxLevel_.IDENTIFIER); alias != nil {
			tnt.Alias = alias.(IdentifierSyntax_)
		}
		tnt.IndexHintList = this.parseMySqlIndexHintListSyntax()
		this.setEndPosDefault(tnt)
		tr_ = tnt
	case Token_.SELECT:
		if Token_.Not(this.prevToken(), Token_.L_PAREN) {
			this.panicByToken("subquery expression must be parenthesized")
		}
		dtt := NewDerivedTableTableReferenceSyntax()
		dtt.Query = this.parseQuerySyntax_(QuerySyntaxLevel_.NORMAL)
		tr_ = dtt
	case Token_.DUAL:
		d := NewDualTableReferenceSyntax()
		this.setBeginPosDefault(d)
		this.nextToken()
		this.setEndPosDefault(d)
		tr_ = d
	default:
		this.panicByUnexpectedToken()
	}
	return
}

func (this *mySqlParser) parsePartitionListSyntax() (pl *PartitionListSyntax) {
	if Token_.Not(this.token(), Token_.PARTITION) {
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

func (this *mySqlParser) parseTableReferenceSyntax_Rest(source TableReferenceSyntax_) (last TableReferenceSyntax_) {
	last = source
	natural := false
	if Token_.Is(this.token(), Token_.NATURAL) {
		natural = true
		this.nextToken()
	}

	var joinType JoinType
	switch this.token() {
	case Token_.LEFT:
		if Token_.Is(this.nextToken(), Token_.OUTER) {
			this.nextToken()
		}
		this.acceptAnyToken(Token_.JOIN)
		joinType = JoinType_.LEFT_OUTER_JOIN
		this.nextToken()
	case Token_.RIGHT:
		if Token_.Is(this.nextToken(), Token_.OUTER) {
			this.nextToken()
		}
		this.acceptAnyToken(Token_.JOIN)
		joinType = JoinType_.RIGHT_OUTER_JOIN
		this.nextToken()
	case Token_.INNER:
		this.nextToken()
		this.acceptAnyToken(Token_.JOIN)
		joinType = JoinType_.INNER_JOIN
		this.nextToken()
	case Token_.JOIN:
		joinType = JoinType_.JOIN
		this.nextToken()
	case Token_.COMMA:
		joinType = JoinType_.COMMA
		this.nextToken()
	case Token_.STRAIGHT_JOIN:
		joinType = JoinType_.STRAIGHT_JOIN
		this.nextToken()
	case Token_.CROSS:
		this.nextToken()
		this.acceptAnyToken(Token_.JOIN)
		joinType = JoinType_.CROSS_JOIN
		this.nextToken()
	}

	if !joinType.Undefined() {
		jt := NewJoinTableReferenceSyntax()
		jt.Left = last
		jt.Natural = natural
		jt.JoinType = joinType
		jt.Right = this.parseTableReferenceSyntax_(TableReferenceSyntaxLevel_.DERIVED)

		if !natural {
			switch this.token() {
			case Token_.ON:
				this.nextToken()
				jt.JoinCondition = this.parseJoinOnSyntax()
			case Token_.USING:
				this.nextToken()
				jt.JoinCondition = this.parseJoinUsingSyntax()
			default:
				if JoinType_.Not(joinType, JoinType_.COMMA, JoinType_.INNER_JOIN, JoinType_.CROSS_JOIN, JoinType_.STRAIGHT_JOIN) {
					this.panicByUnexpectedToken()
				}
			}
		}
		last = this.parseTableReferenceSyntax_Rest(jt)
	}
	return
}

func (this *mySqlParser) parseWhereSyntax() (w *WhereSyntax) {
	if Token_.Not(this.token(), Token_.WHERE) {
		return
	}
	w = NewWhereSyntax()
	this.setBeginPosDefault(w)
	this.nextToken()
	w.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
	this.setEndPosDefault(w)
	return
}

func (this *mySqlParser) parseMySqlGroupBySyntax() (m *MySqlGroupBySyntax) {
	if Token_.Not(this.token(), Token_.GROUP) {
		return
	}
	m = NewMySqlGroupBySyntax()
	this.setBeginPosDefault(m)
	this.nextToken()
	this.acceptAnyToken(Token_.BY)
	this.nextToken()
	// GROUP BY语法的ASC、DESC修饰在8.0中被删除，本解析器依然支持，兼容5.x
	m.OrderingItemList = this.parseOrderingItemListSyntax(OrderingItemSyntaxLevel_.NORNAL)
	if Token_.Is(this.token(), Token_.WITH) {
		this.nextToken()
		this.acceptAnyTokenVal("ROLLUP")
		m.WithRollup = true
		this.nextToken()
	}
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseHavingSyntax() (h *HavingSyntax) {
	if Token_.Not(this.token(), Token_.HAVING) {
		return
	}
	h = NewHavingSyntax()
	this.setBeginPosDefault(h)
	this.nextToken()
	h.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
	this.setEndPosDefault(h)
	return
}

func (this *mySqlParser) parseAssignmentSyntax() (a *AssignmentSyntax) {
	a = NewAssignmentSyntax()
	this.setBeginPosDefault(a)
	a.Column = this.parseColumnItemSyntax_()
	this.acceptAnyToken(Token_.EQ, Token_.COLON_EQ)

	if Token_.Is(this.nextToken(), Token_.DEFAULT) {
		a.Default = true
		this.nextToken()
	} else {
		a.Value = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.BOOLEAN_PREDICATE)
	}
	this.setEndPosDefault(a)
	return
}

func (this *mySqlParser) parseTableNameItemSyntax() (t *TableNameItemSyntax) {
	t = NewTableNameItemSyntax()
	this.setBeginPosDefault(t)
	i := this.parseMySqlIdentifierSyntax(true)
	if Token_.Is(this.token(), Token_.DOT) {
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
	this.acceptAnyToken(Token_.L_PAREN)
	this.nextToken()
	if Token_.Not(this.token(), Token_.R_PAREN) {
		for {
			el.Add(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
			if Token_.Not(this.token(), Token_.COMMA) {
				break
			}
			this.nextToken()
		}
	}
	this.acceptAnyToken(Token_.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(el)
	this.parenthesizingSyntax(el)
	return
}

func (this *mySqlParser) parseMySqlIdentifierSyntax(check bool) (m *MySqlIdentifierSyntax) {
	if check {
		this.acceptAnyToken(Token_.IDENTIFIER)
	}
	m = NewMySqlIdentifierSyntax()
	this.setBeginPosDefault(m)
	m.Name = this.tokenVal()
	m.Qualifier = this.hasQualifier()
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseSingleOperandIExprSyntax(l ExprSyntaxLevel) ExprSyntax_ {
	e := this.parseAnyOperandIExprSyntax(l, 0)
	this.acceptExpectedOperandCount(1, e)
	return e
}

// parseAnyOperandIExprSyntax 解析任意操作数的表达式，operandCount表示若表达式为一个列表时，每个元素的操作数
func (this *mySqlParser) parseAnyOperandIExprSyntax(l ExprSyntaxLevel, listElementOperandCount int) (e ExprSyntax_) {
	if Token_.Is(this.token(), Token_.L_PAREN) {
		beginPos := this.tokenBeginPos()
		this.nextToken()
		e = this.parseAnyOperandIExprSyntax(ExprSyntaxLevel_.EXPR, 0)
		switch this.token() {
		case Token_.R_PAREN:
			this.nextToken()
			this.setBeginPos(e, beginPos)
			this.setEndPosDefault(e)
			this.parenthesizingSyntax(e)
		case Token_.COMMA:
			this.nextToken()
			el := NewExprListSyntax()
			this.setBeginPos(el, beginPos)
			this.acceptExpectedOperandCount(listElementOperandCount, e)
			el.Add(e)
			for {
				e = this.parseAnyOperandIExprSyntax(ExprSyntaxLevel_.EXPR, 0)
				this.acceptExpectedOperandCount(listElementOperandCount, e)
				el.Add(e)
				if Token_.Not(this.token(), Token_.COMMA) {
					break
				}
				this.nextToken()
			}
			this.acceptAnyToken(Token_.R_PAREN)
			this.nextToken()
			this.setEndPosDefault(el)
			this.parenthesizingSyntax(el)
			e = el
		default:
			this.panicByUnexpectedToken()
		}

		if ExprSyntaxLevel_.Not(l, ExprSyntaxLevel_.SINGLE) {
			e = this.parseOperandSyntax_Rest(l, e)
		}
	} else {
		e = this.parseExprSyntax_(l)
	}
	return
}

func (this *mySqlParser) parseExprSyntax_(l ExprSyntaxLevel) (e ExprSyntax_) {
	switch this.token() {
	case Token_.IDENTIFIER:
		i := this.parseMySqlIdentifierSyntax(false)
		switch this.token() {
		case Token_.DOT:
			e = this.parsePropertySyntax(i)
		case Token_.L_PAREN:
			e = this.parseFunctionSyntax_(i)
		case Token_.STRING:
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
				e = this.parseFunctionSyntax_(i)
			default:
				e = i
			}
		}
	case Token_.STRING:
		e = this.parseMySqlStringSyntax(false)
	case Token_.DECIMAL_NUMBER:
		e = this.parseDecimalNumberSyntax()
	case Token_.SELECT:
		if Token_.Not(this.prevToken(), Token_.L_PAREN) {
			this.panicByToken("subquery expression must be parenthesized")
		}
		e = this.parseQuerySyntax_(QuerySyntaxLevel_.NORMAL)
	case Token_.NULL:
		e = this.parseNullSyntax()
	case Token_.CASE:
		e = this.parseCaseSyntax()
	case Token_.EXISTS:
		e = this.parseExistsSyntax()
	case Token_.BINARY, Token_.SUB, Token_.BANG, Token_.TILDE, Token_.PLUS, Token_.NOT:
		e = this.parseMySqlUnarySyntax()
	case Token_.INTERVAL:
		i := this.parseMySqlIdentifierSyntax(false)
		if Token_.Is(this.token(), Token_.L_PAREN) {
			o := this.parseAnyOperandIExprSyntax(ExprSyntaxLevel_.EXPR, 1)
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
	case Token_.TRUE:
		e = this.parseMySqlTrueSyntax()
	case Token_.FALSE:
		e = this.parseMySqlFalseSyntax()
	case Token_.QUES:
		e = this.parseParameterSyntax()
	case Token_.HEXADECIMAL_NUMBER:
		e = this.parseHexadecimalNumberSyntax()
	case Token_.BINARY_NUMBER:
		e = this.parseBinaryNumberSyntax()
	case Token_.VALUES, Token_.CHAR:
		i := this.parseMySqlIdentifierSyntax(false)
		this.acceptAnyToken(Token_.L_PAREN)
		e = this.parseFunctionSyntax_(i)
	case Token_.AT, Token_.AT_AT:
		e = this.parseMySqlVariableSyntax()
	default:
		this.panicByUnexpectedToken()
	}
	if ExprSyntaxLevel_.Not(l, ExprSyntaxLevel_.SINGLE) {
		e = this.parseOperandSyntax_Rest(l, e)
	}
	return
}

func (this *mySqlParser) parseOperandSyntax_Rest(l ExprSyntaxLevel, before ExprSyntax_) (last ExprSyntax_) {
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

func (this *mySqlParser) parseAliasSyntax_(l AliasSyntaxLevel) (a AliasSyntax_) {
	switch this.token() {
	case Token_.AS:
		this.nextToken()
		if a = this.parseAliasSyntax_(l); a == nil {
			this.panicByUnexpectedToken()
		}
	case Token_.STRING:
		if AliasSyntaxLevel_.Is(l, AliasSyntaxLevel_.STRING) {
			a = this.parseMySqlStringSyntax(false)
		}
	case Token_.IDENTIFIER:
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

func (this *mySqlParser) parseSelectItemSyntax_(allowAllColumnSyntax bool) (s SelectItemSyntax_) {
	if Token_.Is(this.token(), Token_.STAR) {
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
	j.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
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

func (this *mySqlParser) parseColumnItemSyntax_() (ci_ ColumnItemSyntax_) {
	i := this.parseMySqlIdentifierSyntax(true)
	if Token_.Is(this.token(), Token_.DOT) {
		ci_ = this.parsePropertySyntax(i)
	} else {
		ci_ = i
	}
	return
}

func (this *mySqlParser) parsePropertySyntax(i *MySqlIdentifierSyntax) (pt *PropertySyntax) {
	pt = NewPropertySyntax()
	this.setBeginPos(pt, i.BeginPos)
	pt.Owner = i
	pt.Value = this.parsePropertyValueSyntax_()
	this.setEndPosDefault(pt)
	return
}

func (this *mySqlParser) parsePropertyValueSyntax_() (pv PropertyValueSyntax_) {
	switch this.nextToken() {
	case Token_.STAR:
		pv = this.parseAllColumnSyntax()
	case Token_.IDENTIFIER:
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
	s.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
	s.Alias = this.parseAliasSyntax_(AliasSyntaxLevel_.STRING)
	this.setEndPosDefault(s)
	return
}

func (this *mySqlParser) parseFunctionSyntax_(functionName *MySqlIdentifierSyntax) (f FunctionSyntax_) {
	upperFunctionName := strings.ToUpper(functionName.Name)
	switch upperFunctionName {
	case "CONVERT":
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		c := NewMySqlConvertFunctionSyntax()
		this.setBeginPos(c, functionName.BeginPos)
		c.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
		switch this.token() {
		case Token_.USING:
			c.UsingTranscoding = true
			this.nextToken()
			c.TranscodingName = this.tokenVal()
			this.nextToken()
		case Token_.COMMA:
			this.nextToken()
			c.DataType = this.parseMySqlCastDataTypeSyntax()
		default:
			this.panicByUnexpectedToken()
		}
		this.acceptAnyToken(Token_.R_PAREN)
		if Token_.Is(this.nextToken(), Token_.COLLATE) {
			this.nextToken()
			c.Collate = this.parseMySqlIdentifierSyntax(true).Name
		}
		this.setEndPosDefault(c)
		f = c
	case "CAST":
		// https://dev.mysql.com/doc/refman/8.0/en/cast-functions.html#function_cast
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		c := NewMySqlCastFunctionSyntax()
		this.setBeginPos(c, functionName.BeginPos)
		c.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
		if this.tokenValUpper() == "AT" {
			this.nextToken()
			this.acceptAnyTokenVal("TIME")
			this.nextToken()
			this.acceptAnyTokenVal("ZONE")
			hasInterval := false
			if Token_.Is(this.nextToken(), Token_.INTERVAL) {
				hasInterval = true
				this.nextToken()
			}
			s := this.parseMySqlStringSyntax(true)
			if !(s.Value() == "+00:00" || (s.Value() == "UTC" && !hasInterval)) {
				this.panicBySyntax(s, "unknown or incorrect time zone: %s", s.Sql())
			}
			c.AtTimeZone = s
		}
		this.acceptAnyToken(Token_.AS)
		this.nextToken()
		c.DataType = this.parseMySqlCastDataTypeSyntax()
		this.acceptAnyToken(Token_.R_PAREN)
		if Token_.Is(this.nextToken(), Token_.COLLATE) {
			this.nextToken()
			c.Collate = this.parseMySqlIdentifierSyntax(true).Name
		}
		this.setEndPosDefault(c)
		f = c
	case "EXTRACT":
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		e := NewMySqlExtractFunctionSyntax()
		this.setBeginPos(e, functionName.BeginPos)
		t := this.parseMySqlTemporalInterval()
		if t.Undefined() {
			this.panicByUnexpectedToken()
		}
		e.Unit = t
		this.acceptAnyToken(Token_.FROM)
		this.nextToken()
		e.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(e)
		f = e
	case "TIMESTAMPADD", "TIMESTAMPDIFF":
		this.acceptAnyToken(Token_.L_PAREN)
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
			this.acceptAnyToken(Token_.COMMA)
			this.nextToken()
			t.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
		}
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(t)
		f = t
	case "TRIM":
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		t := NewMySqlTrimFunctionSyntax()
		this.setBeginPos(t, functionName.BeginPos)
		if Token_.Is(this.token(), Token_.BOTH) {
			t.TrimMode = MySqlTrimMode_.BOTH
			this.nextToken()
		} else if Token_.Is(this.token(), Token_.LEADING) {
			t.TrimMode = MySqlTrimMode_.LEADING
			this.nextToken()
		} else if this.tokenValUpper() == "TRAILING" {
			t.TrimMode = MySqlTrimMode_.TRAILING
			this.nextToken()
		}

		tmpExpr := this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
		if Token_.Is(this.token(), Token_.FROM) {
			t.RemStr = tmpExpr
			this.nextToken()
			t.Str = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
		} else {
			t.Str = tmpExpr
		}

		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(t)
		f = t
	case "CHAR":
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		c := NewMySqlCharFunctionSyntax()
		this.setBeginPos(c, functionName.BeginPos)
		for {
			c.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
			if Token_.Not(this.token(), Token_.COMMA) {
				break
			}
			this.nextToken()
		}
		if Token_.Is(this.token(), Token_.USING) {
			this.nextToken()
			c.CharsetName = this.tokenVal()
			this.nextToken()
		}
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(c)
		f = c
	case "AVG", "BIT_AND", "BIT_OR", "BIT_XOR", "COUNT", "JSON_ARRAYAGG", "JSON_OBJECTAGG", "MAX", "MIN",
		"STD", "STDDEV", "STDDEV_POP", "STDDEV_SAMP", "SUM", "VAR_POP", "VAR_SAMP", "VARIANCE":
		// 通用的聚合函数解析，统一解析为：function_name([DISTINCT ](*|[<expr>[, <expr>]...]))[ <over_syntax>]
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		a := NewAggregateFunctionSyntax()
		this.setBeginPos(a, functionName.BeginPos)
		a.Name = upperFunctionName
		if Token_.Not(this.token(), Token_.R_PAREN) {
			a.AggregateOption = this.parseAggregateOption()
			if Token_.Is(this.token(), Token_.STAR) {
				a.AllColumnParameter = true
				this.nextToken()
			} else {
				for {
					a.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
					if Token_.Not(this.token(), Token_.COMMA) {
						break
					}
					this.nextToken()
				}
			}
		}
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		a.Over = this.parseOverSyntax()
		this.setEndPosDefault(a)
		f = a
	case "GROUP_CONCAT":
		// GROUP_CONCAT聚合函数特殊处理
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		g := NewMySqlGroupConcatFunctionSyntax()
		this.setBeginPos(g, functionName.BeginPos)
		g.AggregateOption = this.parseAggregateOption()
		for {
			g.AddParameter(this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR))
			if Token_.Not(this.token(), Token_.COMMA) {
				break
			}
		}
		g.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevel_.NORNAL)
		if Token_.Is(this.token(), Token_.SEPARATOR) {
			this.nextToken()
			g.Separator = this.parseMySqlStringSyntax(true)
		}
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(g)
		f = g
	case "GET_FORMAT":
		this.acceptAnyToken(Token_.L_PAREN)
		this.nextToken()
		g := NewMySqlGetFormatFunctionSyntax()
		this.setBeginPos(g, functionName.BeginPos)
		d := MySqlGetFormatType_.OfSql(strings.ToUpper(this.tokenVal()))
		if d.Undefined() {
			this.panicByUnexpectedToken()
		}
		g.Type = d
		this.nextToken()
		this.acceptAnyToken(Token_.COMMA)
		this.nextToken()
		g.DateFormat = this.parseMySqlStringSyntax(true)
		this.acceptAnyToken(Token_.R_PAREN)
		this.nextToken()
		this.setEndPosDefault(g)
		f = g
	case "CUME_DIST", "DENSE_RANK", "FIRST_VALUE", "LAG", "LAST_VALUE", "LEAD", "NTH_VALUE", "NTILE", "PERCENT_RANK", "RANK", "ROW_NUMBER":
		wf := NewWindowFunctionSyntax()
		this.setBeginPos(wf, functionName.BeginPos)
		wf.Name = upperFunctionName
		wf.Parameters = this.parseSingleOperandExprListSyntax()
		if Token_.Is(this.token(), Token_.IGNORE) {
			this.nextToken()
			this.acceptAnyToken(Token_.NULL)
			wf.NullTreatment = NullTreatment_.IGNORE_NULLS
			this.nextToken()
		}
		if this.tokenValUpper() == "RESPECT" {
			this.nextToken()
			this.acceptAnyToken(Token_.NULL)
			wf.NullTreatment = NullTreatment_.RESPECT_NULLS
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
	if Token_.Is(this.token(), Token_.COLLATE) {
		this.nextToken()
		t.Collate = this.parseMySqlIdentifierSyntax(true).Name
	}
	this.setEndPosDefault(t)
	return
}

func (this *mySqlParser) parseMySqlDateAndTimeLiteralSyntax(i *MySqlIdentifierSyntax) (d *MySqlDateAndTimeLiteralSyntax) {
	d = NewMySqlDateAndTimeLiteralSyntax()
	this.setBeginPos(d, i.BeginPos)
	d.Type = MySqlDatetimeLiteralType_.OfSql(strings.ToUpper(i.Name))
	d.DateAndTime = this.parseMySqlStringSyntax(true)
	this.setEndPosDefault(d)
	return
}

func (this *mySqlParser) parseMySqlStringSyntax(check bool) (m *MySqlStringSyntax) {
	if check {
		this.acceptAnyToken(Token_.STRING)
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
	if Token_.Not(this.nextToken(), Token_.WHEN) {
		c.ValueExpr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
	}
	this.acceptAnyToken(Token_.WHEN)
	c.WhenItemList = this.parseCaseWhenItemListSyntax()
	if Token_.Is(this.token(), Token_.ELSE) {
		this.nextToken()
		c.ElseExr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
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
		if Token_.Not(this.token(), Token_.WHEN) {
			break
		}
	}
	this.setEndPosDefault(cl)
	return
}

func (this *mySqlParser) parseCaseWhenItemSyntax() (c *CaseWhenItemSyntax) {
	c = NewCaseWhenItem()
	this.setBeginPosDefault(c)
	this.acceptAnyToken(Token_.WHEN)
	this.nextToken()
	c.Condition = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
	this.acceptAnyToken(Token_.THEN)
	this.nextToken()
	c.Result = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
	this.setEndPosDefault(c)
	return
}

func (this *mySqlParser) parseExistsSyntax() (e *ExistsSyntax) {
	e = NewExistsSyntax()
	this.setBeginPosDefault(e)
	this.nextToken()
	this.acceptAnyToken(Token_.L_PAREN)
	e.Query = this.parseQuerySyntax_(QuerySyntaxLevel_.NORMAL)
	this.setEndPosDefault(e)
	return
}

func (this *mySqlParser) parseMySqlUnarySyntax() (m *MySqlUnarySyntax) {
	m = NewMySqlUnarySyntax()
	this.setBeginPosDefault(m)
	uo := this.parseMySqlUnaryOperator()
	m.UnaryOperator = uo
	m.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.SINGLE)
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlIntervalSyntax() (m *MySqlIntervalSyntax) {
	m = NewMySqlIntervalSyntax()
	this.setBeginPos(m, this.prevTokenBeginPos())
	m.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
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

func (this *mySqlParser) parseParameterSyntax() (pa *PlaceholderSyntax) {
	pa = NewPlaceholderSyntax()
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
	if Token_.Is(this.token(), Token_.AT) {
		m.VariableType = MySqlVariableType_.SESSION
	} else if Token_.Is(this.token(), Token_.AT_AT) {
		m.VariableType = MySqlVariableType_.GLOBAL
	}
	this.nextToken()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlBinaryOperator(l ExprSyntaxLevel, leftOperand ExprSyntax_) (bo MySqlBinaryOperator) {
	c := this.saveCursor()
	if Token_.Is(this.token(), Token_.IS) {
		if Token_.Is(this.nextToken(), Token_.NOT) {
			bo = MySqlBinaryOperator_.IS_NOT
			this.nextToken()
		} else {
			bo = MySqlBinaryOperator_.IS
		}
		this.acceptAnyTokenVal("NULL", "TRUE", "FALSE", "Undefined")
	} else if Token_.Is(this.token(), Token_.NOT) {
		switch this.nextToken() {
		case Token_.LIKE:
			bo = MySqlBinaryOperator_.NOT_LIKE
		case Token_.BETWEEN:
			bo = MySqlBinaryOperator_.NOT_BETWEEN
		case Token_.REGEXP:
			bo = MySqlBinaryOperator_.NOT_REGEXP
		case Token_.R_LIKE:
			bo = MySqlBinaryOperator_.NOT_RLIKE
		case Token_.IN:
			bo = MySqlBinaryOperator_.NOT_IN
		default:
			this.panicByUnexpectedToken()
		}
		this.nextToken()
	} else if this.tokenValUpper() == "SOUNDS" {
		if Token_.Is(this.nextToken(), Token_.LIKE) {
			bo = MySqlBinaryOperator_.SOUNDS_LIKE
			this.nextToken()
		} else {
			this.panicByUnexpectedToken()
		}
	} else if this.tokenValUpper() == "MEMBER" {
		if Token_.Is(this.nextToken(), Token_.OF) {
			bo = MySqlBinaryOperator_.MEMBER_OF
			this.nextToken()
		}
	} else {
		bo = mysqlTokenToBinaryOperators[this.token()]
		if !bo.Undefined() {
			this.nextToken()
		}
	}
	if !mysqlExprLevelToBinaryOperators[l].Contains(bo) {
		bo = MySqlBinaryOperator_.Undefined()
	}
	if !bo.Undefined() {
		if !bo.O.AllowMultipleOperand && leftOperand != nil {
			this.acceptExpectedOperandCount(1, leftOperand)
		}
	} else {
		this.rollback(c)
	}
	return
}

func (this *mySqlParser) parseMySqlBinaryOperationSyntax(l ExprSyntaxLevel, leftOperand ExprSyntax_, bo MySqlBinaryOperator) (b *MySqlBinaryOperationSyntax) {
	b = NewMySqlBinaryOperationSyntax()
	this.setBeginPos(b, leftOperand.Syntax_().BeginPos)
	b.LeftOperand = leftOperand
	b.BinaryOperator = bo.O
	if _, ok := leftOperand.(*MySqlIntervalSyntax); ok && MySqlBinaryOperator_.Is(bo, MySqlBinaryOperator_.SUBTRACT) {
		this.panicBySyntax(leftOperand, "for the - operator, INTERVAL expr unit is permitted only on the right side")
	}
	if MySqlBinaryOperator_.Is(bo, MySqlBinaryOperator_.BETWEEN, MySqlBinaryOperator_.NOT_BETWEEN) {
		b.RightOperand = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.CALCULATION)
		this.acceptAnyToken(Token_.AND)
		this.nextToken()
		b.BetweenThirdOperand = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.BOOLEAN_PREDICATE)
	} else {
		var (
			rightOperand ExprSyntax_
			cm           MySqlComparisonMode
		)
		switch bo {
		case MySqlBinaryOperator_.EQUAL_OR_ASSIGNMENT, MySqlBinaryOperator_.GREATER_THAN, MySqlBinaryOperator_.LESS_THAN, MySqlBinaryOperator_.GREATER_THAN_OR_EQUAL, MySqlBinaryOperator_.LESS_THAN_OR_EQUAL, MySqlBinaryOperator_.LESS_THAN_OR_GREATER, MySqlBinaryOperator_.NOT_EQUAL:
			if Token_.Is(this.token(), Token_.ALL) {
				cm = MySqlComparisonMode_.ALL
				this.nextToken()
			} else {
				var (
					tmpCsm MySqlComparisonMode
					c      *cursor
				)
				switch this.tokenValUpper() {
				case "ANY":
					c = this.saveCursor()
					tmpCsm = MySqlComparisonMode_.ANY
				case "SOME":
					c = this.saveCursor()
					tmpCsm = MySqlComparisonMode_.SOME
				}
				// 若下一个记号非运算符，则为特殊语法
				if !tmpCsm.Undefined() {
					this.nextToken()
					if this.parseMySqlBinaryOperator(ExprSyntaxLevel_.EXPR, nil).Undefined() {
						cm = tmpCsm
					} else {
						this.rollback(c)
					}
				}
			}
			if !cm.Undefined() {
				rightOperand = this.parseMySqlComparisonModeRightOperand()
			} else {
				rightOperand = this.parseAnyOperandIExprSyntax(ExprSyntaxLevel_.SINGLE, 0)
			}
		case MySqlBinaryOperator_.IN, MySqlBinaryOperator_.NOT_IN:
			this.acceptAnyToken(Token_.L_PAREN)
			c := this.saveCursor()
			if Token_.Is(this.nextToken(), Token_.SELECT) {
				this.rollback(c)
				rightOperand = this.parseQuerySyntax_(QuerySyntaxLevel_.NORMAL)
			} else {
				this.rollback(c)
				rightOperand = this.parseAnyOperandIExprSyntax(ExprSyntaxLevel_.SINGLE, 0)
			}
		case MySqlBinaryOperator_.MEMBER_OF:
			this.acceptAnyToken(Token_.L_PAREN)
			this.nextToken()
			rightOperand = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.SINGLE)
			this.acceptAnyToken(Token_.R_PAREN)
			this.nextToken()
			this.parenthesizingSyntax(rightOperand)
		default:
			rightOperand = this.parseAnyOperandIExprSyntax(ExprSyntaxLevel_.SINGLE, 0)
		}
		for {
			c := this.saveCursor()
			nextBo := this.parseMySqlBinaryOperator(l, nil)
			success, precedence := this.compareBinaryOperator(bo, nextBo)
			if success && precedence < 0 {
				rightOperand = this.parseMySqlBinaryOperationSyntax(l, rightOperand, nextBo)
			} else {
				this.rollback(c)
				break
			}
		}
		this.acceptEqualOperandCount(leftOperand, rightOperand, MySqlBinaryOperator_.Is(bo, MySqlBinaryOperator_.IN, MySqlBinaryOperator_.NOT_IN))
		b.ComparisonMode = cm
		b.RightOperand = rightOperand
		if MySqlBinaryOperator_.Is(bo, MySqlBinaryOperator_.LIKE, MySqlBinaryOperator_.NOT_LIKE) {
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
func (this *mySqlParser) parseMySqlComparisonModeRightOperand() (e_ ExprSyntax_) {
	this.acceptAnyToken(Token_.L_PAREN)
	switch this.nextToken() {
	case Token_.SELECT:
		e_ = this.parseQuerySyntax_(QuerySyntaxLevel_.NORMAL)
	case Token_.TABLE:
		e_ = this.parseMySqlTablesSyntax()
	default:
		this.panicByUnexpectedToken()
	}
	this.acceptAnyToken(Token_.R_PAREN)
	this.nextToken()
	this.parenthesizingSyntax(e_)
	return
}

func (this *mySqlParser) parseMySqlTablesSyntax() (m *MySqlTableSyntax) {
	this.nextToken()
	m = NewMySqlTableSyntax()
	this.setBeginPosDefault(m)
	m.TableNameItem = this.parseTableNameItemSyntax()
	m.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevel_.IDENTIFIER)
	m.Limit = this.parseMySqlLimitSyntax()
	this.setEndPosDefault(m)
	return
}

func (this *mySqlParser) parseMySqlTemporalInterval() (m MySqlTemporalInterval) {
	m = MySqlTemporalInterval_.OfSql(this.tokenVal())
	if !m.Undefined() {
		this.nextToken()
	}
	return
}

func (this *mySqlParser) parseMySqlUnaryOperator() (m UnaryOperator) {
	m = mysqlTokenToUnaryOperators[this.token()]
	this.nextToken()
	return
}

func (this *mySqlParser) parseMySqlCastDataTypeSyntax() (m *MySqlCastDataTypeSyntax) {
	m = NewMySqlCastDataTypeSyntax()
	this.setBeginPosDefault(m)
	// 通用的处理方式，不对具体类型做检查。https://dev.mysql.com/doc/refman/8.1/en/cast-functions.html#function_cast
	if Token_.Is(this.token(), Token_.IDENTIFIER) && !this.hasQualifier() || this.reserved() {
		m.Name = this.tokenValUpper()
		if Token_.Is(this.nextToken(), Token_.L_PAREN) {
			m.Parameters = this.parseMySqlCastDataTypeParamListSyntax()
		}
		if Token_.Is(this.token(), Token_.CHARACTER) {
			this.nextToken()
			this.acceptAnyToken(Token_.SET)
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
	this.acceptAnyToken(Token_.L_PAREN)
	this.nextToken()
	for {
		this.acceptAnyToken(Token_.DECIMAL_NUMBER)
		m.Add(this.parseDecimalNumberSyntax())
		if Token_.Not(this.token(), Token_.COMMA) {
			break
		}
		this.nextToken()
	}
	this.acceptAnyToken(Token_.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(m)
	this.parenthesizingSyntax(m)
	return
}

func (this *mySqlParser) parseOverSyntax() (o *OverSyntax) {
	if Token_.Not(this.token(), Token_.OVER) {
		return
	}
	o = NewOverSyntax()
	this.setBeginPosDefault(o)
	switch this.nextToken() {
	case Token_.IDENTIFIER:
		o.Window = this.parseMySqlIdentifierSyntax(false)
	case Token_.L_PAREN:
		this.nextToken()
		o.Window = this.parseWindowSpecSyntax()
		this.acceptAnyToken(Token_.R_PAREN)
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
	if Token_.Is(this.token(), Token_.IDENTIFIER) {
		w.Name = this.parseMySqlIdentifierSyntax(false)
	}

	w.PartitionBy = this.parsePartitionBySyntax()
	w.OrderBy = this.parseOrderBySyntax(OrderingItemSyntaxLevel_.NORNAL)
	w.Frame = this.parseWindowFrameSyntax()
	this.setEndPosDefault(w)
	return
}

func (this *mySqlParser) parsePartitionBySyntax() (py *PartitionBySyntax) {
	if Token_.Not(this.token(), Token_.PARTITION) {
		return
	}
	py = NewPartitionBySyntax()
	this.setBeginPosDefault(py)
	this.nextToken()
	this.acceptAnyToken(Token_.BY)
	this.nextToken()
	py.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
	this.setEndPosDefault(py)
	return
}

func (this *mySqlParser) parseWindowFrameSyntax() (w *WindowFrameSyntax) {
	var unit WindowFrameUnit
	switch this.token() {
	case Token_.ROWS:
		unit = WindowFrameUnit_.ROWS
	case Token_.RANGE:
		unit = WindowFrameUnit_.RANGE
	}
	if !unit.Undefined() {
		w = NewWindowFrameSyntax()
		this.setBeginPosDefault(w)
		w.Unit = unit
		if Token_.Is(this.nextToken(), Token_.BETWEEN) {
			this.nextToken()
			w.Extent = this.parseWindowFrameBetweenSyntax()
		} else {
			w.Extent = this.parseWindowFrameStartSyntax_()
		}
		this.setEndPosDefault(w)
	}
	return
}

func (this *mySqlParser) parseWindowFrameBetweenSyntax() (w *WindowFrameBetweenSyntax) {
	w = NewWindowFrameBetweenSyntax()
	this.setBeginPosDefault(w)
	w.Start = this.parseWindowFrameStartSyntax_()
	this.acceptAnyToken(Token_.AND)
	this.nextToken()
	w.End = this.parseWindowFrameStartSyntax_()
	this.setEndPosDefault(w)
	return
}

func (this *mySqlParser) parseWindowFrameStartSyntax_() (w WindowFrameStartEndSyntax_) {
	switch this.tokenValUpper() {
	case "CURRENT":
		wf := NewWindowFrameCurrentRowSyntax()
		this.setBeginPosDefault(wf)
		this.nextToken()
		this.acceptAnyToken(Token_.ROW)
		this.nextToken()
		this.setEndPosDefault(wf)
		w = wf
	case "UNBOUNDED":
		wf := NewWindowFrameUnboundedSyntax()
		this.setBeginPosDefault(wf)
		this.nextToken()
		this.acceptAnyTokenVal("PRECEDING", "FOLLOWING")
		wf.Type = WindowFrameStartEndType_.OfSql(this.tokenVal())
		this.nextToken()
		this.setEndPosDefault(wf)
		w = wf
	default:
		wf := NewWindowFrameExprSyntax()
		this.setBeginPosDefault(wf)
		wf.Expr = this.parseSingleOperandIExprSyntax(ExprSyntaxLevel_.EXPR)
		this.acceptAnyTokenVal("PRECEDING", "FOLLOWING")
		wf.Type = WindowFrameStartEndType_.OfSql(this.tokenVal())
		this.nextToken()
		this.setEndPosDefault(wf)
		w = wf
	}
	return
}

func (this *mySqlParser) parseNamedWindowListSyntax() (nwl *NamedWindowsListSyntax) {
	if Token_.Not(this.token(), Token_.WINDOW) {
		return
	}
	nwl = NewNamedWindowsListSyntax()
	this.setBeginPosDefault(nwl)
	this.nextToken()
	for {
		nwl.Add(this.parseNamedWindowsSyntax())
		if Token_.Not(this.token(), Token_.COMMA) {
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
	this.acceptAnyToken(Token_.AS)
	this.nextToken()
	this.acceptAnyToken(Token_.L_PAREN)
	this.nextToken()
	n.WindowSpec = this.parseWindowSpecSyntax()
	this.acceptAnyToken(Token_.R_PAREN)
	this.nextToken()
	this.setEndPosDefault(n)
	return
}

func newMySqlParser(sql string) *mySqlParser {
	p := &mySqlParser{}
	p.parser__ = extendParser(p, newMySqlLexer(sql))
	return p
}

var mysqlLfBinaryOperators = o.NewSet[MySqlBinaryOperator](
	MySqlBinaryOperator_.BITWISE_AND,
	MySqlBinaryOperator_.BOOLEAN_AND2,
	MySqlBinaryOperator_.BITWISE_OR,
	MySqlBinaryOperator_.BOOLEAN_OR2,
	MySqlBinaryOperator_.BOOLEAN_XOR,
)

var mysqlExprLevelToBinaryOperators = func() map[ExprSyntaxLevel]*o.Set[MySqlBinaryOperator] {
	calculation := o.NewSet[MySqlBinaryOperator]()
	calculation.Add(MySqlBinaryOperator_.MULTIPLY)
	calculation.Add(MySqlBinaryOperator_.ADD)
	calculation.Add(MySqlBinaryOperator_.SUBTRACT)
	calculation.Add(MySqlBinaryOperator_.DIVIDE)
	calculation.Add(MySqlBinaryOperator_.MODULUS)
	calculation.Add(MySqlBinaryOperator_.DIV)
	calculation.Add(MySqlBinaryOperator_.MOD)
	calculation.Add(MySqlBinaryOperator_.BITWISE_XOR)
	calculation.Add(MySqlBinaryOperator_.BITWISE_AND)
	calculation.Add(MySqlBinaryOperator_.BITWISE_OR)
	calculation.Add(MySqlBinaryOperator_.RIGHT_SHIFT)
	calculation.Add(MySqlBinaryOperator_.LEFT_SHIFT)
	calculation.Add(MySqlBinaryOperator_.COLLATE)
	calculation.Add(MySqlBinaryOperator_.JSON_EXTRACT)
	calculation.Add(MySqlBinaryOperator_.JSON_UNQUOTE)
	calculation.Add(MySqlBinaryOperator_.MEMBER_OF)

	booleanPredicate := o.NewSet[MySqlBinaryOperator]()
	booleanPredicate.AddSet(calculation)
	booleanPredicate.Add(MySqlBinaryOperator_.IN)
	booleanPredicate.Add(MySqlBinaryOperator_.NOT_IN)
	booleanPredicate.Add(MySqlBinaryOperator_.IS)
	booleanPredicate.Add(MySqlBinaryOperator_.IS_NOT)
	booleanPredicate.Add(MySqlBinaryOperator_.LIKE)
	booleanPredicate.Add(MySqlBinaryOperator_.NOT_LIKE)
	booleanPredicate.Add(MySqlBinaryOperator_.REGEXP)
	booleanPredicate.Add(MySqlBinaryOperator_.NOT_REGEXP)
	booleanPredicate.Add(MySqlBinaryOperator_.RLIKE)
	booleanPredicate.Add(MySqlBinaryOperator_.NOT_RLIKE)
	booleanPredicate.Add(MySqlBinaryOperator_.BETWEEN)
	booleanPredicate.Add(MySqlBinaryOperator_.NOT_BETWEEN)
	booleanPredicate.Add(MySqlBinaryOperator_.SOUNDS_LIKE)

	booleanPrimary := o.NewSet[MySqlBinaryOperator]()
	booleanPrimary.AddSet(booleanPredicate)
	booleanPrimary.Add(MySqlBinaryOperator_.EQUAL_OR_ASSIGNMENT)
	booleanPrimary.Add(MySqlBinaryOperator_.GREATER_THAN)
	booleanPrimary.Add(MySqlBinaryOperator_.LESS_THAN)
	booleanPrimary.Add(MySqlBinaryOperator_.GREATER_THAN_OR_EQUAL)
	booleanPrimary.Add(MySqlBinaryOperator_.LESS_THAN_OR_EQUAL)
	booleanPrimary.Add(MySqlBinaryOperator_.LESS_THAN_OR_GREATER)
	booleanPrimary.Add(MySqlBinaryOperator_.NOT_EQUAL)
	booleanPrimary.Add(MySqlBinaryOperator_.LESS_THAN_OR_EQUAL_OR_GREATER_THAN)

	booleanLogical := o.NewSet[MySqlBinaryOperator]()
	booleanLogical.AddSet(booleanPrimary)
	booleanLogical.Add(MySqlBinaryOperator_.BOOLEAN_AND)
	booleanLogical.Add(MySqlBinaryOperator_.BOOLEAN_OR)
	booleanLogical.Add(MySqlBinaryOperator_.BOOLEAN_AND2)
	booleanLogical.Add(MySqlBinaryOperator_.BOOLEAN_OR2)
	booleanLogical.Add(MySqlBinaryOperator_.BOOLEAN_XOR)

	return map[ExprSyntaxLevel]*o.Set[MySqlBinaryOperator]{
		ExprSyntaxLevel_.CALCULATION:       calculation,
		ExprSyntaxLevel_.BOOLEAN_PREDICATE: booleanPredicate,
		ExprSyntaxLevel_.BOOLEAN_PRIMARY:   booleanPrimary,
		ExprSyntaxLevel_.EXPR:              booleanLogical,
	}
}()

var mysqlTokenToBinaryOperators = map[Token]MySqlBinaryOperator{
	Token_.CARET:        MySqlBinaryOperator_.BITWISE_XOR,
	Token_.STAR:         MySqlBinaryOperator_.MULTIPLY,
	Token_.SLASH:        MySqlBinaryOperator_.DIVIDE,
	Token_.PERCENT:      MySqlBinaryOperator_.MODULUS,
	Token_.SUB:          MySqlBinaryOperator_.SUBTRACT,
	Token_.PLUS:         MySqlBinaryOperator_.ADD,
	Token_.LT_LT:        MySqlBinaryOperator_.LEFT_SHIFT,
	Token_.GT_GT:        MySqlBinaryOperator_.RIGHT_SHIFT,
	Token_.AMP:          MySqlBinaryOperator_.BITWISE_AND,
	Token_.BAR:          MySqlBinaryOperator_.BITWISE_OR,
	Token_.EQ:           MySqlBinaryOperator_.EQUAL_OR_ASSIGNMENT,
	Token_.LT_EQ_GT:     MySqlBinaryOperator_.LESS_THAN_OR_EQUAL_OR_GREATER_THAN,
	Token_.GT_EQ:        MySqlBinaryOperator_.GREATER_THAN_OR_EQUAL,
	Token_.GT:           MySqlBinaryOperator_.GREATER_THAN,
	Token_.LT:           MySqlBinaryOperator_.LESS_THAN,
	Token_.LT_EQ:        MySqlBinaryOperator_.LESS_THAN_OR_EQUAL,
	Token_.LT_GT:        MySqlBinaryOperator_.LESS_THAN_OR_GREATER,
	Token_.BANG_EQ:      MySqlBinaryOperator_.NOT_EQUAL,
	Token_.AMP_AMP:      MySqlBinaryOperator_.BOOLEAN_AND2,
	Token_.BAR_BAR:      MySqlBinaryOperator_.BOOLEAN_OR2,
	Token_.COLON_EQ:     MySqlBinaryOperator_.ASSIGN,
	Token_.REGEXP:       MySqlBinaryOperator_.REGEXP,
	Token_.R_LIKE:       MySqlBinaryOperator_.RLIKE,
	Token_.DIV:          MySqlBinaryOperator_.DIV,
	Token_.JSON_EXTRACT: MySqlBinaryOperator_.JSON_EXTRACT,
	Token_.JSON_UNQUOTE: MySqlBinaryOperator_.JSON_UNQUOTE,
	Token_.MOD:          MySqlBinaryOperator_.MOD,
	Token_.IN:           MySqlBinaryOperator_.IN,
	Token_.LIKE:         MySqlBinaryOperator_.LIKE,
	Token_.BETWEEN:      MySqlBinaryOperator_.BETWEEN,
	Token_.AND:          MySqlBinaryOperator_.BOOLEAN_AND,
	Token_.XOR:          MySqlBinaryOperator_.BOOLEAN_XOR,
	Token_.OR:           MySqlBinaryOperator_.BOOLEAN_OR,
}

var mysqlTokenToUnaryOperators = map[Token]UnaryOperator{
	Token_.BINARY: UnaryOperator_.BINARY,
	Token_.TILDE:  UnaryOperator_.COMPL,
	Token_.PLUS:   UnaryOperator_.POSITIVE,
	Token_.SUB:    UnaryOperator_.NEGATIVE,
	Token_.BANG:   UnaryOperator_.NOT,
	Token_.NOT:    UnaryOperator_.NOTSTR,
}
