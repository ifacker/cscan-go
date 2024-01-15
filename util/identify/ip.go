package identify

import (
	"cscan/util/regex2"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"net"
	"strconv"
	"strings"
	"time"
)

// IpRange 识别 IP 范围，如：192.168.1.1-20、192.168.1.1/24、192.168.1.1
func IpRange(srcIp string) ([]string, error) {

	// 确认 IP 格式
	results, err := regex2.Regexp2SimpleUse(srcIp, "(\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}-\\d{1,3})|(\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}/\\d{1,3})|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3})")
	if err != nil || len(results) <= 0 {
		return nil, err
	}
	if len(results) <= 0 {
		return nil, errors.New("输入的 IP 异常 或 未识别到 IP ！")
	}

	var destIps []string

	ips := strings.Split(srcIp, ",")
	for _, ip := range ips {

		ip = strings.TrimSpace(ip)

		if strings.Contains(ip, "-") {
			// 192.168.1.1-20

			pointIndex := strings.LastIndex(ip, ".")
			prefix := ip[:pointIndex+1]
			suffix := ip[pointIndex+1:]
			//fmt.Println(prefix, suffix)

			ns := strings.Split(suffix, "-")
			if len(ns) != 2 {
				return nil, errors.New("输入的 IP 范围格式错误！应该是：192.168.1.1-20")
			}
			min, err := strconv.Atoi(ns[0])
			if err != nil {
				return nil, err
			}
			max, err := strconv.Atoi(ns[1])
			if err != nil {
				return nil, err
			}
			for i := min; i <= max; i++ {
				destIps = append(destIps, fmt.Sprintf("%s%d", prefix, i))
			}
		} else if strings.Contains(ip, "/") {
			// 192.168.1.1/24

			inc := func(ip net.IP) {
				for j := len(ip) - 1; j >= 0; j-- {
					ip[j]++
					if ip[j] > 0 {
						break
					}
				}
			}

			iptmp, ipnet, err := net.ParseCIDR(ip) // 解析IP段
			if err != nil {
				color.C256(196).Printf("[!] IP: %s 存在异常，请及时处理！\n\n", ip)
				time.Sleep(1500 * time.Millisecond)
				return nil, errors.New(fmt.Sprintf("IP: %s 存在异常，请及时处理！", ip))
			}
			for iptmp := iptmp.Mask(ipnet.Mask); ipnet.Contains(iptmp); inc(iptmp) { // 通过掩码来遍历这个IP段
				destIps = append(destIps, fmt.Sprintf("%s", iptmp))
				//fmt.Println(ip)
			}

		} else {
			// 192.168.1.1
			if strings.Count(ip, ".") == 3 {
				destIps = append(destIps, ip)
			}
		}
	}
	return destIps, nil
}

func DomainRegex(srcDomain string) ([]string, error) {
	// 确认域名的格式
	results, err := regex2.Regexp2SimpleUse(srcDomain, "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?")
	if err != nil || len(results) <= 0 {
		return nil, err
	}
	if len(results) <= 0 {
		return nil, errors.New("输入的域名异常 或 未识别到域名 ！")
	}
	return results, nil
}
