package redis

import "time"

type MailRedis struct {
}

func (tr MailRedis) GetMail(username string) interface{} {
	token, _ := rdb.Do("Get", username+"mail").Result()
	return token
}

func (tr MailRedis) SetMail(username string, code string) {
	rdb.Set(username+"mail", code, time.Minute*10)
}

func (tr MailRedis) DelMail(username string) {
	rdb.Del(username + "mail")
}
