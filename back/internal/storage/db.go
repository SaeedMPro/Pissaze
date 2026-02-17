package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dataBase *sql.DB
func init(){
	fmt.Println("connecting to database...")
	// Optional: load .env when present (e.g. local dev); in Docker env vars are set by compose
	_ = godotenv.Load()

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

func GetDB()(*sql.DB ){
	return dataBase
}