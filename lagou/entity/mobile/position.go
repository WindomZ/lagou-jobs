package mobile

// Position defines structure of lagou position
type Position struct {
	City            string `json:"city"`
	CompanyID       int    `json:"companyID"`
	CompanyName     string `json:"companyName"`
	CompanyFullName string `json:"companyFullName"`
	PositionID      int    `json:"positionID"`
	PositionName    string `json:"positionName"`
	Salary          string `json:"salary"`
	CreateTime      string `json:"createTime"`
}

// PositionMap defines a map of positions with company name as the key
type PositionMap map[string][]Position

// Add add one or more Position jobs
func (p PositionMap) Add(jobs ...Position) {
	for _, job := range jobs {
		if job.CompanyID != 0 && len(job.CompanyName) != 0 {
			if _, ok := p[job.CompanyName]; !ok {
				p[job.CompanyName] = make([]Position, 0, 3)
			}
			for _, j := range p[job.CompanyName] {
				if j.PositionID == job.PositionID {
					return
				}
			}
			p[job.CompanyName] = append(p[job.CompanyName], job)
		}
	}
}

// Concat concat one or more PositionMap others
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

// Map convert to a map
func (p PositionMap) Map() map[string][]Position {
	return map[string][]Position(p)
}
