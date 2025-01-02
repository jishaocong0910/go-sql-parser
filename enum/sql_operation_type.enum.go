package enum

import o "github.com/jishaocong0910/go-object"

type SqlOperationType struct {
	*o.M_EnumValue
}

type _SqlOperationType struct {
	*o.M_Enum[SqlOperationType]
	SELECT,
	INSERT,
	UPDATE,
	DELETE,
	OTHER SqlOperationType
}

var SqlOperationTypes = o.NewEnum[SqlOperationType](_SqlOperationType{})
