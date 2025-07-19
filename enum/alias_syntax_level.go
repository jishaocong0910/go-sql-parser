package enum

import o "github.com/jishaocong0910/go-object-util"

// 别名解析级别
type AliasSyntaxLevel struct {
	*o.EnumElem__
}

type _AliasSyntaxLevel struct {
	*o.Enum__[AliasSyntaxLevel]
	// 只解析标识符作为别名
	IDENTIFIER,
	// 字符串也可作为别名
	STRING AliasSyntaxLevel
}

var AliasSyntaxLevel_ = o.NewEnum[AliasSyntaxLevel](_AliasSyntaxLevel{})
