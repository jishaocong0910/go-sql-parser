package enum

import (
	"strings"

	o "github.com/jishaocong0910/go-object-util"
)

type MySqlTemporalInterval struct {
	*o.EnumElem__
	Sql string
}

type _MySqlTemporalInterval struct {
	*o.Enum__[MySqlTemporalInterval]
	MICROSECOND,
	SECOND,
	MINUTE,
	HOUR,
	DAY,
	WEEK,
	MONTH,
	QUARTER,
	YEAR,
	SECOND_MICROSECOND,
	MINUTE_MICROSECOND,
	MINUTE_SECOND,
	HOUR_MICROSECOND,
	HOUR_SECOND,
	HOUR_MINUTE,
	DAY_MICROSECOND,
	DAY_SECOND,
	DAY_MINUTE,
	DAY_HOUR,
	YEAR_MONTH MySqlTemporalInterval
}

func (r _MySqlTemporalInterval) OfSql(sql string) (value MySqlTemporalInterval) {
	sql = strings.ToUpper(sql)
	for _, v := range r.Elems() {
		if v.Sql == sql {
			return v
		}
	}
	return
}

var MySqlTemporalInterval_ = o.NewEnum[MySqlTemporalInterval](_MySqlTemporalInterval{
	MICROSECOND:        MySqlTemporalInterval{Sql: "MICROSECOND"},
	SECOND:             MySqlTemporalInterval{Sql: "SECOND"},
	MINUTE:             MySqlTemporalInterval{Sql: "MINUTE"},
	HOUR:               MySqlTemporalInterval{Sql: "HOUR"},
	DAY:                MySqlTemporalInterval{Sql: "DAY"},
	WEEK:               MySqlTemporalInterval{Sql: "WEEK"},
	MONTH:              MySqlTemporalInterval{Sql: "MONTH"},
	QUARTER:            MySqlTemporalInterval{Sql: "QUARTER"},
	YEAR:               MySqlTemporalInterval{Sql: "YEAR"},
	SECOND_MICROSECOND: MySqlTemporalInterval{Sql: "SECOND_MICROSECOND"},
	MINUTE_MICROSECOND: MySqlTemporalInterval{Sql: "MINUTE_MICROSECOND"},
	MINUTE_SECOND:      MySqlTemporalInterval{Sql: "MINUTE_SECOND"},
	HOUR_MICROSECOND:   MySqlTemporalInterval{Sql: "HOUR_MICROSECOND"},
	HOUR_SECOND:        MySqlTemporalInterval{Sql: "HOUR_SECOND"},
	HOUR_MINUTE:        MySqlTemporalInterval{Sql: "HOUR_MINUTE"},
	DAY_MICROSECOND:    MySqlTemporalInterval{Sql: "DAY_MICROSECOND"},
	DAY_SECOND:         MySqlTemporalInterval{Sql: "DAY_SECOND"},
	DAY_MINUTE:         MySqlTemporalInterval{Sql: "DAY_MINUTE"},
	DAY_HOUR:           MySqlTemporalInterval{Sql: "DAY_HOUR"},
	YEAR_MONTH:         MySqlTemporalInterval{Sql: "YEAR_MONTH"},
})
