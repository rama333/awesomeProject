package config

import (
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	ServerPort int `mapstructure:"server_port"`
	ZbxHost string `mapstructure:"zbxHost"`
	ZbxLogin string `mapstructure:"zbxLogin"`
	ZbxPassword string `mapstructure:"zbxPassword"`
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("blueprint")
	v.AutomaticEnv()

	//Config.DSN = v.Get("DSN").(string)
	//Config.ApiKey = v.Get("API_KEY").(string)
	v.SetDefault("server_port", 3000)
	v.SetDefault("zbxHost", "http://192.168.114.49/api_jsonrpc.php")
	v.SetDefault("zbxLogin", "oaist")
	v.SetDefault("zbxPassword", "Nt147")

	//for _, path := range configPaths {
	//	v.AddConfigPath(path)
	//}
	//if err := v.ReadInConfig(); err != nil {
	//	return fmt.Errorf("failed to read the configuration file: %s", err)
	//}
	return v.Unmarshal(&Config)
}
