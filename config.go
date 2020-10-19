package main

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig initializes the application configuration
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("/etc/google-domains-ddns-updater/")
	viper.AddConfigPath("$HOME/.config/google-domains-ddns-updater/")
	viper.AddConfigPath(".")

	viper.SetDefault("http.port", "8000")
	viper.SetDefault("http.path_prefix", "")
	viper.SetDefault("config.toast_timeout", 1000)
	viper.SetDefault("config.default_locale", "en")
	viper.SetDefault("config.json_path", "data/hostnames.json")
	viper.SetDefault("config.cadence", "@hourly")

	viper.BindEnv("http.port", "PORT")
	viper.BindEnv("http.path_prefix", "PATH_PREFIX")
	viper.BindEnv("config.toast_timeout", "CONFIG_TOAST_TIMEOUT")
	viper.BindEnv("config.default_locale", "CONFIG_DEFAULT_LOCALE")
	viper.BindEnv("config.json_path", "JSONPATH")
	viper.BindEnv("config.cadence", "CADENCE")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}
}
