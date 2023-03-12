package main

import (
	"engine"
	"logf"
	"model"
	"session"
)

func init() {
	logf.SetLevel(logf.All)
}

func main() {
	engine, _ := engine.Open("mysql", "root:123321@tcp(localhost:3306)/golang?charset=utf8mb4")
	defer engine.Close()
	sql := engine.NewSession()
	//_, _ = sql.Raw("select id, name, age from go_table").QueryRows()
	//row := sql.Raw("select id, name, age from go_table limit 1").QueryRow()
	//var u = model.User{}
	//// Scan中参数为属性字段地址
	//err := row.Scan(&u.Id, &u.Name, &u.Age)
	//if err != nil {
	//	logf.Error(err)
	//}
	//logf.Info(u)
	//
	//sql = sql.Model(model.User{})
	//_ = sql.CreateTable()
	//logf.Infof("has table %v", sql.HasTable())

	//var u1 = model.User{Id: 5, Name: "ZC", Age: 32}
	//res, _ := sql.Insert(&u1)
	//logf.Infof("insert success? %v", res == 1)

	var us []model.User
	sql.Find(&us)
	for _, v := range us {
		logf.Info(v)
	}

	var us2 []model.User
	sql.Limit(1).Find(&us2)
	logf.Infof("len = %d\n", len(us2))
	logf.Infof("val = %#v\n", us2[0])
	//val = model.User{Id:1, Name:"zs", Age:23}

	engine.Transaction(func(sql *session.Session) (any, error) {
		updates := make(map[string]any)
		updates["age"] = 25
		res, err := sql.Where("name = ?", "zs").Update(updates)
		if err == nil {
			logf.Info(res == 1)
			var us3 []model.User
			sql.Where("name = ?", "zs").Find(&us3)
			logf.Infof("val = %v\n", us3[0])
		}
		return nil, err
	})

	sql.Where("name = ? or age = ?", "ws", 31).Delete()
	num, _ := sql.Count()
	logf.Infof("data number = %d\n", num)
}
