package dialect

import "reflect"

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []any)
}

var dialectMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
