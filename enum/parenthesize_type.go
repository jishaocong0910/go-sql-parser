package enum

import o "github.com/jishaocong0910/go-object-util"

// ParenthesizeType 括号类型（用于标记一个语法对象是否有括号，或者指定该语法对象不能使用括号）
type ParenthesizeType struct {
	*o.EnumElem__
}

var ParenthesizeType_ = o.NewEnum[ParenthesizeType](struct {
	*o.Enum__[ParenthesizeType]
	// 含有括号
	TRUE,
	// 不支持括号
	NOT_SUPPORT ParenthesizeType
}{})
