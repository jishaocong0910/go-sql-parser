package ast

type NamedWindowsSyntax struct {
	*M_Syntax
	Name       I_IdentifierSyntax
	WindowSpec *WindowSpecSyntax
}

func (this *NamedWindowsSyntax) accept(iv I_Visitor) {
	iv.m_Visitor_().visitNamedWindowsSyntax(this)
}

func (this *NamedWindowsSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Name)
	builder.writeStr(" AS ")
	builder.writeSyntax(this.WindowSpec)
}

func NewNamedWindowsSyntax() *NamedWindowsSyntax {
	s := &NamedWindowsSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
