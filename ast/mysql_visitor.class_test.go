package ast_test

import (
	"strings"
	"testing"

	"github.com/jishaocong0910/go-sql-parser/parser"

	"github.com/jishaocong0910/go-sql-parser/ast"

	. "github.com/jishaocong0910/go-sql-parser/enum"

	"github.com/stretchr/testify/require"
)

func visit(r *require.Assertions, sql string) ast.Visitor_ {
	s, err := parser.Parse(Dialect_.MYSQL, sql)
	r.NoError(err)
	v_, err := ast.Visit(s)
	r.NoError(err)
	return v_
}

func validateSet(r *require.Assertions, target []string, caseSensitive bool, contains ...string) {
	r.Equal(len(contains), len(target), "expected: %v, actual: %v", contains, target)
	for _, c := range contains {
		r.Conditionf(func() bool {
			for _, v := range target {
				if caseSensitive {
					if c == v {
						return true
					}
				} else {
					if strings.EqualFold(c, v) {
						return true
					}
				}
			}
			return false
		}, "not contains: %v, actual: %v", c, target)
	}
}

func validateMapSet(r *require.Assertions, target map[string][]string, keyCaseSensitive, valueCaseSensitive bool, contains map[string][]string) {
	actual := make(map[string][]string)
	if keyCaseSensitive {
		for k, v := range target {
			actual[strings.ToUpper(k)] = v
		}
	} else {
		actual = target
	}
	var expectedKeys []string
	for k := range contains {
		expectedKeys = append(expectedKeys, k)
	}
	var actualKeys []string
	for k := range target {
		actualKeys = append(actualKeys, k)
	}

	r.Equal(len(expectedKeys), len(target), "expected: %v, actual: %v", expectedKeys, actualKeys)
	for _, k := range expectedKeys {
		av := target[k]
		r.NotNil(av, "not contains key: %s, expected: %v, actual: %v", k, expectedKeys, actualKeys)
		cv := contains[k]
		r.Equal(len(cv), len(av), "key: %s, expected value: %v, actual value: %v", k, cv, av)
		for _, v := range cv {
			r.Conditionf(func() bool {
				for _, a := range av {
					if valueCaseSensitive {
						if v == a {
							return true
						}
					} else {
						if strings.EqualFold(v, a) {
							return true
						}
					}
				}
				return false
			}, "key: %v, not contains value: %v, actual: %v", k, v, av)
		}
	}
}

func validateTables(r *require.Assertions, v_ ast.Visitor_, contains ...string) {
	validateSet(r, v_.TablesRaw(), v_.Option().TableCaseSensitive, contains...)
}

func validateTableColumns(r *require.Assertions, v_ ast.Visitor_, contains map[string][]string) {
	validateMapSet(r, v_.TableColumnsRaw(), v_.Option().TableCaseSensitive, v_.Option().ColumnCaseSensitive, contains)
}

func validateSelectTableColumns(r *require.Assertions, v_ ast.Visitor_, contains map[string][]string) {
	validateMapSet(r, v_.SelectColumnsRaw(), v_.Option().TableCaseSensitive, v_.Option().ColumnCaseSensitive, contains)
}

func validateWhereColumns(r *require.Assertions, v_ ast.Visitor_, contains map[string][]string) {
	validateMapSet(r, v_.WhereColumnsRaw(), v_.Option().TableCaseSensitive, v_.Option().ColumnCaseSensitive, contains)
}

func validateUpdateTables(r *require.Assertions, v_ ast.Visitor_, contains ...string) {
	validateSet(r, v_.UpdateTablesRaw(), v_.Option().TableCaseSensitive, contains...)
}

func validateError(r *require.Assertions, sql, msg string) {
	is, err := parser.Parse(Dialect_.MYSQL, sql)
	r.NoError(err)
	_, err = ast.Visit(is)
	r.EqualError(err, msg)
}

func validateWarn(r *require.Assertions, sql string, warning string) {
	is, err := parser.Parse(Dialect_.MYSQL, sql)
	r.NoError(err)
	v, err := ast.Visit(is)
	r.NoError(err)
	v.Option()
	r.Equal(warning, v.Warning())
}

func TestAllColumnSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `select t.* from tab1 t`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	r.Equal("tab1", v_.TableOfSimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{"tab1": {"*"}})
	validateSelectTableColumns(r, v_, map[string][]string{"tab1": {"*"}})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestAliasSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `select col1 'c1', col2 c2, col3 "c3", col1 as 'c1', col2 as c2, col3 as "c3" from tab1 as t1`)
	r.True(v_.SimpleSql())
	r.Equal("tab1", v_.TableOfSimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{"tab1": {"col1", "col2", "col3"}})
	validateSelectTableColumns(r, v_, map[string][]string{"tab1": {"col1", "col2", "col3"}})
}

func TestAggregateFunctionSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `select count(distinct t.col1, t.col2) from tab1 t`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	r.Equal("tab1", v_.TableOfSimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{"tab1": {"col1", "col2"}})
	validateSelectTableColumns(r, v_, map[string][]string{"tab1": {"col1", "col2"}})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestAssignmentSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `insert into tab1 set name='nier', status=default`)
	r.Equal(SqlOperationType_.INSERT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	r.Equal("tab1", v_.TableOfSimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{"tab1": {"name", "status"}})
	validateSelectTableColumns(r, v_, map[string][]string{})
}

func TestCaseSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `select
  case
    (select
      t2.level
    from
      t_level t2
    where t2.customer_id = t.id)
    when 1
    then 'gold'
    when 2
    then 'diamond'
    else 'bronze'
  end gender
 from
  t_customer_info t;`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.False(v_.SimpleSql())
	validateTables(r, v_, "t_customer_info", "t_level")
	validateTableColumns(r, v_, map[string][]string{
		"t_customer_info": {"id"},
		"t_level":         {"level", "customer_id"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"t_level": {"level"},
	})
	validateWhereColumns(r, v_, map[string][]string{
		"t_customer_info": {"id"},
		"t_level":         {"customer_id"},
	})
}

func TestConvertFunctionSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT CONVERT(t.name USING utf8),t.age from tab1 t;`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"name", "age"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"name", "age"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestCastFunctionSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT cast(t.name as CHAR),t.age from tab1 t;`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"name", "age"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"name", "age"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestExistsSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT t1.col1 from tab1 t1 where exists(select t2.col1 from tab2 t2 where t2.pid=t1.id);`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.False(v_.SimpleSql())
	validateTables(r, v_, "tab1", "tab2")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"col1", "id"},
		"tab2": {"col1", "pid"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"col1"},
		"tab2": {"col1"},
	})
	validateWhereColumns(r, v_, map[string][]string{
		"tab1": {"id"},
		"tab2": {"pid"},
	})
}

func TestGroupBySyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT col1,count(*) from tab1 t1 group by t1.id desc,t1.col1 asc having col1='abc'`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"id", "col1"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"col1"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestIdentifierSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, "SELECT `age`,t.id,name from tab1 t")
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"name", "age", "id"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"name", "age", "id"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestInsertColumnListSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, "insert into test.tab1(id,col1,col2,col3) values(null,'abc',1,?)")
	r.Equal(SqlOperationType_.INSERT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "test.tab1")
	validateTableColumns(r, v_, map[string][]string{
		"test.tab1": {"id", "col1", "col2", "col3"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestJoinUsingSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, "select t1.col1, t2.col2 from tab1 t1 join tab2 t2 using(id)")
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.False(v_.SimpleSql())
	validateTables(r, v_, "tab1", "tab2")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"id", "col1"},
		"tab2": {"id", "col2"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"col1"},
		"tab2": {"col2"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestPartitionBySyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT
         year, product, profit,
         SUM(profit) OVER(PARTITION BY country) AS country_profit
       FROM sales`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "sales")
	validateTableColumns(r, v_, map[string][]string{
		"sales": {"year", "product", "profit", "country"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"sales": {"year", "product", "profit"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestPropertySyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, "select t.id from t_customer_info t")
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "t_customer_info")
	validateTableColumns(r, v_, map[string][]string{
		"t_customer_info": {"id"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"t_customer_info": {"id"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestTableReferenceSyntax(t *testing.T) {
	r := require.New(t)
	{
		v_ := visit(r, `select 
  * 
from 
  t_customer_info t use index (idx_name) 
  inner join t_level t2 use index for join (idx_customer_id) 
  left join t_favorite t3 force key for order by (idx_customer_id) 
  on t3.id = t.id 
  left join 
    (select 
      * 
    from 
      t_account) t4 
    on t.id = t4.id;`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t_customer_info", "t_favorite", "t_account", "t_level")
		validateTableColumns(r, v_, map[string][]string{
			"t_customer_info": {"*", "id"},
			"t_favorite":      {"*", "id"},
			"t_account":       {"*", "id"},
			"t_level":         {"*"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"t_customer_info": {"*"},
			"t_favorite":      {"*"},
			"t_account":       {"*"},
			"t_level":         {"*"},
		})
		validateWhereColumns(r, v_, map[string][]string{})
	}
	{
		v_ := visit(r, `SELECT * 
  FROM t1 
  LEFT JOIN(t2, t3) 
    ON (t2.a = t1.a AND t3.b = t1.b)`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t1", "t2", "t3")
		validateTableColumns(r, v_, map[string][]string{
			"t1": {"*", "a", "b"},
			"t2": {"*", "a"},
			"t3": {"*", "b"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"t1": {"*"},
			"t2": {"*"},
			"t3": {"*"},
		})
		validateWhereColumns(r, v_, map[string][]string{})
	}
	{
		v_ := visit(r, `select * 
  from t_customer_info t 
 inner join(t_person_info t3 
 inner join t_favorite t4 
    on t3.customer_id = t4.customer_id) on t3.customer_id = t.id 
 inner join (select * from t_level union select * from t_level) t2 
    on t.id = t2.customer_id`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t_customer_info", "t_person_info", "t_favorite", "t_level")
		validateTableColumns(r, v_, map[string][]string{
			"t_customer_info": {"*", "id"},
			"t_person_info":   {"*", "customer_id"},
			"t_favorite":      {"*", "customer_id"},
			"t_level":         {"*", "customer_id"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"t_customer_info": {"*"},
			"t_person_info":   {"*"},
			"t_favorite":      {"*"},
			"t_level":         {"*"},
		})
		validateWhereColumns(r, v_, map[string][]string{})
	}
	{
		v_ := visit(r, `SELECT
	t1.col1,
	t2.col1,
	t2.col2,
	(
	SELECT
		col3
	from
		tab3 t3
	where
		t3.pid = t1.id
		AND EXISTS(
		SELECT
			1
		FROM
			tab4 t4
		where
			t4.col1 = t1.col1
		    and t4.col2 = t2.col2)) col3
FROM
	tab1 t1
JOIN (
	SELECT
	    pid,
		col1,
		col2
	FROM
		tab2) t2 on
	t2.pid = t1.id`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "tab1", "tab2", "tab3", "tab4")
		validateTableColumns(r, v_, map[string][]string{
			"tab1": {"col1", "id"},
			"tab2": {"col1", "col2", "pid"},
			"tab3": {"pid", "col3"},
			"tab4": {"col1", "col2"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"tab1": {"col1"},
			"tab2": {"pid", "col1", "col2"},
			"tab3": {"col3"},
		})
		validateWhereColumns(r, v_, map[string][]string{
			"tab1": {"id", "col1"},
			"tab2": {"col2"},
			"tab3": {"pid"},
			"tab4": {"col1", "col2"},
		})
	}
	{
		v_ := visit(r, `select t.* from (select col1 from tab1) t`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "tab1")
		validateTableColumns(r, v_, map[string][]string{
			"tab1": {"col1"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"tab1": {"col1"},
		})
		validateWhereColumns(r, v_, map[string][]string{})
	}
	{
		v_ := visit(r, `select col1 from (select * from (select * from tab1) t) t2`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "tab1")
		validateTableColumns(r, v_, map[string][]string{
			"tab1": {"col1", "*"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"tab1": {"col1", "*"},
		})
		validateWhereColumns(r, v_, map[string][]string{})
	}
	{
		v_ := visit(r, `select * from tab1 t1 join (select * from tab2) t2 using(id)`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "tab1", "tab2")
		validateTableColumns(r, v_, map[string][]string{
			"tab1": {"id", "*"},
			"tab2": {"id", "*"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"tab1": {"*"},
			"tab2": {"*"},
		})
		validateWhereColumns(r, v_, map[string][]string{})
	}
}

func TestLikeSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT t.username LIKE 'S%',t.first_name not like 'A%',t.name LIKE 'David|_' ESCAPE'|' from t_person t;`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "t_person")
	validateTableColumns(r, v_, map[string][]string{
		"t_person": {"username", "first_name", "name"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"t_person": {"username", "first_name", "name"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestOrderBySyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `select * from t_customer_info t order by t.gender desc,t.level asc`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "t_customer_info")
	validateTableColumns(r, v_, map[string][]string{
		"t_customer_info": {"*", "gender", "level"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"t_customer_info": {"*"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestUnionSyntax(t *testing.T) {
	r := require.New(t)
	{
		v_ := visit(r, `SELECT t.name from t_person t where t.gender = 1 UNION all SELECT t.name from t_person t where t.age between 10 and 20`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t_person")
		validateTableColumns(r, v_, map[string][]string{
			"t_person": {"name", "age", "gender"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"t_person": {"name"},
		})
		validateWhereColumns(r, v_, map[string][]string{
			"t_person": {"age", "gender"},
		})
	}
	{
		v_ := visit(r, `(SELECT * FROM t1 WHERE a=10 AND B=1 order by c LIMIT 10) 
UNION 
(SELECT * FROM t2 WHERE a=11 AND B=2 order by c LIMIT 10) order by a;`)
		r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t1", "t2")
		validateTableColumns(r, v_, map[string][]string{
			"t1": {"*", "a", "B", "c"},
			"t2": {"*", "a", "B", "c"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{
			"t1": {"*"},
			"t2": {"*"},
		})
		validateWhereColumns(r, v_, map[string][]string{
			"t1": {"a", "B"},
			"t2": {"a", "B"},
		})
	}
}

func TestWindowSpecSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT
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
	observations`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "observations")
	validateTableColumns(r, v_, map[string][]string{
		"observations": {"time", "subject", "val"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"observations": {"time", "subject", "val"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestNamedWindowsSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT
  DISTINCT year, country,
  FIRST_VALUE(year) OVER (w ORDER BY year ASC) AS first,
  FIRST_VALUE(year) OVER w AS last
FROM sales
WINDOW w AS (PARTITION BY country),w2 as (w);`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "sales")
	validateTableColumns(r, v_, map[string][]string{
		"sales": {"country", "year"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"sales": {"country", "year"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
}

func TestMySqlDeleteSyntax(t *testing.T) {
	r := require.New(t)
	{
		v_ := visit(r, `DELETE LOW_PRIORITY QUICK IGNORE
FROM
  t_customer_info t1
WHERE t1.gender = 2
ORDER BY
  t1.id
LIMIT 1,2`)
		r.Equal(SqlOperationType_.DELETE, v_.SqlOperationType())
		r.True(v_.SimpleSql())
		validateTables(r, v_, "t_customer_info")
		validateTableColumns(r, v_, map[string][]string{
			"t_customer_info": {"gender", "id"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{})
		validateWhereColumns(r, v_, map[string][]string{
			"t_customer_info": {"gender"},
		})
		validateUpdateTables(r, v_, "t_customer_info")
	}
	{
		v_ := visit(r, `delete
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
  where t3.customer_id = t.id);`)
		r.Equal(SqlOperationType_.DELETE, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t_customer_info", "t_level", "t_communicating_info")
		validateTableColumns(r, v_, map[string][]string{
			"t_customer_info":      {"id", "gender"},
			"t_level":              {"level", "customer_id"},
			"t_communicating_info": {"customer_id"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{})
		validateWhereColumns(r, v_, map[string][]string{
			"t_customer_info":      {"gender", "id"},
			"t_level":              {"level"},
			"t_communicating_info": {"customer_id"},
		})
		validateUpdateTables(r, v_, "t_customer_info", "t_level")
	}
	{
		v_ := visit(r, `DELETE
FROM
  t1.*,
  t2.* USING tab1 t1
             INNER JOIN (tab2 t2)
             INNER JOIN tab3 t3
WHERE t1.id = t2.id
  AND t2.id = t3.id`)
		r.Equal(SqlOperationType_.DELETE, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "tab1", "tab2", "tab3")
		validateTableColumns(r, v_, map[string][]string{
			"tab1": {"id"},
			"tab2": {"id"},
			"tab3": {"id"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{})
		validateWhereColumns(r, v_, map[string][]string{
			"tab1": {"id"},
			"tab2": {"id"},
			"tab3": {"id"},
		})
		validateUpdateTables(r, v_, "tab1", "tab2")
	}
}

func TestMySqlInsertSyntax(t *testing.T) {
	r := require.New(t)
	{
		v_ := visit(r, `INSERT INTO
tab1
  (col1, col2, col3, col4)
VALUES
  ('a', 1, 'b', 'c')`)
		r.Equal(SqlOperationType_.INSERT, v_.SqlOperationType())
		r.True(v_.SimpleSql())
		validateTables(r, v_, "tab1")
		validateTableColumns(r, v_, map[string][]string{
			"tab1": {"col1", "col2", "col3", "col4"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{})
		validateWhereColumns(r, v_, map[string][]string{})
		validateUpdateTables(r, v_, "tab1")
	}
	{
		v_ := visit(r, `INSERT INTO tab1 SET a = 'abc', b = 1, c = ?`)
		r.Equal(SqlOperationType_.INSERT, v_.SqlOperationType())
		r.True(v_.SimpleSql())
		validateTables(r, v_, "tab1")
		validateTableColumns(r, v_, map[string][]string{
			"tab1": {"a", "b", "c"},
		})
		validateSelectTableColumns(r, v_, map[string][]string{})
		validateWhereColumns(r, v_, map[string][]string{})
		validateUpdateTables(r, v_, "tab1")
	}
}

func TestMySqlIntervalSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT DATE_ADD('2018-05-01',INTERVAL t.offset DAY) from tab1 t;`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"offset"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"offset"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
	validateUpdateTables(r, v_)
}

func TestMySqlTrimFunctionSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT TRIM(col1) from tab1`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"col1"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"col1"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
	validateUpdateTables(r, v_)
}

func TestMySqlUnarySyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT BINARY col1 from tab1`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"col1"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"col1"},
	})
	validateWhereColumns(r, v_, map[string][]string{})
	validateUpdateTables(r, v_)
}

func TestMySqlUpdateSyntax(t *testing.T) {
	r := require.New(t)
	{
		v_ := visit(r, `UPDATE LOW_PRIORITY IGNORE 
  t_user a,
  t_account b
SET
  b.username = a.username
WHERE 
  a.id = b.user_id;`)
		r.Equal(SqlOperationType_.UPDATE, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t_user", "t_account")
		validateTableColumns(r, v_, map[string][]string{
			"t_user":    {"id", "username"},
			"t_account": {"username", "user_id"},
		})
		validateWhereColumns(r, v_, map[string][]string{
			"t_user":    {"id"},
			"t_account": {"user_id"},
		})
		validateUpdateTables(r, v_, "t_account")
	}
	{
		v_ := visit(r, `UPDATE
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
limit 3`)
		r.Equal(SqlOperationType_.UPDATE, v_.SqlOperationType())
		r.False(v_.SimpleSql())
		validateTables(r, v_, "t_parent", "t_children", "t_attribute")
		validateTableColumns(r, v_, map[string][]string{
			"t_parent":    {"id", "stat", "type"},
			"t_children":  {"parent_id", "stat", "name"},
			"t_attribute": {"parent_id", "att1"},
		})
		validateWhereColumns(r, v_, map[string][]string{
			"t_parent":    {"type"},
			"t_children":  {"name"},
			"t_attribute": {"att1"},
		})
		validateUpdateTables(r, v_, "t_parent", "t_children")
	}
}

func TestMySqlTableSyntax(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `SELECT
  id
FROM
  tab1
WHERE col1 = 'aa'
  AND col2 > SOME (TABLE tab2)`)
	r.Equal(SqlOperationType_.SELECT, v_.SqlOperationType())
	r.True(v_.SimpleSql())
	validateTables(r, v_, "tab1", "tab2")
	validateTableColumns(r, v_, map[string][]string{
		"tab1": {"id", "col1", "col2"},
		"tab2": {"*"},
	})
	validateSelectTableColumns(r, v_, map[string][]string{
		"tab1": {"id"},
		"tab2": {"*"},
	})
	validateWhereColumns(r, v_, map[string][]string{
		"tab1": {"col1", "col2"},
	})
	validateUpdateTables(r, v_)
}

func TestMySqlVisitorError(t *testing.T) {
	r := require.New(t)
	validateError(r,
		`select t1.col,t2.col1 from tab1 t1 left join tab2 t1 using(id)`,
		`table alias 't1' is duplicate
select t1.col,t2.col1 from tab1 t1 left join ↪tab2 t1↩ using(id)`)
	{
		// SELECT * FROM tab1 t1 UNION SELECT * FROM tab2 t2 ORDER BY t1.id
		sil1 := ast.NewSelectItemListSyntax()
		sil1.Add(ast.NewAllColumnSyntax())
		tn1 := ast.NewMySqlIdentifierSyntax()
		tn1.Name = "tab1"
		tni1 := ast.NewTableNameItemSyntax()
		tni1.TableName = tn1
		alias1 := ast.NewMySqlIdentifierSyntax()
		alias1.Name = "t1"
		nt1 := ast.NewMySqlNameTableReferenceSyntax()
		nt1.TableNameItem = tni1
		nt1.Alias = alias1
		left := ast.NewMySqlSelectSyntax()
		left.SelectItemList = sil1
		left.TableReference = nt1

		sil2 := ast.NewSelectItemListSyntax()
		sil2.Add(ast.NewAllColumnSyntax())
		tn2 := ast.NewMySqlIdentifierSyntax()
		tn2.Name = "tab2"
		tni2 := ast.NewTableNameItemSyntax()
		tni2.TableName = tn2
		alias2 := ast.NewMySqlIdentifierSyntax()
		alias2.Name = "t2"
		nt2 := ast.NewMySqlNameTableReferenceSyntax()
		nt2.TableNameItem = tni2
		nt2.Alias = alias2
		right := ast.NewMySqlSelectSyntax()
		right.SelectItemList = sil2
		right.TableReference = nt2

		owner := ast.NewMySqlIdentifierSyntax()
		owner.Name = "t1"
		value := ast.NewMySqlIdentifierSyntax()
		value.Name = "id"
		p := ast.NewPropertySyntax()
		p.Owner = owner
		p.Value = value

		oi := ast.NewOrderingItemSyntax()
		oi.Column = p
		oil := ast.NewOrderingItemListSyntax()
		oil.Add(oi)
		ob := ast.NewOrderBySyntax()
		ob.OrderByItemList = oil

		ms := ast.NewMySqlMultisetSyntax()
		ms.LeftQuery = left
		ms.RightQuery = right
		ms.OrderBy = ob

		_, err := ast.Visit(ms)
		r.EqualError(err,
			`cannot be used table alias in global clause of multiset syntax
`)
	}
	validateWarn(r,
		`select col1 from tab1 t1 left join tab2 t2 using(id)`,
		`[0]column 'col1' is ambiguous
select ↪col1↩ from tab1 t1 left join tab2 t2 using(id)`)
	validateError(r,
		`select col from dual`,
		`unknown column 'col'
select ↪col↩ from dual`)
	validateWarn(r,
		`select *
from
 tab1
 join (select *
       from
         tab2
         join (select * from tab3) t3 using(id)) t using(id)`,
		`[0]column 'id' is ambiguous
select *
from
 tab1
 join (select *
       from
         tab2
         join (select * from tab3) t3 using(id)) t using(↪id↩)`)
	validateError(r,
		`SELECT t3.name FROM
(SELECT
   t1.col_1 AS name,
   t2.col_1 AS name
 FROM
   tab_1 t1
   JOIN tab_2 t2
     ON t1.id = t2.id) t3`,
		`column 't3.name' is ambiguous
SELECT ↪t3.name↩ FROM
(SELECT
   t1.col_1 AS name,
   t2.col_1 AS name
 FROM
   tab_1 t1
   JOIN tab_2 t2
     ON t1.id = t2.id) t3`)
	validateError(r,
		`SELECT col2 FROM
(SELECT
   t1.col_1
 FROM
   tab_1 t1) t`,
		`unknown column 'col2'
SELECT ↪col2↩ FROM
(SELECT
   t1.col_1
 FROM
   tab_1 t1) t`)
	validateError(r,
		`DELETE a1, a3 FROM t1 AS a1 INNER JOIN t2 AS a2
WHERE a1.id=a2.id;`,
		`unknown table of alias 'a3' in MULTI DELETE
DELETE a1, ↪a3↩ FROM t1 AS a1 INNER JOIN t2 AS a2
WHERE a1.id=a2.id;`)

}

func TestHintContent(t *testing.T) {
	r := require.New(t)
	v_ := visit(r, `UPDATE /*+ NO_MERGE(discounted) */ items,
       (SELECT id FROM items
        WHERE retail / wholesale >= 1.3 AND quantity < 100)
        AS discounted
    SET items.retail = items.retail * 0.9
    WHERE items.id = discounted.id;`)
	r.Equal(" NO_MERGE(discounted) ", v_.HintContent())
}
