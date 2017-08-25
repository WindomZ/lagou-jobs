package lagou

import (
	"fmt"

	"github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
)

// ResultPosition defines structure of lagou position result
type ResultPosition struct {
	PositionID   int    `json:"positionID"`
	PositionName string `json:"positionName"`
	PositionURL  string `json:"positionURL"`
	Salary       string `json:"salary"`
	CreateTime   string `json:"createTime"`
}

// ResultCompany defines structure of lagou company result
type ResultCompany struct {
	CompanyID       int              `json:"companyID"`
	CompanyName     string           `json:"companyName"`
	CompanyFullName string           `json:"companyFullName"`
	CompanyURL      string           `json:"companyURL"`
	Positions       []ResultPosition `json:"positions"`
}

// Results defines all lagou results
type Results struct {
	Companies []ResultCompany `json:"companies"`
}

func (r *Results) fromPositionMap(p mobile.PositionMap) error {
	r.Companies = make([]ResultCompany, 0, len(p))
	for _, v := range p.Map() {
		if len(v) == 0 {
			continue
		}
		c := ResultCompany{
			CompanyID:       v[0].CompanyID,
			CompanyName:     v[0].CompanyName,
			CompanyFullName: v[0].CompanyFullName,
			CompanyURL:      fmt.Sprintf("https://www.lagou.com/gongsi/%v.html", v[0].CompanyID),
			Positions:       make([]ResultPosition, len(v)),
		}
		for i, p := range v {
			c.Positions[i] = ResultPosition{
				PositionID:   p.PositionID,
				PositionName: p.PositionName,
				PositionURL:  fmt.Sprintf("https://www.lagou.com/jobs/%v.html", p.PositionID),
				Salary:       p.Salary,
				CreateTime:   p.CreateTime,
			}
		}
		r.Companies = append(r.Companies, c)
	}
	return nil
}
