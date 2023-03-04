package config

import (
	"bytes"
	"strings"

	"log"

	"github.com/spf13/viper"
)

type Base struct {
	HTTPAddress int `yaml:"http_address" mapstructure:"http_address"`
	GRPCAddress int `mapstructure:"grpc_address"`
	Environment string
}

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Fatal("Failed to read viper config ", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal config", err)
	}

	return cfg
}
