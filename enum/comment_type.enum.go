package enum

import o "github.com/jishaocong0910/go-object"

type CommentType struct {
	*o.M_EnumValue
}

type _CommentType struct {
	*o.M_Enum[CommentType]
	SINGLE_LINE,
	MULTI_LINE CommentType
}

var CommentTypes = o.NewEnum[CommentType](_CommentType{})
