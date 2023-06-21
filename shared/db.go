package shared

import (
	"fmt"
	"os"
)

func GetDBConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	sslCertFile := os.Getenv("DB_SSL_CERT_FILE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, dbName, sslMode)
	if sslMode != "disable" {
		connStr += fmt.Sprintf("&sslrootcert=%s", sslCertFile)
	}

	return connStr
}
