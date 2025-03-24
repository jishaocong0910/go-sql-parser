package parser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jishaocong0910/go-sql-parser/parser"

	"github.com/jishaocong0910/go-sql-parser/ast"

	. "github.com/jishaocong0910/go-sql-parser/enum"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	r := require.New(t)
	sql := "SELECT 1+1"
	s, err := parser.Parse(Dialects.MYSQL, sql)
	r.NoError(err)
	fmt.Println(ast.BuildSql(s, false))
}

func validateSql(t *testing.T, sql, notFormatSql, formatSql string) {
	r := require.New(t)
	s, err := parser.Parse(Dialects.MYSQL, sql)
	r.NoError(err)
	if notFormatSql != "" {
		r.Equal(notFormatSql, ast.BuildSql(s, false))
	}
	if formatSql != "" {
		r.Equal(formatSql, ast.BuildSql(s, true))
	}
}

func validateError(r *require.Assertions, sql, msg string) {
	_, err := parser.Parse(Dialects.MYSQL, sql)
	r.EqualError(err, msg)
}

func TestAggregateFunctionSyntax(t *testing.T) {
	validateSql(t, "select count(DISTINCT name,level) FROM tab1",
		"SELECT COUNT(DISTINCT name, level) FROM tab1",
		`SELECT
  COUNT(DISTINCT name, level)
FROM
  tab1`)
}

func TestAllColumnSyntax(t *testing.T) {
	validateSql(t, "select * FROM tab1",
		"SELECT * FROM tab1",
		`SELECT
  *
FROM
  tab1`)
}

func TestAliasSyntax(t *testing.T) {
	validateSql(t, `select col1 'c1', col2 c2, col3 "c3", col1 as 'c1', col2 as c2, col3 as "c3" from tab1 as t1`,
		"SELECT col1 AS 'c1', col2 AS c2, col3 AS \"c3\", col1 AS 'c1', col2 AS c2, col3 AS \"c3\" FROM tab1 t1",
		`SELECT
  col1 AS 'c1',
  col2 AS c2,
  col3 AS "c3",
  col1 AS 'c1',
  col2 AS c2,
  col3 AS "c3"
FROM
  tab1 t1`)
}

func TestBinaryNumberSyntax(t *testing.T) {
	validateSql(t, "select 0b11010 FROM tab1",
		"SELECT 0b11010 FROM tab1",
		`SELECT
  0b11010
FROM
  tab1`)
}

func TestCaseSyntax(t *testing.T) {
	validateSql(t, `select
  case
    (select
      t2.level
    FROM
      t_level t2
    WHERE t2.customer_id = t.id)
    WHEN 1
    THEN 'gold'
    WHEN 2
    THEN 'diamond'
    else 'bronze'
  end level
FROM
  t_customer_info t;`,
		"SELECT CASE (SELECT t2.level FROM t_level t2 WHERE t2.customer_id = t.id) WHEN 1 THEN 'gold' WHEN 2 THEN 'diamond' ELSE 'bronze' END AS level FROM t_customer_info t",
		`SELECT
  CASE (SELECT
          t2.level
        FROM
          t_level t2
        WHERE t2.customer_id = t.id)
  WHEN 1 THEN 'gold'
  WHEN 2 THEN 'diamond'
  ELSE 'bronze' END AS level
FROM
  t_customer_info t`)
	validateSql(t, `update user set status=case id when 1 then 2 when 2 then 3 else status end,phone=case id when 1 then '1234' else phone end,email=null where id in(1,2) and level<2`,
		`UPDATE user SET status = CASE id WHEN 1 THEN 2 WHEN 2 THEN 3 ELSE status END, phone = CASE id WHEN 1 THEN '1234' ELSE phone END, email = NULL WHERE id IN (1, 2) AND level < 2`,
		`UPDATE
  user
SET
  status = CASE id
           WHEN 1 THEN 2
           WHEN 2 THEN 3
           ELSE status END,
  phone = CASE id
          WHEN 1 THEN '1234'
          ELSE phone END,
  email = NULL
WHERE id IN (1, 2)
  AND level < 2`)
}

func TestColumnItemAssignmentSyntax(t *testing.T) {
	validateSql(t, "UPDATE tab1 SET col1='abc',col2=2 WHERE id=1",
		"UPDATE tab1 SET col1 = 'abc', col2 = 2 WHERE id = 1",
		`UPDATE
  tab1
SET
  col1 = 'abc',
  col2 = 2
WHERE id = 1`)
}

func TestDecimalNumberSyntax(t *testing.T) {
	validateSql(t, "SELECT 48,234.2,6583e12,6583e+12,6583E-24 FROM dual;",
		"SELECT 48, 234.2, 6583e12, 6583e+12, 6583E-24 FROM DUAL",
		`SELECT
  48,
  234.2,
  6583e12,
  6583e+12,
  6583E-24
FROM
  DUAL`)
}

func TestDerivedTableTableReferenceSyntax(t *testing.T) {
	validateSql(t, "SELECT a.col1,b.col2 FROM tab1 a JOIN (select col2 FROM tab2) b ON a.id=b.pid;",
		"SELECT a.col1, b.col2 FROM tab1 a JOIN (SELECT col2 FROM tab2) b ON a.id = b.pid",
		`SELECT
  a.col1,
  b.col2
FROM
  tab1 a
  JOIN (SELECT
          col2
        FROM
          tab2) b
    ON a.id = b.pid`)
}

func TestExistsSyntax(t *testing.T) {
	validateSql(t, "SELECT * FROM tab1 a WHERE EXISTS(select 1 FROM tab2 b WHERE a.id=b.pid);",
		"SELECT * FROM tab1 a WHERE EXISTS(SELECT 1 FROM tab2 b WHERE a.id = b.pid)",
		`SELECT
  *
FROM
  tab1 a
WHERE EXISTS(SELECT
               1
             FROM
               tab2 b
             WHERE a.id = b.pid)`)
}

func TestExprListSyntax(t *testing.T) {
	validateSql(t, "select (1,2)=(1,2)",
		"SELECT (1, 2) = (1, 2)",
		`SELECT
  (1, 2) = (1, 2)`)
}

func TestHavingSyntax(t *testing.T) {
	validateSql(t, "select name FROM student WHERE gender='female' GROUP BY name HAVING sum(score)>210;",
		"SELECT name FROM student WHERE gender = 'female' GROUP BY name HAVING SUM(score) > 210",
		`SELECT
  name
FROM
  student
WHERE gender = 'female'
GROUP BY
  name
HAVING
  SUM(score) > 210`)
	validateSql(t, "select name FROM student WHERE gender='female' HAVING score>210;",
		"SELECT name FROM student WHERE gender = 'female' HAVING score > 210",
		`SELECT
  name
FROM
  student
WHERE gender = 'female'
HAVING
  score > 210`)
}

func TestHexadecimalNumberSyntax(t *testing.T) {
	validateSql(t, "SELECT 0x4D7953514C",
		"SELECT 0x4D7953514C",
		"")
}

func TestIdentifierListSyntax(t *testing.T) {
	validateSql(t, "SELECT a.col1,b.col1 FROM tab1 a JOIN tab2 b USING(id,type)",
		"SELECT a.col1, b.col1 FROM tab1 a JOIN tab2 b USING(id, type)",
		`SELECT
  a.col1,
  b.col1
FROM
  tab1 a
  JOIN tab2 b
    USING(id, type)`)
}

