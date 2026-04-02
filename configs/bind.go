package configs

import (
	"strings"

	"github.com/spf13/viper"
)

func BindAllKeys() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	keys := []string{
		// Server
		ServerPort,
		ServerEnv,

		// Database
		DBURL,
	}

	for _, key := range keys {
		_ = viper.BindEnv(key)
	}
}
