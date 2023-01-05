package connectLib

import (
	"cscan/config"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

// 连接 redis 协议
func (crack *IpCrack) ConnectRedis() {
	option := redis.Options{
		Addr:        fmt.Sprintf("%s:%d", crack.Ip, crack.Port),
		Password:    crack.Password,
		DB:          0,
		DialTimeout: config.TimeOut,
	}
	client := redis.NewClient(&option)
	defer client.Close()
	_, err := client.Ping().Result()
	if err == nil {
		crack.CrackStatus = true
	} else {
		if config.Debug {
			log.Println(err)
		}
	}
}
