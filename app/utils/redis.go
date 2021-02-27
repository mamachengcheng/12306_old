package utils

import (
	"context"
	"github.com/go-basic/uuid"
	"github.com/go-redis/redis/v8"
	"github.com/mamachengcheng/12306/app/static"
	"gopkg.in/ini.v1"
	"time"
)

var (
	RedisDB    *redis.Client
	RedisDBCtx context.Context
)

func init() {
	cfg, _ := ini.Load(static.ConfFilePath)

	redisCfg := cfg.Section("redis")
	address := redisCfg.Key("host").String() + ":" + redisCfg.Key("port").String()
	db, _ := redisCfg.Key("db").Int()

	RedisDB = redis.NewClient(&redis.Options{
		Addr: address,
		DB:   db,
	})
	RedisDBCtx = context.Background()
}

func AcquireLockWithTimeout(lockName string) bool {
	identifier := uuid.New()
	lockName = "lock: " + lockName

	if res, err := RedisDB.SetNX(RedisDBCtx, lockName, identifier, 30*time.Minute).Result(); err == nil && res {
		return true
	}

	return false
}

func ReleaseLock(lockName string) bool {
	lockName = "lock: " + lockName

	txf := func(tx *redis.Tx) error {
		_, err := RedisDB.Get(RedisDBCtx, lockName).Result()

		if err != redis.Nil && err != nil {
			return err
		}

		_, err = tx.TxPipelined(RedisDBCtx, func(pipe redis.Pipeliner) error {
			pipe.Del(RedisDBCtx, lockName)
			return nil
		})

		return err
	}

	err := RedisDB.Watch(RedisDBCtx, txf, lockName)

	if err == nil {
		return true
	}

	return false
}

func GetLockTTL(lockName string) time.Duration {
	lockName = "lock: " + lockName
	ttl := RedisDB.TTL(RedisDBCtx, lockName).Val()
	return ttl
}