package odbc

// Creates an ODBC Connection, connects to it, and returns it for use
func CreateConnection(config ConnectionConfig, options ConnectionOptions) (ODBCConnection, error) {
	conn := ODBCConnection{}
	err := conn.Connect(config, options)
	return conn, err
}
