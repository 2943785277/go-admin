package bootstrap

import (
	"context"
	"fmt"
	"go/global"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.ConfigYml.GetString("redis.host") + ":" + global.ConfigYml.GetString("redis.port"),
		Password: global.ConfigYml.GetString("redis.password"), // no password set
		DB:       global.ConfigYml.GetInt("redis.db"),          // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis连接失败", err)
		// global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	fmt.Println("redis连接success", err)
	errs := client.Set(context.Background(), "key", 200, time.Minute*1).Err()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(errs)
	return client
}
