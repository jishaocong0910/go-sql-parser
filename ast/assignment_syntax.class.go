package ast

type AssignmentSyntax struct {
	*Syntax__
	Column  ColumnItemSyntax_
	Value   ExprSyntax_
	Default bool
}

func (this *AssignmentSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitAssignmentSyntax(this)
}

func (this *AssignmentSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Column)
	builder.writeStr(" = ")
	if this.Default {
		builder.writeStr("DEFAULT")
	} else {
		builder.writeSyntax(this.Value)
	}
}

func NewAssignmentSyntax() *AssignmentSyntax {
	s := &AssignmentSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
