package clause

import "testing"

var cla Clause

func Test_Clause_Insert(t *testing.T) {
	cla.Set(INSERT, "USER", []string{"id", "age", "name"})
	cla.Set(VALUES, []any{1002, 23, "ZS"})
	sql, vars := cla.Build(INSERT, VALUES)
	t.Log(sql, vars)
}

func Test_Clause_Select(t *testing.T) {
	cla.Set(SELECT, "USER", []string{"id", "age", "name"})
	cla.Set(WHERE, "NAME = ?", "Tom")
	cla.Set(LIMIT, 5)
	cla.Set(ORDER, "id asc")
	sql, vars := cla.Build(SELECT, WHERE, LIMIT, ORDER)
	t.Log(sql, vars)
}
