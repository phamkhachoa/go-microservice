package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-ecommerce-backend-api/global"
	"go.uber.org/zap"
	"strconv"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Password: r.Password,
		DB:       0,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error", zap.Error(err))
	}

	global.Logger.Info("Redis initialization success")
	global.Rdb = rdb
}
