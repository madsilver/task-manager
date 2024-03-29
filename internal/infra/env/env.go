package env

import "os"

const (
	MysqlUser     = "silver"
	MysqlPassword = "silver"
	MysqlDatabase = "silverlabs"
	MysqlHost     = "localhost"
	MysqlPort     = "3306"
)

func GetString(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
