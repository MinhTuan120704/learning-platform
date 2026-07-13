package config

type JWTConfig struct {
	Secret              string `env:"JWT_SECRET,required"`
	AccessExpireMinutes int    `env:"JWT_ACCESS_EXPIRE_MINUTES" envDefault:"15"`
	RefreshExpireDays   int    `env:"JWT_REFRESH_EXPIRE_DAYS" envDefault:"7"`
	Issuer              string `env:"JWT_ISSUER" envDefault:"identity-service"`
}
