package config

const (
	MODE_PORT  = 1
	MODE_WEB   = 2
	MODE_CRACK = 3
)

// WebInfo web 信息
type WebInfo struct {
	Url        string //网站 URL
	Title      string // 网站标题
	Len        int64  // 返回 body 的长度
	Code       int    // 返回的状态码
	Url302Jump string // 302跳转后的链接
	Server     string // header 头中的 server 字段
	XPoweredBy string // header 头中的 X-Powered-By 字段
}

// CrackInfo 要爆破的端口的信息
type CrackInfo struct {
	UserName string
	Password string
	Deal     string
}

type IpOption struct {
	Ip             string
	Port           int
	PortOpenStatus bool // 端口开放状态
	WebInfo        WebInfo
	CrackInfo      CrackInfo
}

type IpOptions struct {
	IpOption []*IpOption
	Ips      []string
	Ports    []int
}
