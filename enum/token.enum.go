package enum

import o "github.com/jishaocong0910/go-object"

// Token
// 收集了各种数据库SQL方言解析需要的Token，不同数据库解析SQL时根据自己方言的需要而使用，其中EOI为输入结束标识，其他大致可分以下类型。
// 1.标点符号（如：comma、lParen、rParen等）
// 2.不确定内容的（如：string、comment、identifier等）
// 3.SQL标准或大部分数据库都使用的公共保留字（如：select、create、when等）
// 4.数据类型保留字（如：bit、int、varchar等）
// 5.某数据库方言的特有保留字（如：sqlBigResult、highPriority、match、geometry这些是MySQL方言中特有的，当然不排除其他数据库将来也使用这些）
type Token struct {
	*o.M_EnumValue
	Sql string
}

type _Token struct {
	*o.M_Enum[Token]
	EOI, // end of input
	COMMENT,
	STRING,
	DECIMAL_NUMBER,
	BINARY_NUMBER,
	HEXADECIMAL_NUMBER,
	IDENTIFIER,
	AT,
	AT_AT,
	COMMA,
	L_PAREN,
	R_PAREN,
	L_BRACKET,
	R_BRACKET,
	L_BRACE,
	R_BRACE,
	SEMI,
	DOT,
	STAR,
	SLASH,
	PLUS,
	SUB,
	PERCENT,
	EQ,
	EQ_EQ,
	GT,
	GT_GT,
	GT_EQ,
	LT,
	LT_LT,
	LT_EQ,
	LT_GT,
	LT_EQ_GT,
	BANG,
	BANG_EQ,
	BANG_GT, // MySql不支持
	BANG_LT, // MySql不支持
	AMP,
	AMP_AMP,
	BAR,
	BAR_BAR,
	TILDE,
	COLON_EQ,
	CARET,
	QUES,
	DIV, // MySql有使用
	MOD, // MySql有使用
	JSON_EXTRACT, // MySql有使用
	JSON_UNQUOTE, // MySql有使用
	SELECT,
	INSERT,
	UPDATE,
	DELETE,
	CREATE,
	ALTER,
	CASE,
	WHEN,
	THEN,
	ELSE,
	DISTINCT,
	DISTINCTROW, // MySql、Oracle有使用
	ALL,
	VALUES,
	TABLE,
	ROW,
	ROWS,
	RANGE,
	AS,
	FROM,
	CROSS,
	INNER,
	LEFT,
	RIGHT,
	NATURAL,
	OUTER,
	JOIN,
	ON,
	USING,
	USE,
	INDEX,
	FOR,
	FORCE,
	ORDER,
	GROUP,
	BY,
	WHERE,
	AND,
	OR,
	ASC,
	DESC,
	COLLATE,
	SET,
	LOCK,
	IN,
	WITH,
	HAVING,
	BETWEEN,
	LIKE,
	IS,
	NULL,
	NOT,
	XOR,
	R_LIKE,
	UNION,
	EXCEPT, // MySQL有使用
	INTERSECT, // MySql、Oracle有使用
	INTO,
	EXISTS,
	ZEROFILL,
	REGEXP,
	CHARACTER,
	LIMIT,
	HIGHP_RIORITY,
	STRAIGHT_JOIN,
	SQL_SMALL_RESULT,
	SQL_BIG_RESULT,
	SQL_CALC_FOUND_ROWS,
	UNSIGNED,
	LEADING,
	BOTH,
	TRAILING,
	TRUE,
	FALSE,
	INTERVAL,
	LOW_PRIORITY,
	IGNORE,
	PARTITION,
	DELAYED,
	KEY,
	SECOND_MICROSECOND,
	MINUTE_MICROSECOND,
	MINUTESECOND,
	HOUR_MICROSECOND,
	HOUR_SECOND,
	HOUR_MINUTE,
	DAY_MICROSECOND,
	DAY_SECOND,
	DAY_MINUTE,
	DAY_HOUR,
	YEAR_MONTH,
	SEPARATOR,
	PROCEDURE,
	OF,
	INTEGER,
	INT1,
	INT2,
	INT3,
	INT4,
	INT8,
	FLOAT4,
	FLOAT8,
	BIT,
	TINYINT,
	SMALLINT,
	MEDIUMINT,
	INT,
	BIGINT,
	FLOAT,
	DOUBLE,
	DECIMAL,
	VARCHAR,
	CHAR,
	BINARY,
	VARBINARY,
	MEDIUMTEXT,
	TINYTEXT,
	TEXT,
	LONGTEXT,
	ENUM,
	TINYBLOB,
	BLOB,
	MEDIUMBLOB,
	LONGBLOB,
	GEOMETRY,
	POINT,
	LINESTRING,
	POLYGON,
	GEOMETRYCOLLECTION,
	MULTIPOINT,
	MULTILINESTRING,
	MULTIPOLYGON,
	JSON,
	OVER,
	WINDOW,
	DEFAULT,
	DUAL Token
}

