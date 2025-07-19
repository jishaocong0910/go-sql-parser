package enum

import o "github.com/jishaocong0910/go-object-util"

type Dialect struct {
	*o.EnumElem__
	Name string
}

type _Dialect struct {
	*o.Enum__[Dialect]
	MYSQL,
	ORACLE,
	SQLSERVER Dialect
}

var Dialect_ = o.NewEnum[Dialect](_Dialect{
	MYSQL:     Dialect{Name: "MySQL"},
	ORACLE:    Dialect{Name: "Oracle"},
	SQLSERVER: Dialect{Name: "SQL Server"},
})
