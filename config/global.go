package config

import (
	"github.com/spf13/viper"
)

var (

	// DefaultConfigName default config name for viper
	DefaultConfigName = "default"

	vInstances = make(map[string]*viper.Viper)
)
