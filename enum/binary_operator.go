package enum

import o "github.com/jishaocong0910/go-object-util"

// 二元操作符
type BinaryOperator struct {
	*o.EnumElem__
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
	*o.Enum__[BinaryOperator]
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
var BinaryOperator_ = o.NewEnum[BinaryOperator](_BinaryOperator{
	JSON_EXTRACT:                       BinaryOperator{nil, "->", OperatorType_.SPECIAL, SymbolType_.PUNCTUATION, false},
	JSON_UNQUOTE:                       BinaryOperator{nil, "->>", OperatorType_.SPECIAL, SymbolType_.PUNCTUATION, false},
	MEMBER_OF:                          BinaryOperator{nil, "MEMBER OF", OperatorType_.SPECIAL, SymbolType_.IDENTIFIER, false},
	COLLATE:                            BinaryOperator{nil, "COLLATE", OperatorType_.SPECIAL, SymbolType_.IDENTIFIER, false},
	BITWISE_XOR:                        BinaryOperator{nil, "^", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	MULTIPLY:                           BinaryOperator{nil, "*", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	DIVIDE:                             BinaryOperator{nil, "/", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	MODULUS:                            BinaryOperator{nil, "%", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	DIV:                                BinaryOperator{nil, "DIV", OperatorType_.BITWISE, SymbolType_.IDENTIFIER, false},
	MOD:                                BinaryOperator{nil, "MOD", OperatorType_.BITWISE, SymbolType_.IDENTIFIER, false},
	ADD:                                BinaryOperator{nil, "+", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	SUBTRACT:                           BinaryOperator{nil, "-", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	RIGHT_SHIFT:                        BinaryOperator{nil, ">>", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	LEFT_SHIFT:                         BinaryOperator{nil, "<<", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	BITWISE_AND:                        BinaryOperator{nil, "&", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	BITWISE_OR:                         BinaryOperator{nil, "|", OperatorType_.BITWISE, SymbolType_.PUNCTUATION, false},
	EQUAL_OR_ASSIGNMENT:                BinaryOperator{nil, "=", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	LESS_THAN_OR_EQUAL_OR_GREATER_THAN: BinaryOperator{nil, "<=>", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	GREATER_THAN:                       BinaryOperator{nil, ">", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	LESS_THAN:                          BinaryOperator{nil, "<", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	GREATER_THAN_OR_EQUAL:              BinaryOperator{nil, ">=", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	LESS_THAN_OR_EQUAL:                 BinaryOperator{nil, "<=", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	LESS_THAN_OR_GREATER:               BinaryOperator{nil, "<>", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	NOT_EQUAL:                          BinaryOperator{nil, "!=", OperatorType_.COMPARISON, SymbolType_.PUNCTUATION, true},
	IN:                                 BinaryOperator{nil, "IN", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, true},
	NOT_IN:                             BinaryOperator{nil, "NOT IN", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, true},
	IS:                                 BinaryOperator{nil, "IS", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	IS_NOT:                             BinaryOperator{nil, "IS NOT", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	LIKE:                               BinaryOperator{nil, "LIKE", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	NOT_LIKE:                           BinaryOperator{nil, "NOT LIKE", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	SOUNDS_LIKE:                        BinaryOperator{nil, "SOUNDS LIKE", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	REGEXP:                             BinaryOperator{nil, "REGEXP", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	NOT_REGEXP:                         BinaryOperator{nil, "NOT REGEXP", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	RLIKE:                              BinaryOperator{nil, "RLIKE", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	NOT_RLIKE:                          BinaryOperator{nil, "NOT RLIKE", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	BETWEEN:                            BinaryOperator{nil, "BETWEEN", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	NOT_BETWEEN:                        BinaryOperator{nil, "NOT BETWEEN", OperatorType_.PREDICATE, SymbolType_.IDENTIFIER, false},
	BOOLEAN_AND:                        BinaryOperator{nil, "AND", OperatorType_.LOGICAL, SymbolType_.IDENTIFIER, false},
	BOOLEAN_AND2:                       BinaryOperator{nil, "&&", OperatorType_.LOGICAL, SymbolType_.PUNCTUATION, false},
	BOOLEAN_XOR:                        BinaryOperator{nil, "XOR", OperatorType_.LOGICAL, SymbolType_.IDENTIFIER, false},
	BOOLEAN_OR:                         BinaryOperator{nil, "OR", OperatorType_.LOGICAL, SymbolType_.IDENTIFIER, false},
	BOOLEAN_OR2:                        BinaryOperator{nil, "||", OperatorType_.LOGICAL, SymbolType_.PUNCTUATION, false},
	ASSIGN:                             BinaryOperator{nil, ":=", OperatorType_.ASSIGNMENT, SymbolType_.PUNCTUATION, false},
})
