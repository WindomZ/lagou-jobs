package main

import (
	"fmt"

	"github.com/WindomZ/lagou-jobs/lagou"
)

func main() {
	spider := lagou.New()
	err := spider.ReadConfig("./asset/config.json")
	if err != nil {
		panic(err)
	}

	println(spider.UserAgent)

	ps, err := spider.SearchPositions("深圳", "go")
	if err != nil {
		panic(err)
	}
	for k, v := range ps.Map() {
		println("-------------------------------------")
		println(fmt.Sprintf("公司：%v", k))
		for _, p := range v {
			println(fmt.Sprintf("职位：%v    薪资：%v    时间：%v", p.PositionName, p.Salary, p.CreateTime))
		}
	}
}