func TestInsertColumnListSyntax(t *testing.T) {
	validateSql(t, "INSERT tab1(col1,col2,col3,col4) VALUES ('a',1,'b','c')",
		"INSERT INTO tab1 (col1, col2, col3, col4) VALUES ('a', 1, 'b', 'c')",
		`INSERT INTO
tab1
  (col1, col2, col3, col4)
VALUES
  ('a', 1, 'b', 'c')`)
}

func TestJoinOnSyntax(t *testing.T) {
	validateSql(t, "select a.col1,b.col1 FROM tab1 a JOIN tab2 b ON a.id=b.pid AND a.col2=b.col2 WHERE a.id=1",
		"SELECT a.col1, b.col1 FROM tab1 a JOIN tab2 b ON a.id = b.pid AND a.col2 = b.col2 WHERE a.id = 1",
		`SELECT
  a.col1,
  b.col1
FROM
  tab1 a
  JOIN tab2 b
    ON a.id = b.pid
   AND a.col2 = b.col2
WHERE a.id = 1`)
}

func TestJoinUsingSyntax(t *testing.T) {
	validateSql(t, "select a.col1,b.col1 FROM tab1 a JOIN tab2 b USING(col2,col3) WHERE a.col4=1",
		"SELECT a.col1, b.col1 FROM tab1 a JOIN tab2 b USING(col2, col3) WHERE a.col4 = 1",
		`SELECT
  a.col1,
  b.col1
FROM
  tab1 a
  JOIN tab2 b
    USING(col2, col3)
WHERE a.col4 = 1`)
}

func TestNStringSyntax(t *testing.T) {
	validateSql(t, "select N'abc'",
		"SELECT N'abc'",
		"")
}

func TestNullSyntax(t *testing.T) {
	validateSql(t, "select 1 IS NULL",
		"SELECT 1 IS NULL",
		"")
}

func TestOrderBySyntax(t *testing.T) {
	validateSql(t, "select * FROM tab1 ORDER BY col1 DESC, col2 ASC",
		"SELECT * FROM tab1 ORDER BY col1 DESC, col2 ASC",
		`SELECT
  *
FROM
  tab1
ORDER BY
  col1 DESC,
  col2 ASC`)
}

func TestFunctionSyntax(t *testing.T) {
	validateSql(t, "select myfunc('a',1)",
		"SELECT myfunc('a', 1)",
		`SELECT
  myfunc('a', 1)`)
	validateSql(t, "select CURRENT_TIMESTAMP",
		"SELECT CURRENT_TIMESTAMP",
		`SELECT
  CURRENT_TIMESTAMP`)

}

func TestParameterSyntax(t *testing.T) {
	validateSql(t, "INSERT INTO tab1 VALUES(?,?,?,?,?)",
		"INSERT INTO tab1 VALUES (?, ?, ?, ?, ?)",
		`INSERT INTO
tab1
VALUES
  (?, ?, ?, ?, ?)`)
}

func TestPropertySyntax(t *testing.T) {
	validateSql(t, "select a.col1,a.col2 FROM tab1 a",
		"SELECT a.col1, a.col2 FROM tab1 a",
		`SELECT
  a.col1,
  a.col2
FROM
  tab1 a`)
}

func TestSelectColumnSyntax(t *testing.T) {
	validateSql(t, "select *,abc FROM tab1",
		"SELECT *, abc FROM tab1",
		"")
}

func TestTableNameItemSyntax(t *testing.T) {
	validateSql(t, "select * FROM db1.tab1 t1",
		"SELECT * FROM db1.tab1 t1",
		"")
}

func TestValueListSyntax(t *testing.T) {
	validateSql(t, "INSERT INTO tab1 (col1,col2,col3) VALUES (?,?,?)",
		"INSERT INTO tab1 (col1, col2, col3) VALUES (?, ?, ?)",
		`INSERT INTO
tab1
  (col1, col2, col3)
VALUES
  (?, ?, ?)`)
}

func TestValueListListSyntax(t *testing.T) {
	validateSql(t, "INSERT INTO tab1 (col1,col2,col3) VALUES (?,?,?),(?,?,?),(?,?,?)",
		"INSERT INTO tab1 (col1, col2, col3) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?)",
		`INSERT INTO
tab1
  (col1, col2, col3)
VALUES
  (?, ?, ?),
  (?, ?, ?),
  (?, ?, ?)`)
}

func TestWhereSyntax(t *testing.T) {
	validateSql(t, "select * FROM tab1 WHERE col1='aa' AND col2=2",
		"SELECT * FROM tab1 WHERE col1 = 'aa' AND col2 = 2",
		`SELECT
  *
FROM
  tab1
WHERE col1 = 'aa'
  AND col2 = 2`)
}

func TestOverSyntax(t *testing.T) {
	validateSql(t, `SELECT
         year, country, product , profit,
         SUM(profit) OVER() AS total_profit,
         SUM(profit) OVER(PARTITION BY country) AS country_profit
       FROM sales
       ORDER BY country, year, product, profit;`,
		"SELECT year, country, product, profit, SUM(profit) OVER() AS total_profit, SUM(profit) OVER(PARTITION BY country) AS country_profit FROM sales ORDER BY country, year, product, profit",
		`SELECT
  year,
  country,
  product,
  profit,
  SUM(profit) OVER() AS total_profit,
  SUM(profit) OVER(PARTITION BY country) AS country_profit
FROM
  sales
ORDER BY
  country,
  year,
  product,
  profit`)
}

func TestWindowSpecSyntax(t *testing.T) {
	validateSql(t, `SELECT
	time,
	subject,
	val,
	ROW_NUMBER() OVER(PARTITION BY subject) AS row_num1,
	ROW_NUMBER() OVER(PARTITION BY subject ORDER BY time desc, val asc) AS row_num2,
	SUM(val) OVER (PARTITION BY subject ORDER BY time ROWS CURRENT ROW) AS running_total1,
	SUM(val) OVER (PARTITION BY subject ORDER BY time RANGE UNBOUNDED PRECEDING) AS running_total2,
	SUM(val) OVER (PARTITION BY subject ORDER BY time rows UNBOUNDED FOLLOWING) AS running_total3,
	SUM(val) OVER (PARTITION BY subject ORDER BY time rows 10 PRECEDING) AS running_total4,
	AVG(val) OVER (PARTITION BY subject ORDER BY time range BETWEEN current row AND UNBOUNDED FOLLOWING) AS running_average1,
	AVG(val) OVER (PARTITION BY subject ORDER BY time range between UNBOUNDED PRECEDING and UNBOUNDED FOLLOWING) AS running_average2,
	AVG(val) OVER (PARTITION BY subject ORDER BY time RANGE BETWEEN UNBOUNDED FOLLOWING AND current row) AS running_average3,
	AVG(val) OVER (PARTITION BY subject ORDER BY time RANGE BETWEEN 10 PRECEDING AND 10 FOLLOWING) AS running_average4,
	AVG(val) OVER (PARTITION BY subject ORDER BY time RANGE BETWEEN 10 FOLLOWING AND 10 PRECEDING) AS running_average4
FROM
	observations`,
		`SELECT time, subject, val, ROW_NUMBER() OVER(PARTITION BY subject) AS row_num1, ROW_NUMBER() OVER(PARTITION BY subject ORDER BY time DESC, val ASC) AS row_num2, SUM(val) OVER(PARTITION BY subject ORDER BY time ROWS CURRENT ROW) AS running_total1, SUM(val) OVER(PARTITION BY subject ORDER BY time RANGE UNBOUNDED PRECEDING) AS running_total2, SUM(val) OVER(PARTITION BY subject ORDER BY time ROWS UNBOUNDED PRECEDING) AS running_total3, SUM(val) OVER(PARTITION BY subject ORDER BY time ROWS 10 PRECEDING) AS running_total4, AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN CURRENT ROW AND UNBOUNDED PRECEDING) AS running_average1, AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED PRECEDING) AS running_average2, AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS running_average3, AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN 10 PRECEDING AND 10 PRECEDING) AS running_average4, AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN 10 PRECEDING AND 10 PRECEDING) AS running_average4 FROM observations`,
		`SELECT
  time,
  subject,
  val,
  ROW_NUMBER() OVER(PARTITION BY subject) AS row_num1,
  ROW_NUMBER() OVER(PARTITION BY subject ORDER BY time DESC, val ASC) AS row_num2,
  SUM(val) OVER(PARTITION BY subject ORDER BY time ROWS CURRENT ROW) AS running_total1,
  SUM(val) OVER(PARTITION BY subject ORDER BY time RANGE UNBOUNDED PRECEDING) AS running_total2,
  SUM(val) OVER(PARTITION BY subject ORDER BY time ROWS UNBOUNDED PRECEDING) AS running_total3,
  SUM(val) OVER(PARTITION BY subject ORDER BY time ROWS 10 PRECEDING) AS running_total4,
  AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN CURRENT ROW AND UNBOUNDED PRECEDING) AS running_average1,
  AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED PRECEDING) AS running_average2,
  AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS running_average3,
  AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN 10 PRECEDING AND 10 PRECEDING) AS running_average4,
  AVG(val) OVER(PARTITION BY subject ORDER BY time RANGE BETWEEN 10 PRECEDING AND 10 PRECEDING) AS running_average4
FROM
  observations`)
}

