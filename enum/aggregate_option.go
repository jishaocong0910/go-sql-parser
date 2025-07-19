package enum

import o "github.com/jishaocong0910/go-object-util"

// 聚合选项
type AggregateOption struct {
	*o.EnumElem__
	Sql string
}

type _AggregateOption struct {
	*o.Enum__[AggregateOption]
	ALL,
	DISTINCT AggregateOption
}

var AggregateOption_ = o.NewEnum[AggregateOption](_AggregateOption{
	ALL:      AggregateOption{Sql: "ALL"},
	DISTINCT: AggregateOption{Sql: "DISTINCT"},
})
