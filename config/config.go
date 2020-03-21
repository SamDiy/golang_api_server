package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config configuration for app
type Config struct {
	HostDb     string
	PortDb     string
	UserDb     string
	NameDb     string
	PasswordDb string
}

// GetConfig func returns config
func GetConfig(fileName string) Config {
	data, _ := ioutil.ReadFile(fileName)
	var config Config
	err := json.Unmarshal(data, &config)
	_ = err
	return config
}
