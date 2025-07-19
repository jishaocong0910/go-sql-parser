package enum

import (
	"strings"

	o "github.com/jishaocong0910/go-object-util"
)

type WindowFrameStartEndType struct {
	*o.EnumElem__
	Sql string
}

type _WindowFrameStartEndType struct {
	*o.Enum__[WindowFrameStartEndType]
	PRECEDING,
	FOLLOWING WindowFrameStartEndType
}

func (r _WindowFrameStartEndType) OfSql(sql string) (value WindowFrameStartEndType) {
	sql = strings.ToUpper(sql)
	for _, v := range r.Elems() {
		if strings.EqualFold(v.Sql, sql) {
			return v
		}
	}
	return
}

var WindowFrameStartEndType_ = o.NewEnum[WindowFrameStartEndType](_WindowFrameStartEndType{
	PRECEDING: WindowFrameStartEndType{Sql: "PRECEDING"},
	FOLLOWING: WindowFrameStartEndType{Sql: "FOLLOWING"},
})
