package tchat

import (
	"errors"
	"fmt"
	"github.com/aybabtme/color/brush"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSprintColor(t *testing.T) {
	actualRed := SprintColor("test", 31)
	expectedRed := brush.DarkRed("test").String()
	assert.Equal(t, expectedRed, actualRed)

	actualGreen := SprintColor("test", 32)
	expectedGreen := brush.DarkGreen("test").String()
	assert.Equal(t, expectedGreen, actualGreen)

	actualYellow := SprintColor("test", 33)
	expectedYellow := brush.DarkYellow("test").String()
	assert.Equal(t, expectedYellow, actualYellow)

	actualBlue := SprintColor("test", 34)
	expectedBlue := brush.DarkBlue("test").String()
	assert.Equal(t, expectedBlue, actualBlue)

	actualPurple := SprintColor("test", 35)
	expectedPurple := brush.DarkPurple("test").String()
	assert.Equal(t, expectedPurple, actualPurple)

	actualCyan := SprintColor("test", 36)
	expectedCyan := brush.DarkCyan("test").String()
	assert.Equal(t, expectedCyan, actualCyan)

	actualOther := SprintColor("test", 37)
	expectedOther := "test"
	assert.Equal(t, expectedOther, actualOther)

}

func TestGetTime(t *testing.T) {
	actual := getTime()
	expected := time.Now().Format("2006-01-02 15:04")
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
	actual := MakeBuffer()
	expected := make([]byte, 560)
	assert.Equal(t, expected, actual)
}

func TestGetColor(t *testing.T) {
	actual := getColor()
	assert.Contains(t, [...]int{32, 33, 34, 35, 36}, actual)
}

func TestRead(t *testing.T) {
	connection := NewConnection(new(MockConn))
	actual, actualErr := Read(connection.Conn)
	expectedErr := errors.New("dummy")

	assert.Equal(t, "", actual)
	assert.Equal(t, expectedErr, actualErr)

}
