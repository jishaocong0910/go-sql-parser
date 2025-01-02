# go-sql-parser

用于Golang的SQL解析器，功能如下：

* 解析SQL字符串为语法对象，分析语法对象信息。
* 手动创建语法对象，构建SQL字符串，可格式化。

![coverage](https://raw.githubusercontent.com/jishaocong0910/go-sql-parser/b6780c369b7ba1c19d37d570893503172bbeaf3a/.badges/main/coverage.svg?token=BKK7GPSL7NYJFNEKEBYNYH3HO3CPS)

> [!NOTE]
>
> 本项目采用**go-object风格**编写：https://github.com/jishaocong0910/go-object

# 支持与限制

| 数据库        | 版本    |
|------------|-------|
| MySQL      | 8.4   |
| Oracle     | 🚧计划中 |
| PostgreSQL | 🚧计划中 |
| SQLserver  | 🚧计划中 |

## 限制

* 只支持SELECT、UPDATE、INSERT、DELETE语句。
* 明确不支持项在以下列出，若发现其他不支持请报告作者。

### MySQL

* 不支持`SELECT ... INTO `
* 不支持`WITH ... SELECT`
* 不支持`INSERT ... VALUES ROW(...)[,ROW(...)]`
* 不支持`INSERT ... SELECT`
* 一些特殊语法的函数不支持，如：</br>
  `CAST(expr AS type [ARRAY])`（其中的`[ARRAY]`）</br>
  `POSITION`（可使用`LOCATE`代替）</br>
  `JSON_TABLE`</br>
  `ST_Collect`</br>
  `WEIGHT_STRING`</br>

# 用法

```go
package main

import (
    "fmt"
    "go-sql-parser/ast"
    "go-sql-parser/enum"
    "go-sql-parser/parser"
    "log"
)

func main() {
    sql := "select id, col1, col2 from tab1 where col3=1"
    fmt.Printf("SQL: %s\n", sql)
    // 解析SQL为语法对象
    s, err := parser.Parse(enum.Dialects.MYSQL, sql)
    if err != nil {
      log.Fatalf("%+v", err)
    }
    // 将语法对象构建成SQL并格式化
    fmt.Printf("formatted:\n----------------------\n%s\n", ast.BuildSql(s, true))
    // 访问语法对象
    v, err := ast.Visit(s)
    if err != nil {
      log.Fatalf("%+v", err)
    }
    fmt.Println("----------------------")
    // SQL操作类型
    fmt.Printf("operation type: %v\n", v.SqlOperationType().ID())
    // 是否单表SQL语句
    fmt.Printf("is single table SQL: %v\n", v.SingleTableSql())
    // 单表SQL语句中的表名（v.SingleTableSql()为true时有值）
    fmt.Printf("table of single table sql: %v\n", v.TableOfSingleTableSql())
    // 所有涉及的表
    fmt.Printf("all tables: %v\n", v.TablesRaw())
    // 所有涉及的表字段
    fmt.Printf("all columns: %v\n", v.TableColumnsRaw())
    // 所有涉及的表的查询字段
    fmt.Printf("select columns: %v\n", v.SelectColumnsRaw())
    // 所有涉及的表的条件字段
    fmt.Printf("where columns: %v\n", v.WhereColumnsRaw())
    // Output:
    // SQL: select id, col1, col2 from tab1 where col3=1
    // formatted:
    // ----------------------
    // SELECT
    //   id,
    //   col1,
    //   col2
    // FROM
    //   tab1
    // WHERE col3 = 1
    // ----------------------
    // operation type: SELECT
    // is single table SQL: true
    // table of single table sql: tab1
    // all tables: [tab1]
    // all columns: map[tab1:[col2 id col3 col1]]
    // select columns: map[tab1:[col2 id col1]]
    // where columns: map[tab1:[col3]]
}
```
















