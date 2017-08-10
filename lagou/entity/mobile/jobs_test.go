package mobile

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestJobDetail_MinSalary_MaxSalary(t *testing.T) {
	j := &JobDetail{
		Salary: "15K-30K",
	}

	assert.Equal(t, 15000, j.MinSalary())
	assert.Equal(t, 30000, j.MaxSalary())

	j.Salary = "15K起"
	assert.Equal(t, 15000, j.MinSalary())
	assert.Equal(t, math.MaxInt32, j.MaxSalary())

	j.Salary = "一千"
	assert.Equal(t, 0, j.MinSalary())
	assert.Equal(t, math.MaxInt32, j.MaxSalary())
}
