package redis

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string
	Port     int16
	Password string
	DB       int
}

//	cfg := RedisConfig{
//		Addr:     "localhost",
//		Port:     6379,
//		Password: "",
//		DB:       0,
//	}
type RedisKey string

var cli *redis.Client

func _defer_close_conn(conn *redis.Conn) {
	conn.Close()
}
func RedisInitConnection(cfg *RedisConfig, ctx context.Context) bool {
	if cfg == nil {
		panic("")
	}
	for running := true; running; running = false {
		if cli == nil {
			break
		}
		if err := cli.Close(); err != nil {
			log.Printf("close redis cli error %+v\n", err)
		}
	}

	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	cli = redis.NewClient(&opt)
	return cli.Conn(ctx) != nil
}

func RedisMakeKey(key string, prefix ...string) RedisKey {
	var ps []string
	ps = append(ps, prefix...)

	ret := fmt.Sprintf("%s:%s", strings.Join(ps, ":"), key)
	return RedisKey(ret)
}
func RedisSet(k RedisKey, v string, expire time.Duration, ctx context.Context) error {
	conn := cli.Conn(ctx)
	defer _defer_close_conn(conn)

	sttsCmd := conn.Set(ctx, string(k), v, expire)
	return sttsCmd.Err()
}
func RedisGet(k RedisKey, ctx context.Context) (string, error) {
	conn := cli.Conn(ctx)
	defer _defer_close_conn(conn)

	strCmd := conn.Get(ctx, string(k))
	return strCmd.Val(), strCmd.Err()
}
func RedisGetAndDel(k RedisKey, ctx context.Context) (string, error) {
	conn := cli.Conn(ctx)
	defer _defer_close_conn(conn)

	strCmd := conn.GetDel(ctx, string(k))
	return strCmd.Val(), strCmd.Err()
}
func RedisDel(k RedisKey, ctx context.Context) (bool, error) {
	conn := cli.Conn(ctx)
	defer _defer_close_conn(conn)

	intCmd := conn.Del(ctx, string(k))
	return intCmd.Val() == 1, intCmd.Err()
}
