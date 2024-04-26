package identify

import (
	"cscan/config"
	"fmt"
	"testing"
)

func TestIpRange(t *testing.T) {
	config.LogConfigInit()
	config.Debug = true
	//IpRange("192.168.11.1-254")
	//IpRange("192.168.11.1")
	result, _ := IpRange(" 192.168.12.1-20, 192.168.3.3, 192.168.1.1/24")
	fmt.Println(len(result), result)

}
