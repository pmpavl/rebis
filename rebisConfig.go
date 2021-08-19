package rebis

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Size              int           `yaml:"size"`
	Backup            Backup        `yaml:"backup"`
	DefaultExpiration time.Duration `yaml:"defaultExpiration"`
	CleanupInterval   time.Duration `yaml:"cleanupInterval"`
	Evicted           Evicted       `yaml:"evicted"`
}
type Backup struct {
	Path     string        `yaml:"path,omitempty"`
	Interval time.Duration `yaml:"interval,omitempty"`
	InUse    bool          `yaml:"inUse"`
}
type Evicted struct {
	Path  string `yaml:"path,omitempty"`
	InUse bool   `yaml:"inUse"`
}

func configDefault() *Config {
	return &Config{
		Size: 1024,
		Backup: Backup{
			InUse: false,
		},
		DefaultExpiration: time.Duration(-1),
		CleanupInterval:   time.Duration(time.Second * 5),
		Evicted: Evicted{
			InUse: false,
		},
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

	err := yaml.NewEncoder(buf).Encode(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

/*
	Parse config from yaml file configuration.
*/
func ConfigFrom(filename string) (c *Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
