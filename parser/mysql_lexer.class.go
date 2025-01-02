package parser

import (
	. "go-sql-parser/enum"
	"strings"
)

type mySqlLexer struct {
	*m_Lexer
	// 当前Token是标识符时，是否具有限定符『`』
	qualifier *attr[bool]
}

func (this *mySqlLexer) setHasQualifier() {
	this.qualifier.set(this.m_Lexer, true)
}

func (this *mySqlLexer) hasQualifier() bool {
	return this.qualifier.GetOfDefault(this.m_Lexer, false)
}

func (this *mySqlLexer) nextToken() Token {
	return this.nextTokenInner(false)
}

func (this *mySqlLexer) nextTokenIncludeComment() Token {
	return this.nextTokenInner(true)
}

func (this *mySqlLexer) nextTokenInner(includeComment bool) Token {
	this.beforeNextToken()
	for {
		this.setTokenBegin()
		if Character.IsFirstIdentifierChar(this.char()) {
			return this.nextIdentifier()
		} else if Character.IsWhitespaceChar(this.char()) {
			if this.char() == eoi {
				return this.nextEoi()
			}
			this.nextChar()
			continue
		} else {
			switch this.char() {
			case '`':
				return this.nextIdentifier()
			case '0':
				c := this.previewChar(1)
				if c == 'b' {
					return this.nextBinaryDecimal()
				} else if c == 'x' {
					return this.nextHexDecimal()
				}
				fallthrough
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				return this.nextDecimalNumber()
			case '\'', '"':
				return this.nextString()
			case ',':
				return this.nextLiteralToken(Tokens.COMMA)
			case '.':
				return this.nextLiteralToken(Tokens.DOT)
			case '(':
				return this.nextLiteralToken(Tokens.L_PAREN)
			case ')':
				return this.nextLiteralToken(Tokens.R_PAREN)
			case '=':
				return this.nextLiteralToken(Tokens.EQ)
			case ';':
				return this.nextLiteralToken(Tokens.SEMI)
			case '+':
				return this.nextLiteralToken(Tokens.PLUS)
			case '*':
				return this.nextLiteralToken(Tokens.STAR)
			case '~':
				return this.nextLiteralToken(Tokens.TILDE)
			case '%':
				return this.nextLiteralToken(Tokens.PERCENT)
			case '^':
				return this.nextLiteralToken(Tokens.CARET)
			case '?':
				return this.nextLiteralToken(Tokens.QUES)
			case '>':
				c := this.previewChar(1)
				if c == '>' {
					this.nextChar()
					return this.nextLiteralToken(Tokens.GT_GT)
				} else if c == '=' {
					this.nextChar()
					return this.nextLiteralToken(Tokens.GT_EQ)
				} else {
					return this.nextLiteralToken(Tokens.GT)
				}
			case '<':
				c := this.previewChar(1)
				if c == '>' {
					this.nextChar()
					return this.nextLiteralToken(Tokens.LT_GT)
				} else if c == '<' {
					this.nextChar()
					return this.nextLiteralToken(Tokens.LT_LT)
				} else if c == '=' {
					this.nextChar()
					c = this.previewChar(1)
					if c == '>' {
						this.nextChar()
						return this.nextLiteralToken(Tokens.LT_EQ_GT)
					} else {
						return this.nextLiteralToken(Tokens.LT_EQ)
					}
				} else {
					return this.nextLiteralToken(Tokens.LT)
				}
			case ':':
				this.nextChar()
				this.accept('=')
				return this.nextLiteralToken(Tokens.COLON_EQ)
			case '!':
				if this.previewChar(1) == '=' {
					this.nextChar()
					return this.nextLiteralToken(Tokens.BANG_EQ)
				} else {
					return this.nextLiteralToken(Tokens.BANG)
				}
			case '@':
				return this.nextVariable()
			case '#':
				if includeComment {
					return this.nextComment()
				} else {
					this.nextComment()
					continue
				}
			case '&':
				if this.previewChar(1) == '&' {
					this.nextChar()
					return this.nextLiteralToken(Tokens.AMP_AMP)
				} else {
					return this.nextLiteralToken(Tokens.AMP)
				}
			case '|':
				if this.previewChar(1) == '|' {
					this.nextChar()
					return this.nextLiteralToken(Tokens.BAR_BAR)
				} else {
					return this.nextLiteralToken(Tokens.BAR)
				}
			case '-':
				c := this.previewChar(1)
				if c == '-' {
					if includeComment {
						return this.nextComment()
					} else {
						this.nextComment()
						continue
					}
				} else if c == '>' {
					this.nextChar()
					if this.previewChar(1) == '>' {
						this.nextChar()
						return this.nextLiteralToken(Tokens.JSON_UNQUOTE)
					} else {
						return this.nextLiteralToken(Tokens.JSON_EXTRACT)
					}
				} else {
					return this.nextLiteralToken(Tokens.SUB)
				}
			case '/':
				if this.previewChar(1) == '*' {
					if includeComment {
						return this.nextComment()
					} else {
						this.nextComment()
						continue
					}
				}
				return this.nextLiteralToken(Tokens.SLASH)
			default:
				this.panicByChar("illegal character " + Character.CharDesc(this.char()))
			}
		}
	}
}

