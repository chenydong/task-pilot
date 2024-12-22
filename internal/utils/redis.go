package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func InitRedis(addr, pwd string, db int) *redis.Client {

	rd_cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})

	_, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	_, err := rd_cli.Ping().Result()
	if err != nil {
		panic("连接redis失败, error=" + err.Error())
	} else {
		fmt.Println("连接redis成功")
	}

	return rd_cli
}
