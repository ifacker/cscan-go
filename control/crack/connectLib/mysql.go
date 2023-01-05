package connectLib

import (
	"cscan/config"
	"database/sql"
	"fmt"
	"log"
)

// 连接 mysql 协议
func (crack *IpCrack) ConnectMysql() {
	Host, Port, Username, Password := crack.Ip, crack.Port, crack.UserName, crack.Password
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=utf8&timeout=%v", Username, Password, Host, Port, config.TimeOut)
	db, err := sql.Open("mysql", dataSourceName)
	if err == nil {
		db.SetConnMaxLifetime(config.TimeOut)
		db.SetConnMaxIdleTime(config.TimeOut)
		db.SetMaxIdleConns(0)
		defer db.Close()
		err = db.Ping()
		if err == nil {
			crack.CrackStatus = true
		} else {
			if config.Debug {
				log.Println(err)
			}
		}
	} else {
		if config.Debug {
			log.Println(err)
		}
	}
}
