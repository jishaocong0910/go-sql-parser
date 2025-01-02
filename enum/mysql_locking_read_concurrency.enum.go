package enum

import o "github.com/jishaocong0910/go-object"

type MySqlLockingReadConcurrency struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlLockingReadConcurrency struct {
	*o.M_Enum[MySqlLockingReadConcurrency]
	NO_WAIT,
	SKIP_LOCKED MySqlLockingReadConcurrency
}

var MySqlLockingReadConcurrencys = o.NewEnum[MySqlLockingReadConcurrency](_MySqlLockingReadConcurrency{
	NO_WAIT:     MySqlLockingReadConcurrency{Sql: "NOWAIT"},
	SKIP_LOCKED: MySqlLockingReadConcurrency{Sql: "SKIP LOCKED"},
})
