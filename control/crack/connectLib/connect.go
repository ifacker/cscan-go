package connectLib

type IpCrack struct {
	Ip          string
	Port        int
	UserName    string
	Password    string
	CrackStatus bool // 爆破是否成功
}

type Connect interface {
	ConnectSSH()
	ConnectMssql()
	ConnectMysql()
	ConnectFtp()
	ConnectSmb()
	ConnectMemcached()
	ConnectMongodb()
	ConnectOracle()
	ConnectPostgres()
	ConnectRedis()
	//ConnectRdp()
}
