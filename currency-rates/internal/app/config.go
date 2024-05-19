package app

type Config struct {
	Address string `envconfig:"ADDRESS"`
	DBUrl string `envconfig:"DB_URL"`
}
