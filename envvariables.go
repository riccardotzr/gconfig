package gconfig

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// EnvConfig struct to access env variables
type EnvConfig struct {
	Key          string
	Variable     string
	DefaultValue string
	Required     bool
}

// GetEnvVariables extracts configured environment variables and unmarshals them in provided `output` interface.
func GetEnvVariables(envVariablesConfig []EnvConfig, output interface{}) error {
	v := viper.New()
	v.AutomaticEnv()

	for _, config := range envVariablesConfig {
		if err := setVariable(v, config); err != nil {
			return err
		}
	}

	if err := v.UnmarshalExact(&output); err != nil {
		return fmt.Errorf("Unable to decode into struct: %s", err.Error())
	}

	return nil
}

func setVariable(v *viper.Viper, env EnvConfig) error {
	if err := v.BindEnv(env.Variable, env.Key); err != nil {
		return fmt.Errorf("Unable to bind environment variables: %s", err.Error())
	}

	if env.DefaultValue != "" {
		v.SetDefault(env.Variable, env.DefaultValue)
	}

	if env.Required {
		_, ok := os.LookupEnv(env.Key)

		if !ok {
			return fmt.Errorf("")
		}
	}

	return nil
}
