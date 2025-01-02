package enum

import o "github.com/jishaocong0910/go-object"

// 多结果集操作符
type MultisetOperator struct {
	*o.M_EnumValue
	Sql string
}

type _MultisetOperator struct {
	*o.M_Enum[MultisetOperator]
	UNION,
	EXCEPT,
	INTERSECT MultisetOperator
}

var MultisetOperators = o.NewEnum[MultisetOperator](_MultisetOperator{
	UNION:     MultisetOperator{Sql: "UNION"},
	EXCEPT:    MultisetOperator{Sql: "EXCEPT"},
	INTERSECT: MultisetOperator{Sql: "INTERSECT"},
})
