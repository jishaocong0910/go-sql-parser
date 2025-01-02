package enum

import o "github.com/jishaocong0910/go-object"

type _MySqlUnaryOperator struct {
	*o.M_Enum[UnaryOperator]
	POSITIVE,
	NEGATIVE,
	COMPL,
	BINARY,
	NOT,
	NOTSTR UnaryOperator
}

var MySqlUnaryOperators = o.NewEnum[UnaryOperator](_MySqlUnaryOperator{
	POSITIVE: UnaryOperator{nil, "+", SymbolTypes.PUNCTUATION},
	NEGATIVE: UnaryOperator{nil, "-", SymbolTypes.PUNCTUATION},
	COMPL:    UnaryOperator{nil, "~", SymbolTypes.PUNCTUATION},
	BINARY:   UnaryOperator{nil, "BINARY", SymbolTypes.IDENTIFIER},
	NOT:      UnaryOperator{nil, "!", SymbolTypes.PUNCTUATION},
	NOTSTR:   UnaryOperator{nil, "NOT", SymbolTypes.IDENTIFIER},
})
