package config

// Config holds all settings
var defaultConfig = []byte(`
grpc_address: 10000
http_address: 9000
postgre:
  address: 127.0.0.1:5432
  database: crypto_watch
  username: loc.truong
  password: ""
  protocol: tcp
  parse_time: true
`)

type Config struct {
	HTTPAddress int         `yaml:"http_address" mapstructure:"http_address"`
	GRPCAddress int         `mapstructure:"grpc_address"`
	PostgreSQL  *PostgreSQL `yaml:"postgre" mapstructure:"postgre"`
}

type PostgreSQL struct {
	Username  string `yaml:"username" mapstructure:"username"`
	Password  string `yaml:"password" mapstructure:"password"`
	Protocol  string `yaml:"protocol" mapstructure:"protocol"`
	Address   string `yaml:"address" mapstructure:"address"`
	Database  string `yaml:"database" mapstructure:"database"`
	ParseTime bool   `yaml:"parse_time" mapstructure:"parse_time"`
}
