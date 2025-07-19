package parser

import (
	"strings"

	. "github.com/jishaocong0910/go-sql-parser/enum"
)

type lexer_ interface {
	lexer_() *lexer__
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

type lexer__ struct {
	i             lexer_
	sql           string
	chars         []rune
	reservedWords map[string]Token
	cursor        *cursor
}

func (this *lexer__) lexer_() *lexer__ {
	return this
}

func (this *lexer__) reserved() bool {
	return this.cursor.tokenInfo.reserved
}

func (this *lexer__) token() Token {
	return this.cursor.tokenInfo.token
}

func (this *lexer__) tokenVal() string {
	return this.cursor.tokenInfo.tokenVal
}

func (this *lexer__) tokenValUpper() string {
	if this.cursor.tokenInfo.tokenValUpper == "" {
		this.cursor.tokenInfo.tokenValUpper = strings.ToUpper(this.cursor.tokenInfo.tokenVal)
	}
	return this.cursor.tokenInfo.tokenValUpper
}

func (this *lexer__) char() rune {
	return this.cursor.c
}

func (this *lexer__) setTokenBegin() {
	this.cursor.tokenInfo.tokenBeginPos = this.cursor.pos
}

func (this *lexer__) tokenBeginPos() int {
	return this.cursor.tokenInfo.tokenBeginPos
}

func (this *lexer__) tokenEndPos() int {
	return this.cursor.tokenInfo.tokenEndPos
}

func (this *lexer__) prevToken() Token {
	return this.cursor.prevTokenInfo.token
}

func (this *lexer__) prevTokenBeginPos() int {
	return this.cursor.prevTokenInfo.tokenBeginPos
}

func (this *lexer__) prevTokenEndPos() int {
	return this.cursor.prevTokenInfo.tokenEndPos
}

func (this *lexer__) saveCursor() *cursor {
	return this.cursor.copy()
}

func (this *lexer__) rollback(c *cursor) {
	this.cursor = c
}

func (this *lexer__) nextEoi() Token {
	this.cursor.tokenInfo.token = Token_.EOI
	this.cursor.tokenInfo.tokenVal = string(Eoi)
	this.setTokenEnd()
	return Token_.EOI
}

func (this *lexer__) setToken(token Token) {
	this.cursor.tokenInfo.token = token
	this.cursor.tokenInfo.tokenVal = string(this.chars[this.cursor.tokenInfo.tokenBeginPos:this.cursor.pos])
	this.setTokenEnd()
}

func (this *lexer__) setTokenEnd() {
	this.cursor.tokenInfo.tokenEndPos = this.cursor.pos
}

func (this *lexer__) beforeNextToken() {
	if Token_.Is(this.cursor.tokenInfo.token, Token_.EOI) {
		return
	}
	this.cursor.tokenInfo, this.cursor.prevTokenInfo = this.cursor.prevTokenInfo, this.cursor.tokenInfo
	this.cursor.tokenInfo.reset()
}

func (this *lexer__) nextChar() rune {
	if this.cursor.pos == len(this.chars) {
		return this.cursor.c
	}
	this.cursor.pos++
	if this.cursor.pos == len(this.chars) {
		this.cursor.c = Eoi
		return this.cursor.c
	}
	this.cursor.c = this.chars[this.cursor.pos]
	// 非末尾但为EOI字符则继续下个字符
	if this.cursor.c == Eoi {
		return this.nextChar()
	} else {
		return this.cursor.c
	}
}

func (this *lexer__) previewChar(offset int) rune {
	pos := this.cursor.pos + offset
	if pos < len(this.chars) {
		return this.chars[pos]
	}
	return Eoi
}

func (this *lexer__) accept(c rune) {
	if this.cursor.c != c {
		this.panicByChar("expected char " + Character.CharDesc(c) + ", actual char " + Character.CharDesc(this.cursor.c))
	}
}

func (this *lexer__) panicByNeedChar(msg string) {
	this.panic(msg, this.cursor.pos, this.cursor.pos)
}

func (this *lexer__) panicByChar(msg string) {
	this.panic(msg, this.cursor.pos, this.cursor.pos+1)
}

func (this *lexer__) panicByToken(msg string) {
	this.panic(msg, this.cursor.tokenInfo.tokenBeginPos, this.cursor.pos)
}

func (this *lexer__) panic(msg string, beginPos, endPos int) {
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

func extendLexer(i lexer_, sql string, reservedWords map[string]Token) *lexer__ {
	l := &lexer__{i: i, sql: sql, chars: []rune(sql), reservedWords: reservedWords, cursor: newCursor()}
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
func (r *attr[T]) set(l *lexer__, v T) {
	l.cursor.tokenInfo.attributes[r] = v
}

// GetOfDefault 获取词法器属性值，若无属性则返回指定默认值
func (r *attr[T]) GetOfDefault(l *lexer__, defaultValue T) T {
	v := l.cursor.tokenInfo.attributes[r]
	if v != nil {
		return v.(T)
	}
	return defaultValue
}
