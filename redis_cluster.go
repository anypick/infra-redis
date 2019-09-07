package baseredis

import (
	"github.com/anypick/infra"
	redisConfig "github.com/anypick/infra-redis/config"
	"github.com/go-redis/redis"
)

var redisCluster *redis.ClusterClient

func GetRedisCluster() *redis.ClusterClient {
	return redisCluster
}

type RedisClusterStarter struct {
	infra.BaseStarter
}

func (r *RedisClusterStarter) Setup(ctx infra.StarterContext) {
	config := ctx.Yaml().OtherConfig[redisConfig.ClusterPrefix].(*redisConfig.RedisClusterConfig)
	redisCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        config.Addrs,
		ReadOnly:     config.ReadOnly,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})
}
