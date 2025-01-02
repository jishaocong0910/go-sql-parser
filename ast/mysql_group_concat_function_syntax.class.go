package ast

type MySqlGroupConcatFunctionSyntax struct {
	*M_Syntax
	*M_ExprSyntax
	*M_FunctionSyntax
	*M_AggregateFunctionSyntax
	OrderBy   *OrderBySyntax
	Separator *MySqlStringSyntax
}

func (this *MySqlGroupConcatFunctionSyntax) accept(iv I_Visitor) {
	this.M_AggregateFunctionSyntax.accept(iv)
}

func (this *MySqlGroupConcatFunctionSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.Name)
	builder.writeStr("(")
	if !this.AggregateOption.Undefined() {
		builder.writeStr(this.AggregateOption.Sql)
		builder.writeSpace()
	}
	builder.writeSyntax(this.Parameters)
	if this.OrderBy != nil {
		builder.writeSpace()
		builder.writeSyntaxWithFormat(this.OrderBy, false)
	}
	if this.Separator != nil {
		builder.writeStr(" SEPARATOR ")
		builder.writeSyntax(this.Separator)
	}
	builder.writeStr(")")
}

func NewMySqlGroupConcatFunctionSyntax() *MySqlGroupConcatFunctionSyntax {
	s := &MySqlGroupConcatFunctionSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ExprSyntax = ExtendExprSyntax(s)
	s.M_FunctionSyntax = ExtendFunctionSyntax(s)
	s.M_AggregateFunctionSyntax = ExtendAggregateFunctionSyntax(s)
	return s
}
