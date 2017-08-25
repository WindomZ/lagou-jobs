package lagou

import (
	"encoding/json"
	"io/ioutil"
)

// Config the config.json file structure
type Config struct {
	UserAgent       string `json:"userAgent"`       // recommend UA from Chrome
	RequestTimeout  int    `json:"requestTimeout"`  // seconds
	RequestInterval int    `json:"requestInterval"` // milliseconds
	Search          struct {
		City     string   `json:"city"`     // the city
		Keywords []string `json:"keywords"` // how to search from lagou
		Company  struct {
			ExcludeID []int `json:"excludeID"` // must exclude CompanyID
		} `json:"company"`
		Position struct {
			ExcludeID []int `json:"excludeID"` // must exclude PositionID
			Filter    struct {
				Include []string `json:"include"` // must include words
				Exclude []string `json:"exclude"` // must exclude words
			} `json:"filter"`
			Salary struct {
				Min int `json:"min"` // min salary
				Max int `json:"max"` // max salary
			} `json:"salary"`
		} `json:"position"`
	} `json:"search"`
	Output struct {
		Files struct {
			JSON string `json:"json"`
		} `json:"files"`
		HTTP struct {
			Port int `json:"port"`
		} `json:"http"`
	} `json:"output"`
}

// ReadConfig Read the filename path file, and parse to Config.
func (c *Config) ReadConfig(filename string) error {
	if b, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}
