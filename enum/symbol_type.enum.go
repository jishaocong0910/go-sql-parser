package enum

import o "github.com/jishaocong0910/go-object"

// 操作符的符号类型
type SymbolType struct {
	*o.M_EnumValue
}

type _SymbolType struct {
	*o.M_Enum[SymbolType]
	// 标点（如：, . + -等）
	PUNCTUATION,
	// 标识符（如：AND、LIKE、IN）
	IDENTIFIER SymbolType
}

var SymbolTypes = o.NewEnum[SymbolType](_SymbolType{})
