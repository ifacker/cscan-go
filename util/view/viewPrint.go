package view

import "github.com/gookit/color"

// 打印成功的结果
func PrintlnSuccess(content any) {
	color.C256(46).Printf("[+] %v\n", content)
}

// 打印失败信息
func PrintlnFailed(content any) {
	color.C256(1).Printf("[-] %v\n", content)
}

// 打印基础信息
func PrintlnInfo(content any) {
	color.C256(247).Printf("[*] %v\n", content)
}

// 打印错误信息
func PrintlnError(content any) {
	color.Errorf("[Err] %v\n", content)
}
