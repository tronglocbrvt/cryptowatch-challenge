package config

// Config holds all settings
var defaultConfig = []byte(`
grpc_address: 10000
http_address: 9000
secret_key_access_jwt: c3Ryb25nc3RvcHBlZG1peGFscGhhYmV0c3BhY2VmaWdodHdhc3RlcHJvbWlzZWRpbmM=
secret_key_refresh_jwt: ZGVjaWRlY29tcG9zZWRzZXJpZXNzb3V0aGFjdHVhbGJhY2tzZWNvbmRsZWF2aW5nbWE=
jwt_access_token_expiration_minutes: 5
jwt_refresh_token_expiration_hours: 240
google_oauth_redirect_url: http://localhost:9000/auth/callback
google_oauth_client_id: 143556138247-huq355tmrjci95c8rjfmrbglrt98tcr5.apps.googleusercontent.com
google_oauth_client_secret: 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
google_oauth_scope: https://www.googleapis.com/auth/userinfo.email
google_oauth_api_url: https://www.googleapis.com/oauth2/v2/userinfo?access_token=
crypto_watch_api_key: Z1SO5DO6KG6XF6ECNAN4
crypto_watch_url: wss://stream.cryptowat.ch/connect?apikey=
time_interval_call_crypto_watch_second: 5
limit_get_prices_chart: 1500
postgre:
  address: ec2-3-230-122-20.compute-1.amazonaws.com
  port: 5432
  database: d9dj74icn7ml60
  username: hklvuymhxhwihu
  password: dfe9f4d933f1930a209af24dc4b46545670176e63e45d1f91ac56d9658adfdc8
`)

type Config struct {
	HTTPAddress                       int         `yaml:"http_address" mapstructure:"http_address"`
	GRPCAddress                       int         `yaml:"grpc_address" mapstructure:"grpc_address"`
	SecretKeyAccessJwt                string      `yaml:"secret_key_access_jwt" mapstructure:"secret_key_access_jwt"`
	SecretKeyRefreshJwt               string      `yaml:"secret_key_refresh_jwt" mapstructure:"secret_key_refresh_jwt"`
	JwtAccessTokenExpirationMinutes   int         `yaml:"jwt_access_token_expiration_minutes" mapstructure:"jwt_access_token_expiration_minutes"`
	JwtRefreshTokenExpirationHours    int         `yaml:"jwt_refresh_token_expiration_hours" mapstructure:"jwt_refresh_token_expiration_hours"`
	GoogleOauthRedirectUrl            string      `yaml:"google_oauth_redirect_url" mapstructure:"google_oauth_redirect_url"`
	GoogleOauthClientID               string      `yaml:"google_oauth_client_id" mapstructure:"google_oauth_client_id"`
	GoogleOauthClientSecret           string      `yaml:"google_oauth_client_secret" mapstructure:"google_oauth_client_secret"`
	GoogleOauthScope                  string      `yaml:"google_oauth_scope" mapstructure:"google_oauth_scope"`
	GoogleOauthApiUrl                 string      `yaml:"google_oauth_api_url" mapstructure:"google_oauth_api_url"`
	CryptoWatchApiKey                 string      `yaml:"crypto_watch_api_key" mapstructure:"crypto_watch_api_key"`
	CryptoWatchUrl                    string      `yaml:"crypto_watch_url" mapstructure:"crypto_watch_url"`
	TimeIntervalCallCryptoWatchSecond int         `yaml:"time_interval_call_crypto_watch_second" mapstructure:"time_interval_call_crypto_watch_second"`
	LimitGetPricesChart               int         `yaml:"limit_get_prices_chart" mapstructure:"limit_get_prices_chart"`
	PostgreSQL                        *PostgreSQL `yaml:"postgre" mapstructure:"postgre"`
}

type PostgreSQL struct {
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Address  string `yaml:"address" mapstructure:"address"`
	Database string `yaml:"database" mapstructure:"database"`
}
