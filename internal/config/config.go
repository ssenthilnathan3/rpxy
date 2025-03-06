package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type resource struct {
	Name            string
	Endpoint        string
	Destination_url string
}

type configuration struct {
	Server struct {
		Host        string
		Listen_port string
	}
	Resources []resource
}

var Configuration *configuration

func NewConfiguration() (*configuration, error) {
	viper.AddConfigPath("data")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("error occurred when reading config %s", err)
	}

	err = viper.Unmarshal(&Configuration)

	if err != nil {
		return nil, fmt.Errorf("error occurred when unmarshaling %s", err)
	}
	return Configuration, nil
}
