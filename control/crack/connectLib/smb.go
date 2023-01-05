package connectLib

import (
	"cscan/config"
	"github.com/stacktitan/smb/smb"
	"log"
)

// 连接 smb 协议
func (crack *IpCrack) ConnectSmb() {
	Host, Port, Username, Password := crack.Ip, crack.Port, crack.UserName, crack.Password
	options := smb.Options{
		Host:        Host,
		Port:        Port,
		User:        Username,
		Password:    Password,
		Domain:      "",
		Workstation: "",
	}

	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			crack.CrackStatus = true
		}
	} else {
		if config.Debug {
			log.Println(err)
		}
	}
}
