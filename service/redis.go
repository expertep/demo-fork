package service

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
)

type RedisService interface {
	GetRedis(key string, obj interface{}) error
	SetRedis(key string, timeSet time.Duration, obj interface{}) error
	DelRedis(key string) error
	CloseRedis()
}
type redisInformation struct {
	rdb *redis.Client
}

func (info *redisInformation) DelRedis(key string) error {
	var ctx = context.Background()
	err := info.rdb.Del(ctx, key).Err()
	color.Green("del redis data::" + key + "\n")

	if err != nil {
		return err
	}
	return nil
}

func (info *redisInformation) GetRedis(key string, obj interface{}) error {
	var ctx = context.Background()
	// defer info.rdb.Close()

	byteObj, err := info.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// color.Red("Redis data does not exist::" + key + "\n")
		}
		return err
	}
	json.Unmarshal([]byte(byteObj), obj)
	return nil
}

/* func (info *redisInformation) ExistRedis(key string) isExist {
	var ctx = context.Background()
	// defer info.rdb.Close()

	byteObj, err := info.rdb.Exists(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			color.Red("Redis data does not exist::" + key + "\n")
		}
		return err
	}
	json.Unmarshal([]byte(byteObj), obj)
	return nil
} */

func (info *redisInformation) SetRedis(key string, timeSet time.Duration, obj interface{}) error {
	var ctx = context.Background()
	// defer info.rdb.Close()
	byteObj, _ := json.Marshal(obj)
	// var a = time.Duration(timeSet * int64(time.Hour))
	err := info.rdb.SetEX(ctx, key, byteObj, timeSet).Err()
	if err != nil {
		return err
	}
	return nil
}

func RedisConnect() RedisService {
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_URL"),
		Password:    "",      // no password set
		DB:          redisDB, // use default DB
		IdleTimeout: 1 * time.Minute,
	})
	redisService := &redisInformation{rdb: rdb}
	return redisService
}

func (info *redisInformation) CloseRedis() {
	info.rdb.Close()
}
