package tchat

import (
	"net"
)

type Client struct {
	Name  string
	Conn  net.Conn
	Color int
}

// CreateClientはクライアント情報を設定する。
func NewClient(conn net.Conn, name string) *Client {
	return &Client{
		Name:  name,
		Conn:  conn,
		Color: getColor(),
	}
}

//TODO Quit時はroomから削除しておく必要がある。
