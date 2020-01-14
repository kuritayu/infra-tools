package tchat

import (
	"fmt"
	"net"
)

type Client struct {
	Name  string
	conn  net.Conn
	color int
}

// CreateClientはクライアント情報を設定する。
func CreateClient(conn net.Conn, name string) *Client {
	return &Client{
		Name:  name,
		conn:  conn,
		color: getColor(),
	}
}

//TODO Readがデータの読み込み、ルームメンバに対しての配信を担当しているため、わかりにくい。
// Readがroomを引数として必要としている点からもわかる。ReadはあくまでもReadし、文字列を返すことに特化させる。
func (c *Client) Read(r *room) {
	ch := make(chan []byte)
	buf := makeBuffer()
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			go r.Send(ch)
			ch <- MakeMsg("Quit.", c.Name, RED)
			fmt.Println("User left. name: ", c.Name)
			break
		}
		go r.Send(ch)
		ch <- MakeMsg(string(buf[:n]), c.Name, c.color)
		buf = makeBuffer()
	}
}
