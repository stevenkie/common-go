package config

import (
	"github.com/sirupsen/logrus"
	log "github.com/sleey/common-go/log"
)

// Get get configuration value from an instance type interface
func Get(key string, configName ...string) (value interface{}) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = vInstances[name].Get(key)
	return
}

// GetInt get configuration value from an instance type int
func GetInt(key string, configName ...string) (value int) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = vInstances[name].GetInt(key)
	return
}

// GetInt64 get configuration value from an instance type int64
func GetInt64(key string, configName ...string) (value int64) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = int64(vInstances[name].GetInt(key))
	return
}

// GetFloat64 get configuration value from an instance type float64
func GetFloat64(key string, configName ...string) (value float64) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = vInstances[name].GetFloat64(key)
	return
}

// GetString get configuration value from an instance type string
func GetString(key string, configName ...string) (value string) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = vInstances[name].GetString(key)
	return
}

// GetBool get configuration value from an instance type bool
func GetBool(key string, configName ...string) (value bool) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = vInstances[name].GetBool(key)
	return
}

// GetStringSlice get configuration value from an instance type slice of string
func GetStringSlice(key string, configName ...string) (value []string) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = vInstances[name].GetStringSlice(key)
	return
}

// GetStringMapString get configuration value from an instance type map string of string
func GetStringMapString(key string, configName ...string) (value map[string][]string) {
	name := DefaultConfigName

	if len(configName) == 1 {
		name = configName[0]
	}

	if vInstances[name] == nil {
		log.WithFields(logrus.Fields{
			"name": name,
		}).Error("Viper instance isn't initialized")
		return
	}

	value = vInstances[name].GetStringMapStringSlice(key)
	return
}
