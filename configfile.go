package gconfig

import (
	"fmt"

	"github.com/knadh/koanf"
	kjson "github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"
)

// GetConfigFromFile func read configuration from file and save in output interface.
func GetConfigFromFile(configName string, configPath, output interface{}) error {
	var k = koanf.New(".")

	currentPath := fmt.Sprintf("%s/%s.json", configPath, configName)

	if err := k.Load(file.Provider(currentPath), kjson.Parser()); err != nil {
		return fmt.Errorf("error loading config file: %s", err.Error())
	}

	if err := k.UnmarshalWithConf("", &output, koanf.UnmarshalConf{
		DecoderConfig: &mapstructure.DecoderConfig{
			Metadata:         nil,
			Result:           &output,
			WeaklyTypedInput: true,
			ErrorUnused:      true,
		},
	}); err != nil {
		return fmt.Errorf("error unmarshalling file: %s", err.Error())
	}

	return nil
}
