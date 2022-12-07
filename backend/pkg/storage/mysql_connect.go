package storage

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MysqlConnect struct {
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_DB_NAME  string
	db             *gorm.DB
}

type DBConnector interface {
	readEnv()
	// getConnectionString() string
	SetupConnection()
	// GetDBInstance() *gorm.DB
}

func (c *MysqlConnect) GetDb() *gorm.DB {
	return c.db
}

// reads environmental variables needed to connect to the database
func (c *MysqlConnect) readEnv() {

	// var errMessage string

	if os.Getenv("MYSQL_USER") != "" {
		c.MYSQL_USER = os.Getenv("MYSQL_USER")
	} else {
		log.Fatal("Couldn't read MYSQL_USER env variable")
	}

	if os.Getenv("MYSQL_PASSWORD") != "" {
		c.MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	} else {
		log.Fatal("Couldn't read MYSQL_PASSWORD env variable")
	}

	if os.Getenv("MYSQL_HOST") != "" {
		c.MYSQL_HOST = os.Getenv("MYSQL_HOST")
	} else {
		log.Fatal("Couldn't read MYSQL_HOST env variable")
	}

	if os.Getenv("MYSQL_PORT") != "" {
		c.MYSQL_PORT = os.Getenv("MYSQL_PORT")
	} else {
		log.Fatal("Couldn't read MYSQL_PORT env variable")
	}

	if os.Getenv("MYSQL_DB_NAME") != "" {
		c.MYSQL_DB_NAME = os.Getenv("MYSQL_DB_NAME")
	} else {
		log.Fatal("Couldn't read MYSQL_DB_NAME env variable")
	}
}

// form a connection string
func (c *MysqlConnect) getConnectionString() string {

	c.readEnv()
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		c.MYSQL_USER,
		c.MYSQL_PASSWORD,
		c.MYSQL_HOST,
		c.MYSQL_PORT,
		c.MYSQL_DB_NAME)
	return connString
}

// connect to Mysql
func (mc *MysqlConnect) SetupConnection() {
	var err error
	connString := mc.getConnectionString()
	fmt.Println(mc)
	fmt.Println(*mc)
	log.Print(connString)

	mc.db, err = gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("MYSQL CONNECTED")
	}
	fmt.Println(mc)
}

func (c *MysqlConnect) GetDBInstance() *gorm.DB {
	return c.db
}
