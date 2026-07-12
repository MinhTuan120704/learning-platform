package config

type JWTConfig struct {
	Secret              string `env:"JWT_SECRET,required"`
	AccessExpireMinutes string `env:"JWT_ACCESS_EXPIRE_MINUTES" envDefault:"15"`
	RefreshExpireDays   string `env:"JWT_REFRESH_EXPIRE_DAYS" envDefault:"7"`
}
