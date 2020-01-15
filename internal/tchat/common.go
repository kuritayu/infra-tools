package tchat

import (
	"fmt"
	"github.com/aybabtme/color/brush"
	"math/rand"
	"time"
)

const (
	BUFFERLENGTH = 560
	RED          = 31
	GREEN        = 32
	YELLOW       = 33
	BLUE         = 34
	PURPLE       = 35
	CYAN         = 36
)

func SprintColor(msg string, color int) string {
	switch color {
	case RED:
		return brush.DarkRed(msg).String()
	case GREEN:
		return brush.DarkGreen(msg).String()
	case YELLOW:
		return brush.DarkYellow(msg).String()
	case BLUE:
		return brush.DarkBlue(msg).String()
	case PURPLE:
		return brush.DarkPurple(msg).String()
	case CYAN:
		return brush.DarkCyan(msg).String()
	default:
		return msg
	}
}

func getTime() string {
	return time.Now().Format("15:04")
}

func MakeMsg(msg string, name string, color int) []byte {
	template := fmt.Sprintf("%s[%s] %s", getTime(), name, msg)
	return []byte(SprintColor(template, color))
}

func MakeBuffer() []byte {
	return make([]byte, BUFFERLENGTH)
}

func getColor() int {
	var colorList = [5]int{GREEN, YELLOW, BLUE, PURPLE, CYAN}
	rand.Seed(time.Now().UnixNano())
	return colorList[rand.Intn(5)]
}
