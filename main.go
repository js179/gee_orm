package main

import (
	"engine"
	"logf"
	"model"
)

func init() {
	logf.SetLevel(logf.All)
}

type user struct {
	id   int
	name string
	Age  int
}

func main() {
	engine, _ := engine.Open("mysql", "root:123321@tcp(localhost:3306)/golang?charset=utf8mb4")
	defer engine.Close()
	sql := engine.NewSession()
	_, _ = sql.Raw("select id, name, age from go_table").QueryRows()
	row := sql.Raw("select id, name, age from go_table limit 1").QueryRow()
	var u = user{}
	// Scan中参数为属性字段地址
	err := row.Scan(&u.id, &u.name, &u.Age)
	if err != nil {
		logf.Error(err)
	}
	logf.Info(u)

	sql = sql.Model(model.User{})
	_ = sql.CreateTable()
	logf.Infof("has table %v", sql.HasTable())
}
