package enum

import o "github.com/jishaocong0910/go-object-util"

type MySqlIndexHintFor struct {
	*o.EnumElem__
	Sql string
}

type _MySqlIndexHintFor struct {
	*o.Enum__[MySqlIndexHintFor]
	JOIN,
	ORDER_BY,
	GROUP_BY MySqlIndexHintFor
}

var MySqlIndexHintFor_ = o.NewEnum[MySqlIndexHintFor](_MySqlIndexHintFor{
	JOIN:     MySqlIndexHintFor{Sql: "JOIN"},
	ORDER_BY: MySqlIndexHintFor{Sql: "ORDER BY"},
	GROUP_BY: MySqlIndexHintFor{Sql: "GROUP BY"},
})
