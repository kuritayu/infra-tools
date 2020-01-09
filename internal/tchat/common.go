package tchat

import (
	"fmt"
	"github.com/aybabtme/color/brush"
	"math/rand"
	"net"
	"os"
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

func makeMsg(msg string, name string, color int) []byte {
	template := fmt.Sprintf("%s[%s] %s", getTime(), name, msg)
	return []byte(SprintColor(template, color))
}

func makeBuffer() []byte {
	return make([]byte, BUFFERLENGTH)
}

func getColor() int {
	var colorList = [5]int{32, 33, 34, 35, 36}
	rand.Seed(time.Now().UnixNano())
	return colorList[rand.Intn(5)]
}

func getName(conn net.Conn) (string, error) {
	buf := makeBuffer()
	n, err := conn.Read(buf)
	if err != nil {
		return "unknown", err
	}
	return string(buf[:n]), nil
}

func ChkErr(err error, place string) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "(%s) %s", place, err.Error())
		os.Exit(1)
	}
}
