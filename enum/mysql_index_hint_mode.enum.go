package enum

import o "github.com/jishaocong0910/go-object"

type MySqlIndexHintMode struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlIndexHintMode struct {
	*o.M_Enum[MySqlIndexHintMode]
	USE,
	IGNORE,
	FORCE MySqlIndexHintMode
}

var MySqlIndexHintModes = o.NewEnum[MySqlIndexHintMode](_MySqlIndexHintMode{
	USE:    MySqlIndexHintMode{Sql: "USE"},
	IGNORE: MySqlIndexHintMode{Sql: "IGNORE"},
	FORCE:  MySqlIndexHintMode{Sql: "FORCE"},
})
