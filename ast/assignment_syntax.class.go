package ast

type AssignmentSyntax struct {
	*M_Syntax
	Column  I_ColumnItemSyntax
	Value   I_ExprSyntax
	Default bool
}

func (this *AssignmentSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitAssignmentSyntax(this)
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
	s.M_Syntax = ExtendSyntax(s)
	return s
}
