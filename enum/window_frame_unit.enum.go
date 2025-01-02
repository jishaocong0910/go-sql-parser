package enum

import o "github.com/jishaocong0910/go-object"

type WindowFrameUnit struct {
	*o.M_EnumValue
	Sql string
}

type _WindowFrameUnit struct {
	*o.M_Enum[WindowFrameUnit]
	ROWS,
	RANGE WindowFrameUnit
}

var WindowFrameUnits = o.NewEnum[WindowFrameUnit](_WindowFrameUnit{
	ROWS:  WindowFrameUnit{Sql: "ROWS"},
	RANGE: WindowFrameUnit{Sql: "RANGE"},
})
