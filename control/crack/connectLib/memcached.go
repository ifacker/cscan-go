package connectLib

import (
	"cscan/config"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

// 连接 memcached 协议
func (crack *IpCrack) ConnectMemcached() {
	connectData := fmt.Sprintf("%s:%d", crack.Ip, crack.Port)
	mc := memcache.New(connectData)
	err := mc.Ping()
	if err == nil {
		crack.CrackStatus = true
	} else {
		if config.Debug {
			log.Println(err)
		}
	}
}
