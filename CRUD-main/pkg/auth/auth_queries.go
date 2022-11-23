package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/SzymekN/CRUD/pkg/storage"
)

var ctx = context.Background()

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
