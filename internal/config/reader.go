package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var (
	cfg *Config
)

func ReadFile(cfg *Config) (err error) {
	cfgPath := "./build/config/config.yaml"
	file, err := os.Open(cfgPath)
	if err != nil {
		return err
	}
	defer func() {
		closeerr := file.Close()
		if closeerr != nil {
			err = closeerr
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}
	return nil
}

func ReadEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}
func SetConfig(c *Config) {
	cfg = c
}
