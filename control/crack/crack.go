package crack

import (
	"cscan/config"
	"cscan/control/crack/connectLib"
	"cscan/control/output"
	"cscan/util/file"
	"cscan/util/view"
	"embed"
	"github.com/gookit/color"
	"log"
	"strings"
	"sync"
)

//go:embed dic
var Dic embed.FS

var (
	usernamePath string
	passwordPath string
)

// 初始化 Dic 字典目录
func init() {
	ents, err := Dic.ReadDir("dic")
	if err != nil {
		if config.Debug {
			log.Println(err)
		}
		return
	}
	for _, ent := range ents {
		path := ent.Name()
		if strings.Contains(strings.ToLower(path), "username") {
			usernamePath = path
		}
		if strings.Contains(strings.ToLower(path), "password") {
			passwordPath = path
		}
	}
}

// 破解子模块
//func subCrack(deal string, ipOption *config.IpOption, passwords []string, wg *sync.WaitGroup, ThreadMaxChan chan int) {
//	ipOption.CrackInfo.Deal = deal
//	for _, username := range config.Userdict[deal] {
//		for _, passwordTmp := range passwords {
//			var password string
//			// 替换 superkey
//			if strings.Contains(passwordTmp, "{{USER}}") {
//				password = strings.ReplaceAll(passwordTmp, "{{USER}}", username)
//			} else {
//				password = passwordTmp
//			}
//			connect := &connectLib.IpCrack{
//				Ip:          ipOption.Ip,
//				Port:        ipOption.Port,
//				UserName:    username,
//				Password:    password,
//				CrackStatus: false,
//			}
//			wg.Add(1)
//			ThreadMaxChan <- 1
//			go func(connect *connectLib.IpCrack, ipOption *config.IpOption) {
//
//				connect.ConnectSSH()
//
//				ipOption.CrackInfo.UserName = connect.UserName
//				ipOption.CrackInfo.Password = connect.Password
//				if connect.CrackStatus {
//
//					// 打印输出
//					ipAndPort := color.C256(46).Sprintf("%s:%d\t", ipOption.Ip, ipOption.Port)
//					userName := color.C256(226).Sprintf("%s\t", ipOption.CrackInfo.UserName)
//					passWord := color.C256(51).Sprintf("%s\t", ipOption.CrackInfo.Password)
//					deal := color.C256(219).Sprintf("%s\n", ipOption.CrackInfo.Deal)
//					tmpprint := ipAndPort + userName + passWord + deal
//					view.PrintlnSuccess(tmpprint)
//
//					// 导出
//					if config.OutPutType != "" {
//						if strings.Contains(config.OutPutType, ".csv") {
//							path := config.OutPutType[:strings.Index(config.OutPutType, ".")] + "_Crack.csv"
//							output.OutputFile(path, ipOption, config.MODE_CRACK)
//						} else {
//							output.OutputFile(config.OutPutType, ipOption, config.MODE_CRACK)
//						}
//					}
//				}
//				wg.Done()
//				<-ThreadMaxChan
//			}(connect, ipOption)
//		}
//	}
//}

