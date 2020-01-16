package tchat

import (
	"github.com/comail/colog"
	"log"
	"net"
)

func Define() {
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

}

func ChkErr(err error) {
	if err != nil {
		log.Printf("error: エラーが発生しました。 [%s]", err)
	}
}

func PrintResolveTCPAddr() {
	log.Printf("trace: URIの解決が完了しました。")
}

func PrintListenTCP() {
	log.Printf("info: リッスンを開始しました。")
}

func PrintNewRoom(name string) {
	log.Printf("info: ルーム[%s]を作成しました。", name)
}

func PrintEstablishedConnection(addr net.Addr) {
	log.Printf("info: コネクションが確立されました。接続元: %s", addr)
}

func PrintGetClientName(name string) {
	log.Printf("trace: クライアントの名前を取得しました。名前: %s", name)
}

func PrintNewClient(name string) {
	log.Printf("trace: クライアント情報を生成しました。名前: %s", name)
}

func PrintAddClient(username string, roomname string) {
	log.Printf("trace: クライアント[%s]をルーム[%s]に追加しました。", username, roomname)
}

func PrintSendJoinInfo(username string, roomname string) {
	log.Printf("info: クライアント[%s]がルーム[%s]に入室した情報を配信しました。", username, roomname)
}

func PrintReceiveData(name string) {
	log.Printf("trace: クライアント[%s]からデータを受信しました。", name)
}

func PrintReceiveErrorMessage(name string, err error) {
	log.Printf("trace: クライアント[%s]からエラーメッセージ(%s)を受信しました。", name, err)
}

func PrintLeaveClient(name string) {
	log.Printf("info: クライアント[%s]が退室しました。", name)
}

func PrintDeleteClient(username string, roomname string) {
	log.Printf("trace: クライアント[%s]をルーム[%s]から削除しました。", username, roomname)
}

func PrintSendMessage(username string, roomname string) {
	log.Printf("trace: クライアント[%s]から受信したメッセージをルーム[%s]に配信しました。", username, roomname)
}
