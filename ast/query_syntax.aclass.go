package ast

// 查询类语法，例如SELECT和UNION、EXCEPT、INTERSECT
type QuerySyntax_ interface {
	QuerySyntax_() *QuerySyntax__
	ExprSyntax_
	StatementSyntax_
}

type QuerySyntax__ struct {
	I QuerySyntax_
}

func (this *QuerySyntax__) QuerySyntax_() *QuerySyntax__ {
	return this
}

func ExtendQuerySyntax(i QuerySyntax_) *QuerySyntax__ {
	return &QuerySyntax__{I: i}
}
