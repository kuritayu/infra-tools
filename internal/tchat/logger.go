package tchat

import (
	"github.com/comail/colog"
	"log"
	"os"
)

func logDefine() {
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
	logDefine()
	if err != nil {
		log.Printf("error: エラーが発生しました。 [%s]", err)
	}
}

func saveMessage(b []byte) error {
	f, err := os.OpenFile("message.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(string(b))
	return nil
}
