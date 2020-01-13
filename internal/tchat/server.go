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

const PORT = ":7777"

func createClient(conn net.Conn, name string) *Client {
	return &Client{
		name:  name,
		conn:  conn,
		color: getColor(),
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

		//TODO createClientをシンプルにしたので、Executeの処理が多くなっている
		//TODO 関数から別関数をgoroutineしているため、非常にわかりにくい、テストしにくい
		//TODO 標準出力に出力する情報とログに残すフォーマットはあわせたい
		//TODO ログハンドラもしたい
		//TODO メッセージの記録
		name, err := getName(conn)
		if err != nil {
			conn.Close()
			ChkErr(err, "getName")
		}
		cl := createClient(conn, name)
		room.add(cl)

		ch := make(chan []byte)
		go room.send(ch)
		ch <- makeMsg("joined!!", cl.name, RED)
		fmt.Println("User joined. name: ", cl.name)

		go func() {
			buf := makeBuffer()
			for {
				n, err := cl.conn.Read(buf)
				if err != nil {
					go room.send(ch)
					ch <- makeMsg("Quit.", cl.name, RED)
					fmt.Println("User left. name: ", cl.name)
					break
				}
				go room.send(ch)
				ch <- makeMsg(string(buf[:n]), cl.name, cl.color)
				buf = makeBuffer()
			}

		}()
	}
}
