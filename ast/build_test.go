package ast_test

import (
	"testing"

	"github.com/jishaocong0910/go-sql-parser/ast"
	"github.com/jishaocong0910/go-sql-parser/enum"
	"github.com/stretchr/testify/require"
)

func TestBuildSql(t *testing.T) {
	r := require.New(t)
	o1 := ast.NewMySqlIdentifierSyntax()
	o1.Name = "col_1"
	o2 := ast.NewMySqlStringSyntax()
	o2.SetValue("abc")
	b1 := ast.NewBinaryOperationSyntax()
	b1.LeftOperand = o1
	b1.RightOperand = o2
	b1.BinaryOperator = enum.BinaryOperators.EQUAL_OR_ASSIGNMENT

	o3 := ast.NewIdentifierSyntax()
	o3.Name = "col_2"
	o4 := ast.NewBinaryNumberSyntax()
	o4.Sql = "0"
	b2 := ast.NewBinaryOperationSyntax()
	b2.LeftOperand = o3
	b2.RightOperand = o4
	b2.BinaryOperator = enum.BinaryOperators.GREATER_THAN

	b3 := ast.NewBinaryOperationSyntax()
	b3.LeftOperand = b1
	b3.RightOperand = b2
	b3.BinaryOperator = enum.BinaryOperators.BOOLEAN_AND

	w := ast.NewWhereSyntax()
	w.Condition = b3
	sql := ast.BuildSql(w, false)
	r.Equal("WHERE col_1 = 'abc' AND col_2 > 0", sql)
}
