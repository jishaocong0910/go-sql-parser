package enum

import o "github.com/jishaocong0910/go-object"

// 二元操作符
type BinaryOperator struct {
	*o.M_EnumValue
	// 优先级，越小优先级越高
	Precedence int
	// 符号
	Symbol string
	// 操作符类型
	OperatorType OperatorType
	// 符号类型
	SymbolType SymbolType
	// 是否允许多个操作数
	AllowMultipleOperand bool
}

// 比较操作符优先级，若b比other优先级大则返回值大于0，小则返回值小于0，相等则返回值等于0，若other为nil默认返回1
func (r BinaryOperator) Compare(other BinaryOperator) int {
	if other.Undefined() {
		return 1
	}
	return other.Precedence - r.Precedence
}
