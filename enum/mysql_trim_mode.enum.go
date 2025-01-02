package enum

import o "github.com/jishaocong0910/go-object"

type MySqlTrimMode struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlTrimMode struct {
	*o.M_Enum[MySqlTrimMode]
	BOTH,
	LEADING,
	TRAILING MySqlTrimMode
}

var MySqlTrimModes = o.NewEnum[MySqlTrimMode](_MySqlTrimMode{
	BOTH:     MySqlTrimMode{Sql: "BOTH"},
	LEADING:  MySqlTrimMode{Sql: "LEADING"},
	TRAILING: MySqlTrimMode{Sql: "TRAILING"},
})
