package enum

import o "github.com/jishaocong0910/go-object"

// 聚合选项
type AggregateOption struct {
	*o.M_EnumValue
	Sql string
}

type _AggregateOption struct {
	*o.M_Enum[AggregateOption]
	ALL,
	DISTINCT AggregateOption
}

var AggregateOptions = o.NewEnum[AggregateOption](_AggregateOption{
	ALL:      AggregateOption{Sql: "ALL"},
	DISTINCT: AggregateOption{Sql: "DISTINCT"},
})
