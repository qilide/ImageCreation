package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type SearchApiRedis struct {
}

// SetApiInfo 设置Api信息
func (sar SearchApiRedis) SetApiInfo() ([]string, error) {
	// 获取当前时间
	currentTime := time.Now()
	// 获取当前小时的时间戳
	currentHour := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), 0, 0, 0, currentTime.Location()).Unix()
	// 设置api信息
	apiKeys := viper.GetStringSlice("apiKeys")

	var useApiKeys []string
	for _, api := range apiKeys {
		// 获取 Redis 中对应 API 密钥的哈希表信息
		apiInfo, err := rdb.HGetAll(api).Result()
		if err != nil && err != redis.Nil {
			return useApiKeys, err
		}
		// 检查是否存在该 API 密钥的信息，若不存在则初始化
		if len(apiInfo) == 0 {
			// 更新 Redis 中对应 API 密钥的哈希表信息
			apiInfo = map[string]string{
				"count":     "0",
				"timestamp": fmt.Sprintf("%d", currentHour),
			}
		}
		// 获取当前存储的时间戳
		storedTimestamp, _ := strconv.ParseInt(apiInfo["timestamp"], 10, 64)
		// 如果当前时间与存储的时间戳在不同的小时内，重置计数器为0，并更新时间戳
		if storedTimestamp != currentHour {
			apiInfo["count"] = "0"
			apiInfo["timestamp"] = fmt.Sprintf("%d", currentHour)
		}
		// 检查计数器是否达到限制
		currentCount, _ := strconv.Atoi(apiInfo["count"])
		if currentCount >= 50 {
			continue
		}
		// 更新 Redis 中对应 API 密钥的哈希表信息
		err = rdb.HSet(api, "count", apiInfo["count"]).Err()
		err = rdb.HSet(api, "timestamp", apiInfo["timestamp"]).Err()
		if err != nil {
			return useApiKeys, err
		}
		useApiKeys = append(useApiKeys, api)
	}
	return useApiKeys, nil
}

// UpdateApiUseCount 更新Api使用次数+1
func (sar SearchApiRedis) UpdateApiUseCount(key string) error {
	// 获取 Redis 中对应 API 密钥的哈希表信息
	apiInfo, err := rdb.HGetAll(key).Result()
	if err != nil {
		return err
	}
	currentCount, _ := strconv.Atoi(apiInfo["count"])
	err = rdb.HSet(key, "count", fmt.Sprintf("%d", currentCount+1)).Err()
	return err
}
