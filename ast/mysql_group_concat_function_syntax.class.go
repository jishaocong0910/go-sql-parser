package ast

type MySqlGroupConcatFunctionSyntax struct {
	*Syntax__
	*ExprSyntax__
	*FunctionSyntax__
	*AggregateFunctionSyntax__
	OrderBy   *OrderBySyntax
	Separator *MySqlStringSyntax
}

func (this *MySqlGroupConcatFunctionSyntax) accept(v_ Visitor_) {
	this.AggregateFunctionSyntax__.accept(v_)
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
	s.Syntax__ = ExtendSyntax(s)
	s.ExprSyntax__ = ExtendExprSyntax(s)
	s.FunctionSyntax__ = ExtendFunctionSyntax(s)
	s.AggregateFunctionSyntax__ = ExtendAggregateFunctionSyntax(s)
	return s
}
