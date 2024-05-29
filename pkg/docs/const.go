package docs

const VERSION = "3.1.0"

type DataType string

const (
	STRING  DataType = "string"
	INTEGER DataType = "integer"
	NUMBER  DataType = "number"
)

type Format string

const (
	INT32    Format = "int32"
	INT64    Format = "int64"
	FLOAT    Format = "float"
	DOUBLE   Format = "double"
	PASSWORD Format = "password"
)
