package lagou

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	OutputJsonFile  string `json:"outputJsonFile"`
	UserAgent       string `json:"userAgent"`       // recommend UA from Chrome
	RequestInterval int    `json:"requestInterval"` // millisecond
	Search          struct {
		City     string   `json:"city"`     // the city
		Keywords []string `json:"keywords"` // how to search from lagou
		Filter   struct {
			Include []string `json:"include"` // must include words
			Exclude []string `json:"exclude"` // must exclude words
		} `json:"filter"`
		Company struct {
			ExcludeId []string `json:"excludeId"` // must exclude CompanyId
		} `json:"company"`
		Position struct {
			ExcludeId []string `json:"excludeId"` // must exclude PositionId
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
