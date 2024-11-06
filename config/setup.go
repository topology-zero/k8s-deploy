package config

import (
	"image/color"
	"log"
	"strings"

	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
)

var (
	cfgCnf     = Cfg{}
	ServerConf = Server{}
	JwtConf    = Jwt{}
	LogConf    = Log{}
	MysqlConf  = Mysql{}
	K8sConf    = K8S{}
	Captcha    *base64Captcha.Captcha
)

func init() {
	Captcha = base64Captcha.NewCaptcha(base64Captcha.NewDriverString(
		50,
		150,
		100,
		3,
		4,
		"234567890abcdefghjkmnpqrstuvwxyz",
		&color.RGBA{R: 240, G: 240, B: 246, A: 246},
		nil,
		[]string{"wqy-microhei.ttc"},
	), base64Captcha.DefaultMemStore)
}

// Setup 可以使用环境变量 SERVER_PORT=6666 覆盖
func Setup(fileName string) {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("读取配置文件错误", err)
	}

	if err := viper.Unmarshal(&cfgCnf); err != nil {
		log.Panic("配置文件格式错误", err)
	}

	ServerConf = cfgCnf.Server
	JwtConf = cfgCnf.Jwt
	MysqlConf = cfgCnf.Mysql
	LogConf = cfgCnf.Log
	K8sConf = cfgCnf.K8S
}
