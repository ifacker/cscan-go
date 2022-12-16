package config

type IpOption struct {
	Ip     string
	Port   int
	Status bool
}

type IpOptions struct {
	IpOption []*IpOption
	Ips      []string
	Ports    []int
}
