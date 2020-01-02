package decotail

import (
	"github.com/aybabtme/color/brush"
	"github.com/hpcloud/tail"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
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

func TestConvertToMultiColorOK(t *testing.T) {
	tw := New("/var/log/system.log", true, "A B C D E F G")
	actual := tw.convertToColor("A B C D E F G")
	expected := strings.Join([]string{
		brush.Red("A").String(),
		brush.Blue("B").String(),
		brush.Yellow("C").String(),
		brush.Green("D").String(),
		brush.Purple("E").String(),
		brush.Cyan("F").String(),
		"G",
	}, " ")
	assert.Equal(t, expected, actual)
}

func TestDecisionWithTimeStamp(t *testing.T) {
	tw := New("/var/log/system.log", true, "B")
	line := &tail.Line{
		Text: "AAA",
		Time: time.Date(2020, 1, 1, 11, 59, 0, 0, time.UTC),
		Err:  nil,
	}
	actual := tw.decision(line)
	expected := "2020-01-01 11:59:00.000 1577879940 AAA\n"
	assert.Equal(t, expected, actual)

}

func TestDecisionWithoutTimeStamp(t *testing.T) {
	tw := New("/var/log/system.log", false, "B")
	line := &tail.Line{
		Text: "AAA",
		Time: time.Date(2020, 1, 1, 11, 59, 0, 0, time.UTC),
		Err:  nil,
	}
	actual := tw.decision(line)
	expected := "AAA\n"
	assert.Equal(t, expected, actual)

}

func TestExecute(t *testing.T) {
	tw := New("/var/log/system.log", true, "A")
	var err error
	go func() {
		err = tw.Execute()
	}()
	assert.Equal(t, nil, err)
}