func TestWindowFunctionSyntax(t *testing.T) {
	validateSql(t, `SELECT
         val,
         ROW_NUMBER()  ignore null OVER w AS row_number,
         CUME_DIST()  RESPECT null  OVER w AS cume_dist
       FROM numbers
       WINDOW w AS (ORDER BY val);`,
		`SELECT val, ROW_NUMBERIGNORE NULLS() OVER w AS row_number, CUME_DISTRESPECT NULLS() OVER w AS cume_dist FROM numbers WINDOW w AS (ORDER BY val)`,
		`SELECT
  val,
  ROW_NUMBERIGNORE NULLS() OVER w AS row_number,
  CUME_DISTRESPECT NULLS() OVER w AS cume_dist
FROM
  numbers
WINDOW w AS (ORDER BY val)`)
}

func TestNamedWindowsSyntax(t *testing.T) {
	validateSql(t, `SELECT
  DISTINCT year, country,
  FIRST_VALUE(year) OVER (w ORDER BY year ASC) AS first,
  FIRST_VALUE(year) OVER w AS last
FROM sales
WINDOW w AS (PARTITION BY country),w2 as (w);`,
		`SELECT DISTINCT year, country, FIRST_VALUE(year) OVER(w ORDER BY year ASC) AS first, FIRST_VALUE(year) OVER w AS last FROM sales WINDOW w AS (PARTITION BY country), w2 AS (w)`,
		`SELECT DISTINCT
  year,
  country,
  FIRST_VALUE(year) OVER(w ORDER BY year ASC) AS first,
  FIRST_VALUE(year) OVER w AS last
FROM
  sales
WINDOW w AS (PARTITION BY country), w2 AS (w)`)
}

func TestHintSyntax(t *testing.T) {
	validateSql(t, `UPDATE /*+ NO_MERGE(discounted) */ items,
       (SELECT id FROM items
        WHERE retail / wholesale >= 1.3 AND quantity < 100)
        AS discounted
    SET items.retail = items.retail * 0.9
    WHERE items.id = discounted.id;`,
		`UPDATE /*+ NO_MERGE(discounted) */ items , (SELECT id FROM items WHERE retail / wholesale >= 1.3 AND quantity < 100) discounted SET items.retail = items.retail * 0.9 WHERE items.id = discounted.id`,
		`UPDATE /*+ NO_MERGE(discounted) */
  items , (SELECT
             id
           FROM
             items
           WHERE retail / wholesale >= 1.3
             AND quantity < 100) discounted
SET
  items.retail = items.retail * 0.9
WHERE items.id = discounted.id`)
}

func TestMySqlBinaryOperationSyntax(t *testing.T) {
	validateSql(t, `select
	*
FROM
	tab1
WHERE
	col1 = 'aa'
	AND col2 = 2
	AND (col3 = 1
		OR col4 = 'bb')
	OR col5 = 'tt'
	AND col6 BETWEEN 100 AND 200
	AND col7 LIKE '%mysql\_%' escape '\'
	AND col8 > any (select level FROM tab2 WHERE type=1)
	AND col10 > some (TABLE tab3)
	AND ((1,2))=(1,2)
	AND ((select col1 FROM tab2 WHERE id=1),2)=(1,2)
	and col11 > some +2
    and col12 <> all (select c1 from tab2)`,
		`SELECT * FROM tab1 WHERE col1 = 'aa' AND col2 = 2 AND (col3 = 1 OR col4 = 'bb') OR col5 = 'tt' AND col6 BETWEEN 100 AND 200 AND col7 LIKE '%mysql\_%' ESCAPE '\' AND col8 > ANY (SELECT level FROM tab2 WHERE type = 1) AND col10 > SOME (TABLE tab3) AND (1, 2) = (1, 2) AND ((SELECT col1 FROM tab2 WHERE id = 1), 2) = (1, 2) AND col11 > some + 2 AND col12 <> ALL (SELECT c1 FROM tab2)`,
		`SELECT
  *
FROM
  tab1
WHERE col1 = 'aa'
  AND col2 = 2
  AND (col3 = 1 OR col4 = 'bb')
   OR col5 = 'tt'
  AND col6 BETWEEN 100 AND 200
  AND col7 LIKE '%mysql\_%' ESCAPE '\'
  AND col8 > ANY (SELECT
                    level
                  FROM
                    tab2
                  WHERE type = 1)
  AND col10 > SOME (TABLE tab3)
  AND (1, 2) = (1, 2)
  AND ((SELECT
          col1
        FROM
          tab2
        WHERE id = 1), 2) = (1, 2)
  AND col11 > some + 2
  AND col12 <> ALL (SELECT
                      c1
                    FROM
                      tab2)`)
	validateSql(t, `select * from tab1 where col1 in(1,2,3)`,
		`SELECT * FROM tab1 WHERE col1 IN (1, 2, 3)`,
		`SELECT
  *
FROM
  tab1
WHERE col1 IN (1, 2, 3)`)
	validateSql(t, `select 1 in(1,2,3)`, `SELECT 1 IN (1, 2, 3)`, ``)
	validateSql(t, `select (1,2,3) in (select * from tab1)`, `SELECT (1, 2, 3) IN (SELECT * FROM tab1)`, ``)
	validateSql(t, `select (select * from tab1) in (1,2,3) `, `SELECT (SELECT * FROM tab1) IN (1, 2, 3)`, ``)
	validateSql(t, `select (1,2,3) in ((1,2,3),(2,3,4),(5,4,2))`, `SELECT (1, 2, 3) IN ((1, 2, 3), (2, 3, 4), (5, 4, 2))`, ``)
	validateSql(t, `select * from tab1 where 
col1 is not true
and col2 not like '%abc%'
and col3 not between 1 and 2
and col4 not regexp 'abc'
and col5 not rlike '%abc'
and col6 not in(1,2)
and col7 sounds like 'abc'
and col8 MEMBER OF('[23, "abc", 17, "ab", 10]')`,
		`SELECT * FROM tab1 WHERE col1 IS NOT TRUE AND col2 NOT LIKE '%abc%' AND col3 NOT BETWEEN 1 AND 2 AND col4 NOT REGEXP 'abc' AND col5 NOT RLIKE '%abc' AND col6 NOT IN (1, 2) AND col7 SOUNDS LIKE 'abc' AND col8 MEMBER OF ('[23, "abc", 17, "ab", 10]')`,
		`SELECT
  *
FROM
  tab1
WHERE col1 IS NOT TRUE
  AND col2 NOT LIKE '%abc%'
  AND col3 NOT BETWEEN 1 AND 2
  AND col4 NOT REGEXP 'abc'
  AND col5 NOT RLIKE '%abc'
  AND col6 NOT IN (1, 2)
  AND col7 SOUNDS LIKE 'abc'
  AND col8 MEMBER OF ('[23, "abc", 17, "ab", 10]')`)
}

