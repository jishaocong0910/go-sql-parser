package parser

import (
	"strings"

	. "github.com/jishaocong0910/go-sql-parser/enum"
)

type i_Lexer interface {
	m_EC05053E2C60() *m_Lexer
	reserved() bool
	nextToken() Token
	nextTokenIncludeComment() Token
	token() Token
	tokenVal() string
	tokenValUpper() string
	tokenBeginPos() int
	tokenEndPos() int
	prevToken() Token
	prevTokenBeginPos() int
	prevTokenEndPos() int
	saveCursor() *cursor
	rollback(cursor *cursor)
}

type m_Lexer struct {
	I             i_Lexer
	sql           string
	chars         []rune
	reservedWords map[string]Token
	cursor        *cursor
}

func (this *m_Lexer) m_EC05053E2C60() *m_Lexer {
	return this
}

func (this *m_Lexer) reserved() bool {
	return this.cursor.tokenInfo.reserved
}

func (this *m_Lexer) token() Token {
	return this.cursor.tokenInfo.token
}

func (this *m_Lexer) tokenVal() string {
	return this.cursor.tokenInfo.tokenVal
}

func (this *m_Lexer) tokenValUpper() string {
	if this.cursor.tokenInfo.tokenValUpper == "" {
		this.cursor.tokenInfo.tokenValUpper = strings.ToUpper(this.cursor.tokenInfo.tokenVal)
	}
	return this.cursor.tokenInfo.tokenValUpper
}

func (this *m_Lexer) char() rune {
	return this.cursor.c
}

func (this *m_Lexer) setTokenBegin() {
	this.cursor.tokenInfo.tokenBeginPos = this.cursor.pos
}

func (this *m_Lexer) tokenBeginPos() int {
	return this.cursor.tokenInfo.tokenBeginPos
}

func (this *m_Lexer) tokenEndPos() int {
	return this.cursor.tokenInfo.tokenEndPos
}

func (this *m_Lexer) prevToken() Token {
	return this.cursor.prevTokenInfo.token
}

func (this *m_Lexer) prevTokenBeginPos() int {
	return this.cursor.prevTokenInfo.tokenBeginPos
}

func (this *m_Lexer) prevTokenEndPos() int {
	return this.cursor.prevTokenInfo.tokenEndPos
}

func (this *m_Lexer) saveCursor() *cursor {
	return this.cursor.copy()
}

func (this *m_Lexer) rollback(c *cursor) {
	this.cursor = c
}

func (this *m_Lexer) nextEoi() Token {
	this.cursor.tokenInfo.token = Tokens.EOI
	this.cursor.tokenInfo.tokenVal = string(Eoi)
	this.setTokenEnd()
	return Tokens.EOI
}

func (this *m_Lexer) setToken(token Token) {
	this.cursor.tokenInfo.token = token
	this.cursor.tokenInfo.tokenVal = string(this.chars[this.cursor.tokenInfo.tokenBeginPos:this.cursor.pos])
	this.setTokenEnd()
}

func (this *m_Lexer) setTokenEnd() {
	this.cursor.tokenInfo.tokenEndPos = this.cursor.pos
}

func (this *m_Lexer) beforeNextToken() {
	if Tokens.Is(this.cursor.tokenInfo.token, Tokens.EOI) {
		return
	}
	this.cursor.tokenInfo, this.cursor.prevTokenInfo = this.cursor.prevTokenInfo, this.cursor.tokenInfo
	this.cursor.tokenInfo.reset()
}

func (this *m_Lexer) nextChar() rune {
	if this.cursor.pos < len(this.chars) {
		this.cursor.pos++
		if this.cursor.pos == len(this.chars) {
			this.cursor.c = Eoi
		} else {
			this.cursor.c = this.chars[this.cursor.pos]
			// 非末尾而有eoi字符则跳过，防止误判
			if this.cursor.c == Eoi {
				this.nextChar()
			}
		}
	}
	return this.cursor.c
}

func (this *m_Lexer) previewChar(offset int) rune {
	pos := this.cursor.pos + offset
	if pos < len(this.chars) {
		return this.chars[pos]
	}
	return Eoi
}

func (this *m_Lexer) accept(c rune) {
	if this.cursor.c != c {
		this.panicByChar("expected char " + Character.CharDesc(c) + ", actual char " + Character.CharDesc(this.cursor.c))
	}
}

