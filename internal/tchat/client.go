package tchat

import (
	"net"
)

type Connection struct {
	Conn   net.Conn
	Status bool
}

// NewConnectionはコネクション状態を保持する。
func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn:   conn,
		Status: true,
	}
}

// Senderはchatサーバに対してメッセージを送信する。
func (c *Connection) SendToServer(b []byte) error {
	_, err := c.Conn.Write(b)
	return err
}

// Reflectorはchatサーバから受信したデータをstring型に変換する。
func (c *Connection) ReceiveFromServer() (string, error) {
	buf := makeBuffer()
	n, err := c.Conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}