func TestMySqlBinaryLiteralSyntax(t *testing.T) {
	r := require.New(t)
	sql := "select b'01100001'"
	is, err := parser.Parse(Dialects.MYSQL, sql)
	r.NoError(err)
	r.Equal("SELECT b'01100001'", ast.BuildSql(is, false))
	m := is.(*ast.MySqlSelectSyntax)
	bls := m.SelectItemList.Get(0).(*ast.SelectColumnSyntax).Expr.(*ast.MySqlBinaryLiteralSyntax)
	r.Equal("01100001", bls.BinStr())
	bls.SetBinStr("0110000101100001")
	r.Equal("b'0110000101100001'", bls.Sql())
}

func TestCastFunctionSyntax(t *testing.T) {
	validateSql(t, "SELECT CAST('9.5' AS DECIMAL(10,2)),CAST('test' as char character set utf8),CAST('test' AS CHAR CHARACTER SET utf8) collate utf8_bin",
		"",
		`SELECT
  CAST('9.5' AS DECIMAL(10, 2)),
  CAST('test' AS CHAR CHARACTER SET utf8),
  CAST('test' AS CHAR CHARACTER SET utf8) COLLATE utf8_bin`)
	validateSql(t, `SELECT CAST(c AT TIME ZONE 'UTC' AS DATETIME) AS u FROM tz;`,
		`SELECT CAST(c AT TIME ZONE 'UTC' AS DATETIME) AS u FROM tz`,
		``)
	validateSql(t, `SELECT CAST(c AT TIME ZONE interval '+00:00' AS DATETIME) AS u FROM tz;`,
		`SELECT CAST(c AT TIME ZONE '+00:00' AS DATETIME) AS u FROM tz`,
		``)
}

func TestMySqlCharFunctionSyntax(t *testing.T) {
	validateSql(t, "SELECT CHAR(77,121,83,81,'76')",
		"SELECT CHAR(77, 121, 83, 81, '76')",
		"")
	validateSql(t, "SELECT CHAR(77,121,83,81,'76' USING utf8mb4)",
		"SELECT CHAR(77, 121, 83, 81, '76' USING utf8mb4)",
		"")
}

func TestMySqlConvertFunctionSyntax(t *testing.T) {
	validateSql(t, "select convert('test' USING utf8),convert('test', char(5) character set utf8),CONVERT('test', CHAR CHARACTER SET utf8) collate utf8_bin",
		"",
		`SELECT
  CONVERT('test' USING utf8),
  CONVERT('test', CHAR(5) CHARACTER SET utf8),
  CONVERT('test', CHAR CHARACTER SET utf8) COLLATE utf8_bin`)
}

func TestMySqlDateAndTimeLiteralSyntax(t *testing.T) {
	validateSql(t, "select DATE '2023-09-18',TIME '02:12:00',TIMESTAMP '2023-09-18 02:12:00'",
		"SELECT DATE'2023-09-18', TIME'02:12:00', TIMESTAMP'2023-09-18 02:12:00'",
		"")
}

func TestMySqlDeleteSyntax(t *testing.T) {
	validateSql(t, `delete low_priority quick ignore from test.t_customer_info t1 where t1.gender = 2 order by t1.id limit 1,2`,
		`DELETE LOW_PRIORITY QUICK IGNORE FROM test.t_customer_info t1 WHERE t1.gender = 2 ORDER BY t1.id LIMIT 1,2`,
		`DELETE LOW_PRIORITY QUICK IGNORE
FROM
  test.t_customer_info t1
WHERE t1.gender = 2
ORDER BY
  t1.id
LIMIT 1,2`)
	validateSql(t, `delete
  t.*,t2.*
 from
  t_customer_info t
  straight_join t_level t2
    on t.id = t2.customer_id
where t.gender = 1
  and t2.level = 2
  and not exists
  (select
    1
  from
    t_communicating_info t3
  where t3.customer_id = t.id);`,
		`DELETE t.*, t2.* FROM t_customer_info t STRAIGHT_JOIN t_level t2 ON t.id = t2.customer_id WHERE t.gender = 1 AND t2.level = 2 AND NOT EXISTS(SELECT 1 FROM t_communicating_info t3 WHERE t3.customer_id = t.id)`,
		`DELETE
  t.*,
  t2.*
FROM
  t_customer_info t
  STRAIGHT_JOIN t_level t2
    ON t.id = t2.customer_id
WHERE t.gender = 1
  AND t2.level = 2
  AND NOT EXISTS(SELECT
                   1
                 FROM
                   t_communicating_info t3
                 WHERE t3.customer_id = t.id)`)
	validateSql(t, `delete LOW_PRIORITY QUICK IGNORE
from
  t1.*,
  t2.* using tab1 t1
  inner join (tab2 t2)
  inner join tab3 t3
where t1.id = t2.id
  and t2.id = t3.id`,
		`DELETE LOW_PRIORITY QUICK IGNORE FROM t1.*, t2.* USING tab1 t1 INNER JOIN (tab2 t2) INNER JOIN tab3 t3 WHERE t1.id = t2.id AND t2.id = t3.id`,
		`DELETE LOW_PRIORITY QUICK IGNORE
FROM
  t1.*,
  t2.* USING tab1 t1
             INNER JOIN (tab2 t2)
             INNER JOIN tab3 t3
WHERE t1.id = t2.id
  AND t2.id = t3.id`)
	validateSql(t, `delete
from
  t1,
  t2 using tab1 t1
  inner join tab2 t2
  inner join tab3 t3
where t1.id = t2.id
  and t2.id = t3.id`,
		`DELETE FROM t1, t2 USING tab1 t1 INNER JOIN tab2 t2 INNER JOIN tab3 t3 WHERE t1.id = t2.id AND t2.id = t3.id`,
		`DELETE
FROM
  t1,
  t2 USING tab1 t1
           INNER JOIN tab2 t2
           INNER JOIN tab3 t3
WHERE t1.id = t2.id
  AND t2.id = t3.id`)
}

func TestMySqlExtractFunctionSyntax(t *testing.T) {
	validateSql(t, `SELECT EXTRACT(YEAR FROM t.create_time) AS order_year, abc from t_order t where t.order_no='234234'`,
		`SELECT EXTRACT(YEAR FROM t.create_time) AS order_year, abc FROM t_order t WHERE t.order_no = '234234'`,
		`SELECT
  EXTRACT(YEAR FROM t.create_time) AS order_year,
  abc
FROM
  t_order t
WHERE t.order_no = '234234'`)
}

