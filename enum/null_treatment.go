package enum

import o "github.com/jishaocong0910/go-object-util"

type NullTreatment struct {
	*o.EnumElem__
	Sql string
}

type _NullTreatment struct {
	*o.Enum__[NullTreatment]
	IGNORE_NULLS,
	RESPECT_NULLS NullTreatment
}

var NullTreatment_ = o.NewEnum[NullTreatment](_NullTreatment{
	IGNORE_NULLS:  NullTreatment{Sql: "IGNORE NULLS"},
	RESPECT_NULLS: NullTreatment{Sql: "RESPECT NULLS"},
})
