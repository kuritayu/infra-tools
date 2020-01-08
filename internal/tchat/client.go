package tchat

import (
	"bufio"
	"fmt"
	"github.com/edo1z/go_simple_chat/util"
	"io"
	"net"
	"os"
	"time"
)

var running = true

func sender(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, _ := reader.ReadLine()
		if string(input) == "\\q" {
			running = false
			break
		}
		_, err := conn.Write(input)
		util.ChkErr(err, "sender write")
	}
}

func reflector(conn net.Conn) {
	buf := makeBuffer()
	for running == true {
		n, err := conn.Read(buf)
		util.ChkErr(err, "Receiver read")
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

func ClientExecute() {
	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()

	host := "127.0.0.1:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	util.ChkErr(err, "tcpAddr")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.ChkErr(err, "DialTCP")
	defer Teardown(conn)

	_, err = conn.Write(name)
	util.ChkErr(err, "Write name")

	go reflector(conn)
	go sender(conn)

	for running {
		time.Sleep(1 * 1e9)
	}
}
