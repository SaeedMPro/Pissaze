package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dataBase *sql.DB
func InitDB() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	dbHost, dbUser, dbPass, dbName, dbPort)

	db ,err := sql.Open("postgres" , connectionInfo)

	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	}

	dataBase = db
	return nil
}

func GetDB()(*sql.DB , error){
	if dataBase == nil {
		return nil , errors.New("pointer is null")
	}
	return dataBase , nil
}