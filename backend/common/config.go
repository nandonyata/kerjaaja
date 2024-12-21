package common

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var c Config

type Config struct {
	General General  `mapstructure:"general"`
	DB      DBConfig `mapstructure:"db"`
}

type General struct {
	Env  string `mapstructure:"env"`
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	Schema        string `mapstructure:"schema"`
	MigrateSchema bool   `mapstructure:"migrate_schema"`
	Debug         bool   `mapstructure:"debug"`
}

func Load(path string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)     // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("can not load config: %s", err)
	}
}

func Get() Config {
	return c
}
