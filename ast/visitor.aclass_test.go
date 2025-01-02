package ast_test

import (
	"fmt"
	"go-sql-parser/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotSupportedDialect(t *testing.T) {
	r := require.New(t)
	st := ast.NewNotSupportedDialectStatement()
	_, err := ast.Visit(st)
	r.EqualError(err, fmt.Sprintf("not supported dialect of '%s' yet", st.Dialect().Name))
}
