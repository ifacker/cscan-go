package scan

import (
	"cscan/config"
	"cscan/control/http"
	"cscan/control/output"
	"cscan/util/view"
	"fmt"
	"log"
	http2 "net/http"
	"strings"
	"sync"

	"github.com/antchfx/htmlquery"
	"github.com/gookit/color"
	"golang.org/x/net/html"
)

// 获取关键信息 title
func getMeta(n *html.Node) (title string) {
	if n == nil {
		return ""
	}
	// 获取title
	titletmp := htmlquery.FindOne(n, "/html/head/title/text()")
	if titletmp == nil {
		return ""
	}
	title = htmlquery.InnerText(titletmp)
	return
}

func reqfunc(url string) *config.WebInfo {
	webinfo := &config.WebInfo{}
	client := http.NewClient()
	// 禁止 302 跳转
	client.CheckRedirect = func(req *http2.Request, via []*http2.Request) error {
		return http2.ErrUseLastResponse
	}
	req, err := http2.NewRequest("GET", url, nil)
	if err != nil && config.Debug {
		log.Println(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		if config.Debug {
			log.Println(err)
		}
		return nil
	}
	//defer resp.Body.Close()
	defer func() {
		if err = resp.Body.Close(); err != nil && config.Debug {
			log.Println(err)
		}
	}()

	webinfo.Code = resp.StatusCode
	webinfo.Len = resp.ContentLength
	webinfo.Url = url
	var doc *html.Node
	webinfo.Title = getMeta(doc)
	url302Jump := resp.Header.Get("Location")
	if url302Jump != "" {
		webinfo.Url302Jump = url + url302Jump
	}
	webinfo.Server = resp.Header.Get("Server")
	webinfo.XPoweredBy = resp.Header.Get("X-Powered-By")

	return webinfo
}

// 检测 web 服务
func WebScan(ipOption *config.IpOption) {
	urls := []string{
		fmt.Sprintf("https://%s:%d", ipOption.Ip, ipOption.Port),
		fmt.Sprintf("http://%s:%d", ipOption.Ip, ipOption.Port),
	}

	// 我这边 http 和 https 协议全部都做检测
	for _, url := range urls {
		webInfo := reqfunc(url)
		if webInfo != nil {
			ipOption.WebInfo = *webInfo
		}
	}

	// 备注：这里记得处理一下 http 和 https 两个协议，原版 cscan 中使用的方法是，如果端口不是 443，8443，就走http，否则就走 https

}

// 批量检测 web 服务
func WebScans(ipOptions *config.IpOptions) {
	var ThreadMaxChan = make(chan int, config.ThreadMax)
	wg := &sync.WaitGroup{}

	contains := func(s []int, e int) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}

	for _, option := range ipOptions.IpOption {
		if !option.PortOpenStatus {
			continue
		}
		wg.Add(1)
		ThreadMaxChan <- 1
		go func(ipOption *config.IpOption) {
			WebScan(ipOption)

			if ipOption.WebInfo.Code != 0 {

				// 打印输出
				// filters 为空代表没有使用 -filter 参数
				if contains(config.Filters, ipOption.WebInfo.Code) || config.Filters == nil {
					codeTmp := ""
					if 200 <= ipOption.WebInfo.Code && ipOption.WebInfo.Code < 300 {
						codeTmp += color.C256(46).Sprintf("[") + color.C256(196).Sprintf("%d", ipOption.WebInfo.Code) + color.C256(46).Sprintf("]")
					} else if 300 <= ipOption.WebInfo.Code && ipOption.WebInfo.Code < 400 {
						codeTmp += color.C256(46).Sprintf("[") + color.C256(154).Sprintf("%d", ipOption.WebInfo.Code) + color.C256(46).Sprintf("]")
					} else {
						codeTmp += color.C256(46).Sprintf("[") + color.C256(249).Sprintf("%d", ipOption.WebInfo.Code) + color.C256(46).Sprintf("]")
					}

					titleTmp := color.C256(46).Sprintf("[") + color.C256(165).Sprintf("%s", ipOption.WebInfo.Title) + color.C256(46).Sprintf("]")
					lenTmp := color.C256(46).Sprintf("[") + color.C256(200).Sprintf("%d", ipOption.WebInfo.Len) + color.C256(46).Sprintf("]")
					serverTmp := color.C256(46).Sprintf("[") + color.C256(226).Sprintf("%s", ipOption.WebInfo.Server) + color.C256(46).Sprintf("]")
					xPoweredByTmp := color.C256(46).Sprintf("[") + color.C256(51).Sprintf("%s", ipOption.WebInfo.XPoweredBy) + color.C256(46).Sprintf("]")
					url302JumpTmp := color.C256(46).Sprintf(" --> ") + color.C256(46).Sprintf("%s", ipOption.WebInfo.Url302Jump)
					tmpprint := ipOption.WebInfo.Url + "\t" + codeTmp

					if ipOption.WebInfo.Title != "" {
						tmpprint += titleTmp
					}
					tmpprint += " "
					if ipOption.WebInfo.Server != "" {
						tmpprint += serverTmp
					}
					if ipOption.WebInfo.XPoweredBy != "" {
						tmpprint += xPoweredByTmp
					}
					tmpprint += lenTmp
					if ipOption.WebInfo.Url302Jump != "" {
						tmpprint += url302JumpTmp
					}
					view.PrintlnSuccess(tmpprint)
				}

				// 导出
				if config.OutPutType != "" {
					if strings.Contains(config.OutPutType, ".csv") {
						path := config.OutPutType[:strings.Index(config.OutPutType, ".")] + "_webScan.csv"
						output.OutputFile(path, ipOption, config.MODE_WEB)
					} else {
						output.OutputFile(config.OutPutType, ipOption, config.MODE_WEB)
					}
				}
			}
			wg.Done()
			<-ThreadMaxChan
		}(option)
	}
	close(ThreadMaxChan)
	wg.Wait()
}
