package newScan

import (
	"cscan/config"
	"cscan/control/output"
	"cscan/util/view"
	"fmt"
	"github.com/gookit/color"
	"strings"
	"sync"
)

// 批量扫描端口
func scan(ipOptions *config.IpOptions) {
	var ThreadMaxChan = make(chan int, config.ThreadMax)
	wg := &sync.WaitGroup{}
	for _, port := range ipOptions.Ports {
		for _, ip := range ipOptions.Ips {
			wg.Add(1)
			ThreadMaxChan <- 1
			go func(ip string, port int) {
				ipOption := config.IpOption{Ip: ip, Port: port}
				PortScan(&ipOption)
				if ipOption.PortOpenStatus {
					ipOptions.IpOption = append(ipOptions.IpOption, &ipOption)
					// 对 web 进行扫描
					WebScan(&ipOption)
				}

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

// 对 IP 进行端口扫描，http检测，以及暴力破解
func Scans(ipOptions *config.IpOptions) {
	scan(ipOptions)

	// 扫描好后的资产进行打印
	color.C256(226).Printf("\n\n------ web 页面识别结果：------\n\n")
	for _, ipOption := range ipOptions.IpOption {
		webScanPrint(ipOption)
	}

	crackStart(ipOptions)

}
