package ast

type DualTableReferenceSyntax struct {
	*M_Syntax
	*M_TableReferenceSyntax
}

func (this *DualTableReferenceSyntax) accept(I_Visitor) {}

func (this *DualTableReferenceSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("DUAL")
}

func NewDualTableReferenceSyntax() *DualTableReferenceSyntax {
	s := &DualTableReferenceSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_TableReferenceSyntax = ExtendTableReferenceSyntax(s)
	return s
}
