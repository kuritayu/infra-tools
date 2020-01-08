package tchat

import (
	"fmt"
	"github.com/edo1z/go_simple_chat/util"
	"io"
	"math/rand"
	"net"
	"os"
	"time"
)

type Client struct {
	name  []byte
	conn  net.Conn
	color int
}

var clientList []*Client
var colorList = [5]int{32, 33, 34, 35, 36}

func send(msg []byte) {
	for _, cl := range clientList {
		_, err := cl.conn.Write(msg)
		if err != nil {
			continue
		}
	}
}

func receiver(cl *Client) {
	buf := make([]byte, 560)
	for {
		n, err := cl.conn.Read(buf)
		if err != nil {
			go send(makeMsgForAdmin(string(cl.name) + " Quit."))
			break
		}
		go send(makeMsg(buf[:n], cl.name, cl.color))
		buf = make([]byte, 560)
	}
}

func createClient(conn net.Conn) *Client {
	return &Client{
		name:  getName(conn),
		conn:  conn,
		color: getColor(),
	}
}

func getName(conn net.Conn) []byte {
	buf := make([]byte, 560)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Fail get name")
		Close(conn)
		os.Exit(1)
	}
	return buf[:n]
}

func getColor() int {
	rand.Seed(time.Now().UnixNano())
	return colorList[rand.Intn(5)]
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
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
		cl := createClient(conn)
		clientList = append(clientList, cl)
		send(makeMsgForAdmin(string(cl.name) + " joined!!"))
		go receiver(cl)
	}
}
