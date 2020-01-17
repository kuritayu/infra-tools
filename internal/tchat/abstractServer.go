package tchat

import (
	"fmt"
	"net"
)

const (
	ESCAPESTRING = "\\q"
	ROOMNAME     = "PUBLIC"
)

func ServerExecute() {
	// URIの解決
	tcpAddr, err := net.ResolveTCPAddr("tcp4", URI)
	ChkErr(err)

	// リッスン開始
	li, err := net.ListenTCP("tcp", tcpAddr)
	ChkErr(err)

	// ルーム作成
	room := NewRoom(ROOMNAME)

	for {
		// コネクション確立
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}

		// 確立後の最初のデータからクライアントの名前を取得する。
		name, err := getName(conn)
		if err != nil {
			_ = conn.Close()
		}

		// クライアント情報を生成する。
		cl := NewClient(conn, name)

		// クライアント情報をルームに格納する。
		room.Add(cl)

		// クライアントが参加した旨をルームの参加者全員に配信する。
		ch := make(chan []byte)
		go room.Send(ch)
		ch <- MakeMsg("joined!!", cl.Name, RED)

		// クライアントからのデータ受信を待つ。
		go func() {
			ch := make(chan []byte)
			for {
				go room.Send(ch)
				msg, err := Read(cl.Conn)
				if err != nil {
					ch <- MakeMsg("Quit.", cl.Name, RED)
					room.Delete(cl)
					break
				}

				ch <- MakeMsg(msg, cl.Name, cl.Color)
			}
		}()

	}
}

func getName(conn net.Conn) (string, error) {
	buf := MakeBuffer()
	n, err := conn.Read(buf)
	if err != nil {
		return "unknown", err
	}
	return string(buf[:n]), nil
}
