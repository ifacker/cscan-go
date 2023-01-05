package scan

import (
	"cscan/config"
	"fmt"
	"testing"
)

func TestWebScan(t *testing.T) {
	//config.LogConfigInit()
	//config.Debug = true
	ipOption := &config.IpOption{
		Ip:             "www.baidu.com",
		Port:           80,
		PortOpenStatus: true,
		WebInfo:        config.WebInfo{},
	}
	WebScan(ipOption)
	fmt.Println(ipOption)
}

func TestWebScans(t *testing.T) {
	ipOptions := &config.IpOptions{
		IpOption: []*config.IpOption{
			{
				Ip:             "www.baidu.com",
				Port:           443,
				PortOpenStatus: true,
				WebInfo:        config.WebInfo{},
			},
			{
				Ip:             "www.weibo.com",
				Port:           80,
				PortOpenStatus: true,
				WebInfo:        config.WebInfo{},
			},
		},
	}
	WebScans(ipOptions)
}
