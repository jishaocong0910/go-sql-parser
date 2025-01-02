package enum

import o "github.com/jishaocong0910/go-object"

type MySqlGetFormatType struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlGetFormatType struct {
	*o.M_Enum[MySqlGetFormatType]
	DATE,
	TIME,
	DATETIME MySqlGetFormatType
}

func (r _MySqlGetFormatType) OfSql(sql string) (value MySqlGetFormatType) {
	for _, v := range r.Values() {
		if v.Sql == sql {
			return v
		}
	}
	return
}

var MySqlGetFormatTypes = o.NewEnum[MySqlGetFormatType](_MySqlGetFormatType{
	DATE:     MySqlGetFormatType{Sql: "DATE"},
	TIME:     MySqlGetFormatType{Sql: "TIME"},
	DATETIME: MySqlGetFormatType{Sql: "DATETIME"},
})
