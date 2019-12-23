package lstar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcCheckSumForFile(t *testing.T) {
	actual := CalcCheckSumForFile("../../test/test.tar")
	expected := "b441b2f9a3e8a6154f60a1ef6509e9bf"
	assert.Equal(t, expected, actual)
}
