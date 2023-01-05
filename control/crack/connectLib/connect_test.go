package connectLib

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	var connect Connect
	connect = &IpCrack{
		Ip:       "192.168.222.109",
		Port:     22,
		UserName: "parrot",
		Password: "1345",
	}
	connect.ConnectSSH()
	fmt.Println(connect)
}
