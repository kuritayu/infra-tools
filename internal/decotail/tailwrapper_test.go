package decotail

import (
	"github.com/aybabtme/color/brush"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToColorOK(t *testing.T) {
	tw := New("/var/log/system.log", true, "AAAA")
	actual := tw.convertToColor("AAAA")
	expected := brush.Red("AAAA").String()
	assert.Equal(t, expected, actual)
}

func TestConvertToColorNG(t *testing.T) {
	tw := New("/var/log/system.log", true, "AAAA")
	actual := tw.convertToColor("AAAB")
	expected := "AAAB"
	assert.Equal(t, expected, actual)
}
