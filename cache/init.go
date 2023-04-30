package cache

import (
	"imgo/pkg/util"

	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
)

var (
	redisDBName   int
	redisAddr     string
	redisPassword string
	RedisClient   *redis.Client
)

func init() {
	file, err := ini.Load("conf/conf.ini")
	if err != nil {
		util.LogInstance.Info("load conf.ini failed", err)
	}
	loadRedis(file)
}

func loadRedis(file *ini.File) {
	s := file.Section("redis")
	redisDBName, _ = s.Key("redisDBName").Int()
	redisAddr = s.Key("redisAddr").String()
	redisPassword = s.Key("redisPassword").String()
}

// 连接redis数据库
func RedisInit() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDBName,
	})

	RedisClient = rdb
}
