package ast_test

import (
	"fmt"
	"testing"

	"github.com/jishaocong0910/go-sql-parser/ast"

	"github.com/stretchr/testify/require"
)

func TestNotSupportedDialect(t *testing.T) {
	r := require.New(t)
	st := ast.NewNotSupportedDialectStatement()
	_, err := ast.Visit(st)
	r.EqualError(err, fmt.Sprintf("not supported dialect of '%s' yet", st.Dialect().Name))
}
