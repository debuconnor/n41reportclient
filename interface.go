package n41reportclient

import "database/sql"

type ConnectionInfo struct {
	Host   string
	Port   string
	Dbname string
	UserId string
	UserPw string
}

type Database struct {
	Db          *sql.DB
	QueryString string
	Credentials ConnectionInfo
}

type Connector interface {
	Connect() error
	Disconnect() error
	Select(string) error
}
