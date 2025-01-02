package ast

// MySqlValueListListSyntax
//
//	@Description: https://dev.mysql.com/doc/refman/8.0/en/insert.html
type MySqlValueListListSyntax struct {
	*M_Syntax
	*M_ListSyntax[*ValueListSyntax]
	*M_ValueListListSyntax
	RowConstructorList bool
}

func (this *MySqlValueListListSyntax) writeSql(builder *sqlBuilder) {
	if len(this.elements) > 0 {
		if this.RowConstructorList {
			builder.writeStr("ROW")
		}
		builder.writeSyntax(this.elements[0])
		for i := 1; i < len(this.elements); i++ {
			builder.writeStr(this.separator)
			builder.writeSpaceOrLf(this, false)
			if this.RowConstructorList {
				builder.writeStr("ROW")
			}
			builder.writeSyntax(this.elements[i])
		}
	}
}

func NewMySqlValueListListSyntax() *MySqlValueListListSyntax {
	s := &MySqlValueListListSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	s.M_ListSyntax = ExtendListSyntax[*ValueListSyntax](s)
	s.M_ValueListListSyntax = ExtendValueListListSyntax(s)
	return s
}
