package schema

import (
	"dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Type string
	Name string
	Tag  string
}

type Schema struct {
	Model      any
	Name       string
	Fields     []*Field
	FieldNames []string
	FieldMap   map[string]*Field
}

func (s *Schema) Keys() (keys []string) {
	for k, _ := range s.FieldMap {
		keys = append(keys, k)
	}
	return
}

func (s *Schema) Values() (val []*Field) {
	for _, v := range s.FieldMap {
		val = append(val, v)
	}
	return
}

func (s *Schema) GetField(name string) *Field {
	return s.FieldMap[name]
}

func Parse(dest any, d dialect.Dialect) (schema *Schema) {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema = &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		FieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 非匿名字段并且可导出
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("orm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.FieldMap[p.Name] = field
		}
	}

	return
}

func (s *Schema) RecordValues(dest any) []any {
	destValue := reflect.Indirect(reflect.ValueOf(dest))

	var fieldValues []any
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
