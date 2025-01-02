package parser_test

import (
	"go-sql-parser/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func validatePanic(t *testing.T, sql, msg string) {
	r := require.New(t)
	l := parser.NewMySqlLexer(sql)
	r.PanicsWithValue(msg, func() {
		parser.NextToken(l)
	})
}

func TestBaseLexer_nextChar(t *testing.T) {
	r := require.New(t)
	sql := "" +
		"SELECT *\n" +
		"  FROM tab_1;"
	l := parser.NewMySqlLexer(sql)
	for i, c := range sql {
		r.Equal(c, parser.GetCursorC(l))
		r.Equal(i, parser.GetCursorPos(l))
		parser.NextChar(l)
	}
	r.Equal(parser.Eoi, parser.NextChar(l))
}
