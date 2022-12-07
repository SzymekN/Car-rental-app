package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

type RedisConnect struct {
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	rdb            *redis.Client
}

// var RDB *redis.Client

func (rc *RedisConnect) GetRDB() *redis.Client {
	return rc.rdb
}
func (rc *RedisConnect) readEnv() {

	if os.Getenv("REDIS_HOST") != "" {
		rc.REDIS_HOST = os.Getenv("REDIS_HOST")
	} else {
		log.Fatal("Couldn't read REDIS_HOST env variable")
	}

	if os.Getenv("REDIS_PORT") != "" {
		rc.REDIS_PORT = os.Getenv("REDIS_PORT")
	} else {
		log.Fatal("Couldn't read REDIS_PORT env variable")
	}

	// TODO: ustawić w docker compose hasło redisa
	if os.Getenv("REDIS_PASSWORD") != "" {
		// rc.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	} else {
		log.Fatal("Couldn't read REDIS_PASSWORD env variable")
	}
}

func (rc *RedisConnect) SetupConnection() {
	rc.readEnv()
	rc.rdb = redis.NewClient(&redis.Options{
		Addr:     rc.REDIS_HOST + ":" + rc.REDIS_PORT,
		Password: rc.REDIS_PASSWORD,
		DB:       0,
	})

	// ping db to check if connection is established
	pong, err := rc.rdb.Ping(context.Background()).Result()
	fmt.Println(rc)
	fmt.Println(pong, err)

}
