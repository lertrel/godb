package godb

// ----------------------------------------------------------------------------
// SPI ------------------------------------------------------------------------
// ----------------------------------------------------------------------------

//[ConnectionDriver]
type ConnectionDriver interface {

	Name() string
	Get(config map[string]string) (Connection, error)
}

//[DBConnection]
type Connection interface {

	SQL(sql string) Statement
	MappedSQL(sql string) Statement
	DML(dml string) Statement
	MappedDML(dml string) Statement
}

//[DBStatement]
type Statement interface {
	IsSQL() bool
	IsDML() bool
	IsMappable() bool
	Execute() (ResultSet, error)
	Next() error
	SetString(index int, val string)
	SetInt(index int, val int)
	SetFloat(index int, val float64)
	SetBytes(index int, val []byte)
}

//[ResultSet]
type ResultSet interface {

	GetColumnCount() int
	GetColumnName(index int) string
	GetColumnType(index int) string
	Next() bool
	GetInt(index int) int
	GetFloat(index int) float64
	GetString(index int) string
	GetBytes(index int) []byte

}
