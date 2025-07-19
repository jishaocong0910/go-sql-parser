package ast

// JOIN语法的条件条件
type JoinConditionSyntax_ interface {
	JoinConditionSyntax_() *JoinConditionSyntax__
	Syntax_
}

type JoinConditionSyntax__ struct {
	I JoinConditionSyntax_
}

func (this *JoinConditionSyntax__) JoinConditionSyntax_() *JoinConditionSyntax__ {
	return this
}

func ExtendJoinConditionSyntax(i JoinConditionSyntax_) *JoinConditionSyntax__ {
	return &JoinConditionSyntax__{I: i}
}
