module engine

go 1.20

replace (
	logf v1.0.0 => ../log
	session v1.0.0 => ../session
)

require (
	github.com/go-sql-driver/mysql v1.7.0
	logf v1.0.0
	session v1.0.0
)
