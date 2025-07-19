package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlIndexHintMode struct {
	*o.EnumElem__
	Sql string
}

type _MySqlIndexHintMode struct {
	*o.Enum__[MySqlIndexHintMode]
	USE,
	IGNORE,
	FORCE MySqlIndexHintMode
}

var MySqlIndexHintMode_ = o.NewEnum[MySqlIndexHintMode](_MySqlIndexHintMode{
	USE:    MySqlIndexHintMode{Sql: "USE"},
	IGNORE: MySqlIndexHintMode{Sql: "IGNORE"},
	FORCE:  MySqlIndexHintMode{Sql: "FORCE"},
})
