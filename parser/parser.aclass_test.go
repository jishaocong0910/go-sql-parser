package parser_test

import (
	"testing"

	"github.com/jishaocong0910/go-sql-parser/parser"

	. "github.com/jishaocong0910/go-sql-parser/enum"

	"github.com/stretchr/testify/require"
)

func TestNewParserPanic(t *testing.T) {
	r := require.New(t)
	_, err := parser.Parse(Dialect_.ORACLE, "")
	r.EqualError(err, "not supported database type for 'Oracle'")
}
