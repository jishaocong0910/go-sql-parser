package parser_test

import (
	. "go-sql-parser/parser"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCharDesc(t *testing.T) {
	r := require.New(t)
	r.Equal("(EOI)", Character.CharDesc(0))
	r.Equal("(ASCII: 1)", Character.CharDesc(1))
	r.Equal("(SPACE)", Character.CharDesc(' '))
	r.Equal("(COMMA)", Character.CharDesc(','))
	r.Equal("(SINGLE QUOTE)", Character.CharDesc('\''))
	r.Equal("(DOUBLE QUOTE)", Character.CharDesc('"'))
	r.Equal("(NON-BREAKING SPACE)", Character.CharDesc(160))
	r.Equal("'A'", Character.CharDesc('A'))
}
