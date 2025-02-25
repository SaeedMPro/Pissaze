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
func init(){
	fmt.Println("connecting to database...")
	err := godotenv.Load()

	if err != nil {
		panic(err)
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
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	dataBase = db
	fmt.Println("connected to database")
}

func GetDB()(*sql.DB , error){
	if dataBase == nil {
		return nil , errors.New("pointer is null")
	}
	return dataBase , nil
}