package services

import (
	"github.com/spf13/viper"
)

var initiated = false

func init() {
	if !initiated {
		viper.SetConfigName("conf")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		viper.ReadInConfig()
		initiated = true
	}
}

type DefaultService struct{}

func (s *DefaultService) Init() {

}
