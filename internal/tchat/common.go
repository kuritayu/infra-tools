package tchat

import (
	"fmt"
	"time"
)

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
