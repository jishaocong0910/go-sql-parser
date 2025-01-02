package enum

import o "github.com/jishaocong0910/go-object"

// 解析查询语法级别
type QuerySyntaxLevel struct {
	*o.M_EnumValue
}

type _QuerySyntaxLevel struct {
	*o.M_Enum[QuerySyntaxLevel]
	QUERY_OPERAND,
	NORMAL QuerySyntaxLevel
}

var QuerySyntaxLevels = o.NewEnum[QuerySyntaxLevel](_QuerySyntaxLevel{})
