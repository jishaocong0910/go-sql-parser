package enum

import o "github.com/jishaocong0910/go-object"

// 解析表引用语法级别
type TableReferenceSyntaxLevel struct {
	*o.M_EnumValue
}

type _TableReferenceSyntaxLevel struct {
	*o.M_Enum[TableReferenceSyntaxLevel]
	// 表名称引用
	TABLE_NAME,
	// 派生表
	DERIVED,
	// 连接的表
	JOIN TableReferenceSyntaxLevel
}

var TableReferenceSyntaxLevels = o.NewEnum[TableReferenceSyntaxLevel](_TableReferenceSyntaxLevel{})
