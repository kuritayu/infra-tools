package tchat

import (
	"fmt"
	"net"
)

type Client struct {
	name  []byte
	conn  net.Conn
	color int
}

const PORT = ":7777"

var clientList []*Client

func send(msg []byte) {
	for _, cl := range clientList {
		_, err := cl.conn.Write(msg)
		if err != nil {
			continue
		}
	}
}

func receiver(cl *Client) {
	buf := makeBuffer()
	for {
		n, err := cl.conn.Read(buf)
		if err != nil {
			go send(makeMsgForAdmin(string(cl.name) + " Quit."))
			break
		}
		//TODO チャネル化
		go send(makeMsg(buf[:n], cl.name, cl.color))
		buf = makeBuffer()
	}
}

func createClient(conn net.Conn, name []byte) *Client {
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
		//TODO stringとbyteが混在していて見にくい、送信するときだけbyte、それ以外は常にstringでやりたい
		//TODO 関数から別関数をgoroutineしているため、非常にわかりにくい、テストしにくい
		//TODO getNameはここで実行して、createClientの引数に渡したい
		name := getName(conn)
		cl := createClient(conn, name)
		clientList = append(clientList, cl)
		send(makeMsgForAdmin(string(cl.name) + " joined!!"))
		go receiver(cl)
	}
}
