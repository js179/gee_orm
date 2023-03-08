module engine

go 1.20

replace (
	clause v1.0.0 => ../clause
	dialect v1.0.0 => ../dialect
	logf v1.0.0 => ../log
	model v1.0.0 => ../model
	schema v1.0.0 => ../schema
	session v1.0.0 => ../session
)

require (
	dialect v1.0.0
	github.com/go-sql-driver/mysql v1.7.0
	logf v1.0.0
	session v1.0.0
)

require (
	clause v1.0.0 //indirect
	schema v1.0.0 // indirect
)
