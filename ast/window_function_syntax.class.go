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
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
	NullTreatment enum.NullTreatment
	Over          *OverSyntax
}

func (this *WindowFunctionSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitWindowFunctionSyntax(this)
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	return s
}
