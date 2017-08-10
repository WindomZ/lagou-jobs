package lagou

import (
	"math"
	"regexp"
	"strconv"
)

var regSalary = regexp.MustCompile(`\d+`)

func parseSalary(salary string) (min, max int) {
	if s := regSalary.FindAllString(salary, -1); len(s) >= 1 {
		min, _ = strconv.Atoi(s[0])
		min *= 1000
		if len(s) >= 2 {
			max, _ = strconv.Atoi(s[1])
			max *= 1000
		}
	}
	if max <= 0 {
		max = math.MaxInt32
	}
	return
}

func (s Spider) filterSalary(salary string) bool {
	if s.Filter.Position.Salary.Min > 0 {
		min, max := parseSalary(salary)
		if s.Filter.Position.Salary.Max > 0 {
			if max < s.Filter.Position.Salary.Max ||
				min < s.Filter.Position.Salary.Min {
				return false
			}
		} else if max < s.Filter.Position.Salary.Min {
			return false
		}
	}
	return true
}
