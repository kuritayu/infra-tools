package tchat

import (
	"net"
)

type Client struct {
	Name  string
	conn  net.Conn
	Color int
}

// CreateClientはクライアント情報を設定する。
func CreateClient(conn net.Conn, name string) *Client {
	return &Client{
		Name:  name,
		conn:  conn,
		Color: getColor(),
	}
}

//TODO Quit時はroomから削除しておく必要がある。
//TODO Read()がclient.goとserver.goで同じになっている
func (c *Client) Read() (string, error) {
	buf := MakeBuffer()
	n, err := c.conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
