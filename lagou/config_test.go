package lagou

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestConfig_ReadConfig(t *testing.T) {
	c := new(Config)
	if err := c.ReadConfig("../tests/config.json"); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 15, c.RequestTimeout)
	assert.Equal(t, 0, c.RequestInterval)

	assert.Equal(t, "深圳", c.Search.City)
	assert.Equal(t, []string{"go"}, c.Search.Keywords)
	assert.Equal(t, []string{"go"}, c.Search.Position.Filter.Include)

	assert.Equal(t, "./tests/output/result.json", c.Output.Files.JSON)
}
