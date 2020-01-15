package tchat

import (
	"net"
)

type Connection struct {
	Conn   net.Conn
	Status bool
}

// NewConnectionはコネクション状態を設定する。
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
