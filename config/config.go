package config

type Cfg struct {
	Server Server
	Jwt    Jwt
	Log    Log
	Mysql  Mysql
	K8S    K8S
}

type Server struct {
	Name string
	Host string
	Port int
}

type Jwt struct {
	Secret string
	Expire int
}

type Log struct {
	Dir   string
	Level string
}

type Mysql struct {
	Host string
	Port int
	User string
	Pwd  string
	Db   string
}

type K8S struct {
	WaitPod int
}
