package mysql

import "fmt"

type Config struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DbName string

	Params map[string]string
}

func (c *Config) GetDriverSourceName() string {
	var params string
	if len(params) > 0 {
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
