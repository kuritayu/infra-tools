package tchat

import (
	"fmt"
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
//TODO 今の実装では、標準入力からの読み込み、メッセージ送信(Connへの書き込み)の2つを処理している(関心を分離)
func (c *Connection) SendToServer(b []byte) error {
	_, err := c.Conn.Write(b)
	return err
}

// Reflectorはchatサーバから受信したデータを標準出力に書き込む。
func (c *Connection) ReflectFromServer() {
	buf := makeBuffer()
	for c.Status {
		n, err := c.Conn.Read(buf)
		ChkErr(err, "Receiver read")
		fmt.Println(string(buf[:n]))
		buf = makeBuffer()
	}
}
