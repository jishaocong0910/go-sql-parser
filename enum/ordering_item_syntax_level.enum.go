package enum

import o "github.com/jishaocong0910/go-object"

type OrderingItemSyntaxLevel struct {
	*o.M_EnumValue
}

type _OrderingItemSyntaxLevel struct {
	*o.M_Enum[OrderingItemSyntaxLevel]
	IDENTIFIER,
	NORNAL OrderingItemSyntaxLevel
}

var OrderingItemSyntaxLevels = o.NewEnum[OrderingItemSyntaxLevel](_OrderingItemSyntaxLevel{})
