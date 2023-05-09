package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	// context timeout in seconds
	CtxTimeout int

	SiginKey string

	LogLevel string
	HTTPPort string

	AuthConfigPath string

	AdminServiceHost string
	AdminServicePort string

	PostServiceHost string
	PostServicePort string

	LigaServiceHost string
	LigaServicePort string
}

func Load() Config {
	c := Config{}

	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "ravshan"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "r"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "admindb"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.Environment = cast.ToString(getOrReturnDefault("R", "develop"))

	c.SiginKey = cast.ToString(getOrReturnDefault("SIGNING_KEY", "ravshanSignIn"))

	c.AuthConfigPath = cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./config/rback_model.conf"))

	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	c.AdminServiceHost = cast.ToString(getOrReturnDefault("ADMIN_SERVICE_HOST", "localhost"))
	c.AdminServicePort = cast.ToString(getOrReturnDefault("ADMIN_SERVICE_PORT", "8000"))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToString(getOrReturnDefault("POST_SERVICE_PORT", "8010"))

	c.LigaServiceHost = cast.ToString(getOrReturnDefault("LIGA_SERVICE_HOST", "localhost"))
	c.LigaServicePort = cast.ToString(getOrReturnDefault("LIGA_SERVICE_PORT", "8020"))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
