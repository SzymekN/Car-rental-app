package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/server"
)

type JWTQueryExecutor struct {
	Svr *server.Server
	Ctx context.Context
}

// execute querry to Redis ti get the key needed for signing and validating jwt tokens
func (j JWTQueryExecutor) getSigningKey() (string, error) {
	rdb := j.Svr.GetRedisDB()

	res, err := rdb.Get(j.Ctx, "key").Result()

	if err != nil {
		producer.ProduceMessage("REDIS read", "ERROR reading key:"+err.Error())
		fmt.Println("ERROR reading key:", err.Error())
		return "", err
	}

	return res, nil
}

// generate key for signing jwt tokens
func (j JWTQueryExecutor) generateKey() string {
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
func (j JWTQueryExecutor) setSigningKey() (string, error) {
	rdb := j.Svr.GetRedisDB()

	key := j.generateKey()
	err := rdb.Set(j.Ctx, "key", key, 0).Err()

	if err != nil {
		producer.ProduceMessage("REDIS write", "ERROR writing key:"+err.Error())
		fmt.Println("ERROR writing key:", err.Error())
		return "", err

	}

	producer.ProduceMessage("REDIS write", "Key set:"+key)
	return key, nil
}

// set black listed jwt token in the Redis
func (j JWTQueryExecutor) SetToken(token string, expireTime time.Duration) error {
	rdb := j.Svr.GetRedisDB()

	err := rdb.Set(j.Ctx, token, "0", expireTime*time.Second).Err()
	if err != nil {
		producer.ProduceMessage("REDIS write", "ERROR writing token:"+err.Error())
		return err
	}

	producer.ProduceMessage("REDIS write", "Token:"+token+" set")
	return nil
}

// try to get jwt token from Redis
func (j JWTQueryExecutor) GetToken(token string) (bool, error) {

	rdb := j.Svr.GetRedisDB()

	fmt.Println(rdb)
	_, err := rdb.Get(j.Ctx, token).Result()
	if err != nil {
		producer.ProduceMessage("REDIS read", "ERROR reading token:"+token+", err: "+err.Error())
		return false, err
	}

	producer.ProduceMessage("REDIS write", "Token: "+token+" get")
	return true, nil
}
