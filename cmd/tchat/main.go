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

const (
	ESCAPESTRING = "\\q"
	ROOMNAME     = "PUBLIC"
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
	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	chkErr(err, "tcpaddr")

	// リッスン開始
	li, err := net.ListenTCP("tcp", tcpAddr)
	chkErr(err, "tcpaddr")
	fmt.Println("Listen start.")

	// ルーム作成
	// 現時点ではサーバ起動時に"PUBLIC"ルームを常に作成している。
	// 最終形は、クライアントからルーム名を受け取り、該当のルームがなければ作成、
	// あれば既存のルームに参加するようにする
	room := tchat.NewRoom(ROOMNAME)

	for {
		// コネクション確立
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}
		fmt.Println("Established connection. from: ", conn.RemoteAddr())

		// 確立後の最初のデータからクライアントの名前を取得する。
		name, err := getName(conn)
		if err != nil {
			_ = conn.Close()
			chkErr(err, "getName")
		}

		//TODO ここでルーム名をクライアントから受け取る。

		// クライアント情報を生成する。
		cl := tchat.NewClient(conn, name)

		// クライアント情報をルームに格納する。
		room.Add(cl)

		// クライアントが参加した旨をルームの参加者全員に配信する。
		ch := make(chan []byte)
		go room.Send(ch)
		ch <- tchat.MakeMsg("joined!!", cl.Name, tchat.RED)
		fmt.Println("User joined. name: ", cl.Name)

		// クライアントからのデータ受信を待つ。
		go func() {
			ch := make(chan []byte)
			for {
				go room.Send(ch)
				msg, err := tchat.Read(cl.Conn)
				if err != nil {
					go room.Send(ch)
					ch <- tchat.MakeMsg("Quit.", cl.Name, tchat.RED)
					fmt.Println("User left. name: ", cl.Name)
					room.Delete(cl)
					break
				}

				//TODO 特殊な文字(ex. %L)を受信すると、ルームに在席中のメンバ一覧を表示する。
				ch <- tchat.MakeMsg(msg, cl.Name, cl.Color)
			}
		}()

	}
}

func clientExecute() {
	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	chkErr(err, "tcpAddr")

	// chatサーバへの接続
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	chkErr(err, "DialTCP")
	defer conn.Close()

	// 接続状態を構造体にセット
	connection := tchat.NewConnection(conn)

	// クライアントの名前を標準入力から取得
	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()

	// chatサーバへのデータ送信(クライアントの名前)
	err = connection.SendToServer(name)
	chkErr(err, "Write name")

	// chatサーバからメッセージを受信すると、標準出力に反映するためのゴルーチン
	go func() {
		// chatサーバからデータを受信
		for connection.Status {
			msg, err := tchat.Read(connection.Conn)
			if err != nil {
				connection.Status = false
			}

			// 標準出力に書き込み
			fmt.Println(msg)
		}
	}()

	// chatサーバにメッセージを送信するためにゴルーチン
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {

			// 標準入力からメッセージを取得
			input, _, _ := reader.ReadLine() // ここがクローズされていないからと思われる
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

	for connection.Status {
		// メッセージ送信、受信用の待ち処理
		time.Sleep(time.Microsecond)
	}
}

func getName(conn net.Conn) (string, error) {
	buf := tchat.MakeBuffer()
	n, err := conn.Read(buf)
	if err != nil {
		return "unknown", err
	}
	return string(buf[:n]), nil
}

func chkErr(err error, place string) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "(%s) %s", place, err.Error())
		os.Exit(1)
	}
}
