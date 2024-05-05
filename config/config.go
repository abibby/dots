package config

import (
	"os"

	"github.com/spf13/viper"
)

func getString(key string) string {
	return os.ExpandEnv(viper.GetString(key))
}

func DotDir() string {
	return getString("dot_dir")
}
func HomeDir() string {
	return getString("home_dir")
}
