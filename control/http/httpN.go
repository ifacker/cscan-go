package http

import (
	"crypto/tls"
	"cscan/config"
	"cscan/util/log"
	"github.com/ifacker/myutil"
	"net/http"
)

func NewClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 500,
		MaxConnsPerHost:     500,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	err := myutil.InitProxy(tr, config.Proxy)
	if err != nil {
		log.Print(err)
	}
	httpClient := &http.Client{
		Transport: tr,
		Timeout:   config.TimeOut,
	}
	return httpClient
}
