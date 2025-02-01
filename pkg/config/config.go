package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var runtimeConfig *MallRuntimeConfig

const (
	EnvMallCoreConfigDir     = "Mall_CORE_CONFIG_DIR"
	DefaultMallCoreConfigDir = ""
)

type MySQLConfig struct {
	MySQLUsername string `mapstructure:"Mall_MYSQL_USERNAME"`
	MySQLPassword string `mapstructure:"Mall_MYSQL_PASSWORD"`
	MySQLDatabase string `mapstructure:"Mall_MYSQL_DATABASE"`
	MySQLHost     string `mapstructure:"Mall_MYSQL_HOST"`
	MySQLPort     string `mapstructure:"Mall_MYSQL_PORT"`
}

type GatewayConfig struct {
	GateWayHost string `mapstructure:"Mall_GATEWAY_HOST"`
	GateWayPort string `mapstructure:"Mall_GATEWAY_PORT"`
}

type UserServiceConfig struct {
	UserServiceHost string `mapstructure:"Mall_USER_SERVICE_HOST"`
	UserServicePort string `mapstructure:"Mall_USER_SERVICE_PORT"`
}

type MallRuntimeConfig struct {
	MySQLConfig       MySQLConfig       `mapstructure:"Mall_MYSQL_CONFIG"`
	GatewayConfig     GatewayConfig     `mapstructure:"Mall_GATEWAY_CONFIG"`
	UserServiceConfig UserServiceConfig `mapstructure:"Mall_USER_SERVICE_CONFIG"`
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetRuntimeConfig() *MallRuntimeConfig {
	v := viper.New()
	confDir := getEnv(EnvMallCoreConfigDir, DefaultMallCoreConfigDir)
	v.SetConfigName("conf.local")
	v.SetConfigType("yml")
	v.AddConfigPath(confDir)

	// important! do not miss this step
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	runtimeConfig = &MallRuntimeConfig{}
	err := v.Unmarshal(runtimeConfig)
	if err != nil {
		panic(err)
	}
	return runtimeConfig
}
