package parser

import "strconv"

var (
	// 标识符第一个识别字符集
	firstIdentifierChars = func() [256]bool {
		var arr [256]bool
		for c := range arr {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
				arr[c] = true
			}
		}
		arr['_'] = true
		return arr
	}()
	// 标识符字符集
	identifierChars = func() [256]bool {
		var arr [256]bool
		for c := range arr {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
				arr[c] = true
			}
		}
		arr['_'] = true
		return arr
	}()
	// 空白字符集
	whitespaceChars = func() [256]bool {
		var arr [256]bool
		for c := 0; c <= 32; c++ {
			arr[c] = true
		}
		for c := 0x7F; c <= 0xA0; c++ {
			arr[c] = true
		}
		arr[160] = true
		return arr
	}()
	// 16进制数字符集
	hexChars = func() [256]bool {
		var arr [256]bool
		for c := range arr {
			if (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') || (c >= '0' && c <= '9') {
				arr[c] = true
			}
		}
		return arr
	}()

	Character character
)

type character struct{}

func (character) IsFirstIdentifierChar(c rune) bool {
	if int(c) <= len(firstIdentifierChars) {
		return firstIdentifierChars[c]
	}
	// 圆角空格和中文逗号
	return c != 12288 && c != '，'
}

func (character) IsIdentifierChar(c rune) bool {
	if int(c) <= len(identifierChars) {
		return identifierChars[c]
	}
	// 圆角空格
	return c != 12288
}

func (character) IsWhitespaceChar(c rune) bool {
	if int(c) <= len(whitespaceChars) {
		return whitespaceChars[c]
	}
	// 圆角空格
	return c == 12288
}

func (character) IsHexChar(c rune) bool {
	return int(c) < len(hexChars) && hexChars[c]
}

func (character) IsBinaryChar(c rune) bool {
	return c == '0' || c == '1'
}

func (character) CharDesc(c rune) string {
	if c < 32 {
		if c == 0 {
			return "(EOI)"
		}
		return "(ASCII: " + strconv.Itoa(int(c)) + ")"
	}
	switch c {
	case ' ':
		return "(SPACE)"
	case ',':
		return "(COMMA)"
	case '\'':
		return "(SINGLE QUOTE)"
	case '"':
		return "(DOUBLE QUOTE)"
	case 160:
		return "(NON-BREAKING SPACE)"
	default:
		return "'" + string(c) + "'"
	}
}
