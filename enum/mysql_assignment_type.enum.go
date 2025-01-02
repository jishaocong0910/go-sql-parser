package enum

import o "github.com/jishaocong0910/go-object"

type MySqlAssignmentType struct {
	*o.M_EnumValue
}

type _MySqlAssignmentType struct {
	*o.M_Enum[MySqlAssignmentType]
	VALUES_LIST,
	ASSIGNMENT_LIST MySqlAssignmentType
}

var MySqlAssignmentTypes = o.NewEnum[MySqlAssignmentType](_MySqlAssignmentType{})
