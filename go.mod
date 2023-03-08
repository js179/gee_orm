module gee_orm

go 1.20

replace (
	dialect v1.0.0 => ./dialect
	engine v1.0.0 => ./engine
	logf v1.0.0 => ./log
	model v1.0.0 => ./model
	schema v1.0.0 => ./schema
	session v1.0.0 => ./session
)

require (
	engine v1.0.0
	logf v1.0.0
	model v1.0.0
)

require (
	dialect v1.0.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	schema v1.0.0 // indirect
	session v1.0.0 // indirect
)
