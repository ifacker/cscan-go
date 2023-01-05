package connectLib

import (
	"cscan/config"
	"database/sql"
	"fmt"
	"log"
)

// 连接 postgres 协议
func (crack *IpCrack) ConnectPostgres() {
	Host, Port, Username, Password := crack.Ip, crack.Port, crack.UserName, crack.Password
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", Username, Password, Host, Port, "postgres", "disable")
	db, err := sql.Open("postgres", dataSourceName)
	if err == nil {
		db.SetConnMaxLifetime(config.TimeOut)
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
