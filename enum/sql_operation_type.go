package enum

import o "github.com/jishaocong0910/go-object-util"

type SqlOperationType struct {
	*o.EnumElem__
}

type _SqlOperationType struct {
	*o.Enum__[SqlOperationType]
	SELECT,
	INSERT,
	UPDATE,
	DELETE,
	OTHER SqlOperationType
}

var SqlOperationType_ = o.NewEnum[SqlOperationType](_SqlOperationType{})
