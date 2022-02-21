package odbc

var DEFAULT_CONNECTION_OPTIONS = ConnectionOptions{
	TrimEndOfResults:      true,
	UseCustomInsertParams: true,
}

type databaseConnectionStringKeys struct {
	DriverKey   string
	ServerKey   string
	DatabaseKey string
	UserKey     string
	PasswordKey string
}

const (
	DEFAULT_PROVIDER   = PROVIDER_ACTIANZEN
	PROVIDER_ACTIANZEN = "actianzen"
	PROVIDER_MYSQL     = "mysql"
	PROVIDER_PG        = "pg"
	PROVIDER_MSSQL     = "mssql"
)

var DATABASE_CONNECTION_STRING_KEYS = map[string]databaseConnectionStringKeys{
	PROVIDER_ACTIANZEN: {
		DriverKey:   "Driver",
		ServerKey:   "ServerName",
		DatabaseKey: "DBQ",
		UserKey:     "UID",
		PasswordKey: "PWD",
	},
	PROVIDER_MYSQL: {
		DriverKey:   "DRIVER",
		ServerKey:   "SERVER",
		DatabaseKey: "DATABASE",
		UserKey:     "USER",
		PasswordKey: "PASSWORD",
	},
	PROVIDER_PG: {
		DriverKey:   "DRIVER",
		ServerKey:   "SERVER",
		DatabaseKey: "DATABASE",
		UserKey:     "UID",
		PasswordKey: "PWD",
	},
	PROVIDER_MSSQL: {
		DriverKey:   "DRIVER",
		ServerKey:   "SERVER",
		DatabaseKey: "DATABASE",
		UserKey:     "UID",
		PasswordKey: "PWD",
	},
}
