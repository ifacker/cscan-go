package connectLib

import (
	"cscan/config"
	"database/sql"
	"fmt"
	"log"
)

// 连接 oracle 协议
func (crack *IpCrack) ConnectOracle() {
	Host, Port, Username, Password := crack.Ip, crack.Port, crack.UserName, crack.Password
	dataSourceName := fmt.Sprintf("oracle://%s:%s@%s:%d/orcl", Username, Password, Host, Port)
	db, err := sql.Open("oracle", dataSourceName)
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
