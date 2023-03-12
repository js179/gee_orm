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

type TxFunc func(*session.Session) (any, error)

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

func (e *Engine) Transaction(f TxFunc) (result any, err error) {
	s := e.NewSession()
	if err = s.Begin(); err != nil {
		logf.Error(err)
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = s.RollBack()
			panic(r)
		} else if err != nil {
			_ = s.RollBack()
		} else {
			err = s.Commit()
		}
	}()

	return f(s)
}
