package engine

import (
	"database/sql"
	"dialect"
	_ "github.com/go-sql-driver/mysql"
	"logf"
	"session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func Open(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		logf.Error(err)
		return
	}

	err = db.Ping()
	if err != nil {
		logf.Error(err)
		return
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		logf.Error("dialect %s not found", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}
	logf.Info("database connect success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		logf.Error("database closed fail")
	}
	logf.Info("database closed success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
