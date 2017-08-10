package main

import "github.com/WindomZ/lagou-jobs/lagou"

func main() {
	spider := lagou.New()

	if err := spider.ReadConfig("./tests/asset/config.json"); err != nil {
		panic(err)
	}

	if err := spider.Start(); err != nil {
		panic(err)
	}
}
