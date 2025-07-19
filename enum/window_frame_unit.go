package enum

import o "github.com/jishaocong0910/go-object-util"

type WindowFrameUnit struct {
	*o.EnumElem__
	Sql string
}

type _WindowFrameUnit struct {
	*o.Enum__[WindowFrameUnit]
	ROWS,
	RANGE WindowFrameUnit
}

var WindowFrameUnit_ = o.NewEnum[WindowFrameUnit](_WindowFrameUnit{
	ROWS:  WindowFrameUnit{Sql: "ROWS"},
	RANGE: WindowFrameUnit{Sql: "RANGE"},
})
