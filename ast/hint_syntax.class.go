package ast

import "github.com/jishaocong0910/go-sql-parser/enum"

type HintSyntax struct {
	*Syntax__
	CommentType enum.CommentType
	Content     string
}

func (this *HintSyntax) accept(Visitor_) {}

func (this *HintSyntax) writeSql(builder *sqlBuilder) {
	switch this.CommentType.ID() {
	case enum.CommentType_.SINGLE_LINE.ID():
		builder.writeStr("--+ ")
		builder.writeStr(this.Content)
		builder.writeLf()
	case enum.CommentType_.MULTI_LINE.ID():
		builder.writeStr("/*+")
		builder.writeStr(this.Content)
		builder.writeStr("*/")
	}
}

func NewHintSyntax() *HintSyntax {
	s := &HintSyntax{}
	s.Syntax__ = ExtendSyntax(s)
	return s
}
