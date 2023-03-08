module session

go 1.20

replace (
	dialect v1.0.0 => ../dialect
	logf v1.0.0 => ../log
	model v1.0.0 => ../model
	schema v1.0.0 => ../schema
)

require (
	dialect v1.0.0
	logf v1.0.0
	schema v1.0.0
)
