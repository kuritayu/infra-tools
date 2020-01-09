package tchat

import (
	"fmt"
	"github.com/aybabtme/color/brush"
	"io"
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
		return brush.Red(msg).String()
	case GREEN:
		return brush.Green(msg).String()
	case YELLOW:
		return brush.Yellow(msg).String()
	case BLUE:
		return brush.Blue(msg).String()
	case PURPLE:
		return brush.Purple(msg).String()
	case CYAN:
		return brush.Cyan(msg).String()
	default:
		return msg
	}
}

func getTime() string {
	return time.Now().Format("15:04")
}

func makeMsg(msg []byte, name []byte, color int) []byte {
	template := fmt.Sprintf("%s[%s] %s", getTime(), name, string(msg))
	return []byte(SprintColor(template, color))
}

func makeMsgForAdmin(msg string) []byte {
	template := fmt.Sprintf("(%s) %s", getTime(), msg)
	return []byte(SprintColor(template, 31))
}

func makeBuffer() []byte {
	return make([]byte, BUFFERLENGTH)
}

func getColor() int {
	var colorList = [5]int{32, 33, 34, 35, 36}
	rand.Seed(time.Now().UnixNano())
	return colorList[rand.Intn(5)]
}

func getName(conn net.Conn) []byte {
	buf := makeBuffer()
	n, err := conn.Read(buf)
	if err != nil {
		//TODO 失敗したらerrorを返したほうがよい
		fmt.Println("Fail get name")
		Close(conn)
		os.Exit(1)
	}
	return buf[:n]
}

//TODO このメソッド不要の可能性高い
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
	}
}

func ChkErr(err error, place string) {
	if err != nil {
		fmt.Printf("(%s)", place)
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(0)
	}
}
