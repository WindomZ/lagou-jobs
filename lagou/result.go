package lagou

import (
	"fmt"

	. "github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
)

type ResultPosition struct {
	PositionId   int    `json:"positionId"`
	PositionName string `json:"positionName"`
	PositionURL  string `json:"positionURL"`
	Salary       string `json:"salary"`
	CreateTime   string `json:"createTime"`
}

type ResultCompany struct {
	CompanyId       int              `json:"companyId"`
	CompanyName     string           `json:"companyName"`
	CompanyFullName string           `json:"companyFullName"`
	CompanyURL      string           `json:"companyURL"`
	Positions       []ResultPosition `json:"positions"`
}

type Results struct {
	Companies []ResultCompany `json:"companies"`
}

func (r *Results) fromPositionMap(p PositionMap) error {
	r.Companies = make([]ResultCompany, 0, len(p))
	for _, v := range p.Map() {
		if len(v) == 0 {
			continue
		}
		c := ResultCompany{
			CompanyId:       v[0].CompanyId,
			CompanyName:     v[0].CompanyName,
			CompanyFullName: v[0].CompanyFullName,
			CompanyURL:      fmt.Sprintf("https://www.lagou.com/gongsi/%v.html", v[0].CompanyId),
			Positions:       make([]ResultPosition, len(v)),
		}
		for i, p := range v {
			c.Positions[i] = ResultPosition{
				PositionId:   p.PositionId,
				PositionName: p.PositionName,
				PositionURL:  fmt.Sprintf("https://www.lagou.com/jobs/%v.html", p.PositionId),
				Salary:       p.Salary,
				CreateTime:   p.CreateTime,
			}
		}
		r.Companies = append(r.Companies, c)
	}
	return nil
}
