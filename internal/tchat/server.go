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

var clientList []*Client

func send(ch <-chan []byte) {
	msg := <-ch
	for _, cl := range clientList {
		_, err := cl.conn.Write(msg)
		if err != nil {
			continue
		}
	}
}

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
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}

		//TODO createClientをシンプルにしたので、Executeの処理が多くなっている
		//TODO 関数から別関数をgoroutineしているため、非常にわかりにくい、テストしにくい
		name, err := getName(conn)
		if err != nil {
			conn.Close()
			ChkErr(err, "getName")
		}
		cl := createClient(conn, name)
		clientList = append(clientList, cl)

		ch := make(chan []byte)
		go send(ch)
		ch <- makeMsg("joined!!", cl.name, RED)

		go func() {
			buf := makeBuffer()
			for {
				n, err := cl.conn.Read(buf)
				if err != nil {
					go send(ch)
					ch <- makeMsg("Quit.", cl.name, RED)
					break
				}
				go send(ch)
				ch <- makeMsg(string(buf[:n]), cl.name, cl.color)
				buf = makeBuffer()
			}

		}()
	}
}
