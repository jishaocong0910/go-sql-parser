package enum

import o "github.com/jishaocong0910/go-object-util"

type CommentType struct {
	*o.EnumElem__
}

type _CommentType struct {
	*o.Enum__[CommentType]
	SINGLE_LINE,
	MULTI_LINE CommentType
}

var CommentType_ = o.NewEnum[CommentType](_CommentType{})
