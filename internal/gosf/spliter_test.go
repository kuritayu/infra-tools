package gosf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	actual := NewConfig().separator
	expected := " "
	assert.Equal(t, expected, actual)
}

func TestSelectField(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "1")
	expected := "A"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithNF(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "NF")
	expected := "C"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWith0(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "0")
	expected := "A B C"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithNFMinus1(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "NF-1")
	expected := "B"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithNFMinusA(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "NF-A")
	expected := ""
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithSeparator(t *testing.T) {
	config := NewConfig()
	config.separator = ","
	actual, _ := Concat("A,B,C", config, "1")
	expected := "A"
	assert.Equal(t, expected, actual)
}

func TestConcat(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "1", "2")
	expected := "A B"
	assert.Equal(t, expected, actual)
}

func TestConcatWithNF(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "1", "NF-1", "NF")
	expected := "A B C"
	assert.Equal(t, expected, actual)
}

func TestConcatArrange(t *testing.T) {
	config := NewConfig()
	actual, _ := Concat("A B C", config, "2", "1")
	expected := "B A"
	assert.Equal(t, expected, actual)
}
