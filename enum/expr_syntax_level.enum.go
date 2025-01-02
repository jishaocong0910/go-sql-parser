package enum

import o "github.com/jishaocong0910/go-object"

// 解析表达式语法级别。本解析器使用MySQL的分级方法：https://dev.mysql.com/doc/refman/8.0/en/expressions.html
type ExprSyntaxLevel struct {
	*o.M_EnumValue
}

type _ExprSyntaxLevel struct {
	*o.M_Enum[ExprSyntaxLevel]
	//对应MySQL文档中的simple_expr
	SINGLE,
	// 对应MySQL文档中的bit_expr
	CALCULATION,
	// 对应MySQL文档中的predicate
	BOOLEAN_PREDICATE,
	// 对应MySQL文档中的boolean_primary
	BOOLEAN_PRIMARY,
	// 对应MySQL文档中的expr
	EXPR ExprSyntaxLevel
}

var ExprSyntaxLevels = o.NewEnum[ExprSyntaxLevel](_ExprSyntaxLevel{})
