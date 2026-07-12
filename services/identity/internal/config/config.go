package config

type Config struct {
	App   AppConfig
	HTTP  HTTPConfig
	DB    DatabaseConfig
	Redis RedisConfig
	JWT   JWTConfig
}
