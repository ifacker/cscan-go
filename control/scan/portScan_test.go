package scan

import (
	"cscan/config"
	"fmt"
	"testing"
)

func TestPortScan(t *testing.T) {
	ipOption := config.IpOption{Ip: "localhost", Port: 1080}
	PortScan(&ipOption)
	if ipOption.Status {
		fmt.Println("open")
	} else {
		fmt.Println("close")
	}
}

func TestPortScans(t *testing.T) {
	ipOptions := config.IpOptions{
		Ips:   []string{"www.baidu.com", "www.taobao.com"},
		Ports: []int{443, 80, 88},
	}
	config.ThreadMax = 10
	config.ViewAll = true
	PortScans(&ipOptions)
	fmt.Println("--------")
	for _, option := range ipOptions.IpOption {
		fmt.Printf("%s:%d --> %v\n", option.Ip, option.Port, option.Status)
	}
}
