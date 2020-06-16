package configs

import (
	"os"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string
	Database string
}

type RpcConfig struct {
	Port int
}

type Configurarion struct {
	Mysql MysqlConfig
	Rpc   RpcConfig
}

var Config Configurarion

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	v := viper.New()
	v.SetConfigName(env)
	v.AddConfigPath("./configs")
	v.AutomaticEnv()
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(&Config)
}
