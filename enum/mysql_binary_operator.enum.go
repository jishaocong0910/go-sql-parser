package enum

import o "github.com/jishaocong0910/go-object"

type MySqlBinaryOperator struct {
	*o.M_EnumValue
	O BinaryOperator
	// 优先级，越小优先级越高
	Precedence int
}

type _MySqlBinaryOperator struct {
	*o.M_Enum[MySqlBinaryOperator]
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
	ASSIGN MySqlBinaryOperator
}

var MySqlBinaryOperators = o.NewEnum[MySqlBinaryOperator](_MySqlBinaryOperator{
	JSON_EXTRACT:                       MySqlBinaryOperator{nil, BinaryOperators.JSON_EXTRACT, 0},
	JSON_UNQUOTE:                       MySqlBinaryOperator{nil, BinaryOperators.JSON_UNQUOTE, 0},
	MEMBER_OF:                          MySqlBinaryOperator{nil, BinaryOperators.MEMBER_OF, 1},
	COLLATE:                            MySqlBinaryOperator{nil, BinaryOperators.COLLATE, 1},
	BITWISE_XOR:                        MySqlBinaryOperator{nil, BinaryOperators.BITWISE_XOR, 2},
	MULTIPLY:                           MySqlBinaryOperator{nil, BinaryOperators.MULTIPLY, 3},
	DIVIDE:                             MySqlBinaryOperator{nil, BinaryOperators.DIVIDE, 3},
	MODULUS:                            MySqlBinaryOperator{nil, BinaryOperators.MODULUS, 3},
	DIV:                                MySqlBinaryOperator{nil, BinaryOperators.DIV, 3},
	MOD:                                MySqlBinaryOperator{nil, BinaryOperators.MOD, 3},
	ADD:                                MySqlBinaryOperator{nil, BinaryOperators.ADD, 4},
	SUBTRACT:                           MySqlBinaryOperator{nil, BinaryOperators.SUBTRACT, 4},
	RIGHT_SHIFT:                        MySqlBinaryOperator{nil, BinaryOperators.RIGHT_SHIFT, 5},
	LEFT_SHIFT:                         MySqlBinaryOperator{nil, BinaryOperators.LEFT_SHIFT, 5},
	BITWISE_AND:                        MySqlBinaryOperator{nil, BinaryOperators.BITWISE_AND, 6},
	BITWISE_OR:                         MySqlBinaryOperator{nil, BinaryOperators.BITWISE_OR, 7},
	EQUAL_OR_ASSIGNMENT:                MySqlBinaryOperator{nil, BinaryOperators.EQUAL_OR_ASSIGNMENT, 8},
	LESS_THAN_OR_EQUAL_OR_GREATER_THAN: MySqlBinaryOperator{nil, BinaryOperators.LESS_THAN_OR_EQUAL_OR_GREATER_THAN, 8},
	GREATER_THAN:                       MySqlBinaryOperator{nil, BinaryOperators.GREATER_THAN, 8},
	LESS_THAN:                          MySqlBinaryOperator{nil, BinaryOperators.LESS_THAN, 8},
	GREATER_THAN_OR_EQUAL:              MySqlBinaryOperator{nil, BinaryOperators.GREATER_THAN_OR_EQUAL, 8},
	LESS_THAN_OR_EQUAL:                 MySqlBinaryOperator{nil, BinaryOperators.LESS_THAN_OR_EQUAL, 8},
	LESS_THAN_OR_GREATER:               MySqlBinaryOperator{nil, BinaryOperators.LESS_THAN_OR_GREATER, 8},
	NOT_EQUAL:                          MySqlBinaryOperator{nil, BinaryOperators.NOT_EQUAL, 8},
	IN:                                 MySqlBinaryOperator{nil, BinaryOperators.IN, 8},
	NOT_IN:                             MySqlBinaryOperator{nil, BinaryOperators.NOT_IN, 8},
	IS:                                 MySqlBinaryOperator{nil, BinaryOperators.IS, 8},
	IS_NOT:                             MySqlBinaryOperator{nil, BinaryOperators.IS_NOT, 8},
	LIKE:                               MySqlBinaryOperator{nil, BinaryOperators.LIKE, 8},
	NOT_LIKE:                           MySqlBinaryOperator{nil, BinaryOperators.NOT_LIKE, 8},
	SOUNDS_LIKE:                        MySqlBinaryOperator{nil, BinaryOperators.SOUNDS_LIKE, 8},
	REGEXP:                             MySqlBinaryOperator{nil, BinaryOperators.REGEXP, 8},
	NOT_REGEXP:                         MySqlBinaryOperator{nil, BinaryOperators.NOT_REGEXP, 8},
	RLIKE:                              MySqlBinaryOperator{nil, BinaryOperators.RLIKE, 8},
	NOT_RLIKE:                          MySqlBinaryOperator{nil, BinaryOperators.NOT_RLIKE, 8},
	BETWEEN:                            MySqlBinaryOperator{nil, BinaryOperators.BETWEEN, 9},
	NOT_BETWEEN:                        MySqlBinaryOperator{nil, BinaryOperators.NOT_BETWEEN, 9},
	BOOLEAN_AND:                        MySqlBinaryOperator{nil, BinaryOperators.BOOLEAN_AND, 10},
	BOOLEAN_AND2:                       MySqlBinaryOperator{nil, BinaryOperators.BOOLEAN_AND2, 10},
	BOOLEAN_XOR:                        MySqlBinaryOperator{nil, BinaryOperators.BOOLEAN_XOR, 11},
	BOOLEAN_OR:                         MySqlBinaryOperator{nil, BinaryOperators.BOOLEAN_OR, 12},
	BOOLEAN_OR2:                        MySqlBinaryOperator{nil, BinaryOperators.BOOLEAN_OR2, 12},
	ASSIGN:                             MySqlBinaryOperator{nil, BinaryOperators.ASSIGN, 13},
})
