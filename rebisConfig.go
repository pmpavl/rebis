package rebis

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	DefaultSize              = 1024
	DefaultDefaultExpiration = time.Duration(-1)
	DefaultCleanupInterval   = time.Duration(time.Minute * 5)
)

/*
	Config for create cache.
*/
type Config struct {
	Size              uintptr       `yaml:"size"`              // how many elements should fit into the cache
	Backup            Backup        `yaml:"backup"`            // meta backup
	DefaultExpiration time.Duration `yaml:"defaultExpiration"` // default time of life element
	CleanupInterval   time.Duration `yaml:"cleanupInterval"`   // interval for cleanup
	LogAll            bool          `yaml:"logAll"`            // log in standard out or not
	Evicted           bool          `yaml:"evicted"`           // do standard function with expired item
}

/*
	Backup is configuration for backup logic.
*/
type Backup struct {
	Path     string        `yaml:"path,omitempty"`     // path to save backup, must be like "./backup"
	Interval time.Duration `yaml:"interval,omitempty"` // interval for save backup, its hard operation
	InUse    bool          `yaml:"inUse"`              // use backup save or not
}

func configDefault() *Config {
	return &Config{
		Size: DefaultSize,
		Backup: Backup{
			InUse: false,
		},
		DefaultExpiration: DefaultDefaultExpiration,
		CleanupInterval:   DefaultCleanupInterval,
		LogAll:            false,
		Evicted:           false,
	}
}

/*
	Create default config for rebis cache container.
	Default config is described in the function configDefault().
*/
func ConfigCreateDefault(filename string) error {
	if !strings.Contains(filename, ".yaml") && !strings.Contains(filename, ".yml") {
		return errors.New("config file should be in yaml format")
	}

	buf := new(bytes.Buffer)
	c := configDefault()

	if err := yaml.NewEncoder(buf).Encode(c); err != nil {
		return err
	}

	if err := os.WriteFile(filename, buf.Bytes(), os.FileMode(0644)); err != nil { // nolint
		return err
	}

	return nil
}

/*
	ConfigFrom parse config from yaml file configuration.
*/
func ConfigFrom(filename string) (c *Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(file, &c); err != nil {
		return nil, err
	}

	return c, nil
}
