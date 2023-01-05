package connectLib

import (
	"cscan/config"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

// 连接 mongodb 协议
func (crack *IpCrack) ConnectMongodb() {
	url := ""
	if crack.UserName == "" && crack.Password == "" {
		url = fmt.Sprintf("mongodb://%s:%d/", crack.Ip, crack.Port)
	} else {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%d/", crack.UserName, crack.Password, crack.Ip, crack.Port)
	}
	session, err := mgo.DialWithTimeout(url, config.TimeOut)
	//defer session.Close()
	if err != nil {
		if config.Debug {
			log.Println(err)
		}
	}
	err = session.Ping()
	if err == nil {
		session.Close()
		crack.CrackStatus = true
	} else {
		if config.Debug {
			log.Println(err)
		}
	}
}
