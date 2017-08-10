package lagou

import (
	"math"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestSalary_parseSalary(t *testing.T) {
	var min, max int

	min, max = parseSalary("15K-30K")
	assert.Equal(t, 15000, min)
	assert.Equal(t, 30000, max)

	min, max = parseSalary("15K起")
	assert.Equal(t, 15000, min)
	assert.Equal(t, math.MaxInt32, max)

	min, max = parseSalary("一千")
	assert.Equal(t, 0, min)
	assert.Equal(t, math.MaxInt32, max)
}
