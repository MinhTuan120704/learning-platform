package config

type AppConfig struct {
	Env string `env:"APP_ENV" envDefault:"development"`
}
