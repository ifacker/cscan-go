package connectLib

import (
	"cscan/config"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

// 连接 ssh 协议
func (crack *IpCrack) ConnectSSH() {
	clientconfig := &ssh.ClientConfig{
		User:            crack.UserName,
		Auth:            []ssh.AuthMethod{ssh.Password(crack.Password)},
		Timeout:         config.TimeOut,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	address := fmt.Sprintf("%s:%d", crack.Ip, crack.Port)
	sshClient, err := ssh.Dial("tcp", address, clientconfig)
	time.Sleep(1 * time.Second)
	if err == nil {
		sshClient.Close()
		crack.CrackStatus = true
	} else {
		if config.Debug {
			log.Println(err)
		}
	}
}
