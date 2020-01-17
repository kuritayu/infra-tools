package tchat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func ClientExecute(port int) {
	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", getURI(port))
	ChkErr(err)

	// chatサーバへの接続
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	ChkErr(err)
	defer conn.Close()

	// 接続状態を構造体にセット
	connection := NewConnection(conn)

	// クライアントの名前を標準入力から取得
	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()

	// chatサーバへのデータ送信(クライアントの名前)
	err = connection.SendToServer(name)
	ChkErr(err)

	// chatサーバからメッセージを受信すると、標準出力に反映するためのゴルーチン
	go func() {
		// chatサーバからデータを受信
		for connection.Status {
			msg, err := Read(connection.Conn)
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

	for connection.Status {
		// メッセージ送信、受信用の待ち処理
		time.Sleep(time.Microsecond)
	}
}
