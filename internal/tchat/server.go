package tchat

import (
	"fmt"
	"github.com/edo1z/go_simple_chat/util"
	"net"
)

type Client struct {
	name  []byte
	conn  net.Conn
	color int
}

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
		go send(makeMsg(buf[:n], cl.name, cl.color))
		buf = makeBuffer()
	}
}

func createClient(conn net.Conn) *Client {
	return &Client{
		name:  getName(conn),
		conn:  conn,
		color: getColor(),
	}
}

func ServerExecute() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	util.ChkErr(err, "tcpaddr")
	li, err := net.ListenTCP("tcp", tcpAddr)
	util.ChkErr(err, "tcpaddr")
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}

		//TODO createClientをシンプルにしたので、Executeの処理が多くなっている
		//TODO stringとbyteが混在していて見にくい、送信するときだけbyte、それ以外は常にstringでやりたい
		//TODO 関数から別関数をgoroutineしているため、非常にわかりにくい、テストしにくい
		cl := createClient(conn)
		clientList = append(clientList, cl)
		send(makeMsgForAdmin(string(cl.name) + " joined!!"))
		go receiver(cl)
	}
}
