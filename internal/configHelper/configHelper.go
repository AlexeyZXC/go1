package configHelper

import (
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// port: 8080
// db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable

var prefix string

type Config struct {
	Port   uint   `yaml:"port"`
	Db_url string `yaml:"db_url"`
}

func SetPrefix(value string) {
	prefix = value
}

func (c *Config) empty() {
	c.Db_url = ""
	c.Port = 0
}

func (c *Config) validate() (err error) {
	if c.Port > 65535 {
		err = errors.New("Invlaid port")
		c.empty()
		return
	}

	if _, err = url.ParseRequestURI(c.Db_url); err != nil {
		c.empty()
		return
	}

	return
}

func (c *Config) ReadEnvVar() (err error) {
	var p int
	if p, err = strconv.Atoi(os.Getenv(prefix + "port")); err != nil {
		return err
	}
	c.Port = uint(p)
	c.Db_url = os.Getenv(prefix + "Db_url")

	if err = c.validate(); err != nil {
		c = &Config{}
	}

	return
}

func (c Config) WriteEnvVar() (err error) {
	err = os.Setenv(prefix+"port", strconv.Itoa(int(c.Port)))
	if err != nil {
		return
	}
	err = os.Setenv(prefix+"db_url", c.Db_url)
	return err
}

//--- yaml

func (c *Config) ReadFromYaml() (err error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return
	}

	if err = c.validate(); err != nil {
		c = nil
	}

	return
}
