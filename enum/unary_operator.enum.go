package enum

import o "github.com/jishaocong0910/go-object"

// 一元操作符
type UnaryOperator struct {
	*o.M_EnumValue
	// 符号
	Symbol string
	// 符号类型
	SymbolType SymbolType
}
