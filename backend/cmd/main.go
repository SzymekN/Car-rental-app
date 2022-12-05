package main

import (
	"fmt"
	"os"

	"github.com/SzymekN/Car-rental-app/pkg/controller"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/server"
)

// pobieranie nazwy z headera albo body requesta i tworzenie odpowiedniej zmiennej na jej podstawie

// stwozryć obiekt nadzrędny kontroller - zawiera w sobie mniejsze kontrolery dla każdego z requestów
// funkcje z klas podrzędnych odwołują się do nadrzędnych

func main() {
	svr := &server.Server{}
	// svr.MysqlConn = *storage.New()
	svr.MysqlConn.SetupConnection()
	fmt.Println(svr.MysqlConn)
	svr.RedisConn.SetupConnection()
	fmt.Println(svr.RedisConn)
	controller.SetupRouter(svr)

	// Close() przyjmuje teraz interfejs bazy danych - polimorfizm
	// defer storage.CloseAll()
	producer.SetupKafka()
	fmt.Println(svr.E)

	// Drop old tables, create new and populate them - for test purposes
	// seeder.CreateAndSeed()

	// start server at port=API_PORT
	svr.E.Logger.Fatal(svr.E.Start(":" + os.Getenv("API_PORT")))
}
