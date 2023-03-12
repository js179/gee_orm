package session

import "logf"

func (s *Session) Begin() (err error) {
	logf.Info("transaction begin...")
	if s.tx, err = s.db.Begin(); err != nil {
		logf.Error(err)
	}
	return
}

func (s *Session) Commit() (err error) {
	logf.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		logf.Error(err)
	}
	return
}

func (s *Session) RollBack() (err error) {
	logf.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		logf.Error(err)
	}
	return
}