func (this *mySqlLexer) nextIdentifier() Token {
	if this.char() == '`' {
		this.setHasQualifier()
		this.nextChar()
		for {
			if this.char() == '`' {
				break
			}
			if this.char() == eoi || !Character.IsIdentifierChar(this.char()) {
				this.panicByNeedChar("need character ` to finish")
			}
			this.nextChar()
		}
		this.nextChar()
		this.setTokenIdentifier()
	} else {
		for {
			this.nextChar()
			if !Character.IsIdentifierChar(this.char()) {
				this.setTokenIdentifier()
				break
			}
		}
	}
	return this.token()
}

func (this *mySqlLexer) nextLiteralToken(token Token) Token {
	this.nextChar()
	this.setToken(token)
	return token
}

func (this *mySqlLexer) nextDecimalNumber() Token {
	for {
		c := this.nextChar()
		if '0' <= c && c <= '9' {
			continue
		}
		break
	}

	if this.char() == '.' {
		for {
			c := this.nextChar()
			if '0' <= c && c <= '9' {
				continue
			}
			break
		}
	}

	if c := this.char(); c == 'e' || c == 'E' {
		c = this.nextChar()
		if c == '+' || c == '-' {
			c = this.nextChar()
		}
		for {
			if '0' <= c && c <= '9' {
				c = this.nextChar()
				continue
			}
			break
		}
	}

	this.setToken(Tokens.DECIMAL_NUMBER)
	return this.token()
}

func (this *mySqlLexer) nextBinaryDecimal() Token {
	this.nextChar()
	valid := false
	for {
		if !Character.IsBinaryChar(this.nextChar()) {
			break
		}
		valid = true
	}
	if !valid {
		this.panicByToken("invalid binary number")
	}
	this.setToken(Tokens.BINARY_NUMBER)
	return this.token()

}

func (this *mySqlLexer) nextHexDecimal() Token {
	this.nextChar()
	valid := false
	for {
		if !Character.IsHexChar(this.nextChar()) {
			break
		}
		valid = true
	}
	if !valid {
		this.panicByToken("invalid hexadecimal number")
	}
	this.setToken(Tokens.HEXADECIMAL_NUMBER)
	return this.token()
}

func (this *mySqlLexer) nextString() Token {
	quote := this.char()
	for {
		c := this.nextChar()
		if c == quote && this.previewChar(1) != '\\' {
			this.nextChar()
			this.setToken(Tokens.STRING)
			return this.token()
		}
		if c == eoi {
			this.panicByNeedChar("need character '" + string(quote) + "' to finish")
		}
	}
}

func (this *mySqlLexer) nextVariable() Token {
	c := this.nextChar()
	var t Token
	if c != '@' {
		t = Tokens.AT
	} else {
		t = Tokens.AT_AT
		c = this.nextChar()
	}

	var quote rune
	if c == '\'' || c == '"' {
		quote = c
	}

	if quote != 0 {
		for {
			c = this.nextChar()
			if c == quote {
				this.nextChar()
				break
			}
			if c == eoi {
				this.panicByNeedChar("need character " + string(quote) + " to finish")
			}
		}
	} else {
		for {
			if Character.IsIdentifierChar(c) {
				c = this.nextChar()
				continue
			}
			break
		}
		if c == '.' {
			for {
				c = this.nextChar()
				if Character.IsIdentifierChar(c) {
					continue
				}
				break
			}
		}
	}
	this.setToken(t)
	return this.token()

}

func (this *mySqlLexer) nextComment() Token {
	c := this.char()
	if c == '#' {
		for {
			c = this.nextChar()
			if c == '\n' || c == eoi {
				break
			}
		}
	} else if c == '-' && this.previewChar(1) == '-' && this.previewChar(2) == ' ' {
		this.nextChar()
		this.nextChar()
		for {
			c = this.nextChar()
			if c == '\n' || c == eoi {
				break
			}
		}
	} else if c == '/' && this.previewChar(1) == '*' {
		this.nextChar()
		for {
			c = this.nextChar()
			if c == '*' {
				if this.previewChar(1) == '/' {
					this.nextChar()
					this.nextChar()
					break
				}
			}
			if c == eoi {
				this.panicByNeedChar("need string '*/' to finish")
			}
		}
	}
	this.setToken(Tokens.COMMENT)
	return this.token()
}

func (this *mySqlLexer) setTokenIdentifier() {
	var tokenVal string
	if this.hasQualifier() {
		tokenVal = string(this.chars[this.cursor.tokenInfo.tokenBeginPos+1 : this.cursor.pos-1])
		this.cursor.tokenInfo.token = Tokens.IDENTIFIER
	} else {
		tokenVal = string(this.chars[this.cursor.tokenInfo.tokenBeginPos:this.cursor.pos])
		if rw, ok := this.reservedWords[strings.ToUpper(tokenVal)]; ok {
			this.cursor.tokenInfo.token = rw
			this.cursor.tokenInfo.reserved = true
		} else {
			this.cursor.tokenInfo.token = Tokens.IDENTIFIER
		}
	}
	this.cursor.tokenInfo.tokenVal = tokenVal
	this.setTokenEnd()
}

func newMySqlLexer(sql string) *mySqlLexer {
	l := &mySqlLexer{qualifier: &attr[bool]{}}
	l.m_Lexer = extendLexer(l, sql, MYSQL_RESERVED_WORDS)
	return l
}
