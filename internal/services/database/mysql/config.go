package database

import (
	"errors"
	"fmt"
)

type Config struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DbName string

	Params map[string]string
}

func (c Config) Validate() error {
	if c.Host == "" {
		return errors.New("database host is required")
	}
	if c.Port == 0 {
		return errors.New("database port is required")
	}
	if c.User == "" {
		return errors.New("database user is required")
	}
	if c.DbName == "" {
		return errors.New("database name is required")
	}
	return nil
}

//GetDriverSourceName returns a mySQL driver compatible data source name
func (c Config) GetDriverSourceName() string {
	var params string
	if len(c.Params) > 0 {
		var query string
		for key, val := range c.Params {
			if query != "" {
				query += "&"
			}
			query += key + "=" + val
		}
		params = "?" + query
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.DbName,
		params)
}
