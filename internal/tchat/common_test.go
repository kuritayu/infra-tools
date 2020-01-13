package tchat

import (
	"fmt"
	"github.com/aybabtme/color/brush"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSprintColor(t *testing.T) {
	actual_red := SprintColor("test", 31)
	expected_red := brush.DarkRed("test").String()
	assert.Equal(t, expected_red, actual_red)

	actual_green := SprintColor("test", 32)
	expected_green := brush.DarkGreen("test").String()
	assert.Equal(t, expected_green, actual_green)

	actual_yellow := SprintColor("test", 33)
	expected_yellow := brush.DarkYellow("test").String()
	assert.Equal(t, expected_yellow, actual_yellow)

	actual_blue := SprintColor("test", 34)
	expected_blue := brush.DarkBlue("test").String()
	assert.Equal(t, expected_blue, actual_blue)

	actual_purple := SprintColor("test", 35)
	expected_purple := brush.DarkPurple("test").String()
	assert.Equal(t, expected_purple, actual_purple)

	actual_cyan := SprintColor("test", 36)
	expected_cyan := brush.DarkCyan("test").String()
	assert.Equal(t, expected_cyan, actual_cyan)

	actual_other := SprintColor("test", 37)
	expected_other := "test"
	assert.Equal(t, expected_other, actual_other)

}

func TestGetTime(t *testing.T) {
	actual := getTime()
	expected := time.Now().Format("15:04")
	assert.Equal(t, expected, actual)
}

func TestMakeMsg(t *testing.T) {
	now := getTime()
	name := "TEST-USER"
	msg := "Hello."
	actual := MakeMsg(msg, name, GREEN)
	expected := []byte(brush.DarkGreen(fmt.Sprintf("%s[%s] %s", now, name, msg)).String())
	assert.Equal(t, expected, actual)
}

func TestMakeBuffer(t *testing.T) {
	actual := makeBuffer()
	expected := make([]byte, 560)
	assert.Equal(t, expected, actual)
}

func TestGetColor(t *testing.T) {
	actual := getColor()
	assert.Contains(t, [...]int{32, 33, 34, 35, 36}, actual)
}
