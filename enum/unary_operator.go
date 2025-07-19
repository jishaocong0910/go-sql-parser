package enum

import o "github.com/jishaocong0910/go-object-util"

// 一元操作符
type UnaryOperator struct {
	*o.EnumElem__
	// 符号
	Symbol string
	// 符号类型
	SymbolType SymbolType
}

type _UnaryOperator struct {
	*o.Enum__[UnaryOperator]
	POSITIVE,
	NEGATIVE,
	COMPL,
	BINARY,
	NOT,
	NOTSTR UnaryOperator
}

// 这里搜集了所有数据库的操作符，不同数据库的语法解析器将使用各自方言具有的操作符
var UnaryOperator_ = o.NewEnum[UnaryOperator](_UnaryOperator{
	POSITIVE: UnaryOperator{nil, "+", SymbolType_.PUNCTUATION},
	NEGATIVE: UnaryOperator{nil, "-", SymbolType_.PUNCTUATION},
	COMPL:    UnaryOperator{nil, "~", SymbolType_.PUNCTUATION},
	BINARY:   UnaryOperator{nil, "BINARY", SymbolType_.IDENTIFIER},
	NOT:      UnaryOperator{nil, "!", SymbolType_.PUNCTUATION},
	NOTSTR:   UnaryOperator{nil, "NOT", SymbolType_.IDENTIFIER},
})
