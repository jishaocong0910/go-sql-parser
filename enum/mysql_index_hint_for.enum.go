package enum

import o "github.com/jishaocong0910/go-object"

type MySqlIndexHintFor struct {
	*o.M_EnumValue
	Sql string
}

type _MySqlIndexHintFor struct {
	*o.M_Enum[MySqlIndexHintFor]
	JOIN,
	ORDER_BY,
	GROUP_BY MySqlIndexHintFor
}

var MySqlIndexHintFors = o.NewEnum[MySqlIndexHintFor](_MySqlIndexHintFor{
	JOIN:     MySqlIndexHintFor{Sql: "JOIN"},
	ORDER_BY: MySqlIndexHintFor{Sql: "ORDER BY"},
	GROUP_BY: MySqlIndexHintFor{Sql: "GROUP BY"},
})
