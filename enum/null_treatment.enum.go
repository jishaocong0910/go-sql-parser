package enum

import o "github.com/jishaocong0910/go-object"

type NullTreatment struct {
	*o.M_EnumValue
	Sql string
}

type _NullTreatment struct {
	*o.M_Enum[NullTreatment]
	IGNORE_NULLS,
	RESPECT_NULLS NullTreatment
}

var NullTreatments = o.NewEnum[NullTreatment](_NullTreatment{
	IGNORE_NULLS:  NullTreatment{Sql: "IGNORE NULLS"},
	RESPECT_NULLS: NullTreatment{Sql: "RESPECT NULLS"},
})
