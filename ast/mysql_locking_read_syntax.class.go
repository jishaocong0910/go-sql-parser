package ast

import "go-sql-parser/enum"

type MySqlLockingReadSyntax struct {
	*M_Syntax
	LockingRead            enum.MySqlLockingRead
	OfTableName            *MySqlIdentifierSyntax           // 8.0新增
	LockingReadConcurrency enum.MySqlLockingReadConcurrency // 8.0新增
}

func (this *MySqlLockingReadSyntax) accept(I_Visitor) {}

func (this *MySqlLockingReadSyntax) writeSql(builder *sqlBuilder) {
	builder.writeStr(this.LockingRead.Sql)
	if this.OfTableName != nil {
		builder.writeSpace()
		builder.writeSyntax(this.OfTableName)
	}
	if !this.LockingReadConcurrency.Undefined() {
		builder.writeSpace()
		builder.writeStr(this.LockingReadConcurrency.Sql)
	}
}

func NewMySqlLockingReadSyntax() *MySqlLockingReadSyntax {
	s := &MySqlLockingReadSyntax{}
	s.M_Syntax = ExtendSyntax(s)
	return s
}