func TestMySqlFalseSyntax(t *testing.T) {
	validateSql(t, `select 1>2=false`,
		`SELECT 1 > 2 = FALSE`,
		``)
}

func TestMySqlTrueSyntax(t *testing.T) {
	validateSql(t, `select 1<2=true`,
		`SELECT 1 < 2 = TRUE`,
		``)
}

func TestMySqlGetFormatFunctionSyntax(t *testing.T) {
	validateSql(t, `SELECT GET_FORMAT(DATE,'USA'),get_format(DATETIME,'INTERNAL'),GET_FORMAT(TIME,'ISO')`,
		`SELECT GET_FORMAT(DATE, 'USA'), GET_FORMAT(DATETIME, 'INTERNAL'), GET_FORMAT(TIME, 'ISO')`,
		`SELECT
  GET_FORMAT(DATE, 'USA'),
  GET_FORMAT(DATETIME, 'INTERNAL'),
  GET_FORMAT(TIME, 'ISO')`)
}

func TestMySqlGroupConcatFunctionSyntax(t *testing.T) {
	validateSql(t, `SELECT GROUP_CONCAT(DISTINCT test_score ORDER BY test_score DESC SEPARATOR ' ') from t_student`,
		`SELECT (DISTINCT test_score ORDER BY test_score DESC SEPARATOR ' ') FROM t_student`,
		`SELECT
  (DISTINCT test_score ORDER BY test_score DESC SEPARATOR ' ')
FROM
  t_student`)
}

func TestMySqlHexagonalLiteralSyntax(t *testing.T) {
	validateSql(t, `SELECT x'4D7953514C' from dual`,
		`SELECT x'4D7953514C' FROM DUAL`, ``)
	r := require.New(t)
	m := ast.NewMySqlHexagonalLiteralSyntax()
	m.SetSql("x'4D7953514C'")
	r.Equal("4D7953514C", m.HexStr())
	m.SetHexStr("4552444f53")
	r.Equal("x'4552444f53'", m.Sql())
}

func TestMySqlIdentifierSyntax(t *testing.T) {
	validateSql(t, "SELECT `name`,level from `test`.t_user",
		"SELECT `name`, level FROM `test`.t_user", ``)
}

func TestMySqlInsertSyntax(t *testing.T) {
	validateSql(t, `insert LOW_PRIORITY DELAYED HIGH_PRIORITY ignore into t_customer_info(id,nickname,level,type,address) values (null, 'aname', 1, 2, 'address'),(null, 'jack', 1, 2, 'address'),(null,?,?,?,?) as t1(c1,c2,c3,c4,c5) on duplicate key update level=values(level), type=t1.c4, address=concat(address,values(address))`,
		`INSERT LOW_PRIORITY HIGH_PRIORITY DELAYED IGNORE INTO t_customer_info (id, nickname, level, type, address) VALUES (NULL, 'aname', 1, 2, 'address'), (NULL, 'jack', 1, 2, 'address'), (NULL, ?, ?, ?, ?) AS t1(c1, c2, c3, c4, c5) ON DUPLICATE KEY UPDATE level = values(level), type = t1.c4, address = concat(address, values(address))`,
		`INSERT LOW_PRIORITY HIGH_PRIORITY DELAYED IGNORE INTO
t_customer_info
  (id, nickname, level, type, address)
VALUES
  (NULL, 'aname', 1, 2, 'address'),
  (NULL, 'jack', 1, 2, 'address'),
  (NULL, ?, ?, ?, ?)
AS t1(c1, c2, c3, c4, c5)
ON DUPLICATE KEY UPDATE
  level = values(level),
  type = t1.c4,
  address = concat(address, values(address))`)
	validateSql(t, `insert into tab1 set a='abc',b=1,c=?`,
		`INSERT INTO tab1 SET a = 'abc', b = 1, c = ?`,
		`INSERT INTO
tab1
SET
  a = 'abc',
  b = 1,
  c = ?`)
	validateSql(t, `insert into tbl_name (a,b,c) values ROW(1,2,3), ROW(4,5,6), ROW(7,8,9);`,
		`INSERT INTO tbl_name (a, b, c) VALUES ROW(1, 2, 3), ROW(4, 5, 6), ROW(7, 8, 9)`,
		`INSERT INTO
tbl_name
  (a, b, c)
VALUES
  ROW(1, 2, 3),
  ROW(4, 5, 6),
  ROW(7, 8, 9)`)
	validateSql(t, `insert into tab1(col1,col2,col3) value (1,'aa',now()) as t1(c1,c2,c3)`,
		`INSERT INTO tab1 (col1, col2, col3) VALUES (1, 'aa', now()) AS t1(c1, c2, c3)`,
		`INSERT INTO
tab1
  (col1, col2, col3)
VALUES
  (1, 'aa', now())
AS t1(c1, c2, c3)`)
}

func TestMySqlIntervalSyntax(t *testing.T) {
	validateSql(t, `SELECT DATE_ADD('2023-09-26',INTERVAL (1+1) DAY);`,
		`SELECT DATE_ADD('2023-09-26', INTERVAL (1 + 1) DAY)`,
		``)
	validateSql(t, `SELECT NOW()-INTERVAL 24 HOUR`,
		`SELECT NOW() - INTERVAL 24 HOUR`,
		``)
	validateSql(t, `SELECT INTERVAL(23, 1, 15, 17, 30, 44, 200)`,
		`SELECT INTERVAL(23, 1, 15, 17, 30, 44, 200)`,
		``)
}

