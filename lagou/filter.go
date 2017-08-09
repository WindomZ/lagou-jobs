package lagou

import "strings"

type Filter struct {
	Company struct {
		ExcludeId map[int]bool // must exclude CompanyId
	}
	Position struct {
		ExcludeId map[int]bool // must exclude PositionId
		Filter    struct {
			Include []string // must include words
			Exclude []string // must exclude words
		}
	}
}

func (s *Spider) initFilter() {
	s.Filter.Company.ExcludeId = make(map[int]bool, len(s.Config.Search.Company.ExcludeId))
	for _, id := range s.Config.Search.Company.ExcludeId {
		s.Filter.Company.ExcludeId[id] = true
	}
	s.Filter.Position.ExcludeId = make(map[int]bool, len(s.Config.Search.Position.ExcludeId))
	for _, id := range s.Config.Search.Position.ExcludeId {
		s.Filter.Position.ExcludeId[id] = true
	}
	s.Filter.Position.Filter.Include = make([]string, 0, len(s.Config.Search.Position.Filter.Include))
	for _, include := range s.Config.Search.Position.Filter.Include {
		if include = strings.ToLower(include); len(include) != 0 {
			s.Filter.Position.Filter.Include = append(s.Filter.Position.Filter.Include, include)
		}
	}
	s.Filter.Position.Filter.Exclude = make([]string, 0, len(s.Config.Search.Position.Filter.Exclude))
	for _, exclude := range s.Config.Search.Position.Filter.Exclude {
		if exclude = strings.ToLower(exclude); len(exclude) != 0 {
			s.Filter.Position.Filter.Exclude = append(s.Filter.Position.Filter.Exclude, exclude)
		}
	}
}

func (s Spider) filterCompanyId(id int) bool {
	if id <= 0 {
		return false
	}
	if _, ok := s.Filter.Company.ExcludeId[id]; ok {
		return false
	}
	return true
}

func (s Spider) filterPositionId(id int) bool {
	if id <= 0 {
		return false
	}
	if _, ok := s.Filter.Position.ExcludeId[id]; ok {
		return false
	}
	return true
}

func (s Spider) filterString(str string) bool {
	if len(str) == 0 {
		return false
	}
	str = strings.ToLower(str)
	for _, include := range s.Filter.Position.Filter.Include {
		if strings.Contains(str, include) {
			for _, exclude := range s.Filter.Position.Filter.Exclude {
				if strings.Contains(str, exclude) {
					return false
				}
			}
			return true
		}
	}
	return false
}
