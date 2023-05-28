package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x := 4
	y := 5
	expected := 9
	actual := Add(x, y)
	assert.Equal(t, expected, actual, "they should be equal")
}
