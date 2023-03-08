package clause

import (
	"fmt"
	"strings"
)

type generatorFunc func(values ...any) (string, []any)

var generatorMap map[Type]generatorFunc

func init() {
	generatorMap = make(map[Type]generatorFunc)
	generatorMap[INSERT] = _insert
	generatorMap[VALUES] = _values
	generatorMap[SELECT] = _select
	generatorMap[WHERE] = _where
	generatorMap[LIMIT] = _limit
	generatorMap[ORDER] = _orderBy
	generatorMap[UPDATE] = _update
	generatorMap[DELETE] = _delete
	generatorMap[COUNT] = _count
}

// 生成 ?, ?, ? 带填充字符串
func genBindVars(num int) string {
	if num < 1 {
		return ""
	}
	var vars strings.Builder
	for i := 0; i < num-1; i++ {
		vars.WriteString("?, ")
	}
	vars.WriteString("?")
	return vars.String()
}

func _insert(values ...any) (sql string, vars []any) {
	l := len(values)
	if l != 2 && l != 1 {
		panic("the insert param number must be equal 1 or 2")
	}

	table := values[0]
	if l == 2 {
		var fields = strings.Join(values[1].([]string), ", ")
		sql = fmt.Sprintf("INSERT INTO %s(%v)", table, fields)
	} else {
		sql = fmt.Sprintf("INSERT INTO %s", table)
	}
	return
}

func _values(values ...any) (string, []any) {

	if len(values) < 1 {
		panic("the insert param number must be greater than or equal 1")
	}

	firstVal := values[0].([]any)
	end := len(values) - 1
	// 生成 ?, ?, ? 带填充字符串
	var bindStr = genBindVars(len(firstVal))

	var sql strings.Builder
	sql.WriteString("VALUES ")
	var vars []any

	for i := 0; i < end; i++ {
		v := values[i].([]any)
		sql.WriteString(fmt.Sprintf("(%v), ", bindStr))
		vars = append(vars, v...)
	}
	// 最后一个不加,
	sql.WriteString(fmt.Sprintf("(%v);", bindStr))
	endVal := values[end].([]any)
	vars = append(vars, endVal...)

	return sql.String(), vars
}

func _select(values ...any) (sql string, vars []any) {
	l := len(values)
	if l != 2 && l != 1 {
		panic("the select param number must be equal 1 or 2")
	}

	tableName := values[0]
	if l == 2 {
		fields := strings.Join(values[1].([]string), ", ")
		sql = fmt.Sprintf("SELECT %v FROM %s", fields, tableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s", tableName)
	}
	return
}

func _where(values ...any) (sql string, vars []any) {
	if len(values) < 1 {
		panic("the where param number must bu greater than or equal 1")
	}

	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

func _limit(values ...any) (sql string, vars []any) {
	if len(values) < 1 {
		return
	}
	sql = "LIMIT ?"
	vars = values
	return
}

func _orderBy(values ...any) (sql string, vars []any) {
	if len(values) < 1 {
		return
	}
	sql = fmt.Sprintf("ORDER BY %s", values[0])
	return
}

func _update(values ...any) (sql string, vars []any) {
	if len(values) != 2 {
		panic("the update param number must be 2")
	}

	tableName := values[0]
	params := values[1].(map[string]any)

	var keys []string

	for key, val := range params {
		keys = append(keys, key+" = ?")
		vars = append(vars, val)
	}
	sql = fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(keys, ", "))
	return
}

func _delete(values ...any) (sql string, vars []any) {
	if len(values) < 1 {
		panic("the count param number must be greater than 1")
	}
	sql = fmt.Sprintf("DELETE FROM %s", values[0])
	return
}

func _count(values ...any) (sql string, vars []any) {
	if len(values) < 1 {
		panic("the count param number must be greater than 1")
	}
	sql, vars = _select(values[0], []string{"COUNT(0)"})
	return
}
