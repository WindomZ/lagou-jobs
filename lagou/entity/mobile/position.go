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

type PositionMap map[string][]Position

func (p PositionMap) Add(jobs ...Position) {
	for _, job := range jobs {
		if job.CompanyId != 0 && len(job.CompanyName) != 0 {
			if _, ok := p[job.CompanyName]; !ok {
				p[job.CompanyName] = make([]Position, 0, 3)
			}
			for _, j := range p[job.CompanyName] {
				if j.PositionId == job.PositionId {
					return
				}
			}
			p[job.CompanyName] = append(p[job.CompanyName], job)
		}
	}
}

func (p PositionMap) Concat(others ...PositionMap) {
	for _, other := range others {
		for k, v := range other {
			if _, ok := p[k]; ok {
				p[k] = append(p[k], v...)
			} else {
				p[k] = v
			}
		}
	}
}

func (p PositionMap) Map() map[string][]Position {
	return map[string][]Position(p)
}
