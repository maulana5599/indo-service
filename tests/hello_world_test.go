package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHalo(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func TestString(t *testing.T) {
	assert.Equal(t, "Hello Eko", "Hello Eko", "Result must be 'Hello Eko'")
}
