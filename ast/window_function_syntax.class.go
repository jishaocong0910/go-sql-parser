package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

// WindowFunctionSyntax
// Window函数，统一解析为
//
//	<window_function>: function_name([<expr>[, <expr>]...]) [<null_treatment>] <over_syntax>
//	<null_treatment>: IGNORE NULLS|RESPECT NULLS
//	<over_syntax>: 省略...
//
// 即不校验窗口函数的参数、是否支持<null_treatment>，这样解析能覆盖已有的所有Window函数。
type WindowFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	NullTreatment enum.NullTreatment
	Over          *OverSyntax
}

func (this *WindowFunctionSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitWindowFunctionSyntax(this)
}

func (this *WindowFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	if !this.NullTreatment.Undefined() {
		builder.writeStr(this.NullTreatment.Sql)
	}
	builder.writeSyntax(this.Parameters)
	builder.writeSpace()
	builder.writeSyntax(this.Over)
}

func NewWindowFunctionSyntax() *WindowFunctionSyntax {
	s := &WindowFunctionSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	return s
}
