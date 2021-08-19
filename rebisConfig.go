package rebis

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LoggerPath        string        `yaml:"loggerPath"`
	LoggerLevel       int8          `yaml:"loggerLevel"`
	DefaultExpiration time.Duration `yaml:"defaultExpiration"`
	CleanupInterval   time.Duration `yaml:"cleanupInterval"`
}

func configDefault() *Config {
	return &Config{
		LoggerPath:        "-1",
		LoggerLevel:       -1,
		DefaultExpiration: time.Duration(-1),
		CleanupInterval:   time.Duration(time.Second * 60),
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
	err = ioutil.WriteFile(filename, []byte(buf.String()), 0644)
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
	err = c.validConfig()
	if err != nil {
		return nil, err
	}
	return c, nil
}

/*
	Check the correctness of the compiled config.
*/
func (c *Config) validConfig() error {
	if c.LoggerLevel < -1 || c.LoggerLevel > 5 {
		return errors.New(fmt.Sprintf("%d is wrong logger level", c.LoggerLevel))
	}
	return nil
}