func (this *m_Lexer) panicByNeedChar(msg string) {
	this.panic(msg, this.cursor.pos, this.cursor.pos)
}

func (this *m_Lexer) panicByChar(msg string) {
	this.panic(msg, this.cursor.pos, this.cursor.pos+1)
}

func (this *m_Lexer) panicByToken(msg string) {
	this.panic(msg, this.cursor.tokenInfo.tokenBeginPos, this.cursor.pos)
}

func (this *m_Lexer) panic(msg string, beginPos, endPos int) {
	var builder strings.Builder
	if msg != "" {
		builder.WriteString(msg)
	}
	builder.WriteString("\n")

	if beginPos < len(this.chars) {
		if beginPos == endPos {
			for i := range this.chars {
				if i == beginPos {
					builder.WriteString("↪↩")
				}
				builder.WriteRune(this.chars[i])
			}
		} else {
			for i := range this.chars {
				if i == beginPos {
					builder.WriteString("↪")
				}
				builder.WriteRune(this.chars[i])
				if i == endPos-1 {
					builder.WriteString("↩")
				}
			}
		}
	} else {
		builder.WriteString(this.sql)
		builder.WriteString("↪↩")
	}
	panic(builder.String())
}

func extendLexer(i i_Lexer, sql string, reservedWords map[string]Token) *m_Lexer {
	l := &m_Lexer{I: i, sql: sql, chars: []rune(sql), reservedWords: reservedWords, cursor: newCursor()}
	l.nextChar()
	return l
}

// end of input
const Eoi rune = 0

// cursor 游标
type cursor struct {
	// 当前SQL语句定位的字符
	c rune
	// 当前SQL语句定位的字符位置
	pos int
	// Token信息
	tokenInfo *tokenInfo
	// 上一个Token信息
	prevTokenInfo *tokenInfo
}

// copy 深复制
func (c *cursor) copy() *cursor {
	return &cursor{
		c.c,
		c.pos,
		c.tokenInfo.copy(),
		c.prevTokenInfo.copy(),
	}
}

// tokenInfo Token信息
type tokenInfo struct {
	// Token
	token Token
	// Token值
	tokenVal string
	// Token值大写形式
	tokenValUpper string
	// Token开始位置
	tokenBeginPos int
	// Token结束位置（不包含）
	tokenEndPos int
	// Token是否保留字
	reserved bool
	// 其他属性，key为*attr[T]类型（由于Golang泛型太垃圾只能定为any类型），value不要使用指针类型（因为需要深复制）
	attributes map[any]any
}

// copy 深复制
func (t *tokenInfo) copy() *tokenInfo {
	a := make(map[any]any, len(t.attributes))
	for key, value := range t.attributes {
		a[key] = value
	}
	return &tokenInfo{
		t.token,
		t.tokenVal,
		t.tokenValUpper,
		t.tokenBeginPos,
		t.tokenEndPos,
		t.reserved,
		a,
	}
}

func (t *tokenInfo) reset() {
	t.token = Token{}
	t.tokenVal = ""
	t.tokenValUpper = ""
	t.tokenBeginPos = -1
	t.tokenEndPos = -1
	t.reserved = false
	for k := range t.attributes {
		delete(t.attributes, k)
	}
}

func newCursor() *cursor {
	return &cursor{pos: -1, tokenInfo: newTokenInfo(), prevTokenInfo: newTokenInfo()}
}

func newTokenInfo() *tokenInfo {
	t := &tokenInfo{attributes: make(map[any]any, 1)}
	t.reset()
	return t
}

// 词法器属性Key，可为不同数据库的实现的词法器增加数据库特性的属性
type attr[T any] struct{}

// set 设置词法器属性，值不要使用指针！！！
func (r *attr[T]) set(l *m_Lexer, v T) {
	l.cursor.tokenInfo.attributes[r] = v
}

// GetOfDefault 获取词法器属性值，若无属性则返回指定默认值
func (r *attr[T]) GetOfDefault(l *m_Lexer, defaultValue T) T {
	v := l.cursor.tokenInfo.attributes[r]
	if v != nil {
		return v.(T)
	}
	return defaultValue
}
