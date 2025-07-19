package ast

import (
	"github.com/jishaocong0910/go-sql-parser/enum"
)

// 聚合函数
type AggregateFunctionSyntax_ interface {
	AggregateFunctionSyntax_() *AggregateFunctionSyntax__
	FunctionSyntax_
}

type AggregateFunctionSyntax__ struct {
	I AggregateFunctionSyntax_

	AggregateOption    enum.AggregateOption
	AllColumnParameter bool
	Over               *OverSyntax
}

func (this *AggregateFunctionSyntax__) AggregateFunctionSyntax_() *AggregateFunctionSyntax__ {
	return this
}

func (this *AggregateFunctionSyntax__) accept(v_ Visitor_) {
	v_.visitor_().visitAggregateFunctionSyntax__(this)
}

func (this *AggregateFunctionSyntax__) writeSql(builder *sqlBuilder) {
	functionSyntax := this.I.FunctionSyntax_()
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

func ExtendAggregateFunctionSyntax(i AggregateFunctionSyntax_) *AggregateFunctionSyntax__ {
	return &AggregateFunctionSyntax__{I: i}
}
