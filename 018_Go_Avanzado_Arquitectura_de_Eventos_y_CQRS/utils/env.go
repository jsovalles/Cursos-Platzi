package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

// NewEnv creates a new environment
// Constructs type Env, depends on Logger

func NewEnv() Env {

	// AddConfigPath called multiple times for testing purposes (viper look for config file from the path we call NewEnv)
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	env := Env{}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("☠️ cannot read configuration", err)
		} else {
			fmt.Println("☠️ config file was found but another error was produced", err)
		}
	}

	err := viper.Unmarshal(&env)
	if err != nil {
		fmt.Println("☠️ environment can't be loaded: ", err)
	}

	ForceMapping(&env)

	return env
}

func ForceMapping(env *Env) {

	if env.Environment == "" {
		env.Environment = viper.GetString("ENVIRONMENT")
	}

	if env.DatabaseUrl == "" {
		env.DatabaseUrl = viper.GetString("DATABASE_URL")
	}

}
