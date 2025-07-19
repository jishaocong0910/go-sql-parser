package enum

import o "github.com/jishaocong0910/go-object-util"

// https://dev.mysql.com/doc/refman/8.0/en/delete.html
type MySqlMultiDeleteMode struct {
	*o.EnumElem__
}

type _MySqlMultiDeleteMode struct {
	*o.Enum__[MySqlMultiDeleteMode]
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

var MySqlMultiDeleteMode_ = o.NewEnum[MySqlMultiDeleteMode](_MySqlMultiDeleteMode{})