var Tokens = o.NewEnum[Token](_Token{
	EOI:                 Token{},
	COMMENT:             Token{},
	STRING:              Token{},
	DECIMAL_NUMBER:      Token{},
	BINARY_NUMBER:       Token{},
	HEXADECIMAL_NUMBER:  Token{},
	IDENTIFIER:          Token{},
	AT:                  Token{},
	AT_AT:               Token{},
	COMMA:               Token{Sql: ","},
	L_PAREN:             Token{Sql: "("},
	R_PAREN:             Token{Sql: "},"},
	L_BRACKET:           Token{Sql: "["},
	R_BRACKET:           Token{Sql: "]"},
	L_BRACE:             Token{Sql: "{"},
	R_BRACE:             Token{Sql: "}"},
	SEMI:                Token{Sql: ";"},
	DOT:                 Token{Sql: "."},
	STAR:                Token{Sql: "*"},
	SLASH:               Token{Sql: "/"},
	PLUS:                Token{Sql: "+"},
	SUB:                 Token{Sql: "-"},
	PERCENT:             Token{Sql: "%"},
	EQ:                  Token{Sql: "="},
	EQ_EQ:               Token{Sql: "=="},
	GT:                  Token{Sql: ">"},
	GT_GT:               Token{Sql: ">>"},
	GT_EQ:               Token{Sql: ">="},
	LT:                  Token{Sql: "<"},
	LT_LT:               Token{Sql: "<<"},
	LT_EQ:               Token{Sql: "<="},
	LT_GT:               Token{Sql: "<>"},
	LT_EQ_GT:            Token{Sql: "<=>"},
	BANG:                Token{Sql: "!"},
	BANG_EQ:             Token{Sql: "!="},
	BANG_GT:             Token{Sql: "!>"},
	BANG_LT:             Token{Sql: "!<"},
	AMP:                 Token{Sql: "&"},
	AMP_AMP:             Token{Sql: "&&"},
	BAR:                 Token{Sql: "|"},
	BAR_BAR:             Token{Sql: "||"},
	TILDE:               Token{Sql: "~"},
	COLON_EQ:            Token{Sql: ":="},
	CARET:               Token{Sql: "^"},
	QUES:                Token{Sql: "?"},
	DIV:                 Token{Sql: "DIV"},
	MOD:                 Token{Sql: "MOD"},
	JSON_EXTRACT:        Token{Sql: "->"},
	JSON_UNQUOTE:        Token{Sql: "->>"},
	SELECT:              Token{Sql: "SELECT"},
	INSERT:              Token{Sql: "INSERT"},
	UPDATE:              Token{Sql: "UPDATE"},
	DELETE:              Token{Sql: "DELETE"},
	CREATE:              Token{Sql: "CREATE"},
	ALTER:               Token{Sql: "ALTER"},
	CASE:                Token{Sql: "CASE"},
	WHEN:                Token{Sql: "WHEN"},
	THEN:                Token{Sql: "THEN"},
	ELSE:                Token{Sql: "ELSE"},
	DISTINCT:            Token{Sql: "DISTINCT"},
	DISTINCTROW:         Token{Sql: "DISTINCTROW"},
	ALL:                 Token{Sql: "ALL"},
	VALUES:              Token{Sql: "VALUES"},
	TABLE:               Token{Sql: "TABLE"},
	ROW:                 Token{Sql: "ROW"},
	ROWS:                Token{Sql: "ROWS"},
	RANGE:               Token{Sql: "RANGE"},
	AS:                  Token{Sql: "AS"},
	FROM:                Token{Sql: "FROM"},
	CROSS:               Token{Sql: "CROSS"},
	INNER:               Token{Sql: "INNER"},
	LEFT:                Token{Sql: "LEFT"},
	RIGHT:               Token{Sql: "RIGHT"},
	NATURAL:             Token{Sql: "NATURAL"},
	OUTER:               Token{Sql: "OUTER"},
	JOIN:                Token{Sql: "JOIN"},
	ON:                  Token{Sql: "ON"},
	USING:               Token{Sql: "USING"},
	USE:                 Token{Sql: "USE"},
	INDEX:               Token{Sql: "INDEX"},
	FOR:                 Token{Sql: "FOR"},
	FORCE:               Token{Sql: "FORCE"},
	ORDER:               Token{Sql: "ORDER"},
	GROUP:               Token{Sql: "GROUP"},
	BY:                  Token{Sql: "BY"},
	WHERE:               Token{Sql: "WHERE"},
	AND:                 Token{Sql: "AND"},
	OR:                  Token{Sql: "OR"},
	ASC:                 Token{Sql: "ASC"},
	DESC:                Token{Sql: "DESC"},
	COLLATE:             Token{Sql: "COLLATE"},
	SET:                 Token{Sql: "SET"},
	LOCK:                Token{Sql: "LOCK"},
	IN:                  Token{Sql: "IN"},
	WITH:                Token{Sql: "WITH"},
	HAVING:              Token{Sql: "HAVING"},
	BETWEEN:             Token{Sql: "BETWEEN"},
	LIKE:                Token{Sql: "LIKE"},
	IS:                  Token{Sql: "IS"},
	NULL:                Token{Sql: "NULL"},
	NOT:                 Token{Sql: "NOT"},
	XOR:                 Token{Sql: "XOR"},
	R_LIKE:              Token{Sql: "RLIKE"},
	UNION:               Token{Sql: "UNION"},
	EXCEPT:              Token{Sql: "EXCEPT"},
	INTERSECT:           Token{Sql: "INTERSECT"},
	INTO:                Token{Sql: "INTO"},
	EXISTS:              Token{Sql: "EXISTS"},
	ZEROFILL:            Token{Sql: "ZEROFILL"},
	REGEXP:              Token{Sql: "REGEXP"},
	CHARACTER:           Token{Sql: "CHARACTER"},
	LIMIT:               Token{Sql: "LIMIT"},
	HIGHP_RIORITY:       Token{Sql: "HIGH_PRIORITY"},
	STRAIGHT_JOIN:       Token{Sql: "STRAIGHT_JOIN"},
	SQL_SMALL_RESULT:    Token{Sql: "SQL_SMALL_RESULT"},
	SQL_BIG_RESULT:      Token{Sql: "SQL_BIG_RESULT"},
	SQL_CALC_FOUND_ROWS: Token{Sql: "SQL_CALC_FOUND_ROWS"},
	UNSIGNED:            Token{Sql: "UNSIGNED"},
	LEADING:             Token{Sql: "LEADING"},
	BOTH:                Token{Sql: "BOTH"},
	TRAILING:            Token{Sql: "TRAILING"},
	TRUE:                Token{Sql: "TRUE"},
	FALSE:               Token{Sql: "FALSE"},
	INTERVAL:            Token{Sql: "INTERVAL"},
	LOW_PRIORITY:        Token{Sql: "LOW_PRIORITY"},
	IGNORE:              Token{Sql: "IGNORE"},
	PARTITION:           Token{Sql: "PARTITION"},
	DELAYED:             Token{Sql: "DELAYED"},
	KEY:                 Token{Sql: "KEY"},
	SECOND_MICROSECOND:  Token{Sql: "SECOND_MICROSECOND"},
	MINUTE_MICROSECOND:  Token{Sql: "MINUTE_MICROSECOND"},
	MINUTESECOND:        Token{Sql: "MINUTE_SECOND"},
	HOUR_MICROSECOND:    Token{Sql: "HOUR_MICROSECOND"},
	HOUR_SECOND:         Token{Sql: "HOUR_SECOND"},
	HOUR_MINUTE:         Token{Sql: "HOUR_MINUTE"},
	DAY_MICROSECOND:     Token{Sql: "DAY_MICROSECOND"},
	DAY_SECOND:          Token{Sql: "DAY_SECOND"},
	DAY_MINUTE:          Token{Sql: "DAY_MINUTE"},
	DAY_HOUR:            Token{Sql: "DAY_HOUR"},
	YEAR_MONTH:          Token{Sql: "YEAR_MONTH"},
	SEPARATOR:           Token{Sql: "SEPARATOR"},
	PROCEDURE:           Token{Sql: "PROCEDURE"},
	OF:                  Token{Sql: "OF"},
	INTEGER:             Token{Sql: "INTEGER"},
	INT1:                Token{Sql: "INT1"},
	INT2:                Token{Sql: "INT2"},
	INT3:                Token{Sql: "INT3"},
	INT4:                Token{Sql: "INT4"},
	INT8:                Token{Sql: "INT8"},
	FLOAT4:              Token{Sql: "FLOAT4"},
	FLOAT8:              Token{Sql: "FLOAT8"},
	BIT:                 Token{Sql: "BIT"},
	TINYINT:             Token{Sql: "TINYINT"},
	SMALLINT:            Token{Sql: "SMALLINT"},
	MEDIUMINT:           Token{Sql: "MEDIUMINT"},
	INT:                 Token{Sql: "INT"},
	BIGINT:              Token{Sql: "BIGINT"},
	FLOAT:               Token{Sql: "FLOAT"},
	DOUBLE:              Token{Sql: "DOUBLE"},
	DECIMAL:             Token{Sql: "DECIMAL"},
	VARCHAR:             Token{Sql: "VARCHAR"},
	CHAR:                Token{Sql: "CHAR"},
	BINARY:              Token{Sql: "BINARY"},
	VARBINARY:           Token{Sql: "VARBINARY"},
	MEDIUMTEXT:          Token{Sql: "MEDIUMTEXT"},
	TINYTEXT:            Token{Sql: "TINYTEXT"},
	TEXT:                Token{Sql: "TEXT"},
	LONGTEXT:            Token{Sql: "LONGTEXT"},
	ENUM:                Token{Sql: "ENUM"},
	TINYBLOB:            Token{Sql: "TINYBLOB"},
	BLOB:                Token{Sql: "BLOB"},
	MEDIUMBLOB:          Token{Sql: "MEDIUMBLOB"},
	LONGBLOB:            Token{Sql: "LONGBLOB"},
	GEOMETRY:            Token{Sql: "GEOMETRY"},
	POINT:               Token{Sql: "POINT"},
	LINESTRING:          Token{Sql: "LINESTRING"},
	POLYGON:             Token{Sql: "POLYGON"},
	GEOMETRYCOLLECTION:  Token{Sql: "GEOMETRYCOLLECTION"},
	MULTIPOINT:          Token{Sql: "MULTIPOINT"},
	MULTILINESTRING:     Token{Sql: "MULTILINESTRING"},
	MULTIPOLYGON:        Token{Sql: "MULTIPOLYGON"},
	JSON:                Token{Sql: "JSON"},
	OVER:                Token{Sql: "OVER"},
	WINDOW:              Token{Sql: "WINDOW"},
	DEFAULT:             Token{Sql: "DEFAULT"},
	DUAL:                Token{Sql: "DUAL"},
})
