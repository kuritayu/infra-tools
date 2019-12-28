package decotail

import (
	"fmt"
	"testing"
)

func TestConvertToColorOK(t *testing.T) {
	tw := New("/var/log/system.log", true, "AAAA")
	fmt.Println(tw.convertToColor("AAAA"))
}

func TestConvertToColorNG(t *testing.T) {
	tw := New("/var/log/system.log", true, "AAAA")
	fmt.Println(tw.convertToColor("AAAB"))
}
