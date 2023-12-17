package redis

import (
	"time"
)

type TokenRedis struct {
}

func (tr TokenRedis) GetToken(username string) interface{} {
	token, _ := rdb.Do("Get", username).Result()
	return token
}

func (tr TokenRedis) SetToken(username string, token string) {
	rdb.Set(username, token, time.Minute*30)
}

func (tr TokenRedis) DelToken(username string) {
	rdb.Del(username)
}
