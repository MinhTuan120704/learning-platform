package config

type AppConfig struct {
	Name string `env:"APP_NAME" envDefault:"identity-service"`
	Env  string `env:"APP_ENV" envDefault:"development"`
}
