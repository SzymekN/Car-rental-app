package main

import (
	"os"

	"github.com/SzymekN/CRUD/pkg/controller"
	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/SzymekN/CRUD/pkg/seeder"
	"github.com/SzymekN/CRUD/pkg/storage"
)

func main() {

	e := controller.SetupRouter()
	storage.SetupPostgresConnection()
	// storage.SetupCassandraConnection()
	storage.SetupRedisConnection()
	producer.SetupKafka()

	// defer storage.CloseAll()

	seeder.CreateAndSeed()
	// go grpc.CreateGRPCServer()
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))

}
