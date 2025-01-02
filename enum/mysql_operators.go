package enum

import o "github.com/jishaocong0910/go-object"

var MYSQL_LF_BINARY_OPERATORS = o.NewSet[BinaryOperator](
	MySqlBinaryOperators.BITWISE_AND,
	MySqlBinaryOperators.BOOLEAN_AND2,
	MySqlBinaryOperators.BITWISE_OR,
	MySqlBinaryOperators.BOOLEAN_OR2,
	MySqlBinaryOperators.BOOLEAN_XOR,
)

var MYSQL_EXPR_LEVEL_TO_BINARY_OPERATORS = func() map[ExprSyntaxLevel]*o.Set[BinaryOperator] {
	calculation := o.NewSet[BinaryOperator]()
	calculation.Add(MySqlBinaryOperators.MULTIPLY)
	calculation.Add(MySqlBinaryOperators.ADD)
	calculation.Add(MySqlBinaryOperators.SUBTRACT)
	calculation.Add(MySqlBinaryOperators.DIVIDE)
	calculation.Add(MySqlBinaryOperators.MODULUS)
	calculation.Add(MySqlBinaryOperators.DIV)
	calculation.Add(MySqlBinaryOperators.MOD)
	calculation.Add(MySqlBinaryOperators.BITWISE_XOR)
	calculation.Add(MySqlBinaryOperators.BITWISE_AND)
	calculation.Add(MySqlBinaryOperators.BITWISE_OR)
	calculation.Add(MySqlBinaryOperators.RIGHT_SHIFT)
	calculation.Add(MySqlBinaryOperators.LEFT_SHIFT)
	calculation.Add(MySqlBinaryOperators.COLLATE)
	calculation.Add(MySqlBinaryOperators.JSON_EXTRACT)
	calculation.Add(MySqlBinaryOperators.JSON_UNQUOTE)
	calculation.Add(MySqlBinaryOperators.MEMBER_OF)

	booleanPredicate := o.NewSet[BinaryOperator]()
	booleanPredicate.AddSet(calculation)
	booleanPredicate.Add(MySqlBinaryOperators.IN)
	booleanPredicate.Add(MySqlBinaryOperators.NOT_IN)
	booleanPredicate.Add(MySqlBinaryOperators.IS)
	booleanPredicate.Add(MySqlBinaryOperators.IS_NOT)
	booleanPredicate.Add(MySqlBinaryOperators.LIKE)
	booleanPredicate.Add(MySqlBinaryOperators.NOT_LIKE)
	booleanPredicate.Add(MySqlBinaryOperators.REGEXP)
	booleanPredicate.Add(MySqlBinaryOperators.NOT_REGEXP)
	booleanPredicate.Add(MySqlBinaryOperators.RLIKE)
	booleanPredicate.Add(MySqlBinaryOperators.NOT_RLIKE)
	booleanPredicate.Add(MySqlBinaryOperators.BETWEEN)
	booleanPredicate.Add(MySqlBinaryOperators.NOT_BETWEEN)
	booleanPredicate.Add(MySqlBinaryOperators.SOUNDS_LIKE)

	booleanPrimary := o.NewSet[BinaryOperator]()
	booleanPrimary.AddSet(booleanPredicate)
	booleanPrimary.Add(MySqlBinaryOperators.EQUAL_OR_ASSIGNMENT)
	booleanPrimary.Add(MySqlBinaryOperators.GREATER_THAN)
	booleanPrimary.Add(MySqlBinaryOperators.LESS_THAN)
	booleanPrimary.Add(MySqlBinaryOperators.GREATER_THAN_OR_EQUAL)
	booleanPrimary.Add(MySqlBinaryOperators.LESS_THAN_OR_EQUAL)
	booleanPrimary.Add(MySqlBinaryOperators.LESS_THAN_OR_GREATER)
	booleanPrimary.Add(MySqlBinaryOperators.NOT_EQUAL)
	booleanPrimary.Add(MySqlBinaryOperators.LESS_THAN_OR_EQUAL_OR_GREATER_THAN)

	booleanLogical := o.NewSet[BinaryOperator]()
	booleanLogical.AddSet(booleanPrimary)
	booleanLogical.Add(MySqlBinaryOperators.BOOLEAN_AND)
	booleanLogical.Add(MySqlBinaryOperators.BOOLEAN_OR)
	booleanLogical.Add(MySqlBinaryOperators.BOOLEAN_AND2)
	booleanLogical.Add(MySqlBinaryOperators.BOOLEAN_OR2)
	booleanLogical.Add(MySqlBinaryOperators.BOOLEAN_XOR)

	return map[ExprSyntaxLevel]*o.Set[BinaryOperator]{
		ExprSyntaxLevels.CALCULATION:       calculation,
		ExprSyntaxLevels.BOOLEAN_PREDICATE: booleanPredicate,
		ExprSyntaxLevels.BOOLEAN_PRIMARY:   booleanPrimary,
		ExprSyntaxLevels.EXPR:              booleanLogical,
	}
}()

