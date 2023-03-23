// Package redis 工具包
package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	u "gin/pkg/util"

	redis "github.com/go-redis/redis/v8"
)

// RedisClient Redis 服务
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用 db 1
var RedisCli *RedisClient

// Redis const

const (
	USER_STATISTC   string = "userStatistic:"
	NOTIFI_STATISTC string = "notifStatistic:"
)

func SetupRedis() {
	var (
		address, host, port, password string
		db                            int
	)

	host = u.GetEnv("REDIS_HOST")
	port = u.GetEnv("REDIS_PORT")
	password = u.GetEnv("REDIS_PASSWORD")
	db, _ = strconv.Atoi(u.GetEnv("REDIS_DATABASE"))

	address = fmt.Sprintf("%v:%v", host, port)
	// fmt.Println("redis address:", address)

	ConnectRedis(address, "", password, db)
}

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		RedisCli = NewClient(address, username, password, db)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, username string, password string, db int) *RedisClient {

	// 初始化自定的 RedisClient 实例
	rds := &RedisClient{}
	// 使用默认的 context
	rds.Context = context.Background()

	// 使用 redis 库里的 NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	err := rds.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return rds
}

// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Fatal(err)
		}
		return ""
	}
	return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
		}
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		return false
	}
	return true
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		return false
	}
	return true
}

// Increment 当参数只有 1 个时，为 key，其值增加 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值 int64 类型。
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			return false
		}
	default:
		return false
	}
	return true
}

// Decrement 当参数只有 1 个时，为 key，其值减去 1。
// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值 int64 类型。
func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			return false
		}
	default:
		return false
	}
	return true
}

func (rds RedisClient) SAdd(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.SAdd(rds.Context, key, value, expiration).Err(); err != nil {
		log.Fatal(err)

		return false
	}
	return true
}

func (rds RedisClient) RPush(key string, value interface{}) bool {
	if err := rds.Client.RPush(rds.Context, key, value).Err(); err != nil {
		log.Fatal(err)

		return false
	}
	return true
}

func (rds RedisClient) Incr(key string) bool {
	if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
		log.Fatal(err)

		return false
	}
	return true
}

func (rds RedisClient) HGet(key string, key2 string) string {
	result, err := rds.Client.HGet(rds.Context, key, key2).Result()
	if err != nil {
		return ""
	}
	return result
}

func (rds RedisClient) HSet(key string, key2 string, val interface{}) bool {
	if err := rds.Client.HSet(rds.Context, key, key2, val).Err(); err != nil {
		return false
	}
	return true
}

func (rds RedisClient) HSetNX(key string, key2 string, val interface{}) bool {
	if err := rds.Client.HSetNX(rds.Context, key, key2, val).Err(); err != nil {
		return false
	}
	return true
}

func (rds RedisClient) HExists(key string, key2 string) bool {
	_, err := rds.Client.HExists(rds.Context, key, key2).Result()
	if err != nil {
		if err != redis.Nil {
		}
		return false
	}
	return true
}

func (rds RedisClient) HIncrBy(key string, key2 string, val int64) bool {
	_, err := rds.Client.HIncrBy(rds.Context, key, key2, val).Result()
	if err != nil {
		return false
	}
	return true
}

func (rds RedisClient) LRange(key string, start int, end int) interface{} {
	result, err := rds.Client.LRange(rds.Context, key, int64(start), int64(end)).Result()
	if err != nil {
		return nil
	}
	return result
}
