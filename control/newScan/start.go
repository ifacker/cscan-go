package newScan

import (
	"cscan/config"
	"cscan/util/file"
	"cscan/util/identify"
	"cscan/util/view"
	"fmt"
	"github.com/ifacker/myutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// 加载输入的参数
func loadConfig() *config.IpOptions {
	ipOptions := &config.IpOptions{}

	// 端口识别代码块
	var ports []int

	// 添加端口的内部通用方法
	dfPort := func() {

		tmpPorts := ""
		if config.DefaultPorts != "" {
			tmpPorts = config.DefaultPorts
		} else if config.ForbidPorts != "" {
			tmpPorts = config.ForbidPorts
		}

		// 识别分隔符关键字符的内部方法
		identifyKeywords := func(keyword string) {
			portsStr := strings.Split(strings.TrimSpace(tmpPorts), keyword)
			for _, s := range portsStr {
				s = strings.TrimSpace(s)
				// 判断 s 里是否含有 "-"
				if strings.Contains(s, "-") {
					stmps := strings.Split(s, "-")
					min, err := strconv.Atoi(stmps[0])
					max, err := strconv.Atoi(stmps[1])
					if err != nil {
						if config.Debug {
							log.Println(err)
							view.PrintlnError("输入的 port 格式有误，请重试！案例：1-65535")
						}
						os.Exit(1)
					}
					for i := min; i <= max; i++ {
						ports = append(ports, i)
					}
				} else {
					port, err := strconv.Atoi(s)
					if err != nil {
						if config.Debug {
							log.Println(err)
							view.PrintlnError("输入的 port 格式有误，请重试！")
						}
						os.Exit(1)
					}
					ports = append(ports, port)
				}
			}
		}

		// 识别分隔符
		if strings.Contains(tmpPorts, ",") {
			identifyKeywords(",")
		} else if strings.Contains(tmpPorts, ";") {
			identifyKeywords(";")
		} else {
			identifyKeywords(" ")
		}
	}

	if config.DefaultPorts != "" {
		// 指定端口，加默认端口扫描

		dfPort()

		// 添加默认端口
		ports = append(ports, config.Ports...)

		// 去重函数
		uniq := func(s []int) []int {
			m := make(map[int]interface{}) // 用来记录元素
			for _, v := range s {
				m[v] = nil
			}

			ret := []int{} // 结果slice
			for k := range m {
				ret = append(ret, k)
			}
			return ret
		}
		// 对 ports 进行去重
		ports = uniq(ports)

		// 测试代码：展示所有要扫描的端口
		//fmt.Println(ports)

	} else if config.ForbidPorts != "" {
		// 指定端口，并禁用默认端口扫描

		dfPort()

		// 测试代码：展示所有要扫描的端口
		fmt.Println(ports)
	} else {
		// 未输入端口参数，使用默认端口
		ports = config.Ports
	}

	ipOptions.Ports = ports

	// IP识别代码块
	//var ips []string
	var ips = make(myutil.Set)

	// -l 识别 IP 文件
	if config.IpFilePath != "" {
		//c := make(chan bool)
		textb := file.ReadFile(config.IpFilePath)
		texts := strings.Split(string(textb), "\n")
		//wg := &sync.WaitGroup{}
		//wg.Add(1)
		// 显示加载界面
		//go func(c chan bool) {
		//	color.C256(226).Print("  Loading")
		//out:
		//	for {
		//		select {
		//		case <-c:
		//			break out
		//		default:
		//			color.C256(226).Print(".")
		//			time.Sleep(1 * time.Second)
		//		}
		//	}
		//	color.C256(226).Println("OK!\n")
		//	wg.Done()
		//}(c)
		for _, text := range texts {
			resultDomain, err := identify.DomainRegex(text)
			if err != nil && config.Debug {
				log.Println(err)
			}
			ips.AddAll(resultDomain)
			resultIp, err := identify.IpRange(text)
			if err != nil && config.Debug {
				log.Println(err)
			}
			//ips = append(ips, resultIp...)
			ips.AddAll(resultIp)
		}
		if len(ips) <= 0 {
			os.Exit(1)
		}
		//c <- true
		//wg.Wait()
	}

	// -i 识别 IP、IP段或IP区间
	// 如：192.168.1.1 | 192.168.1.1/24 | 192.168.1.1-20
	if config.IpInfo != "" {

		resultDomain, err := identify.DomainRegex(config.IpInfo)
		if err != nil && config.Debug {
			log.Println(err)
		}
		//ips = append(ips, resultDomain...)
		ips.AddAll(resultDomain)

		resultIp, err := identify.IpRange(config.IpInfo)
		if err != nil && config.Debug {
			log.Println(err)
		}
		//ips = append(ips, resultIp...)
		ips.AddAll(resultIp)

		if len(ips) <= 0 {
			os.Exit(1)
		}
	}
	var ipsTmp []string
	for ip, _ := range ips {
		ipsTmp = append(ipsTmp, ip)
	}
	ipOptions.Ips = ipsTmp
	return ipOptions
}

// 初始化过滤输出端口
func filterInit() {

	if config.Filter == "" {
		return
	}

	// 识别分隔符关键字符的内部方法
	identifyKeywords := func(keyword string) {
		filterStatusCode := strings.Split(strings.TrimSpace(config.Filter), keyword)
		for _, s := range filterStatusCode {
			s = strings.TrimSpace(s)
			// 判断 s 里是否含有 "-"
			if strings.Contains(s, "-") {
				stmps := strings.Split(s, "-")
				min, err := strconv.Atoi(stmps[0])
				max, err := strconv.Atoi(stmps[1])
				if err != nil {
					if config.Debug {
						log.Println(err)
						view.PrintlnError("输入的 filter 格式有误，请重试！案例：300-500")
					}
					os.Exit(1)
				}
				for i := min; i <= max; i++ {
					config.Filters = append(config.Filters, i)
				}
			} else {
				port, err := strconv.Atoi(s)
				if err != nil {
					if config.Debug {
						log.Println(err)
						view.PrintlnError("输入的 filter 格式有误，请重试！")
					}
					os.Exit(1)
				}
				config.Filters = append(config.Filters, port)
			}
		}
	}

	if strings.Contains(config.Filter, ",") {
		identifyKeywords(",")
	} else if strings.Contains(config.Filter, ";") {
		identifyKeywords(";")
	} else {
		identifyKeywords(" ")
	}
}

func NewStartScans() {
	ipOptions := loadConfig()
	filterInit()
	// 对 IP 进行端口扫描，http检测，以及暴力破解
	Scans(ipOptions)
}
