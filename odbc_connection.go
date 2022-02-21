package odbc

import (
	"errors"
	"fmt"

	_ "github.com/alexbrainman/odbc"
	"github.com/jmoiron/sqlx"
)

var execution_type_row uint8 = 0
var execution_type_rows uint8 = 1

type ODBCConnection struct {
	conn    *sqlx.DB
	options ConnectionOptions
}

// func (o *ODBCConnection) UseExisitingConnection(conn *sqlx.DB, ops ConnectionOptions) {
// 	o.conn = conn
// 	o.setOptions(ops)
// }

func (o *ODBCConnection) Connect(config ConnectionConfig, ops ConnectionOptions) error {
	// If trying to connect on this connection and already connected
	// then close the current connection
	if o.conn != nil {
		o.Close()
	}
	// Create the new connection and assign it
	connStr, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	conn, err := sqlx.Connect("odbc", connStr)
	if err != nil {
		return err
	}

	o.conn = conn
	o.setOptions(ops)

	return nil
}

func (o *ODBCConnection) Query(sql string, params []QueryParameter, dest interface{}) error {
	return o.QueryRows(sql, params, dest)
}

func (o *ODBCConnection) QueryRows(sql string, params []QueryParameter, dest interface{}) error {
	return o.executeQuery(execution_type_rows, sql, params, dest)
}

func (o *ODBCConnection) QueryRow(sql string, params []QueryParameter, dest interface{}) error {
	return o.executeQuery(execution_type_row, sql, params, dest)
}

func (o *ODBCConnection) Close() {
	o.conn.Close()
	o.conn = nil
}

// Private

func (o *ODBCConnection) verifyConnection() error {
	if o.conn == nil {
		return errors.New("no active ODBC connection found")
	}

	return nil
}

func (o *ODBCConnection) executeQuery(t uint8, sql string, params []QueryParameter, dest interface{}) error {
	if err := o.verifyConnection(); err != nil {
		return err
	}

	query, err := o.prepareQuery(sql, params)
	if err != nil {
		return err
	}

	var qerr error
	switch t {
	case execution_type_row:
		qerr = o.conn.Get(dest, query)
		if o.options.TrimEndOfResults {
			TrimEndOfResults(dest)
		}
	case execution_type_rows:
		qerr = o.conn.Select(dest, query)
		if o.options.TrimEndOfResults {
			TrimEndOfResults(dest)
		}
	default:
		return fmt.Errorf("incorrect execution type provided: %d", t)
	}

	if qerr != nil {
		return qerr
	}

	return nil
}

func (o *ODBCConnection) setOptions(ops ConnectionOptions) {
	o.options = ops
}

func (o *ODBCConnection) prepareQuery(sql string, params []QueryParameter) (string, error) {
	query := sql

	if o.options.UseCustomInsertParams {
		q, err := customInsertParams(sql, params)
		if err != nil {
			return "", err
		}
		query = q
	}

	return query, nil
}
