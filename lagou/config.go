package lagou

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	UserAgent       string `json:"user-agent"`
	RequestInterval int    `json:"request-interval"`
}

func (c *Config) ReadConfig(filename string) error {
	if b, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}
