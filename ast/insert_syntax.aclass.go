package ast

// INSERT语法
type InsertSyntax_ interface {
	InsertSyntax_() *InsertSyntax__
	StatementSyntax_
}

type InsertSyntax__ struct {
	I                  InsertSyntax_
	NameTableReference NameTableReferenceSyntax_
	InsertColumnList   *InsertColumnListSyntax
	ValueListList      ValueListListSyntax_
	Hint               *HintSyntax
}

func (this *InsertSyntax__) InsertSyntax_() *InsertSyntax__ {
	return this
}

func ExtendInsertSyntax(i InsertSyntax_) *InsertSyntax__ {
	return &InsertSyntax__{I: i}
}
