package database

var conn string

func init() {
	conn = "MySQL"
}

func GetConnection() string {
	return conn
}
