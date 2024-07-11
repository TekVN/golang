package strings

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRand(t *testing.T) {
	src := "1234567890"
	testcase := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	for _, v := range testcase {
		target := Rand(src, uint(v))
		assert.Equal(t, v, len(target))
		num, err := strconv.Atoi(target)
		assert.NoError(t, err)
		assert.True(t, num >= 0)
	}

	target := Rand(src, 0)
	assert.Equal(t, 0, len(target))
}
