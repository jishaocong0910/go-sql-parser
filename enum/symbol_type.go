package enum

import o "github.com/jishaocong0910/go-object-util"

// 操作符的符号类型
type SymbolType struct {
	*o.EnumElem__
}

type _SymbolType struct {
	*o.Enum__[SymbolType]
	// 标点（如：, . + -等）
	PUNCTUATION,
	// 标识符（如：AND、LIKE、IN）
	IDENTIFIER SymbolType
}

var SymbolType_ = o.NewEnum[SymbolType](_SymbolType{})
