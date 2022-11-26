package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
)

var ctx = context.Background()

// execute querry to Redis ti get the key needed for signing and validating jwt tokens
func getSigningKey() (string, error) {
	rdb := storage.GetRDB()

	res, err := rdb.Get(ctx, "key").Result()

	if err != nil {
		producer.ProduceMessage("REDIS read", "ERROR reading key:"+err.Error())
		fmt.Println("ERROR reading key:", err.Error())
		return "", err
	}

	return res, nil
}

// generate key for signing jwt tokens
func generateKey() string {
	//33 - 126 valid ascii characters
	var min int64 = 33  // '!'
	var max int64 = 126 // '~'
	len := 24
	key := make([]byte, len)
	for i := 0; i < len; i++ {
		key[i] = byte(rand.Int63n(max-min) + min)
	}

	return string(key)
}

// set signing key for intance, try reading it from Redis, if not exists generate new
func setSigningKey() (string, error) {
	rdb := storage.GetRDB()

	key := generateKey()
	err := rdb.Set(ctx, "key", key, 0).Err()

	if err != nil {
		producer.ProduceMessage("REDIS write", "ERROR writing key:"+err.Error())
		fmt.Println("ERROR writing key:", err.Error())
		return "", err

	}

	producer.ProduceMessage("REDIS write", "Key set:"+key)
	return key, nil
}

// set black listed jwt token in the Redis
func SetToken(token string, expireTime time.Duration) error {
	rdb := storage.GetRDB()

	err := rdb.Set(ctx, token, "0", expireTime*time.Second).Err()
	if err != nil {
		producer.ProduceMessage("REDIS write", "ERROR writing token:"+err.Error())
		return err
	}

	producer.ProduceMessage("REDIS write", "Token:"+token+" set")
	return nil
}

// try to get jwt token from Redis
func GetToken(token string) (bool, error) {

	rdb := storage.GetRDB()

	fmt.Println(rdb)
	_, err := rdb.Get(ctx, token).Result()
	if err != nil {
		producer.ProduceMessage("REDIS read", "ERROR reading token:"+token+", err: "+err.Error())
		return false, err
	}

	producer.ProduceMessage("REDIS write", "Token: "+token+" get")
	return true, nil
}
