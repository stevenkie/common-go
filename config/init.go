package config

import (
	"bytes"
	"fmt"

	"github.com/fsnotify/fsnotify"
	log "github.com/sleey/common-go/log"
	"github.com/spf13/viper"
)

// NewConfigFromFile instantiate new configuration instance
// Config in this example is read-only, you can't set it on the fly
// so don't return any possible interface to do so
func NewConfigFromFile(configName, configType, filePath string, options NewConfigOptions) (err error) {
	if vInstances[configName] != nil {
		err = fmt.Errorf("Viper instance is already initialized for %s", configName)
		return
	}
	vInstances[configName] = viper.New()
	vInstances[configName].SetConfigType(configType)
	vInstances[configName].SetConfigFile(filePath)

	if options.Defaults != nil {
		defaults := options.Defaults
		for key, value := range defaults {
			vInstances[configName].SetDefault(key, value)
		}
	}

	if options.DefaultName != "" {
		DefaultConfigName = options.DefaultName
	}

	err = vInstances[configName].ReadInConfig()
	if err != nil {
		return
	}

	if options.IsWatch {
		vInstances[configName].WatchConfig()
		vInstances[configName].OnConfigChange(func(e fsnotify.Event) {
			log.WithField("configName", e.Name).Info("Config File Changed")
		})
	}

	return
}

// NewConfigFromString instantiate new configuration instance using string, useful for testing
func NewConfigFromString(configName, value string, options NewConfigOptions) (err error) {
	if vInstances[configName] != nil {
		err = fmt.Errorf("Viper instance is already initialized for %s", configName)
		return
	}
	vInstances[configName] = viper.New()
	valueByte := []byte(value)
	viper.ReadConfig(bytes.NewBuffer(valueByte))

	if options.Defaults != nil {
		defaults := options.Defaults
		for key, value := range defaults {
			vInstances[configName].SetDefault(key, value)
		}
	}
	return
}
