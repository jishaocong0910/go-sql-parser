package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlLockingReadConcurrency struct {
	*o.EnumElem__
	Sql string
}

type _MySqlLockingReadConcurrency struct {
	*o.Enum__[MySqlLockingReadConcurrency]
	NO_WAIT,
	SKIP_LOCKED MySqlLockingReadConcurrency
}

var MySqlLockingReadConcurrency_ = o.NewEnum[MySqlLockingReadConcurrency](_MySqlLockingReadConcurrency{
	NO_WAIT:     MySqlLockingReadConcurrency{Sql: "NOWAIT"},
	SKIP_LOCKED: MySqlLockingReadConcurrency{Sql: "SKIP LOCKED"},
})
