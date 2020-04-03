package config

import (
	"log"
	"strings"

	"github.com/involvestecnologia/mole/models"
	"github.com/spf13/viper"
)

const (
	configFile    = "config"
	configFileExt = "json"
)

//Load - loads application settings
func Load() models.ReadConfig {

	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetConfigName(configFile)
	v.SetConfigType(configFileExt)

	v.AddConfigPath("../../configs/")
	v.AddConfigPath("configs/")

	readConfig := models.ReadConfig{}

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := v.Unmarshal(&readConfig); err != nil {
		log.Fatal(err)
	}

	return readConfig
}