var MYSQL_TOKEN_TO_BINARY_OPERATORS = map[Token]BinaryOperator{
	Tokens.CARET:        MySqlBinaryOperators.BITWISE_XOR,
	Tokens.STAR:         MySqlBinaryOperators.MULTIPLY,
	Tokens.SLASH:        MySqlBinaryOperators.DIVIDE,
	Tokens.PERCENT:      MySqlBinaryOperators.MODULUS,
	Tokens.SUB:          MySqlBinaryOperators.SUBTRACT,
	Tokens.PLUS:         MySqlBinaryOperators.ADD,
	Tokens.LT_LT:        MySqlBinaryOperators.LEFT_SHIFT,
	Tokens.GT_GT:        MySqlBinaryOperators.RIGHT_SHIFT,
	Tokens.AMP:          MySqlBinaryOperators.BITWISE_AND,
	Tokens.BAR:          MySqlBinaryOperators.BITWISE_OR,
	Tokens.EQ:           MySqlBinaryOperators.EQUAL_OR_ASSIGNMENT,
	Tokens.LT_EQ_GT:     MySqlBinaryOperators.LESS_THAN_OR_EQUAL_OR_GREATER_THAN,
	Tokens.GT_EQ:        MySqlBinaryOperators.GREATER_THAN_OR_EQUAL,
	Tokens.GT:           MySqlBinaryOperators.GREATER_THAN,
	Tokens.LT:           MySqlBinaryOperators.LESS_THAN,
	Tokens.LT_EQ:        MySqlBinaryOperators.LESS_THAN_OR_EQUAL,
	Tokens.LT_GT:        MySqlBinaryOperators.LESS_THAN_OR_GREATER,
	Tokens.BANG_EQ:      MySqlBinaryOperators.NOT_EQUAL,
	Tokens.AMP_AMP:      MySqlBinaryOperators.BOOLEAN_AND2,
	Tokens.BAR_BAR:      MySqlBinaryOperators.BOOLEAN_OR2,
	Tokens.COLON_EQ:     MySqlBinaryOperators.ASSIGN,
	Tokens.REGEXP:       MySqlBinaryOperators.REGEXP,
	Tokens.R_LIKE:       MySqlBinaryOperators.RLIKE,
	Tokens.DIV:          MySqlBinaryOperators.DIV,
	Tokens.JSON_EXTRACT: MySqlBinaryOperators.JSON_EXTRACT,
	Tokens.JSON_UNQUOTE: MySqlBinaryOperators.JSON_UNQUOTE,
	Tokens.MOD:          MySqlBinaryOperators.MOD,
	Tokens.IN:           MySqlBinaryOperators.IN,
	Tokens.LIKE:         MySqlBinaryOperators.LIKE,
	Tokens.BETWEEN:      MySqlBinaryOperators.BETWEEN,
	Tokens.AND:          MySqlBinaryOperators.BOOLEAN_AND,
	Tokens.XOR:          MySqlBinaryOperators.BOOLEAN_XOR,
	Tokens.OR:           MySqlBinaryOperators.BOOLEAN_OR,
}

var MYSQL_TOKEN_TO_UNARY_OPERATORS = map[Token]UnaryOperator{
	Tokens.BINARY: MySqlUnaryOperators.BINARY,
	Tokens.TILDE:  MySqlUnaryOperators.COMPL,
	Tokens.PLUS:   MySqlUnaryOperators.POSITIVE,
	Tokens.SUB:    MySqlUnaryOperators.NEGATIVE,
	Tokens.BANG:   MySqlUnaryOperators.NOT,
	Tokens.NOT:    MySqlUnaryOperators.NOTSTR,
}
