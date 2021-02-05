package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mamachengcheng/12306/app/static"
	"gopkg.in/ini.v1"
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
