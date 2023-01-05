package file

import (
	"cscan/config"
	"embed"
	"io"
	"log"
	"os"
	"strings"
)

// 读取文件
func ReadFile(filePath string) []byte {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil && config.Debug {
		log.Println(err)
	}
	body, err := io.ReadAll(file)
	if err != nil && config.Debug {
		log.Println(err)
	}
	return body
}

// 读取文件并按行转化成数组类型的 string 类型
func ReadFile2Strings4embed(dic embed.FS, filePath string) []string {
	resultByte, err := dic.ReadFile("dic/" + filePath)
	if err != nil {
		if config.Debug {
			log.Println(err)
		}
	}
	resultStr := string(resultByte)
	return strings.Split(resultStr, "\n")
}
