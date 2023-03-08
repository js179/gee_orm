package model

type User struct {
	Id   int `json:"id,string" orm:"PRIMARY KEY"`
	Name string
	Age  int `json:"age,string"`
}
