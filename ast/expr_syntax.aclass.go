package ast

// 表达式
type I_ExprSyntax interface {
	I_Syntax
	M_A7A153289C4D() *M_ExprSyntax
	// 每个表达式都是可作为运算式的操作数，此方法返回表达式中包含的操作数数量。一般的，字符串、数字、标量函数的操作数是其本身。
	// 列表、子查询也是表达式，并且他们也可进行运算操作，所以也是操作数。列表的操作数数量是其内部包含的元素数量，子查询的操作数
	// 数量为查询列表中的元素数量。
	// e.g.
	//  SELECT (SELECT nickname, level FROM user WHERE id = 1001) = ('my_name', 10), 1 > 2
	// 对于上面的SQL语句中的等于条件的比较运算式，左操作数为子查询，右操作数为列表，其操作数都为2；SQL中的大于条件的比较运算式，
	// 左右操作数都是一个数字标量，操作数是其本身，即只有一个操作数。
	OperandCount() int
	// 是否为表达式列表
	IsExprList() bool
	// 表达式语法类表长度
	ExprLen() int
	// 获得指定索引的表达式
	GetExpr(int) I_ExprSyntax
}

type M_ExprSyntax struct {
	I I_ExprSyntax
}

func (this *M_ExprSyntax) M_A7A153289C4D() *M_ExprSyntax {
	return this
}

func (this *M_ExprSyntax) OperandCount() int {
	if this.I.IsExprList() {
		return this.I.ExprLen()
	}
	return 1
}

func (this *M_ExprSyntax) IsExprList() bool {
	return false
}

func (this *M_ExprSyntax) ExprLen() int {
	return 0
}

func (this *M_ExprSyntax) GetExpr(int) I_ExprSyntax {
	return nil
}

func ExtendExprSyntax(i I_ExprSyntax) *M_ExprSyntax {
	return &M_ExprSyntax{I: i}
}
