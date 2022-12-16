package http

import (
	"cscan/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	config.Proxy = "socks5://localhost:1080"
	clinet := NewClient()
	req, err := http.NewRequest("GET", "https://baidu.com", nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := clinet.Do(req)
	if err != nil {
		log.Println(err)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
}
