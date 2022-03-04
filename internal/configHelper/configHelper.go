package configHelper

import (
	"net/url"
	"os"
	"strconv"
)

// port: 8080
// db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable

var prefix string

type Config struct {
	Port   uint
	Db_url string
}

func SetPrefix(value string) {
	prefix = value
}

func (c *Config) ReadEnvVar() (err error) {
	var p int
	if p, err = strconv.Atoi(os.Getenv(prefix + "port")); err != nil {
		return err
	}
	c.Port = uint(p)
	c.Db_url = os.Getenv(prefix + "Db_url")

	if _, err := url.ParseRequestURI(c.Db_url); err != nil {
		return err
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
