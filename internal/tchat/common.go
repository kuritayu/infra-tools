package tchat

import (
	"fmt"
	"github.com/aybabtme/color/brush"
	"math/rand"
	"net"
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

const SERVER = "127.0.0.1"

func getURI(port int) string {
	return fmt.Sprintf("%s:%d", SERVER, port)
}

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
	return time.Now().Format("2006-01-02 15:04")
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

func Read(conn net.Conn) (string, error) {
	buf := MakeBuffer()
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
