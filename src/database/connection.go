package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/microsoft/go-mssqldb"
)

// connectionString -> Retorna la cadena de conexion en para sql server
func connectionString() string {
	godotenv.Load()

	server := os.Getenv("SERVER_DB")
	port := os.Getenv("PORT_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	database := os.Getenv("NAME_DB")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", server, user, password, port, database)

	return connString
}

// CreateConnection -> Crea la conexion a la base de datos
func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlserver", connectionString())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
