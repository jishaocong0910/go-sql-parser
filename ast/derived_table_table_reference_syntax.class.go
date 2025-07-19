package ast

type DerivedTableReferenceSyntax struct {
	*Syntax__
	*TableReferenceSyntax__
	Query QuerySyntax_
	Alias IdentifierSyntax_
}

func (this *DerivedTableReferenceSyntax) accept(v_ Visitor_) {
	v_.visitor_().visitDerivedTableTableReferenceSyntax(this)
}

func (this *DerivedTableReferenceSyntax) writeSql(builder *sqlBuilder) {
	builder.writeSyntax(this.Query)
	if this.Alias != nil {
		builder.writeSpace()
		builder.writeSyntax(this.Alias)
	}
}

func NewDerivedTableTableReferenceSyntax() *DerivedTableReferenceSyntax {
	s := &DerivedTableReferenceSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.TableReferenceSyntax__ = ExtendTableReferenceSyntax(s)
	return s
}
