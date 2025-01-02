package ast

// 查询类语法，例如SELECT和UNION、EXCEPT、INTERSECT
type I_QuerySyntax interface {
	I_ExprSyntax
	I_StatementSyntax
	M_E90D7FD2CE68() *M_QuerySyntax
}

type M_QuerySyntax struct {
	I I_QuerySyntax
}

func (this *M_QuerySyntax) M_E90D7FD2CE68() *M_QuerySyntax {
	return this
}

func ExtendQuerySyntax(i I_QuerySyntax) *M_QuerySyntax {
	return &M_QuerySyntax{I: i}
}
