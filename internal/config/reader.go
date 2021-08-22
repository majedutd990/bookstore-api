package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var (
	cfg *Config
	// errorUnknownFilExtension is returned by fileExtension Function
	// when the file extension is not supported by parse function
	errorUnknownFilExtension = errors.New("file extension not allowed")
)

// parses path of the config file
func Parser(path string, cfg *Config) error {
	switch filepath.Ext(path) {
	case "yaml", "yml":
		return parseYaml(path, cfg)
	default:
		return errorUnknownFilExtension
	}
}

//  parse a yaml file
func parseYaml(path string, cfg *Config) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = e
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
