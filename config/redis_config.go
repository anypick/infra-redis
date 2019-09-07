package config

/**
定义Redis的配置信息，映射yaml文件中的配置
*/

import (
	"errors"
	"github.com/anypick/infra/base/props/container"
	"github.com/anypick/infra/utils/common"
	"reflect"
	"time"
)

const (
	ReplicationPrefix = "redis"
	SentinelPrefix    = "sentinel"
	ClusterPrefix     = "redisCluster"
)

func init() {
	container.RegisterYamContainer(&RedisClusterConfig{Prefix: ClusterPrefix})
	container.RegisterYamContainer(&Redis{Prefix: ReplicationPrefix})
	container.RegisterYamContainer(&RedisSentinelConfig{Prefix: SentinelPrefix})
}

type Redis struct {
	Prefix      string
	RedisConfig []RedisConfig `yaml:"redisConfig,flow"`
}

func (r *Redis) ConfigAdd(config map[interface{}]interface{}) {
	tmpConfigs := make([]RedisConfig, 0)
	redisConfigs := config["redisConfig"].([]map[interface{}]interface{})
	for _, redisConfig := range redisConfigs {
		tmpConfig := RedisConfig{}
		tmpConfig.Addr = redisConfig["addr"].(string)
		tmpConfig.Password = redisConfig["passsword"].(string)
		tmpConfig.DB = redisConfig["db"].(int)
		tmpConfig.MaxRetries = redisConfig["maxRetries"].(int)
		tmpConfig.PoolSize = redisConfig["poolSize"].(int)
		tmpConfig.MinIdleConns = redisConfig["minIdleConns"].(int)
		tmpConfig.MaxConnAge = time.Duration(redisConfig["maxConnAge"].(int))
		tmpConfig.ReadOnly = redisConfig["readOnly"].(bool)
		tmpConfigs = append(tmpConfigs, tmpConfig)
	}
	r.RedisConfig = tmpConfigs
}

// 定义Redis需要的配置
type RedisConfig struct {
	Prefix       string
	Addr         string        `yaml:"addr,omitempty"`
	Password     string        `yaml:"password,omitempty"`
	DB           int           `yaml:"db,omitempty"`
	MaxRetries   int           `yaml:"maxRetries,omitempty"`
	PoolSize     int           `yaml:"poolSize,omitempty"`
	MinIdleConns int           `yaml:"minIdleConns,omitempty"`
	MaxConnAge   time.Duration `yaml:"maxConnAge,omitempty"`
	ReadOnly     bool          `yaml:"readOnly,omitempty"`
}

// key为字段名称，大小写要一致
func (t RedisConfig) GetString(key string) (string, error) {
	valueOf := reflect.ValueOf(t)
	value := valueOf.FieldByName(key).Interface().(string)
	if common.StrIsBlank(value) {
		return "", errors.New("please setting" + key)
	}
	return value, nil
}

func (t RedisConfig) GetInt(key string) (int, error) {
	valueOf := reflect.ValueOf(t)
	value := valueOf.FieldByName(key).Interface().(int)
	if value == 0 {
		return value, errors.New("please setting" + key)
	}
	return value, nil
}

func (t RedisConfig) GetBool(key string) (bool, error) {
	valueOf := reflect.ValueOf(t)
	value := valueOf.FieldByName(key).Interface().(bool)
	return value, nil
}

func (t RedisConfig) GetTime(key string) (time.Duration, error) {
	valueOf := reflect.ValueOf(t)
	value := valueOf.FieldByName(key).Interface().(int64)
	return time.Second * time.Duration(value), nil
}

// redis哨兵配置
type RedisSentinelConfig struct {
	Prefix     string
	Addrs      []string `yaml:"addrs"`
	MasterName string   `yaml:"masterName"`
	Password   string   `yaml:"password"`
}

func (r *RedisSentinelConfig) ConfigAdd(config map[interface{}]interface{}) {
	r.Addrs = config["addrs"].([]string)
	r.MasterName = config["readOnly"].(string)
	r.Password = config["poolSize"].(string)
}

// redis集群配置
type RedisClusterConfig struct {
	Prefix       string
	Addrs        []string `yaml:"addrs"`
	ReadOnly     bool     `yaml:"readOnly"`
	PoolSize     int      `yaml:"poolSize"`
	MinIdleConns int      `yaml:"minIdleConns"`
}

func (r *RedisClusterConfig) ConfigAdd(config map[interface{}]interface{}) {
	r.Addrs = config["addrs"].([]string)
	r.ReadOnly = config["readOnly"].(bool)
	r.PoolSize = config["poolSize"].(int)
	r.MinIdleConns = config["minIdleConns"].(int)
}
