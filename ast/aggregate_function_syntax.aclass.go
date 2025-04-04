package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"
)

// 聚合函数
type I_AggregateFunctionSyntax interface {
	I_FunctionSyntax
	M_AggregateFunctionSyntax_() *M_AggregateFunctionSyntax
}

type M_AggregateFunctionSyntax struct {
	I                  I_AggregateFunctionSyntax
	AggregateOption    enum.AggregateOption
	AllColumnParameter bool
	Over               *OverSyntax
}

func (this *M_AggregateFunctionSyntax) M_AggregateFunctionSyntax_() *M_AggregateFunctionSyntax {
	return this
}

func (this *M_AggregateFunctionSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitAggregateFunctionSyntax(this)
}

func (this *M_AggregateFunctionSyntax) writeSql(builder *sqlBuilder) {
	functionSyntax := this.I.M_FunctionSyntax_()
	builder.writeStr(functionSyntax.Name)
	builder.writeStr("(")
	if !this.AggregateOption.Undefined() {
		builder.writeStr(this.AggregateOption.Sql)
		builder.writeSpace()
	}
	if this.AllColumnParameter {
		builder.writeStr("*")
	} else if functionSyntax.Parameters != nil {
		builder.writeSyntaxWithFormat(functionSyntax.Parameters, false)
	}
	builder.writeStr(")")
	if this.Over != nil {
		builder.writeSpace()
		builder.writeSyntax(this.Over)
	}
}

func ExtendAggregateFunctionSyntax(i I_AggregateFunctionSyntax) *M_AggregateFunctionSyntax {
	return &M_AggregateFunctionSyntax{I: i}
}
