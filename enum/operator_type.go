package enum

import o "github.com/jishaocong0910/go-object-util"

// 操作（或运算）符类型
type OperatorType struct {
	*o.EnumElem__
}

type _OperatorType struct {
	*o.Enum__[OperatorType]
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

var OperatorType_ = o.NewEnum[OperatorType](_OperatorType{})
