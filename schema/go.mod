module schema

go 1.20

require (
	dialect v1.0.0
	model v1.0.0
)

replace (
	dialect v1.0.0 => ../dialect
	model v1.0.0 => ../model
)
