package config

// Config holds all settings
var defaultConfig = []byte(`
grpc_address: 10000
http_address: 9000
`)

type Config struct {
	HTTPAddress int `yaml:"http_address" mapstructure:"http_address"`
	GRPCAddress int `mapstructure:"grpc_address"`
}
