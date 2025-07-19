package enum

import o "github.com/jishaocong0910/go-object-util"

type OrderingItemSyntaxLevel struct {
	*o.EnumElem__
}

type _OrderingItemSyntaxLevel struct {
	*o.Enum__[OrderingItemSyntaxLevel]
	IDENTIFIER,
	NORNAL OrderingItemSyntaxLevel
}

var OrderingItemSyntaxLevel_ = o.NewEnum[OrderingItemSyntaxLevel](_OrderingItemSyntaxLevel{})
