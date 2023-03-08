package dialect

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type mysql struct {
}

var _ Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}

func (m *mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int8:
		return "tinyint"
	case reflect.Int16:
		return "smallint"
	case reflect.Int, reflect.Int32:
		return "integer"
	case reflect.Int64:
		return "bigint"
	case reflect.Uint8:
		return "tinyint unsigned"
	case reflect.Uint16:
		return "smallint unsigned"
	case reflect.Uint, reflect.Uint32:
		return "integer unsigned"
	case reflect.Uint64:
		return "bigint unsigned"
	case reflect.Float32, reflect.Float64:
		return "double precision"
	case reflect.String:
		return "varchar(20)"
	case reflect.Array, reflect.Slice:
		return "text"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (m *mysql) TableExistSQL(tableName string) (string, []any) {
	args := []any{tableName}
	return "select TABLE_NAME from INFORMATION_SCHEMA.TABLES where TABLE_NAME = ? ", args
}

func (m *mysql) TableEqual(name1, name2 string) bool {
	return strings.ToLower(name1) == strings.ToLower(name2)
}
