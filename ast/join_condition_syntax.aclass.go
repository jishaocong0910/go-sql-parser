package ast

// JOIN语法的条件条件
type I_JoinConditionSyntax interface {
	I_Syntax
	M_JoinConditionSyntax_() *M_JoinConditionSyntax
}

type M_JoinConditionSyntax struct {
	I I_JoinConditionSyntax
}

func (this *M_JoinConditionSyntax) M_JoinConditionSyntax_() *M_JoinConditionSyntax {
	return this
}

func ExtendJoinConditionSyntax(i I_JoinConditionSyntax) *M_JoinConditionSyntax {
	return &M_JoinConditionSyntax{I: i}
}
