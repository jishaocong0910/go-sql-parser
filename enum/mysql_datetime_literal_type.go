package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlDatetimeLiteralType struct {
	*o.EnumElem__
	Sql string
}

type _MySqlDatetimeLiteralType struct {
	*o.Enum__[MySqlDatetimeLiteralType]
	DATE,
	TIME,
	TIMESTAMP MySqlDatetimeLiteralType
}

func (r _MySqlDatetimeLiteralType) OfSql(sql string) (value MySqlDatetimeLiteralType) {
	for _, v := range r.Elems() {
		if v.Sql == sql {
			return v
		}
	}
	return
}

var MySqlDatetimeLiteralType_ = o.NewEnum[MySqlDatetimeLiteralType](_MySqlDatetimeLiteralType{
	DATE:      MySqlDatetimeLiteralType{Sql: "DATE"},
	TIME:      MySqlDatetimeLiteralType{Sql: "TIME"},
	TIMESTAMP: MySqlDatetimeLiteralType{Sql: "TIMESTAMP"},
})
