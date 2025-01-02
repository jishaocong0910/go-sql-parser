package enum

import o "github.com/jishaocong0910/go-object"

// SQL语句类型
type StatementType struct {
	*o.M_EnumValue
}

type _StatementType struct {
	*o.M_Enum[StatementType]
	SELECT,
	INSERT,
	UPDATE,
	DELETE,
	OTHER StatementType
}

var StatementTypes = o.NewEnum[StatementType](_StatementType{})
