package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type JWTConfig struct {
	Issuer        string `mapstructure:"ISSUER"`
	Secret        string `mapstructure:"SECRET"`
	SecretRefresh string `mapstructure:"SECRET_REFRESH"`
}

type AppConfiguration struct {
	Name    string `mapstructure:"APP_NAME"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort int    `mapstructure:"APP_PORT"`
}

type DatabaseConfig struct {
	Name     string `mapstructure:"DB_NAME"`
	Port     int    `mapstructure:"DB_PORT"`
	Host     string `mapstructure:"DB_ADDRESS"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
}

type Config struct {
	AppCfg *AppConfiguration
	DbCfg  *DatabaseConfig
	JwtCfg *JWTConfig
}

func NewConfig() *Config {
	return &Config{
		AppCfg: getAppConfiguration(),
		DbCfg:  getDatabaseConfig(),
		JwtCfg: getJwtConfig(),
	}
}

var appConfig *AppConfiguration
var dbConfig *DatabaseConfig
var jwtConfig *JWTConfig
var path string = "./"
var configType string = "env"
var configName string = ".env"

var mu sync.Mutex

func getDatabaseConfig() *DatabaseConfig {
	mu.Lock()

	defer mu.Unlock()

	if dbConfig == nil {
		viper.AddConfigPath(path)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		// Init config
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Config] - Failed to read config %v", err)
		}

		var dbCfg DatabaseConfig

		if err := viper.Unmarshal(&dbCfg); err != nil {
			panic("unable to unmarshal config")
		}

		dbConfig = &dbCfg
	}

	return dbConfig
}

func getAppConfiguration() *AppConfiguration {

	mu.Lock()

	defer mu.Unlock()

	if appConfig == nil {
		viper.AddConfigPath(path)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		// Init config
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Config] - Failed to read config %v", err)
		}
		var cfg AppConfiguration

		if err := viper.Unmarshal(&cfg); err != nil {
			panic("unable to unmarshal config")
		}

		appConfig = &cfg
	}

	return appConfig
}

func getJwtConfig() *JWTConfig {
	mu.Lock()

	defer mu.Unlock()

	if jwtConfig == nil {
		viper.AddConfigPath(path)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		// Init config
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Config] - Failed to read config %v", err)
		}
		var cfg JWTConfig

		if err := viper.Unmarshal(&cfg); err != nil {
			panic("unable to unmarshal config")
		}

		jwtConfig = &cfg
	}

	return jwtConfig
}
