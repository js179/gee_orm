module gee_orm

go 1.20

replace (
	engine v1.0.0 => ./engine
	logf v1.0.0 => ./log
	session v1.0.0 => ./session
)

require (
	engine v1.0.0
	logf v1.0.0
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	session v1.0.0 // indirect
)
