package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
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

func TestCat(t *testing.T) {
	f, _ := os.Open("../README.md")
	actual := len(Cat(f))
	expected := 7
	assert.Equal(t, expected, actual)
}

func TestCut(t *testing.T) {
	actual := Cut("1 2 3", " ", 3, 1)
	expected := "3 1"
	assert.Equal(t, expected, actual)
}

func TestWc(t *testing.T) {
	data := []string{"a b c", "a b c", "a b c"}
	actual := Wc(data)
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestUniq(t *testing.T) {
	data := []string{"a b c", "a b c", "a b c"}
	actual := len(Uniq(data))
	expected := 1
	assert.Equal(t, expected, actual)
}
