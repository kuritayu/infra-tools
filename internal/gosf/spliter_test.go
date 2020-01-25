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
	actual, _ := SelectField("A B C", "1", config)
	expected := "A"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithNF(t *testing.T) {
	config := NewConfig()
	actual, _ := SelectField("A B C", "NF", config)
	expected := "C"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWith0(t *testing.T) {
	config := NewConfig()
	actual, _ := SelectField("A B C", "0", config)
	expected := "A B C"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithNFMinus1(t *testing.T) {
	config := NewConfig()
	actual, _ := SelectField("A B C", "NF-1", config)
	expected := "B"
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithNFMinusA(t *testing.T) {
	config := NewConfig()
	actual, _ := SelectField("A B C", "NF-A", config)
	expected := ""
	assert.Equal(t, expected, actual)
}

func TestSelectFieldWithSeparator(t *testing.T) {
	config := NewConfig()
	config.separator = ","
	actual, _ := SelectField("A,B,C", "1", config)
	expected := "A"
	assert.Equal(t, expected, actual)
}
