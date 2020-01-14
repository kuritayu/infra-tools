package tchat

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Connection struct {
	Conn   net.Conn
	Status bool
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn:   conn,
		Status: true,
	}
}

func (c *Connection) Sender() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, _ := reader.ReadLine()
		if string(input) == "\\q" {
			c.Status = false
			break
		}
		_, err := c.Conn.Write(input)
		ChkErr(err, "sender write")
	}
}

func (c *Connection) Reflector() {
	buf := makeBuffer()
	for c.Status {
		n, err := c.Conn.Read(buf)
		ChkErr(err, "Receiver read")
		fmt.Println(string(buf[:n]))
		buf = makeBuffer()
	}
}
