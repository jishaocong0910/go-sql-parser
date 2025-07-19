package enum

import o "github.com/jishaocong0910/go-object-util"

// 解析查询语法级别
type QuerySyntaxLevel struct {
	*o.EnumElem__
}

type _QuerySyntaxLevel struct {
	*o.Enum__[QuerySyntaxLevel]
	QUERY_OPERAND,
	NORMAL QuerySyntaxLevel
}

var QuerySyntaxLevel_ = o.NewEnum[QuerySyntaxLevel](_QuerySyntaxLevel{})
