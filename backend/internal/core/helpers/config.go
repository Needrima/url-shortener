package helpers

import "github.com/spf13/viper"

type Config struct {
	RedisAddr   string `mapstructure:"REDIS_ADDR"`
	RedisPass   string `mapstructure:"REDIS_PASSWORD"`
	LogDir      string `mapstructure:"LOG_DIR"`
	LogFile     string `mapstructure:"LOG_FILE"`
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
	UsageTrials string `mapstructure:"USAGE_TRAILS"`
	Domain      string `mapstructure:"DOMAIN"`
}

func LoadEnv(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
