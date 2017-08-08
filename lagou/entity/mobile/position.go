package mobile

type Position struct {
	City            string `json:"city"`
	CompanyId       int    `json:"companyId"`
	CompanyName     string `json:"companyName"`
	CompanyFullName string `json:"companyFullName"`
	PositionId      int    `json:"positionId"`
	PositionName    string `json:"positionName"`
	Salary          string `json:"salary"`
	CreateTime      string `json:"createTime"`
}

type Positions map[string][]Position

func (p Positions) Add(jobs ...Position) {
	for _, job := range jobs {
		if job.CompanyId != 0 && len(job.CompanyName) != 0 {
			if _, ok := p[job.CompanyName]; !ok {
				p[job.CompanyName] = make([]Position, 0, 3)
			}
			p[job.CompanyName] = append(p[job.CompanyName], job)
		}
	}
}

func (p Positions) Map() map[string][]Position {
	return map[string][]Position(p)
}
