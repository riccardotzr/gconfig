package gconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

// GetConfigFromFile read configuration from file and save in output interface
func GetConfigFromFile(path string, fileName string, output interface{}) error {

	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("Config file not found: %s", err.Error())
		}

		return fmt.Errorf("Error reading config file: %s", err.Error())
	}

	if err := viper.UnmarshalExact(&output); err != nil {
		return fmt.Errorf("Unable to decode into struct: %s", err.Error())
	}

	return nil
}
