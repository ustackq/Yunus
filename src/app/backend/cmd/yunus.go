package cmd

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

// LoadConfig load config from fs
func LoadConfig(fs afero.Fs, relativePath, configFileName string) (*viper.Viper, error) {
	var cfg = viper.New()
	cfg.SetFs(fs)

	if relativePath == "" {
		relativePath = "."
	}

	cfg.SetConfigFile(configFileName)

	if relativePath == "" {
		cfg.AddConfigPath(".")
	} else {
		cfg.AddConfigPath(relativePath)
	}
	err := cfg.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return nil, err
		}
		return nil, fmt.Errorf("Unable to load config file %s/%s,may be need validate your config ", relativePath, configFileName)
	}
	loadDefaultSettings(cfg)
	return cfg, nil
}

func loadDefaultSettings(v *viper.Viper) {
	v.SetDefault("test", "test")
}
