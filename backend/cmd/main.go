package main

import (
	"os"

	"github.com/SzymekN/Car-rental-app/pkg/controller"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
)

// pobieranie nazwy z headera albo body requesta i tworzenie odpowiedniej zmiennej na jej podstawie

// stwozryć obiekt nadzrędny kontroller - zawiera w sobie mniejsze kontrolery dla każdego z requestów
// funkcje z klas podrzędnych odwołują się do nadrzędnych

func main() {

	e := controller.SetupRouter()
	storage.SetupMysqlConnection()
	storage.SetupRedisConnection()
	producer.SetupKafka()

	// defer storage.CloseAll()

	// Drop old tables, create new and populate them - for test purposes
	// seeder.CreateAndSeed()

	// start server at port=API_PORT
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))

}
