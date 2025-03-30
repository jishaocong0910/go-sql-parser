package enum

import o "github.com/jishaocong0910/go-object"

// 二元操作符
type BinaryOperator struct {
	*o.M_EnumValue
	// 符号
	Symbol string
	// 操作符类型
	OperatorType OperatorType
	// 符号类型
	SymbolType SymbolType
	// 是否允许多个操作数
	AllowMultipleOperand bool
}

type _BinaryOperator struct {
	*o.M_Enum[BinaryOperator]
	JSON_EXTRACT,
	JSON_UNQUOTE,
	MEMBER_OF,
	COLLATE,
	BITWISE_XOR,
	MULTIPLY,
	DIVIDE,
	MODULUS,
	DIV,
	MOD,
	ADD,
	SUBTRACT,
	RIGHT_SHIFT,
	LEFT_SHIFT,
	BITWISE_AND,
	BITWISE_OR,
	EQUAL_OR_ASSIGNMENT,
	LESS_THAN_OR_EQUAL_OR_GREATER_THAN,
	GREATER_THAN,
	LESS_THAN,
	GREATER_THAN_OR_EQUAL,
	LESS_THAN_OR_EQUAL,
	LESS_THAN_OR_GREATER,
	NOT_EQUAL,
	IN,
	NOT_IN,
	IS,
	IS_NOT,
	LIKE,
	NOT_LIKE,
	SOUNDS_LIKE,
	REGEXP,
	NOT_REGEXP,
	RLIKE,
	NOT_RLIKE,
	BETWEEN,
	NOT_BETWEEN,
	BOOLEAN_AND,
	BOOLEAN_AND2,
	BOOLEAN_XOR,
	BOOLEAN_OR,
	BOOLEAN_OR2,
	ASSIGN BinaryOperator
}

// 这里搜集了所有数据库的操作符，不同数据库的语法解析器将使用各自方言具有的操作符
var BinaryOperators = o.NewEnum[BinaryOperator](_BinaryOperator{
	JSON_EXTRACT:                       BinaryOperator{nil, "->", OperatorTypes.SPECIAL, SymbolTypes.PUNCTUATION, false},
	JSON_UNQUOTE:                       BinaryOperator{nil, "->>", OperatorTypes.SPECIAL, SymbolTypes.PUNCTUATION, false},
	MEMBER_OF:                          BinaryOperator{nil, "MEMBER OF", OperatorTypes.SPECIAL, SymbolTypes.IDENTIFIER, false},
	COLLATE:                            BinaryOperator{nil, "COLLATE", OperatorTypes.SPECIAL, SymbolTypes.IDENTIFIER, false},
	BITWISE_XOR:                        BinaryOperator{nil, "^", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	MULTIPLY:                           BinaryOperator{nil, "*", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	DIVIDE:                             BinaryOperator{nil, "/", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	MODULUS:                            BinaryOperator{nil, "%", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	DIV:                                BinaryOperator{nil, "DIV", OperatorTypes.BITWISE, SymbolTypes.IDENTIFIER, false},
	MOD:                                BinaryOperator{nil, "MOD", OperatorTypes.BITWISE, SymbolTypes.IDENTIFIER, false},
	ADD:                                BinaryOperator{nil, "+", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	SUBTRACT:                           BinaryOperator{nil, "-", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	RIGHT_SHIFT:                        BinaryOperator{nil, ">>", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	LEFT_SHIFT:                         BinaryOperator{nil, "<<", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	BITWISE_AND:                        BinaryOperator{nil, "&", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	BITWISE_OR:                         BinaryOperator{nil, "|", OperatorTypes.BITWISE, SymbolTypes.PUNCTUATION, false},
	EQUAL_OR_ASSIGNMENT:                BinaryOperator{nil, "=", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	LESS_THAN_OR_EQUAL_OR_GREATER_THAN: BinaryOperator{nil, "<=>", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	GREATER_THAN:                       BinaryOperator{nil, ">", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	LESS_THAN:                          BinaryOperator{nil, "<", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	GREATER_THAN_OR_EQUAL:              BinaryOperator{nil, ">=", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	LESS_THAN_OR_EQUAL:                 BinaryOperator{nil, "<=", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	LESS_THAN_OR_GREATER:               BinaryOperator{nil, "<>", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	NOT_EQUAL:                          BinaryOperator{nil, "!=", OperatorTypes.COMPARISON, SymbolTypes.PUNCTUATION, true},
	IN:                                 BinaryOperator{nil, "IN", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, true},
	NOT_IN:                             BinaryOperator{nil, "NOT IN", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, true},
	IS:                                 BinaryOperator{nil, "IS", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	IS_NOT:                             BinaryOperator{nil, "IS NOT", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	LIKE:                               BinaryOperator{nil, "LIKE", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	NOT_LIKE:                           BinaryOperator{nil, "NOT LIKE", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	SOUNDS_LIKE:                        BinaryOperator{nil, "SOUNDS LIKE", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	REGEXP:                             BinaryOperator{nil, "REGEXP", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	NOT_REGEXP:                         BinaryOperator{nil, "NOT REGEXP", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	RLIKE:                              BinaryOperator{nil, "RLIKE", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	NOT_RLIKE:                          BinaryOperator{nil, "NOT RLIKE", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	BETWEEN:                            BinaryOperator{nil, "BETWEEN", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	NOT_BETWEEN:                        BinaryOperator{nil, "NOT BETWEEN", OperatorTypes.PREDICATE, SymbolTypes.IDENTIFIER, false},
	BOOLEAN_AND:                        BinaryOperator{nil, "AND", OperatorTypes.LOGICAL, SymbolTypes.IDENTIFIER, false},
	BOOLEAN_AND2:                       BinaryOperator{nil, "&&", OperatorTypes.LOGICAL, SymbolTypes.PUNCTUATION, false},
	BOOLEAN_XOR:                        BinaryOperator{nil, "XOR", OperatorTypes.LOGICAL, SymbolTypes.IDENTIFIER, false},
	BOOLEAN_OR:                         BinaryOperator{nil, "OR", OperatorTypes.LOGICAL, SymbolTypes.IDENTIFIER, false},
	BOOLEAN_OR2:                        BinaryOperator{nil, "||", OperatorTypes.LOGICAL, SymbolTypes.PUNCTUATION, false},
	ASSIGN:                             BinaryOperator{nil, ":=", OperatorTypes.ASSIGNMENT, SymbolTypes.PUNCTUATION, false},
})
