package ast

import "go-sql-parser/enum"

type HintSyntax struct {
	*M_Syntax
	CommentType enum.CommentType
	Content     string
}

func (this *HintSyntax) accept(I_Visitor) {}

func (this *HintSyntax) writeSql(builder *sqlBuilder) {
	switch this.CommentType.ID() {
	case enum.CommentTypes.SINGLE_LINE.ID():
		builder.writeStr("--+ ")
		builder.writeStr(this.Content)
		builder.writeLf()
	case enum.CommentTypes.MULTI_LINE.ID():
		builder.writeStr("/*+")
		builder.writeStr(this.Content)
		builder.writeStr("*/")
	}
}

func NewHintSyntax() *HintSyntax {
	s := &HintSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
