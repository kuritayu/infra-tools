package tchat

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"time"
)

const BUFFERLENGTH = 560

func SprintColor(msg string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, msg)
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
