package main

import (
	"bufio"
	"fmt"
	"github.com/comail/colog"
	"github.com/kuritayu/infra-tools/internal/tchat"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
	"time"
)

//TODO ログ定義がサーバとクライアントで重複
//TODO ログ出力が随所にでてみにくい

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
	// ログ定義
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Flag:        log.Ldate | log.Ltime | log.Lshortfile,
		HeaderPlain: nil,
		HeaderColor: nil,
		Colors:      true,
		NoColors:    false,
	})
	colog.Register()

	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	chkErr(err)
	log.Printf("trace: URIの解決が完了しました。")

	// リッスン開始
	li, err := net.ListenTCP("tcp", tcpAddr)
	chkErr(err)
	log.Printf("info: リッスンを開始しました。")

	// ルーム作成
	// 現時点ではサーバ起動時に"PUBLIC"ルームを常に作成している。
	// 最終形は、クライアントからルーム名を受け取り、該当のルームがなければ作成、
	// あれば既存のルームに参加するようにする
	room := tchat.NewRoom(ROOMNAME)
	log.Printf("info: ルーム[%s]を作成しました。", ROOMNAME)

	for {
		// コネクション確立
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}
		log.Printf("info: コネクションが確立されました。接続元: %s", conn.RemoteAddr())

		// 確立後の最初のデータからクライアントの名前を取得する。
		name, err := getName(conn)
		if err != nil {
			_ = conn.Close()
		}
		log.Printf("trace: クライアントの名前を取得しました。名前: %s", name)

		//TODO ここでルーム名をクライアントから受け取る。

		// クライアント情報を生成する。
		cl := tchat.NewClient(conn, name)
		log.Printf("trace: クライアント情報を生成しました。名前: %s", cl.Name)

		// クライアント情報をルームに格納する。
		room.Add(cl)
		log.Printf("trace: クライアント[%s]をルーム[%s]に追加しました。", cl.Name, room.Name)

		// クライアントが参加した旨をルームの参加者全員に配信する。
		ch := make(chan []byte)
		go room.Send(ch)
		ch <- tchat.MakeMsg("joined!!", cl.Name, tchat.RED)
		log.Printf(
			"info: クライアント[%s]がルーム[%s]に入室した情報を配信しました。", cl.Name, room.Name)

		// クライアントからのデータ受信を待つ。
		go func() {
			ch := make(chan []byte)
			for {
				go room.Send(ch)
				msg, err := tchat.Read(cl.Conn)
				log.Printf("trace: クライアント[%s]からデータを受信しました。", cl.Name)
				if err != nil {
					log.Printf(
						"trace: クライアント[%s]からエラーメッセージ(%s)を受信しました。",
						cl.Name, err)
					ch <- tchat.MakeMsg("Quit.", cl.Name, tchat.RED)
					log.Printf("info: クライアント[%s]が退室しました。", cl.Name)
					room.Delete(cl)
					log.Printf("trace: クライアント[%s]をルーム[%s]から削除しました。", cl.Name, room.Name)
					break
				}

				//TODO 特殊な文字(ex. %L)を受信すると、ルームに在席中のメンバ一覧を表示する。
				ch <- tchat.MakeMsg(msg, cl.Name, cl.Color)
				log.Printf(
					"trace: クライアント[%s]から受信したメッセージをルーム[%s]に配信しました。",
					cl.Name, room.Name)
			}
		}()

	}
}

func clientExecute() {
	// ログ定義
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Flag:        log.Ldate | log.Ltime | log.Lshortfile,
		HeaderPlain: nil,
		HeaderColor: nil,
		Colors:      true,
		NoColors:    false,
	})
	colog.Register()

	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	chkErr(err)

	// chatサーバへの接続
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	chkErr(err)
	defer conn.Close()

	// 接続状態を構造体にセット
	connection := tchat.NewConnection(conn)

	// クライアントの名前を標準入力から取得
	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()

	// chatサーバへのデータ送信(クライアントの名前)
	err = connection.SendToServer(name)
	chkErr(err)

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

func chkErr(err error) {
	if err != nil {
		log.Printf("error: エラーが発生しました。 [%s]", err)
		os.Exit(1)
	}
}
