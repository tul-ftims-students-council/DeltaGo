package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBurl string `mapstructure:"DB_URL"`

	SMTPUsername string `mapstructure:"SMTP_USERNAME"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
