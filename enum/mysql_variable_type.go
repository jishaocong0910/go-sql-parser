package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlVariableType struct {
	*o.EnumElem__
}

type _MySqlVariableType struct {
	*o.Enum__[MySqlVariableType]
	SESSION,
	GLOBAL MySqlVariableType
}

var MySqlVariableType_ = o.NewEnum[MySqlVariableType](_MySqlVariableType{})
