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

type _UnaryOperator struct {
	*o.M_Enum[UnaryOperator]
	POSITIVE,
	NEGATIVE,
	COMPL,
	BINARY,
	NOT,
	NOTSTR UnaryOperator
}

// 这里搜集了所有数据库的操作符，不同数据库的语法解析器将使用各自方言具有的操作符
var UnaryOperators = o.NewEnum[UnaryOperator](_UnaryOperator{
	POSITIVE: UnaryOperator{nil, "+", SymbolTypes.PUNCTUATION},
	NEGATIVE: UnaryOperator{nil, "-", SymbolTypes.PUNCTUATION},
	COMPL:    UnaryOperator{nil, "~", SymbolTypes.PUNCTUATION},
	BINARY:   UnaryOperator{nil, "BINARY", SymbolTypes.IDENTIFIER},
	NOT:      UnaryOperator{nil, "!", SymbolTypes.PUNCTUATION},
	NOTSTR:   UnaryOperator{nil, "NOT", SymbolTypes.IDENTIFIER},
})
