package tchat

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var Running = true

func Sender(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, _ := reader.ReadLine()
		if string(input) == "\\q" {
			Running = false
			break
		}
		_, err := conn.Write(input)
		ChkErr(err, "sender write")
	}
}

func Reflector(conn net.Conn) {
	buf := makeBuffer()
	for Running == true {
		n, err := conn.Read(buf)
		ChkErr(err, "Receiver read")
		fmt.Println(string(buf[:n]))
		buf = makeBuffer()
	}
}

func Teardown(c io.Closer) {
	err := c.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
	}
}
