package models

type ServerConfig struct {
	Port      string `json:"port"`
	CoinGecko string `json:"api"`
}

type Config struct {
	Server ServerConfig `json:"server"`
}

type Env struct {
	ServerMode           string `env:"SERVER_MODE,notEmpty"`
	Domain               string `env:"DOMAIN,notEmpty"`
	LogDirectory         string `env:"LOG_DIRECTORY,notEmpty"`
	ConfigPath           string `env:"CONFIG_PATH,notEmpty"`
	IsVerifyDependencies bool   `env:"IS_VERIFY_DEPENDENCIES,notEmpty"`
}
