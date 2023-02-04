package helpers

import "github.com/spf13/viper"

type Config struct {
	RedisAddr string `mapstructure:"REDIS_ADDR"`
	RedisPass string `mapstructure:"REDIS_PASSWORD"`
	LogDir    string `mapstructure:"LOG_DIR"`
	LogFile   string `mapstructure:"LOG_FILE"`
}

func LoadEnv(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
