package auth

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/go-redis/redis/v8"
)

type JWTQueryExecutor struct {
	Svr *server.Server
	Ctx context.Context
}

type JWTQueryExecutorInterface interface {
	GetToken(token string) (bool, error)
	SetToken(token string, expireTime time.Duration) error
	setSigningKey() (string, error)
	getSigningKey() (string, error)
	ProduceMessage(k, val string)
}

func (j JWTQueryExecutor) ProduceMessage(k, val string) {
	j.Svr.Logger.ProduceMessage(k, val)
}

// execute querry to Redis ti get the key needed for signing and validating jwt tokens
func (j JWTQueryExecutor) getSigningKey() (string, error) {
	rdb := j.Svr.GetRedisDB()
	logger := j.Svr.Logger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("REDIS Read  ")

	res, err := rdb.Get(j.Ctx, "key").Result()

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceLog()
	}()

	if err == redis.Nil {
		code := http.StatusNotFound
		msg := fmt.Sprintf("[ERROR]: key doesn't exist, HTTP: %v", code)
		logger.Log.Populate("err", msg, code, err)
		return "", err
	} else if err != nil {
		code := http.StatusInternalServerError
		msg := fmt.Sprintf("[ERROR]: unexpected error reading key, HTTP: %v", code)
		logger.Log.Populate("err", msg, code, err)
		return "", err
	}

	code := http.StatusOK
	msg := fmt.Sprintf("[INFO]: signing key retrieved, HTTP: %v", code)
	logger.Log.Populate("info", msg, code, err)
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
	logger := j.Svr.Logger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("REDIS Write  ")

	key := j.generateKey()
	err := rdb.Set(j.Ctx, "key", key, 0).Err()

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceLog()
	}()

	if err != nil {
		code := http.StatusInternalServerError
		msg := fmt.Sprintf("[ERROR]: Signing key sending failure , HTTP: %v", code)
		logger.Log.Populate("err", msg, code, err)
		return "", err
	}

	code := http.StatusOK
	msg := fmt.Sprintf("[INFO]: Signing key set {%s}, HTTP: %v", key, code)
	logger.Log.Populate("info", msg, code, err)
	return key, nil
}

// set black listed jwt token in the Redis
func (j JWTQueryExecutor) SetToken(token string, expireTime time.Duration) error {
	rdb := j.Svr.GetRedisDB()
	logger := j.Svr.Logger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("REDIS Write  ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceLog()
	}()

	err := rdb.Set(j.Ctx, token, "0", expireTime*time.Second).Err()
	if err != nil {
		code := http.StatusInternalServerError
		msg := fmt.Sprintf("[ERROR]: Sending token failure , HTTP: %v", code)
		logger.Log.Populate("err", msg, code, err)
		return err
	}

	code := http.StatusOK
	msg := fmt.Sprintf("[INFO]: Token sent {%s}, HTTP: %v", token, code)
	logger.Log.Populate("info", msg, code, err)
	go j.ProduceMessage("REDIS write", "Token:"+token+" set")
	return nil
}

// try to get jwt token from Redis
func (j JWTQueryExecutor) GetToken(token string) (bool, error) {
	rdb := j.Svr.GetRedisDB()
	logger := j.Svr.Logger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("REDIS Read  ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceLog()
	}()

	fmt.Println(rdb)
	_, err := rdb.Get(j.Ctx, token).Result()
	if err != nil {
		code := http.StatusInternalServerError
		msg := fmt.Sprintf("[ERROR]: Sending token failure , HTTP: %v", code)
		logger.Log.Populate("err", msg, code, err)
		return false, err
	}

	code := http.StatusOK
	msg := fmt.Sprintf("[INFO]: Token read {%s}, HTTP: %v", token, code)
	logger.Log.Populate("info", msg, code, err)
	return true, nil
}
