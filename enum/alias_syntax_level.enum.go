package enum

import o "github.com/jishaocong0910/go-object"

// 别名解析级别
type AliasSyntaxLevel struct {
	*o.M_EnumValue
}

type _AliasSyntaxLevel struct {
	*o.M_Enum[AliasSyntaxLevel]
	// 只解析标识符作为别名
	IDENTIFIER,
	// 字符串也可作为别名
	STRING AliasSyntaxLevel
}

var AliasSyntaxLevels = o.NewEnum[AliasSyntaxLevel](_AliasSyntaxLevel{})
