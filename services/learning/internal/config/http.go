package config

type HTTPConfig struct {
	Host               string `env:"HTTP_HOST" envDefault:"0.0.0.0"`
	Port               string `env:"HTTP_PORT" envDefault:"8083"`
	IdentityServiceUrl string `env:"IDENTITY_SERVICE_URL,required"`
	InternalApiKey     string `env:"INTERNAL_API_KEY,required"`
}
