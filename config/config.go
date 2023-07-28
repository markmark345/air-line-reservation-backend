package config

import (
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var configOnce sync.Once
var config *Config

type Config struct {
	Server   Server   `mapstructure:"server"`
	Postgres Postgres `mapstructure:"Postgres"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type Postgres struct {
	DbName   string `mapstructure:"db-name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool-size"`
}

func NewConfig() *Config {
	configOnce.Do(func() {
		configPath, ok := os.LookupEnv("API_CONFIG_PATH")
		if !ok {
			configPath = "./config"
		}

		configName, ok := os.LookupEnv("API_CONFIG_NAME")
		if !ok {
			configName = "config"
		}

		viper.SetConfigName(configName)
		viper.AddConfigPath(configPath)

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
		}
		viper.AutomaticEnv()

		GetSecretValue()

		viper.WatchConfig() // Watch for changes to the configuration file and recompile
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})

	return config
}

func GetSecretValue() {
	for _, value := range os.Environ() {
		pair := strings.SplitN(value, "=", 2)
		if strings.Contains(pair[0], "SECRET_") {
			keys := strings.Replace(pair[0], "SECRET_", "secrets.", -1)
			keys = strings.Replace(keys, "_", ".", -1)
			newKey := strings.Trim(keys, " ")
			newValue := strings.Trim(pair[1], " ")
			viper.Set(newKey, newValue)
		}
	}
}
