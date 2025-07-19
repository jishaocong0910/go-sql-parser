package enum

import o "github.com/jishaocong0910/go-object-util"

// 排序方向
type OrderingSequence struct {
	*o.EnumElem__
	Sql string
}

type _OrderingSequence struct {
	*o.Enum__[OrderingSequence]
	ASC,
	DESC OrderingSequence
}

var OrderingSequence_ = o.NewEnum[OrderingSequence](_OrderingSequence{
	ASC:  OrderingSequence{Sql: "ASC"},
	DESC: OrderingSequence{Sql: "DESC"},
})
