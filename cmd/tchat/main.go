package main

import (
	"bufio"
	"fmt"
	"github.com/kuritayu/infra-tools/internal/tchat"
	"github.com/urfave/cli"
	"net"
	"os"
	"time"
)

var (
	SERVER = "127.0.0.1"
	PORT   = "7777"
	URI    = fmt.Sprintf("%s:%s", SERVER, PORT)
)

func main() {
	app := cli.NewApp()
	app.Name = "tchat"
	app.Usage = "chat tool by terminal"
	app.Version = "1.0"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name: "c",
		},
	}

	//TODO port番号は指定できるようにしたい
	app.Action = func(c *cli.Context) error {
		if c.Bool("c") {
			clientExecute()
		} else {
			serverExecute()
		}

		return nil
	}
	_ = app.Run(os.Args)
	os.Exit(0)

}

func serverExecute() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	tchat.ChkErr(err, "tcpaddr")
	li, err := net.ListenTCP("tcp", tcpAddr)
	tchat.ChkErr(err, "tcpaddr")
	fmt.Println("Listen start.")
	room := tchat.NewRoom()
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
		name, err := tchat.GetName(conn)
		if err != nil {
			_ = conn.Close()
			tchat.ChkErr(err, "getName")
		}
		cl := tchat.CreateClient(conn, name)
		room.Add(cl)

		ch := make(chan []byte)
		go room.Send(ch)
		ch <- tchat.MakeMsg("joined!!", cl.Name, tchat.RED)
		fmt.Println("User joined. name: ", cl.Name)
		go cl.Read(room)
	}
}

func clientExecute() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	tchat.ChkErr(err, "tcpAddr")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	tchat.ChkErr(err, "DialTCP")
	defer tchat.Teardown(conn)

	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()
	_, err = conn.Write(name)
	tchat.ChkErr(err, "Write name")

	go tchat.Reflector(conn)
	go tchat.Sender(conn)

	for tchat.Running {
		time.Sleep(1 * 1e9)
	}
}
