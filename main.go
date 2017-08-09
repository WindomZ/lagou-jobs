package main

import (
	"fmt"

	"github.com/WindomZ/lagou-jobs/lagou"
)

func main() {
	spider := lagou.New()

	if err := spider.ReadConfig("./asset/config.json"); err != nil {
		panic(err)
	}

	if err := spider.Start(); err != nil {
		panic(err)
	}

	m, err := spider.SearchPositionMaps("深圳", "go")
	if err != nil {
		panic(err)
	}
	for k, v := range m.Map() {
		println("-------------------------------------")
		println(fmt.Sprintf("公司：%v", k))
		for _, p := range v {
			println(fmt.Sprintf("职位：%v    薪资：%v    时间：%v", p.PositionName, p.Salary, p.CreateTime))
		}
	}
}
