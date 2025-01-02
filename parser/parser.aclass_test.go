package parser_test

import (
	. "go-sql-parser/enum"
	"go-sql-parser/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParserPanic(t *testing.T) {
	r := require.New(t)
	_, err := parser.Parse(Dialects.ORACLE, "")
	r.EqualError(err, "not supported database type for 'Oracle'")
}
