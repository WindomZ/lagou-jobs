package lagou

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	OutputJsonFile  string `json:"outputJsonFile"`
	UserAgent       string `json:"userAgent"`       // recommend UA from Chrome
	RequestTimeout  int    `json:"requestTimeout"`  // seconds
	RequestInterval int    `json:"requestInterval"` // milliseconds
	Search          struct {
		City     string   `json:"city"`     // the city
		Keywords []string `json:"keywords"` // how to search from lagou
		Company  struct {
			ExcludeId []int `json:"excludeId"` // must exclude CompanyId
		} `json:"company"`
		Position struct {
			ExcludeId []int `json:"excludeId"` // must exclude PositionId
			Filter    struct {
				Include []string `json:"include"` // must include words
				Exclude []string `json:"exclude"` // must exclude words
			} `json:"filter"`
		} `json:"position"`
	} `json:"search"`
}

func (c *Config) ReadConfig(filename string) error {
	if b, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}
