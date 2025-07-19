package enum

import o "github.com/jishaocong0910/go-object-util"

// 连接类型
type JoinType struct {
	*o.EnumElem__
	Sql string
}

type _JoinType struct {
	*o.Enum__[JoinType]
	COMMA,
	JOIN,
	INNER_JOIN,
	CROSS_JOIN,
	NATURAL_JOIN,
	NATURAL_INNER_JOIN,
	LEFT_OUTER_JOIN,
	RIGHT_OUTER_JOIN,
	FULL_OUTER_JOIN,
	STRAIGHT_JOIN,
	OUTER_APPLY,
	CROSS_APPLY JoinType
}

var JoinType_ = o.NewEnum[JoinType](_JoinType{
	COMMA:              JoinType{Sql: ","},
	JOIN:               JoinType{Sql: "JOIN"},
	INNER_JOIN:         JoinType{Sql: "INNER JOIN"},
	CROSS_JOIN:         JoinType{Sql: "CROSS JOIN"},
	NATURAL_JOIN:       JoinType{Sql: "NATURAL JOIN"},
	NATURAL_INNER_JOIN: JoinType{Sql: "NATURAL INNER JOIN"},
	LEFT_OUTER_JOIN:    JoinType{Sql: "LEFT JOIN"},
	RIGHT_OUTER_JOIN:   JoinType{Sql: "RIGHT JOIN"},
	FULL_OUTER_JOIN:    JoinType{Sql: "FULL JOIN"},
	STRAIGHT_JOIN:      JoinType{Sql: "STRAIGHT_JOIN"},
	OUTER_APPLY:        JoinType{Sql: "OUTER APPLY"},
	CROSS_APPLY:        JoinType{Sql: "CROSS APPLY"},
})
