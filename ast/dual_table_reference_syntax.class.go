package ast

type DualTableReferenceSyntax struct {
	*Syntax__
	*TableReferenceSyntax__
}

func (this *DualTableReferenceSyntax) accept(Visitor_) {}

func (this *DualTableReferenceSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr("DUAL")
}

func NewDualTableReferenceSyntax() *DualTableReferenceSyntax {
	s := &DualTableReferenceSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	s.TableReferenceSyntax__ = ExtendTableReferenceSyntax(s)
	return s
}
