package mobile

import (
	"strconv"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestPositions_Add_Map(t *testing.T) {
	p := PositionMap{}

	for i := 1; i <= 10; i++ {
		p.Add(Position{
			CompanyId:   i,
			CompanyName: "aaa" + strconv.Itoa(i),
		})
	}

	assert.Equal(t, 10, len(p))

	p.Add(Position{
		CompanyId:   5,
		CompanyName: "aaa" + strconv.Itoa(5),
	})
	assert.Equal(t, 10, len(p))

	assert.NotEmpty(t, p.Map())
}
