package enum

import o "github.com/jishaocong0910/go-object"

type MySqlComparisonMode struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlComparisonMode struct {
	*o.M_Enum[MySqlComparisonMode]
	ALL,
	ANY,
	SOME MySqlComparisonMode
}

var MySqlComparisonModes = o.NewEnum[MySqlComparisonMode](_MySqlComparisonMode{
	ALL:  MySqlComparisonMode{Sql: "ALL"},
	ANY:  MySqlComparisonMode{Sql: "ANY"},
	SOME: MySqlComparisonMode{Sql: "SOME"},
})
