package session

import (
	"clause"
	"database/sql"
	"dialect"
	"logf"
	"schema"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	clause   clause.Clause
	sql      strings.Builder
	sqlVars  []any
}

func New(db *sql.DB, dl dialect.Dialect) *Session {
	return &Session{db: db, dialect: dl}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, vars ...any) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	if vars != nil {
		s.sqlVars = append(s.sqlVars, vars...)
	}
	return s
}

func (s *Session) Exec() (exec sql.Result, err error) {
	defer s.Clear()
	sql := s.sql.String()
	logf.Info(sql, s.sqlVars)
	if exec, err = s.db.Exec(sql, s.sqlVars...); err != nil {
		logf.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	sql := s.sql.String()
	logf.Info(sql, s.sqlVars)
	return s.db.QueryRow(sql, s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	sql := s.sql.String()
	logf.Info(sql, s.sqlVars)
	if rows, err = s.db.Query(sql, s.sqlVars...); err != nil {
		logf.Error(err)
	}
	return
}