func TestMySqlSelectSyntax(t *testing.T) {
	validateSql(t, `select all distinct distinctrow high_priority straight_join sql_small_result sql_big_result sql_buffer_result sql_cache sql_no_cache sql_calc_found_rows col_1,col_2 from t_table`,
		`SELECT DISTINCT HIGH_PRIORITY STRAIGHT_JOIN SQL_SMALL_RESULT SQL_BIG_RESULT SQL_BUFFER_RESULT SQL_CACHE SQL_NO_CACHE SQL_CALC_FOUND_ROWS col_1, col_2 FROM t_table`,
		`SELECT DISTINCT HIGH_PRIORITY STRAIGHT_JOIN SQL_SMALL_RESULT SQL_BIG_RESULT SQL_BUFFER_RESULT SQL_CACHE SQL_NO_CACHE SQL_CALC_FOUND_ROWS
  col_1,
  col_2
FROM
  t_table`)
	validateSql(t, `SELECT col_1,col_2,count(*) FROM tab_1 group by col_1 asc,col_2 desc WITH ROLLUP having col_1='jack' order by col_1 desc, col_2 asc`,
		`SELECT col_1, col_2, COUNT(*) FROM tab_1 GROUP BY col_1 ASC, col_2 DESC WITH ROLLUP HAVING col_1 = 'jack' ORDER BY col_1 DESC, col_2 ASC`,
		`SELECT
  col_1,
  col_2,
  COUNT(*)
FROM
  tab_1
GROUP BY
  col_1 ASC,
  col_2 DESC
WITH ROLLUP
HAVING
  col_1 = 'jack'
ORDER BY
  col_1 DESC,
  col_2 ASC`)
	validateSql(t, `select * from tab1 limit 2,10`,
		`SELECT * FROM tab1 LIMIT 2,10`,
		``)
	validateSql(t, `select * from tab1 limit 10 offset 2`,
		`SELECT * FROM tab1 LIMIT 2,10`,
		``)
	validateSql(t, `select a.col1,a.col1,b.col1,b.col2 from tab1 a join tab2 b on a.id=b.pid for update nowait`,
		`SELECT a.col1, a.col1, b.col1, b.col2 FROM tab1 a JOIN tab2 b ON a.id = b.pid FOR UPDATE NOWAIT`,
		`SELECT
  a.col1,
  a.col1,
  b.col1,
  b.col2
FROM
  tab1 a
  JOIN tab2 b
    ON a.id = b.pid
FOR UPDATE NOWAIT`)
	validateSql(t, `select a.col1,a.col1,b.col1,b.col2 from tab1 a join tab2 b on a.id=b.pid for share skip locked`,
		`SELECT a.col1, a.col1, b.col1, b.col2 FROM tab1 a JOIN tab2 b ON a.id = b.pid FOR SHARE SKIP LOCKED`,
		`SELECT
  a.col1,
  a.col1,
  b.col1,
  b.col2
FROM
  tab1 a
  JOIN tab2 b
    ON a.id = b.pid
FOR SHARE SKIP LOCKED`)
	validateSql(t, `select a.col1,a.col1,b.col1,b.col2 from tab1 a join tab2 b on a.id=b.pid lock in share mode`,
		`SELECT a.col1, a.col1, b.col1, b.col2 FROM tab1 a JOIN tab2 b ON a.id = b.pid LOCK IN SHARE MODE`,
		`SELECT
  a.col1,
  a.col1,
  b.col1,
  b.col2
FROM
  tab1 a
  JOIN tab2 b
    ON a.id = b.pid
LOCK IN SHARE MODE`)
	validateSql(t, `select a.col1,a.col1,b.col1,b.col2 from tab1 a join tab2 b on a.id=b.pid for share of a skip locked for update of b nowait`,
		`SELECT a.col1, a.col1, b.col1, b.col2 FROM tab1 a JOIN tab2 b ON a.id = b.pid FOR SHARE a SKIP LOCKED FOR UPDATE b NOWAIT`,
		`SELECT
  a.col1,
  a.col1,
  b.col1,
  b.col2
FROM
  tab1 a
  JOIN tab2 b
    ON a.id = b.pid
FOR SHARE a SKIP LOCKED
FOR UPDATE b NOWAIT`)
}

func TestMySqlVariableSyntax(t *testing.T) {
	validateSql(t, `select @myParam,@@global.max_connections`, `SELECT @myParam, @@global.max_connections`, "")
	r := require.New(t)
	m := ast.NewMySqlVariableSyntax()
	m.SetSql("@myVar")
	r.Equal("@myVar", m.Sql())
	r.Equal("myVar", m.Name())
	m.SetSql("@@myVar")
	r.Equal("myVar", m.Name())
	m.SetSql("@'myVar'")
	r.Equal("myVar", m.Name())
	m.SetSql("@\"myVar\"")
	r.Equal("myVar", m.Name())
}

func TestMySqlStringSyntax(t *testing.T) {
	at := assert.New(t)
	s := ast.NewMySqlStringSyntax()
	s.SetSql(`'\0\'\"\b\n\r\t\Z\\\%\_'`)
	chars := []rune(s.Value())
	at.Equal(11, len(chars))
	at.Equal(rune(0), chars[0])
	at.Equal('\'', chars[1])
	at.Equal('"', chars[2])
	at.Equal('\b', chars[3])
	at.Equal('\n', chars[4])
	at.Equal('\r', chars[5])
	at.Equal('\t', chars[6])
	at.Equal(rune(26), chars[7])
	at.Equal('\\', chars[8])
	at.Equal('%', chars[9])
	at.Equal('_', chars[10])

	var v strings.Builder
	v.WriteRune(rune(0))
	v.WriteRune('\'')
	v.WriteRune('"')
	v.WriteRune('\b')
	v.WriteRune('\n')
	v.WriteRune('\r')
	v.WriteRune('\t')
	v.WriteRune(rune(26))
	v.WriteRune('\\')
	v.WriteRune('%')
	v.WriteRune('_')
	v.WriteRune('a')
	s.SetValue(v.String())
	at.Equal(`'\0\'\"\b\n\r\t\Z\\\%\_a'`, s.Sql())
}

func TestMySqlNameTableReferenceSyntax(t *testing.T) {
	validateSql(t, `select * from 
test.tab1 PARTITION(p0,p1,p2) t1 use index for join (idx_col1) 
join test.tab2 t2 use key for order by (idx_col2) on t1.id=t2.pid 
join test.tab3 t3 use index for group by () IGNORE INDEX FOR JOIN (idx_col3) on t1.id=t3.pid 
join test.tab4 t4 ignore index for join(idx_col4) on t1.id=t4.pid 
join test.tab4 t5 force index for order by (idx_col4) on t1.id=t5.pid 
join test.tab4 t5 force index for group by (idx_col5) on t1.id=t5.pid`,
		`SELECT * FROM test.tab1 t1 PARTITION(p0, p1, p2) USE INDEX FOR JOIN (idx_col1) JOIN test.tab2 t2 USE INDEX FOR ORDER BY (idx_col2) ON t1.id = t2.pid JOIN test.tab3 t3 USE INDEX FOR GROUP BY () IGNORE INDEX FOR JOIN (idx_col3) ON t1.id = t3.pid JOIN test.tab4 t4 IGNORE INDEX FOR JOIN (idx_col4) ON t1.id = t4.pid JOIN test.tab4 t5 FORCE INDEX FOR ORDER BY (idx_col4) ON t1.id = t5.pid JOIN test.tab4 t5 FORCE INDEX FOR GROUP BY (idx_col5) ON t1.id = t5.pid`,
		`SELECT
  *
FROM
  test.tab1 t1 PARTITION(p0, p1, p2) USE INDEX FOR JOIN (idx_col1)
  JOIN test.tab2 t2 USE INDEX FOR ORDER BY (idx_col2)
    ON t1.id = t2.pid
  JOIN test.tab3 t3 USE INDEX FOR GROUP BY ()
                    IGNORE INDEX FOR JOIN (idx_col3)
    ON t1.id = t3.pid
  JOIN test.tab4 t4 IGNORE INDEX FOR JOIN (idx_col4)
    ON t1.id = t4.pid
  JOIN test.tab4 t5 FORCE INDEX FOR ORDER BY (idx_col4)
    ON t1.id = t5.pid
  JOIN test.tab4 t5 FORCE INDEX FOR GROUP BY (idx_col5)
    ON t1.id = t5.pid`)
}

func TestMySqlTimestampFunctionSyntax(t *testing.T) {
	validateSql(t, `SELECT timestampadd(MINUTE, 1, '2003-01-02')`,
		`SELECT TIMESTAMPADD(MINUTE, 1, '2003-01-02')`,
		``)
	validateSql(t, `SELECT timestampdiff(MONTH, 1, DATE '2024-03-30') AS t1`,
		`SELECT TIMESTAMPDIFF(MONTH, 1, DATE'2024-03-30') AS t1`,
		``)
}

func TestMySqlTranscodingStringSyntax(t *testing.T) {
	validateSql(t, `select _utf8mb4'some text' collate utf8_bin`,
		`SELECT _utf8mb4'some text' COLLATE utf8_bin`,
		``)
}

