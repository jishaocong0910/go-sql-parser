package enum

import o "github.com/jishaocong0910/go-object"

type Dialect struct {
	*o.M_EnumValue
	Name string
}

type _Dialect struct {
	*o.M_Enum[Dialect]
	MYSQL,
	ORACLE,
	SQLSERVER Dialect
}

var Dialects = o.NewEnum[Dialect](_Dialect{
	MYSQL:     Dialect{Name: "MySQL"},
	ORACLE:    Dialect{Name: "Oracle"},
	SQLSERVER: Dialect{Name: "SQL Server"},
})
