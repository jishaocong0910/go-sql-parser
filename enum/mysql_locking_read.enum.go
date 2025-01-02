package enum

import o "github.com/jishaocong0910/go-object"

type MySqlLockingRead struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlLockingRead struct {
	*o.M_Enum[MySqlLockingRead]
	FOR_UPDATE,
	// FOR_SHARE 8.0新增
	FOR_SHARE,
	LOCK_IN_SHARE_MODE MySqlLockingRead
}

var MySqlLockingReads = o.NewEnum[MySqlLockingRead](_MySqlLockingRead{
	FOR_UPDATE:         MySqlLockingRead{Sql: "FOR UPDATE"},
	FOR_SHARE:          MySqlLockingRead{Sql: "FOR SHARE"},
	LOCK_IN_SHARE_MODE: MySqlLockingRead{Sql: "LOCK IN SHARE MODE"},
})
