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

//TODO 標準出力に出力する情報とログに残すフォーマットはあわせたい
//TODO ログハンドラもしたい
//TODO メッセージの記録

var (
	SERVER = "127.0.0.1"
	PORT   = "7777"
	URI    = fmt.Sprintf("%s:%s", SERVER, PORT)
)

const ESCAPESTRING = "\\q"

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
	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	tchat.ChkErr(err, "tcpaddr")

	// リッスン開始
	li, err := net.ListenTCP("tcp", tcpAddr)
	tchat.ChkErr(err, "tcpaddr")
	fmt.Println("Listen start.")

	// ルーム作成
	room := tchat.NewRoom()

	for {
		// コネクション確立
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}
		fmt.Println("Established connection. from: ", conn.RemoteAddr())

		// 確立後の最初のデータからクライアントの名前を取得する。
		//TODO GetNameはmainの中でメソッドにすればよい
		name, err := tchat.GetName(conn)
		if err != nil {
			_ = conn.Close()
			tchat.ChkErr(err, "getName")
		}

		// クライアント情報を生成する。
		cl := tchat.CreateClient(conn, name)

		// クライアント情報をルームに格納する。
		room.Add(cl)

		// クライアントが参加した旨をルームの参加者全員に配信する。
		ch := make(chan []byte)
		go room.Send(ch)
		ch <- tchat.MakeMsg("joined!!", cl.Name, tchat.RED)
		fmt.Println("User joined. name: ", cl.Name)

		// クライアントからのデータ受信を待つ。
		go cl.Read(room)
	}
}

func clientExecute() {
	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	tchat.ChkErr(err, "tcpAddr")

	// chatサーバへの接続
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	tchat.ChkErr(err, "DialTCP")
	defer conn.Close()

	// 接続状態を構造体にセット
	connection := tchat.NewConnection(conn)

	// クライアントの名前を標準入力から取得
	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()

	// chatサーバへのデータ送信(クライアントの名前)
	err = connection.SendToServer(name)
	tchat.ChkErr(err, "Write name")

	for connection.Status {
		// chatサーバからメッセージを受信すると、標準出力に反映するためのゴルーチン
		go func() {
			// chatサーバからデータを受信
			msg, err := connection.ReceiveFromServer()
			if err != nil {
				connection.Status = false
			}

			// 標準出力に書き込み
			fmt.Println(msg)
		}()

		// chatサーバにメッセージを送信するためにゴルーチン
		go func() {
			reader := bufio.NewReader(os.Stdin)
			for {

				// 標準入力からメッセージを取得
				input, _, _ := reader.ReadLine()
				if string(input) == ESCAPESTRING {
					connection.Status = false
					break
				}

				// chatサーバへのデータ送信
				err := connection.SendToServer(input)
				if err != nil {
					connection.Status = false
					break
				}
			}
		}()

		// メッセージ送信、受信用の待ち処理
		time.Sleep(time.Microsecond)
	}
}
