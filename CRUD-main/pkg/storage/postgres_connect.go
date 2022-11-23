package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

var (
	PGUser     = "userapi"
	PGPassword = "userapi"
	PGName     = "userapi"
	PGHost     = "192.168.33.50"
	PGPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func readEnv() {
	if os.Getenv("PG_USER") != "" {
		PGUser = os.Getenv("PG_USER")
	}

	if os.Getenv("PG_PASSWORD") != "" {
		PGPassword = os.Getenv("PG_PASSWORD")
	}

	if os.Getenv("PG_NAME") != "" {
		PGName = os.Getenv("PG_NAME")
	}

	if os.Getenv("PG_HOST") != "" {
		PGHost = os.Getenv("PG_HOST")
	}

	if os.Getenv("PG_PORT") != "" {
		PGPort = os.Getenv("PG_PORT")
	}
}

func GetPostgresConnectionString() string {

	readEnv()
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		PGHost,
		PGPort,
		PGUser,
		PGName,
		PGPassword)
	return dataBase
}

func SetupPostgresConnection(params ...string) *gorm.DB {
	var err error
	conString := GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
