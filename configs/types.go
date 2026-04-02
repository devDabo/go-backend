package configs

type Config struct {
	// Server
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerEnv  string `mapstructure:"SERVER_ENV"`
	// Database
	DBURL string `mapstructure:"DB_URL"`
}