func TestMySqlTrimFunctionSyntax(t *testing.T) {
	validateSql(t, `SELECT TRIM('  bar   ')`,
		`SELECT TRIM('  bar   ')`,
		``)
	validateSql(t, `SELECT trim(leading 'x' from 'xxxbarxxx')`,
		`SELECT TRIM(LEADING 'x' FROM 'xxxbarxxx')`,
		``)
	validateSql(t, `SELECT trim(both 'x' from 'xxxbarxxx')`,
		`SELECT TRIM(BOTH 'x' FROM 'xxxbarxxx')`,
		``)
	validateSql(t, `SELECT trim(trailing 'xyz' from 'barxxyz')`,
		`SELECT TRIM(TRAILING 'xyz' FROM 'barxxyz')`,
		``)
}

func TestMySqlUnarySyntax(t *testing.T) {
	validateSql(t, `select +123`,
		`SELECT +123`,
		``)
	validateSql(t, `select -123`,
		`SELECT -123`,
		``)
	validateSql(t, `select ~123`,
		`SELECT ~123`,
		``)
	validateSql(t, `select binary 'a'`,
		`SELECT BINARY 'a'`,
		``)
	validateSql(t, `select !true`,
		`SELECT !TRUE`,
		``)
	validateSql(t, `select NOT true`,
		`SELECT NOT TRUE`,
		``)
	validateSql(t, `select NOT true`,
		`SELECT NOT TRUE`,
		``)

}

func TestMySqlMultisetSyntax(t *testing.T) {
	validateSql(t, `(SELECT a FROM t1 WHERE a=10 AND B=1 ORDER BY a LIMIT 10)
UNION ALL
SELECT a FROM t2 WHERE a=11 AND B=2
UNION ALL
SELECT a FROM t2 WHERE a=12 AND B=3 ORDER BY a LIMIT 10`,
		`(SELECT a FROM t1 WHERE a = 10 AND B = 1 ORDER BY a LIMIT 10) UNION ALL SELECT a FROM t2 WHERE a = 11 AND B = 2 UNION ALL SELECT a FROM t2 WHERE a = 12 AND B = 3 ORDER BY a LIMIT 10`,
		`(SELECT
   a
 FROM
   t1
 WHERE a = 10
   AND B = 1
 ORDER BY
   a
 LIMIT 10)
UNION ALL
SELECT
  a
FROM
  t2
WHERE a = 11
  AND B = 2
UNION ALL
SELECT
  a
FROM
  t2
WHERE a = 12
  AND B = 3
ORDER BY
  a
LIMIT 10`)
	validateSql(t, `(SELECT a FROM t1 WHERE a=10 AND B=1 ORDER BY a LIMIT 10)
UNION ALL
(SELECT a FROM t2 WHERE a=11 AND B=2 ORDER BY a LIMIT 10)`,
		`(SELECT a FROM t1 WHERE a = 10 AND B = 1 ORDER BY a LIMIT 10) UNION ALL (SELECT a FROM t2 WHERE a = 11 AND B = 2 ORDER BY a LIMIT 10)`,
		`(SELECT
   a
 FROM
   t1
 WHERE a = 10
   AND B = 1
 ORDER BY
   a
 LIMIT 10)
UNION ALL
(SELECT
   a
 FROM
   t2
 WHERE a = 11
   AND B = 2
 ORDER BY
   a
 LIMIT 10)`)
	validateSql(t, `select * from tab1 where type=1 except select * from tab1 where type=2`,
		`SELECT * FROM tab1 WHERE type = 1 EXCEPT SELECT * FROM tab1 WHERE type = 2`,
		`SELECT
  *
FROM
  tab1
WHERE type = 1
EXCEPT
SELECT
  *
FROM
  tab1
WHERE type = 2`)
	validateSql(t, `select * from tab1 where type=1 intersect select * from tab1 where type=2`,
		`SELECT * FROM tab1 WHERE type = 1 INTERSECT SELECT * FROM tab1 WHERE type = 2`,
		`SELECT
  *
FROM
  tab1
WHERE type = 1
INTERSECT
SELECT
  *
FROM
  tab1
WHERE type = 2`)
}

func TestMySqlUpdateSyntax(t *testing.T) {
	validateSql(t, `update LOW_PRIORITY IGNORE 
  t_user a,
  t_account b
set
  b.username = a.username
where 
  a.id = b.user_id;`,
		`UPDATE LOW_PRIORITY IGNORE t_user a , t_account b SET b.username = a.username WHERE a.id = b.user_id`,
		`UPDATE LOW_PRIORITY IGNORE
  t_user a , t_account b
SET
  b.username = a.username
WHERE a.id = b.user_id`)

	validateSql(t, `UPDATE
  t_parent a
  JOIN t_children b
    ON a.id = b.parent_id
    join t_attribute c on
    a.id=c.parent_id
SET
  a.stat = default,
  b.stat = 1
WHERE a.type = 1
and b.name='aname'
and c.att1='n'
ORDER BY
  a.id DESC
limit 3`,
		`UPDATE t_parent a JOIN t_children b ON a.id = b.parent_id JOIN t_attribute c ON a.id = c.parent_id SET a.stat = DEFAULT, b.stat = 1 WHERE a.type = 1 AND b.name = 'aname' AND c.att1 = 'n' ORDER BY a.id DESC LIMIT 3`,
		`UPDATE
  t_parent a
  JOIN t_children b
    ON a.id = b.parent_id
  JOIN t_attribute c
    ON a.id = c.parent_id
SET
  a.stat = DEFAULT,
  b.stat = 1
WHERE a.type = 1
  AND b.name = 'aname'
  AND c.att1 = 'n'
ORDER BY
  a.id DESC
LIMIT 3`)
}

func TestMySqlTableSyntax(t *testing.T) {
	validateSql(t, `SELECT s1 FROM t1 WHERE s1 > ANY (TABLE t2 ORDER BY id LIMIT 10);`,
		`SELECT s1 FROM t1 WHERE s1 > ANY (TABLE t2 ORDER BY id LIMIT 10)`,
		`SELECT
  s1
FROM
  t1
WHERE s1 > ANY (TABLE t2 ORDER BY id LIMIT 10)`)
	validateSql(t, `select t1.*,t2.col1,t2.col2 from tab1 t1 left outer join (select * from tab2 where type=2) t2 on t1.id=t2.pid`,
		`SELECT t1.*, t2.col1, t2.col2 FROM tab1 t1 LEFT JOIN (SELECT * FROM tab2 WHERE type = 2) t2 ON t1.id = t2.pid`,
		`SELECT
  t1.*,
  t2.col1,
  t2.col2
FROM
  tab1 t1
  LEFT JOIN (SELECT
               *
             FROM
               tab2
             WHERE type = 2) t2
    ON t1.id = t2.pid`)
	validateSql(t, `select * from tab1 t1 cross join tab2 t2 on t1.id=t2.id`,
		`SELECT * FROM tab1 t1 CROSS JOIN tab2 t2 ON t1.id = t2.id`,
		`SELECT
  *
FROM
  tab1 t1
  CROSS JOIN tab2 t2
    ON t1.id = t2.id`)
	validateSql(t, `select * from tab1 t1 natural join tab2 t2`,
		`SELECT * FROM tab1 t1 NATURAL JOIN tab2 t2`,
		`SELECT
  *
FROM
  tab1 t1 NATURAL JOIN tab2 t2`)
	validateSql(t, `select * from tab1 t1 right outer join tab2 t2 on t1.id=t2.id`,
		`SELECT * FROM tab1 t1 RIGHT JOIN tab2 t2 ON t1.id = t2.id`,
		`SELECT
  *
FROM
  tab1 t1
  RIGHT JOIN tab2 t2
    ON t1.id = t2.id`)
}

