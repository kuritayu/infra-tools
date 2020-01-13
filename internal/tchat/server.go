package tchat

import (
	"fmt"
	"net"
)

type Client struct {
	name  string
	conn  net.Conn
	color int
}

//TODO ポートは可変(パラメータ化)
const PORT = ":7777"

func createClient(conn net.Conn, name string) *Client {
	return &Client{
		name:  name,
		conn:  conn,
		color: getColor(),
	}
}

func (c *Client) read(r *room) {
	ch := make(chan []byte)
	buf := makeBuffer()
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			go r.send(ch)
			ch <- makeMsg("Quit.", c.name, RED)
			fmt.Println("User left. name: ", c.name)
			break
		}
		go r.send(ch)
		ch <- makeMsg(string(buf[:n]), c.name, c.color)
		buf = makeBuffer()
	}
}

func ServerExecute() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", PORT)
	ChkErr(err, "tcpaddr")
	li, err := net.ListenTCP("tcp", tcpAddr)
	ChkErr(err, "tcpaddr")
	fmt.Println("Listen start.")
	room := newRoom()
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}
		fmt.Println("Established connection. from: ", conn.RemoteAddr())

		//TODO 標準出力に出力する情報とログに残すフォーマットはあわせたい
		//TODO ログハンドラもしたい
		//TODO メッセージの記録
		name, err := getName(conn)
		if err != nil {
			_ = conn.Close()
			ChkErr(err, "getName")
		}
		cl := createClient(conn, name)
		room.add(cl)

		ch := make(chan []byte)
		go room.send(ch)
		ch <- makeMsg("joined!!", cl.name, RED)
		fmt.Println("User joined. name: ", cl.name)
		go cl.read(room)
	}
}
