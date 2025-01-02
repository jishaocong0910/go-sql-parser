package enum

import o "github.com/jishaocong0910/go-object"

// 操作（或运算）符类型
type OperatorType struct {
	*o.M_EnumValue
}

type _OperatorType struct {
	*o.M_Enum[OperatorType]
	// 算数
	ARITHMETIC,
	// 位
	BITWISE,
	// 比较
	COMPARISON,
	// 断言
	PREDICATE,
	// 逻辑
	LOGICAL,
	// 赋值
	ASSIGNMENT,
	// 特殊
	SPECIAL,
	// 拼接
	CONCAT OperatorType
}

var OperatorTypes = o.NewEnum[OperatorType](_OperatorType{})
