package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlComparisonMode struct {
	*o.EnumElem__
	Sql string
}

type _MySqlComparisonMode struct {
	*o.Enum__[MySqlComparisonMode]
	ALL,
	ANY,
	SOME MySqlComparisonMode
}

var MySqlComparisonMode_ = o.NewEnum[MySqlComparisonMode](_MySqlComparisonMode{
	ALL:  MySqlComparisonMode{Sql: "ALL"},
	ANY:  MySqlComparisonMode{Sql: "ANY"},
	SOME: MySqlComparisonMode{Sql: "SOME"},
})
