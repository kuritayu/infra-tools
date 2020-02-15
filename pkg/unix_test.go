package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLs(t *testing.T) {
	actual, err := Ls(".")
	assert.NoError(t, err)
	assert.Contains(t, actual, "unix.go")
}

func TestGrep(t *testing.T) {
	actual := Grep("I'm creating Unix Tools.", "Unix")
	expected := "I'm creating Unix Tools."
	assert.Equal(t, expected, actual)
}

func TestSed(t *testing.T) {
	actual := Sed("I'm creating Unix Tools.", ".nix", "Linux")
	expected := "I'm creating Linux Tools."
	assert.Equal(t, expected, actual)
}
