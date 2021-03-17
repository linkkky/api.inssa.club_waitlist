package test

import (
	"testing"

	"github.com/tj/assert"
)

func Test(t *testing.T) {
	x1 := 1
	assert.Equal(t, true, x1 == 1)
}
