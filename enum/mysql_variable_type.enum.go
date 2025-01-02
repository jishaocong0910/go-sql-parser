package enum

import o "github.com/jishaocong0910/go-object"

type MySqlVariableType struct {
	*o.M_EnumValue
}

type _MySqlVariableType struct {
	*o.M_Enum[MySqlVariableType]
	SESSION,
	GLOBAL MySqlVariableType
}

var MySqlVariableTypes = o.NewEnum[MySqlVariableType](_MySqlVariableType{})
