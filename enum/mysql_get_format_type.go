package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlGetFormatType struct {
	*o.EnumElem__
	Sql string
}

type _MySqlGetFormatType struct {
	*o.Enum__[MySqlGetFormatType]
	DATE,
	TIME,
	DATETIME MySqlGetFormatType
}

func (r _MySqlGetFormatType) OfSql(sql string) (value MySqlGetFormatType) {
	for _, v := range r.Elems() {
		if v.Sql == sql {
			return v
		}
	}
	return
}

var MySqlGetFormatType_ = o.NewEnum[MySqlGetFormatType](_MySqlGetFormatType{
	DATE:     MySqlGetFormatType{Sql: "DATE"},
	TIME:     MySqlGetFormatType{Sql: "TIME"},
	DATETIME: MySqlGetFormatType{Sql: "DATETIME"},
})
