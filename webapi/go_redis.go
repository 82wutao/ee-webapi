package webapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string
	Port     int16
	Password string
	DB       int
}

func getRedisConfig() *RedisConfig {
	cfg := RedisConfig{
		Addr:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
	}
	return &cfg
}

var cli *redis.Client

func connect2redisServer(cfg *RedisConfig, ctx context.Context) {
	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	for running := true; running; running = false {
		if cli == nil {
			break
		}
		if err := cli.Close(); err != nil {
			log.Printf("close redis cli error %+v\n", err)
		}
	}
	cli = redis.NewClient(&opt)
	cli.Conn(ctx)
	log.Printf("debug : redis cli %+v\n", cli)
}
func closeRedisConnection() {
	if cli == nil {
		return
	}
	cli.Close()
	cli = nil
}

func getRedisCli() interface{} {
	return cli
}
func redisTryAcquire(ctx context.Context, k string, v interface{}, expire time.Duration) (bool, error) {
	boolCmd := cli.Conn(context.Background()).SetNX(ctx, k, v, expire)
	return boolCmd.Val(), boolCmd.Err()
}
