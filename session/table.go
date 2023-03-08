package session

import (
	"fmt"
	"logf"
	"reflect"
	"schema"
	"strings"
)

func (s *Session) Model(val any) *Session {

	if s.refTable == nil || reflect.TypeOf(val) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(val, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		logf.Error("model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ", ")
	_, err := s.Raw(fmt.Sprintf("create table if not exists %s (%s);", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	table := s.RefTable()
	_, err := s.Raw(fmt.Sprintf("drop table if exists %s", table.Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	table := s.RefTable()
	sql, args := s.dialect.TableExistSQL(table.Name)
	rows, err := s.Raw(sql, args...).QueryRows()
	return err == nil && rows.Next()
}
