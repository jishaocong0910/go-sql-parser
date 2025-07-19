package enum

import o "github.com/jishaocong0910/go-object-util"

// 解析表引用语法级别
type TableReferenceSyntaxLevel struct {
	*o.EnumElem__
}

type _TableReferenceSyntaxLevel struct {
	*o.Enum__[TableReferenceSyntaxLevel]
	// 表名称引用
	TABLE_NAME,
	// 派生表
	DERIVED,
	// 连接的表
	JOIN TableReferenceSyntaxLevel
}

var TableReferenceSyntaxLevel_ = o.NewEnum[TableReferenceSyntaxLevel](_TableReferenceSyntaxLevel{})
