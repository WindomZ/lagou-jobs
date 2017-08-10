package main

import (
	"fmt"

	"github.com/WindomZ/go-commander"
	"github.com/WindomZ/go-develop-kit/path"
	"github.com/WindomZ/lagou-jobs/lagou"
)

func main() {
	// lagou-jobs
	commander.Program.
		Version("v1.0.0")

	// lagou-jobs <config-file>
	commander.Program.
		Command("<config-file>", "load the configuration file.").
		Description("add user configuration").
		Action(func(c commander.Context) error {
			configPath := c.MustString("<config-file>")
			if configPath == "." {
				configPath = "config.json" // default filename
			}

			if ok, err := path.IsExist(configPath); err != nil {
				return err
			} else if !ok {
				return fmt.Errorf("no \"%v\" file found", configPath)
			}

			spider, err := lagou.New(configPath)
			if err != nil {
				return err
			}

			return spider.Start()
		})

	commander.Program.Parse()

}
