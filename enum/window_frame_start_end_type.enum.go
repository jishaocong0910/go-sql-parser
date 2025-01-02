package enum

import (
	"strings"

	o "github.com/jishaocong0910/go-object"
)

type WindowFrameStartEndType struct {
	*o.M_EnumValue
	Sql string
}

type _WindowFrameStartEndType struct {
	*o.M_Enum[WindowFrameStartEndType]
	PRECEDING,
	FOLLOWING WindowFrameStartEndType
}

func (r _WindowFrameStartEndType) OfSql(sql string) (value WindowFrameStartEndType) {
	sql = strings.ToUpper(sql)
	for _, v := range r.Values() {
		if strings.EqualFold(v.Sql, sql) {
			return v
		}
	}
	return
}

var WindowFrameStartEndTypes = o.NewEnum[WindowFrameStartEndType](_WindowFrameStartEndType{
	PRECEDING: WindowFrameStartEndType{Sql: "PRECEDING"},
	FOLLOWING: WindowFrameStartEndType{Sql: "FOLLOWING"},
})
