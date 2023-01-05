package connectLib

import (
	"cscan/config"
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
)

// 连接 ftp 协议
func (crack *IpCrack) ConnectFtp() {
	Host, Port, Username, Password := crack.Ip, crack.Port, crack.UserName, crack.Password
	conn, err := ftp.DialTimeout(fmt.Sprintf("%s:%d", Host, Port), config.TimeOut)
	if err == nil {
		err = conn.Login(Username, Password)
		defer conn.Logout()
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
