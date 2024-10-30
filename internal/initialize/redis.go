package initialize

import (
	"context"
	"ecommerce/global"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // Số lượng kết nối tối đa
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error: %v", zap.Error(err))
	}

	fmt.Println("Init Redis is running")
	global.Rdb = rdb
	redusExample()
}

func redusExample() {
	if err := global.Rdb.Set(ctx, "ping", "pong", 0).Err(); err != nil {
		fmt.Println("Error redis setting: ", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "ping").Result()
	if err != nil {
		fmt.Println("Error redis setting: ", zap.Error(err))
	}

	global.Logger.Info("Value ping is:: ", zap.String("value", value))
}