func TestMySqlParserError(t *testing.T) {
	r := require.New(t)
	validateError(r, "select * from tab1; a",
		`unexpected token
select * from tab1; ↪a↩`)
	validateError(r, `SELECT  INTERVAL 1 SECOND-'2025-01-01'`,
		`for the - operator, INTERVAL expr unit is permitted only on the right side
SELECT  ↪INTERVAL 1 SECOND↩-'2025-01-01'`)
	validateError(r, `SELECT NOW()-INTERVAL 24 NANOSECOND`,
		`unexpected token
SELECT NOW()-INTERVAL 24 ↪NANOSECOND↩`)
	validateError(r, `alter table tab1 add column col1 varchar(255)`,
		`unexpected token
↪alter↩ table tab1 add column col1 varchar(255)`)
	validateError(r, `select  * from aaa where exists(1)`,
		`unexpected token
select  * from aaa where exists(↪1↩)`)
	validateError(r, `select * from tab1 where col=1 order by col2 union select * from tab1 where col=2`,
		`incorrect usage of ORDER BY
↪select * from tab1 where col=1 order by col2↩ union select * from tab1 where col=2`)
	validateError(r, `select * from tab1 where col=1 limit 5 union select * from tab1 where col=2`,
		`incorrect usage of LIMIT
↪select * from tab1 where col=1 limit 5↩ union select * from tab1 where col=2`)
	validateError(r, `select * from tab1 where col=1 for unknown`,
		`unexpected token
select * from tab1 where col=1 for ↪unknown↩`)
	validateError(r, `SELECT s1 FROM t1 WHERE s1 > ANY ('a')`,
		`unexpected token
SELECT s1 FROM t1 WHERE s1 > ANY (↪'a'↩)`)
	validateError(r, `select CAST(expr AS 'CHAR')`,
		`unexpected token
select CAST(expr AS ↪'CHAR'↩)`)
	validateError(r, `select sum(col1) over 'a' from tab1`,
		`unexpected token
select sum(col1) over ↪'a'↩ from tab1`)
	validateError(r, `insert into tab1 val (1,'aa')`,
		`unexpected token
insert into tab1 ↪val↩ (1,'aa')`)
	validateError(r, `insert into tab1(col1,col2,col3) value (1,'aa',now()) as t1(c1,c2)`,
		`column alias count doesn't match value count, column: 3, alias: 2
insert into tab1(col1,col2,col3) value (1,'aa',now()) as t1↪(c1,c2)↩`)
	validateError(r, `insert into tab1(col1,col2,col3) value (1,'aa',now()),(2,'bb')`,
		`column count doesn't match value count, column: 3, value: 2
insert into tab1(col1,col2,col3) value (1,'aa',now()),↪(2,'bb')↩`)
	validateError(r, `select t1.* from tab1 t1 join (select * from tab2) on tab1.id=tab2.pid`,
		`every derived table must have its own alias
select t1.* from tab1 t1 join ↪(select * from tab2)↩ on tab1.id=tab2.pid`)
	validateError(r, `select t1.* from tab1 t1 join select * from tab2 t2 on t1.id=t2.pid`,
		`subquery expression must be parenthesized
select t1.* from tab1 t1 join ↪select↩ * from tab2 t2 on t1.id=t2.pid`)
	validateError(r, `select t1.* from 'tab1'`,
		`unexpected token
select t1.* from ↪'tab1'↩`)
	validateError(r, `select * from tab1 as t1 use index for aaa (idx_col1)`,
		`unexpected token
select * from tab1 as t1 use index for ↪aaa↩ (idx_col1)`)
	validateError(r, `delete from test.'abc' where id=1`,
		`unexpected token
delete from test.↪'abc'↩ where id=1`)
	validateError(r, `select * from tab1 t1 join tab2 t2 to t1.id=t2.id`,
		`unexpected token
select * from tab1 t1 join tab2 t2 ↪to↩ t1.id=t2.id`)
	validateError(r, `select * from tab1 as`,
		`unexpected token
select * from tab1 as↪↩`)
	validateError(r, `select * from tab1 group col1`,
		`expected token: BY, actual token: IDENTIFIER
select * from tab1 group ↪col1↩`)
	validateError(r, `select 
col1,
case col2
when 1 then 'level1'
when 2 then 'level2',
from tab1`,
		`expected tokenVal: END, actual tokenVal: (COMMA)
select 
col1,
case col2
when 1 then 'level1'
when 2 then 'level2'↪,↩
from tab1`)
	validateError(r, `(delete from tab1 where id=1)`,
		`this syntax cannot be parenthesized
↪(delete from tab1 where id=1)↩`)
	validateError(r, `select (1,2)=(1,2,3)`,
		`operand should contain 2 Column(s), but found 3
select (1,2)=↪(1,2,3)↩`)
	validateError(r, `select (1 abc`,
		`unexpected token
select (1 ↪abc↩`)
	validateError(r, `select 1-alter`,
		`unexpected token
select 1-↪alter↩`)
	validateError(r, `select *,* from tab1`,
		`unexpected token
select *,↪*↩ from tab1`)
	validateError(r, `select t1.'col1' from tab1 t1`,
		`unexpected token
select t1.↪'col1'↩ from tab1 t1`)
	validateError(r, `select 1 not equal 2`,
		`unexpected token
select 1 not ↪equal↩ 2`)
	validateError(r, `select 'abc' sounds as 'abcd'`,
		`unexpected token
select 'abc' sounds ↪as↩ 'abcd'`)
	validateError(r, `select t1.col1, select t2.col1 from tab2 t2 where t1.id=t2.pid from tab1 t1`,
		`subquery expression must be parenthesized
select t1.col1, ↪select↩ t2.col1 from tab2 t2 where t1.id=t2.pid from tab1 t1`)
	validateError(r, `SELECT GET_FORMAT(TIMESTAMP,'USA')`,
		`unexpected token
SELECT GET_FORMAT(↪TIMESTAMP↩,'USA')`)
	validateError(r, `SELECT timestampadd(MIN, 1, '2003-01-02')`,
		`unexpected token
SELECT timestampadd(↪MIN↩, 1, '2003-01-02')`)
	validateError(r, "select convert('test' as utf8)",
		`unexpected token
select convert('test' ↪as↩ utf8)`)
	validateError(r, `SELECT CAST(c AT TIME ZONE interval 'UTC' AS DATETIME) AS u FROM tz;`,
		`unknown or incorrect time zone: 'UTC'
SELECT CAST(c AT TIME ZONE interval ↪'UTC'↩ AS DATETIME) AS u FROM tz;`)
	validateError(r, `SELECT EXTRACT(NANOSECOND FROM t.create_time) AS order_year, abc from t_order t where t.order_no='234234'`,
		`unexpected token
SELECT EXTRACT(↪NANOSECOND↩ FROM t.create_time) AS order_year, abc from t_order t where t.order_no='234234'`)
	validateError(r, "select col1 as , col2 c2 from tab1 as t1",
		`unexpected token
select col1 as ↪,↩ col2 c2 from tab1 as t1`)
	validateError(r, "select '2018-12-31 23:59:59' + INTERVAL(23)",
		`syntax error
select '2018-12-31 23:59:59' + INTERVAL↪(23)↩`)
}
