package ast

type DerivedTableReferenceSyntax struct {
	*M_Syntax
	*M_TableReferenceSyntax
	Query I_QuerySyntax
	Alias I_IdentifierSyntax
}

func (this *DerivedTableReferenceSyntax) accept(iv I_Visitor) {
	iv.m_E61B18189B57().visitDerivedTableTableReferenceSyntax(this)
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
	s.M_Syntax = ExtendSyntax(s)
	s.M_TableReferenceSyntax = ExtendTableReferenceSyntax(s)
	return s
}
