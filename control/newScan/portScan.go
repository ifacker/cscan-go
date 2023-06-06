package newScan

import (
	"cscan/config"
	"fmt"
	"log"
	"net"
)

// 传入 ipOption，检查端口有没有开放
func PortScan(ipOption *config.IpOption) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ipOption.Ip, ipOption.Port), config.TimeOut)
	if err != nil && config.Debug {
		log.Println(err)
		ipOption.PortOpenStatus = false
	} else {
		if conn != nil {
			ipOption.PortOpenStatus = true
		} else {
			ipOption.PortOpenStatus = false
		}
	}
}
