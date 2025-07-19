package ast

type NamedWindowsSyntax struct {
	*Syntax__
	Name       IdentifierSyntax_
	WindowSpec *WindowSpecSyntax
}

func (this *NamedWindowsSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitNamedWindowsSyntax(this)
}

func (this *NamedWindowsSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Name)
	builder.writeStr(" AS ")
	builder.writeSyntax(this.WindowSpec)
}

func NewNamedWindowsSyntax() *NamedWindowsSyntax {
	s := &NamedWindowsSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
