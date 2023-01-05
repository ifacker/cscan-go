package connectLib

import (
	"cscan/config"
	"database/sql"
	"fmt"
	"log"
)

// 连接 mssql 协议
func (crack *IpCrack) ConnectMssql() {
	Host, Port, Username, Password := crack.Ip, crack.Port, crack.UserName, crack.Password
	dataSourceName := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;encrypt=disable;timeout=%v", Host, Username, Password, Port, config.TimeOut)
	db, err := sql.Open("mssql", dataSourceName)
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
