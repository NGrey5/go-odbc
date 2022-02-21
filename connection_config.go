package odbc

import (
	"errors"
	"strings"
)

type ConnectionConfig interface {
	GetConnectionString() (string, error)
}

// A connection config using a data source name
type ConnectionConfigDSN struct {
	DSN string
}

func (c ConnectionConfigDSN) GetConnectionString() (string, error) {
	return "DSN=" + c.DSN, nil
}

// A connection config using a connection string
type ConnectionConfigConnectionString struct {
	ConnectionString string
}

func (c ConnectionConfigConnectionString) GetConnectionString() (string, error) {
	return c.ConnectionString, nil
}

// A connection config using odbc properties
type ConnectionConfigProperties struct {

	// The database provider
	Provider string

	// The odbc driver to use for the connection
	Driver string

	// The server to connect to
	Server string

	// The database on the server to make requests
	Database string

	// Authentication for the database
	Auth dbAuth
}

type dbAuth struct {
	User     string
	Password string
}

func (c ConnectionConfigProperties) GetConnectionString() (string, error) {
	p := c.Provider // Get the provider

	// If no provider, set the default
	if p == "" {
		p = DEFAULT_PROVIDER
	}

	keys, exists := DATABASE_CONNECTION_STRING_KEYS[p] // Get the connection string keys for the provider

	// If there is no struct, then an invalid provider was given
	if !exists {
		return "", errors.New("invalid database provider (" + p + ")")
	}

	drStr := keys.DriverKey + "=" + c.Driver
	srStr := keys.ServerKey + "=" + c.Server
	dbStr := keys.DatabaseKey + "=" + c.Database
	auStr := keys.UserKey + "=" + c.Auth.User + ";" + keys.PasswordKey + "=" + c.Auth.Password

	strs := []string{drStr, srStr, dbStr}
	// If there is auth, then use the auth string
	// Else if only one prop of auth is provided, throw an error
	if c.Auth.User != "" && c.Auth.Password != "" {
		strs = append(strs, auStr)
	} else if (c.Auth.User != "" && c.Auth.Password == "") || (c.Auth.User == "" && c.Auth.Password != "") {
		return "", errors.New("must provide both user and password for auth")
	}

	return strings.Join(strs, ";"), nil
}
