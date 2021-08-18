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
		DefaultExpiration: -1,
		CleanupInterval:   600,
	}
}

func ConfigCreateDefault(path string) error {
	if !strings.Contains(path, ".yaml") || strings.Contains(path, "/") {
		return errors.New(fmt.Sprintf("%s wrong configuration path", path))
	}
	buf := new(bytes.Buffer)
	c := configDefault()

	err := yaml.NewEncoder(buf).Encode(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, []byte(buf.String()), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ConfigFrom(path string) (c *Config, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return nil, err
	}
	err = c.validLoggerLevel()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) validLoggerLevel() error {
	if c.LoggerLevel < -1 || c.LoggerLevel > 5 {
		return errors.New(fmt.Sprintf("%d wrong logger level", c.LoggerLevel))
	}
	return nil
}
