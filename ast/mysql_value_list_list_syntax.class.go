package ast

// MySqlValueListListSyntax
//
//	@Description: https://dev.mysql.com/doc/refman/8.0/en/insert.html
type MySqlValueListListSyntax struct {
	*Syntax__
	*ListSyntax__[*ValueListSyntax]
	*ValueListListSyntax__
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
	s.Syntax__ = ExtendSyntax(s)
	s.ListSyntax__ = ExtendListSyntax[*ValueListSyntax](s)
	s.ValueListListSyntax__ = ExtendValueListListSyntax(s)
	return s
}
