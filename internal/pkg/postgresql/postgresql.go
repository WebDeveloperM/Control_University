package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"os"
)

func ConnectPostgresql() *bun.DB {
	godotenv.Load(".env")

	dbName := os.Getenv("DBNAME")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("PASSWORD")
	log.Println("new db" + dbName)
	log.Println("password  " + password)
	log.Println("username   " + userName)
	dsn := fmt.Sprintf("postgres://%v:%v@localhost:5432/%v?sslmode=disable", userName, password, dbName)
	connect := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(connect, pgdialect.New())

	return db

}
