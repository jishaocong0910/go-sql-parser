package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlTrimMode struct {
	*o.EnumElem__
	Sql string
}

type _MySqlTrimMode struct {
	*o.Enum__[MySqlTrimMode]
	BOTH,
	LEADING,
	TRAILING MySqlTrimMode
}

var MySqlTrimMode_ = o.NewEnum[MySqlTrimMode](_MySqlTrimMode{
	BOTH:     MySqlTrimMode{Sql: "BOTH"},
	LEADING:  MySqlTrimMode{Sql: "LEADING"},
	TRAILING: MySqlTrimMode{Sql: "TRAILING"},
})
