package enum

import o "github.com/jishaocong0910/go-object"

// 排序方向
type OrderingSequence struct {
	*o.M_EnumValue
	Sql string
}

type _OrderingSequence struct {
	*o.M_Enum[OrderingSequence]
	ASC,
	DESC OrderingSequence
}

var OrderingSequences = o.NewEnum[OrderingSequence](_OrderingSequence{
	ASC:  OrderingSequence{Sql: "ASC"},
	DESC: OrderingSequence{Sql: "DESC"},
})
