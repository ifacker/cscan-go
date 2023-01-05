package scan

import (
	"cscan/config"
	"cscan/control/output"
	"cscan/util/view"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
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

// 批量扫描端口
func PortScans(ipOptions *config.IpOptions) {
	var ThreadMaxChan = make(chan int, config.ThreadMax)
	wg := &sync.WaitGroup{}
	for _, port := range ipOptions.Ports {
		for _, ip := range ipOptions.Ips {
			wg.Add(1)
			ThreadMaxChan <- 1
			go func(ip string, port int) {
				ipOption := config.IpOption{Ip: ip, Port: port}
				PortScan(&ipOption)
				ipOptions.IpOption = append(ipOptions.IpOption, &ipOption)

				// 打印输出
				if ipOption.PortOpenStatus {
					view.PrintlnSuccess(fmt.Sprintf("%s:%d --> open", ipOption.Ip, ipOption.Port))
				} else if config.ViewAll {
					view.PrintlnFailed(fmt.Sprintf("%s:%d --> close", ipOption.Ip, ipOption.Port))
				}

				// 导出
				if config.OutPutType != "" {
					if strings.Contains(config.OutPutType, ".csv") {
						path := config.OutPutType[:strings.Index(config.OutPutType, ".")] + "_portScan.csv"
						output.OutputFile(path, &ipOption, config.MODE_PORT)
					} else {
						output.OutputFile(config.OutPutType, &ipOption, config.MODE_PORT)
					}
				}

				// 测试代码
				//time.Sleep(1 * time.Second)
				//fmt.Printf("%s:%d --> %v\n", ipOption.Ip, ipOption.Port, ipOption.PortOpenStatus)

				wg.Done()
				<-ThreadMaxChan
			}(ip, port)
		}
	}
	close(ThreadMaxChan)
	wg.Wait()
}
