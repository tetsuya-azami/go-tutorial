package configs

import (
	"go-tutorial/chapter8/pkg/logger"
	"os"
	"strconv"

	"go.uber.org/zap"
)

type ConfigList struct {
	Env                 string
	DBHost              string
	DBPort              int
	DBDriver            string
	DBName              string
	DBUser              string
	DBPassword          string
	APICorsAllowOrigins []string
}

var Config ConfigList

func init() {
	if err := LoadEnv(); err != nil {
		logger.Error("Failed to load env: ", zap.Error(err))
		panic(err)
	}
}

func LoadEnv() error {
	DBPort, err := strconv.Atoi(GetEnvDefault("MYSQL_PORT", "3306"))
	if err != nil {
		return err
	}
	Config = ConfigList{
		Env:                 GetEnvDefault("APP_ENV", "development"),
		DBHost:              GetEnvDefault("DB_HOST", "0.0.0.0"),
		DBPort:              DBPort,
		DBDriver:            GetEnvDefault("DB_DRIVER", "mysql"),
		DBName:              GetEnvDefault("DB_NAME", "api_database"),
		DBUser:              GetEnvDefault("DB_USER", "app"),
		DBPassword:          GetEnvDefault("DB_PASSWORD", "password"),
		APICorsAllowOrigins: []string{"http://0.0.0.0:8001"},
	}
	return nil
}

func GetEnvDefault(key, defaultValue string) string {
	val, isEnvFound := os.LookupEnv(key)
	if !isEnvFound {
		return defaultValue
	}

	return val
}

func (c *ConfigList) IsDevelopment() bool {
	return c.Env == "development"
}
