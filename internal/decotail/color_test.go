package decotail

import (
	"fmt"
	"github.com/aybabtme/color/brush"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRed(t *testing.T) {
	actual := DefineColor(0, "A")
	expected := fmt.Sprint(brush.Red("A"))
	assert.Equal(t, expected, actual)
}

func TestBlue(t *testing.T) {
	actual := DefineColor(1, "A")
	expected := fmt.Sprint(brush.Blue("A"))
	assert.Equal(t, expected, actual)
}

func TestYellow(t *testing.T) {
	actual := DefineColor(2, "A")
	expected := fmt.Sprint(brush.Yellow("A"))
	assert.Equal(t, expected, actual)
}

func TestGreen(t *testing.T) {
	actual := DefineColor(3, "A")
	expected := fmt.Sprint(brush.Green("A"))
	assert.Equal(t, expected, actual)
}

func TestPurple(t *testing.T) {
	actual := DefineColor(4, "A")
	expected := fmt.Sprint(brush.Purple("A"))
	assert.Equal(t, expected, actual)
}

func TestCyan(t *testing.T) {
	actual := DefineColor(5, "A")
	expected := fmt.Sprint(brush.Cyan("A"))
	assert.Equal(t, expected, actual)
}

func TestNothing(t *testing.T) {
	actual := DefineColor(6, "A")
	expected := "A"
	assert.Equal(t, expected, actual)
}