// 暴力破解模块
func StartCrack(ipOptions *config.IpOptions) {
	// 加载读取字典文件
	usernames := file.ReadFile2Strings4embed(Dic, usernamePath)
	// password 中含有 superkey，后期进行暴力破解的时候，会对 superkey 进行解析适配
	passwords := file.ReadFile2Strings4embed(Dic, passwordPath)

	wg := &sync.WaitGroup{}
	var ThreadMaxChan = make(chan int, config.ThreadMax)

	for _, username := range usernames {
		for _, passwordTmp := range passwords {
			var password string
			// 替换 superkey
			if strings.Contains(passwordTmp, "{{USER}}") {
				password = strings.ReplaceAll(passwordTmp, "{{USER}}", username)
			} else {
				password = passwordTmp
			}
			for _, ipOption := range ipOptions.IpOption {
				if ipOption.PortOpenStatus {
					connect := &connectLib.IpCrack{
						Ip:          ipOption.Ip,
						Port:        ipOption.Port,
						UserName:    username,
						Password:    password,
						CrackStatus: false,
					}

					wg.Add(1)
					ThreadMaxChan <- 1
					go func(connect *connectLib.IpCrack, ipOption *config.IpOption) {
						// 根据端口来选择要使用的爆破协议
						switch ipOption.Port {
						case config.PostsCrackMap["ssh"]:
							ipOption.CrackInfo.Deal = "ssh"
							connect.ConnectSSH()
						case config.PostsCrackMap["mssql"]:
							ipOption.CrackInfo.Deal = "mssql"
							connect.ConnectMssql()
						case config.PostsCrackMap["mysql"]:
							ipOption.CrackInfo.Deal = "mysql"
							connect.ConnectMysql()
						case config.PostsCrackMap["ftp"]:
							ipOption.CrackInfo.Deal = "ftp"
							connect.ConnectFtp()
						case config.PostsCrackMap["smb"]:
							ipOption.CrackInfo.Deal = "smb"
							connect.ConnectSmb()
						case config.PostsCrackMap["memcached"]:
							ipOption.CrackInfo.Deal = "memcached"
							connect.ConnectMemcached()
						case config.PostsCrackMap["mongodb"]:
							ipOption.CrackInfo.Deal = "mongodb"
							connect.ConnectMongodb()
						case config.PostsCrackMap["oracle"]:
							ipOption.CrackInfo.Deal = "oracle"
							connect.ConnectOracle()
						case config.PostsCrackMap["postgres"]:
							ipOption.CrackInfo.Deal = "postgres"
							connect.ConnectPostgres()
						case config.PostsCrackMap["redis"]:
							ipOption.CrackInfo.Deal = "redis"
							connect.ConnectRedis()
							//case config.PostsCrackMap["rdp"]:
							//	ipOption.CrackInfo.Deal = "rdp"
							//	connect.ConnectRdp()
						}
						ipOption.CrackInfo.UserName = connect.UserName
						ipOption.CrackInfo.Password = connect.Password
						if connect.CrackStatus {

							// 打印输出
							ipAndPort := color.C256(46).Sprintf("%s:%d\t", ipOption.Ip, ipOption.Port)
							userName := color.C256(226).Sprintf("%s\t", ipOption.CrackInfo.UserName)
							passWord := color.C256(51).Sprintf("%s\t", ipOption.CrackInfo.Password)
							deal := color.C256(219).Sprintf("%s\n", ipOption.CrackInfo.Deal)
							tmpprint := ipAndPort + userName + passWord + deal
							view.PrintlnSuccess(tmpprint)

							// 导出
							if config.OutPutType != "" {
								if strings.Contains(config.OutPutType, ".csv") {
									path := config.OutPutType[:strings.Index(config.OutPutType, ".")] + "_Crack.csv"
									output.OutputFile(path, ipOption, config.MODE_CRACK)
								} else {
									output.OutputFile(config.OutPutType, ipOption, config.MODE_CRACK)
								}
							}
						}
						wg.Done()
						<-ThreadMaxChan
					}(connect, ipOption)
				}
			}
		}
	}
	close(ThreadMaxChan)
	wg.Wait()

	//// 新的代码
	//// password 中含有 superkey，后期进行暴力破解的时候，会对 superkey 进行解析适配
	//passwords := file.ReadFile2Strings4embed(Dic, passwordPath)
	//
	//wg := &sync.WaitGroup{}
	//var ThreadMaxChan = make(chan int, config.ThreadMax)
	//
	//for _, ipOption := range ipOptions.IpOption {
	//	if ipOption.PortOpenStatus {
	//		switch ipOption.Port {
	//		case config.PostsCrackMap["ssh"]:
	//			subCrack("ssh", ipOption, passwords, wg, ThreadMaxChan)
	//
	//		}
	//	}
	//}
	//close(ThreadMaxChan)
	//wg.Wait()
}
