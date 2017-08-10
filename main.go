package main

import "github.com/WindomZ/lagou-jobs/lagou"

func main() {
	spider, err := lagou.New("./tests/asset/config.json")
	if err != nil {
		panic(err)
	}

	if err := spider.Start(); err != nil {
		panic(err)
	}
}
