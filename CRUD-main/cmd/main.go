package main

import (
	"os"

	"github.com/SzymekN/Car-rental-app/pkg/controller"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/seeder"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
)

func main() {

	e := controller.SetupRouter()
	// storage.SetupPostgresConnection()
	storage.SetupMysqlConnection()
	// storage.SetupCassandraConnection()
	storage.SetupRedisConnection()
	producer.SetupKafka()

	// defer storage.CloseAll()

	seeder.CreateAndSeed()
	// go grpc.CreateGRPCServer()
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))

}
