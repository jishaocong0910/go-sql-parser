package enum

import o "github.com/jishaocong0910/go-object"

// https://dev.mysql.com/doc/refman/8.0/en/delete.html
type MySqlMultiDeleteMode struct {
	*o.M_EnumValue
}

type _MySqlMultiDeleteMode struct {
	*o.M_Enum[MySqlMultiDeleteMode]
	// DELETE [LOW_PRIORITY] [QUICK] [IGNORE]
	// tbl_name[.*] [, tbl_name[.*]] ...
	// FROM table_references
	// [WHERE where_condition]
	MODE1,
	// DELETE [LOW_PRIORITY] [QUICK] [IGNORE]
	// FROM tbl_name[.*] [, tbl_name[.*]] ...
	// USING table_references
	// [WHERE where_condition]
	MODE2 MySqlMultiDeleteMode
}

var MySqlMultiDeleteModes = o.NewEnum[MySqlMultiDeleteMode](_MySqlMultiDeleteMode{})
