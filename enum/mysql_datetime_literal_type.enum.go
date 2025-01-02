package enum

import o "github.com/jishaocong0910/go-object"

type MySqlDatetimeLiteralType struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlDatetimeLiteralType struct {
	*o.M_Enum[MySqlDatetimeLiteralType]
	DATE,
	TIME,
	TIMESTAMP MySqlDatetimeLiteralType
}

func (r _MySqlDatetimeLiteralType) OfSql(sql string) (value MySqlDatetimeLiteralType) {
	for _, v := range r.Values() {
		if v.Sql == sql {
			return v
		}
	}
	return
}

var MySqlDatetimeLiteralTypes = o.NewEnum[MySqlDatetimeLiteralType](_MySqlDatetimeLiteralType{
	DATE:      MySqlDatetimeLiteralType{Sql: "DATE"},
	TIME:      MySqlDatetimeLiteralType{Sql: "TIME"},
	TIMESTAMP: MySqlDatetimeLiteralType{Sql: "TIMESTAMP"},
})
