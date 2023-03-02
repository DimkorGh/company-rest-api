package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
	Logger   Logger
	Kafka    Kafka
}

type Server struct {
	Port              string
	Mode              string
	JwtSecretKey      string
	ReadHeaderTimeout time.Duration
}

type Logger struct {
	Encoding string
	Level    string
}

type Database struct {
	Address        string
	Username       string
	Password       string
	DbName         string
	MigrationsPath string
}

type Kafka struct {
	Address string
	Client  string
}

func NewConfig() *Config {
	return &Config{}
}

func (cnf *Config) Load(configName, configPath string) *Config {
	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file: %s", err.Error())
	}

	if err := v.Unmarshal(cnf); err != nil {
		log.Fatalf("Error while unmarshalling config file into a struct: %s", err.Error())
	}

	return cnf
}
