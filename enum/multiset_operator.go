package enum

import o "github.com/jishaocong0910/go-object-util"

// 多结果集操作符
type MultisetOperator struct {
	*o.EnumElem__
	Sql string
}

type _MultisetOperator struct {
	*o.Enum__[MultisetOperator]
	UNION,
	EXCEPT,
	INTERSECT MultisetOperator
}

var MultisetOperator_ = o.NewEnum[MultisetOperator](_MultisetOperator{
	UNION:     MultisetOperator{Sql: "UNION"},
	EXCEPT:    MultisetOperator{Sql: "EXCEPT"},
	INTERSECT: MultisetOperator{Sql: "INTERSECT"},
})
