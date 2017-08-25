package lagou

import "strings"

// Filter lagou jobs filter
type Filter struct {
	Company struct {
		ExcludeID map[int]bool // must exclude CompanyID
	}
	Position struct {
		ExcludeID map[int]bool // must exclude PositionID
		Filter    struct {
			Include []string // must include words
			Exclude []string // must exclude words
		}
		Salary struct {
			Min int // min salary
			Max int // max salary
		}
	}
}

func (s *Spider) initFilter() error {
	s.Filter.Company.ExcludeID = make(map[int]bool, len(s.Config.Search.Company.ExcludeID))
	for _, id := range s.Config.Search.Company.ExcludeID {
		s.Filter.Company.ExcludeID[id] = true
	}

	s.Filter.Position.ExcludeID = make(map[int]bool, len(s.Config.Search.Position.ExcludeID))
	for _, id := range s.Config.Search.Position.ExcludeID {
		s.Filter.Position.ExcludeID[id] = true
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

	s.Filter.Position.Salary.Min = s.Config.Search.Position.Salary.Min
	s.Filter.Position.Salary.Max = s.Config.Search.Position.Salary.Max

	return nil
}

func (s Spider) filterCompanyID(id int) bool {
	if id <= 0 {
		return false
	}
	if _, ok := s.Filter.Company.ExcludeID[id]; ok {
		return false
	}
	return true
}

func (s Spider) filterPositionID(id int) bool {
	if id <= 0 {
		return false
	}
	if _, ok := s.Filter.Position.ExcludeID[id]; ok {
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
