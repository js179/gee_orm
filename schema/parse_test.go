package schema

import (
	"dialect"
	"model"
	"testing"
)

func Test_Parse(t *testing.T) {
	user := model.User{}
	d, _ := dialect.GetDialect("mysql")
	schema := Parse(user, d)
	t.Log(schema.Name)
	t.Log(schema.FieldNames)
	for _, v := range schema.Fields {
		t.Log(v)
	}
}
