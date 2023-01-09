package main

import (
	"fmt"
	"os"

	"github.com/SzymekN/Car-rental-app/pkg/controller"
	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
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
	svr.Logger.SetupKafka()
	fmt.Println(svr.Logger.KafkaLogger)
	controller.SetupRouter(svr)
	fmt.Println(svr.EchoServ)

	// Close() przyjmuje teraz interfejs bazy danych - polimorfizm
	defer storage.Close(&svr.MysqlConn, svr.Logger)
	defer storage.Close(&svr.RedisConn, svr.Logger)

	// start server at port=API_PORT
	svr.EchoServ.Logger.Fatal(svr.EchoServ.Start(":" + os.Getenv("API_PORT")))
}
