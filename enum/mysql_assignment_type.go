package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlAssignmentType struct {
	*o.EnumElem__
}

type _MySqlAssignmentType struct {
	*o.Enum__[MySqlAssignmentType]
	VALUES_LIST,
	ASSIGNMENT_LIST MySqlAssignmentType
}

var MySqlAssignmentType_ = o.NewEnum[MySqlAssignmentType](_MySqlAssignmentType{})
